package _1_config_pattern

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// 1. NewServer - базовый вариант, где все аргументы обязательны
func NewServer(addr string, port int, timeout time.Duration) (*http.Server, error) {
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", addr, port),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}, nil
}

// 2. NewServerWithConfig - вариант с конфигом. Обязательные аргументы передаем явно.
// Необязательные в структуре Config, это дает нам обратную совместимость, если параметр не указан, то используем значение по умолчанию.
const (
	defaultPort    = 8080
	defaultTimeout = 30 * time.Second
)

type Config struct {
	port    int
	timeout time.Duration
}

func NewServerWithConfig(addr string, cfg Config) (*http.Server, error) {
	if cfg.port == 0 {
		cfg.port = defaultPort
	}
	if cfg.timeout == 0 {
		cfg.timeout = defaultTimeout
	}
	return &http.Server{
		Addr:        fmt.Sprintf("%s:%d", addr, cfg.port),
		ReadTimeout: cfg.timeout,
	}, nil
}

// 3. Паттерн строитель. Позволяет полностью создать конфиг, а затем использовать его для создания сервера.
// - нужно всегда передавать конфиг в функцию создания сервера
// - громоздко создавать конфиг, потом его применять
// - не получится в функции With... вернуть ошибку, если хотим вызывать их друг за другом. Следовательно сложно валидировать данные в функциях With... Нужно переносить эту логику в Build.
type ConfigBuilder struct {
	port    int
	timeout time.Duration
}

func (b *ConfigBuilder) WithPort(port int) *ConfigBuilder {
	b.port = port
	return b
}

func (b *ConfigBuilder) WithTimeout(timeout time.Duration) *ConfigBuilder {
	b.timeout = timeout
	return b
}

func (b *ConfigBuilder) Build() (Config, error) {
	if b.port == 0 {
		b.port = defaultPort
	}
	if b.timeout == 0 {
		b.timeout = defaultTimeout
	}
	return Config{
		port:    b.port,
		timeout: b.timeout,
	}, nil
}

func NewServerWithBuilder(addr string, cfg Config) (*http.Server, error) {
	return &http.Server{
		Addr:        fmt.Sprintf("%s:%d", addr, cfg.port),
		ReadTimeout: cfg.timeout,
	}, nil
}

func TestConfigWithBuilder(t *testing.T) {
	cfgBuilder := (&ConfigBuilder{}).
		WithPort(8081).
		WithTimeout(10 * time.Second)

	cfg, err := cfgBuilder.Build()
	if err != nil {
		panic(err)
	}

	_, _ = NewServerWithBuilder("localhost", cfg)
}

// 4. Паттерн функциональных опций.
// + Идиоматичный способ создания конфигов в Go (в GRPC и других либах используется)
// + В каждую опцию можно добавить валидацию
//

type options struct {
	port    int
	timeout time.Duration
}

type Option func(options *options) error

// Замыкание (closure) — это анонимная
// функция, которая ссылается на переменные вне своего тела, в данном случае
// на переменную port
func WithPort(port int) Option {
	return func(o *options) error {
		o.port = port
		return nil
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *options) error {
		o.timeout = timeout
		return nil
	}
}

func NewServerWithFunctionalOptions(addr string, opts ...Option) (*http.Server, error) {
	cfg := options{
		port:    defaultPort,
		timeout: defaultTimeout,
	}

	for _, opt := range opts {
		err := opt(&cfg)
		if err != nil {
			return nil, err
		}
	}

	return &http.Server{
		Addr:        fmt.Sprintf("%s:%d", addr, cfg.port),
		ReadTimeout: cfg.timeout,
	}, nil
}

func TestConfigWithFunctionalOptions(t *testing.T) {
	_, _ = NewServerWithFunctionalOptions("localhost", WithPort(8081), WithTimeout(10*time.Second))
}
