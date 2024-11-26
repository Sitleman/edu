// Когда использовать:
// * Структуры данных. (если реализуем двоичное дерево, связанный список или кучу)
// * Функции, работающие со срезами, картами и каналами любого типа
// * Факторизация поведения вместо типов (реализовать sort.Interface сразу для нескольких типов)
// Когда не стоит:
// * При вызове метода с аргументом типа
// func foo[T io.Writer](w T) {
// b := getBytes()
// _, _ = w.Write(b)
// }
// * Когда это делает код более сложным.
package __generic

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGeneric(t *testing.T) {
	m1 := map[string]int{"Nikita": 1, "Maria": 2, "Denis": 3}
	m2 := map[int]int{4324: 1, 4234: 2, 2921: 3}
	fmt.Println(getKeys(m1))
	fmt.Println(getKeysWithGeneric[string, int](m1))
	fmt.Println(getKeysWithGeneric(m2))
	fmt.Println(getKeysWithGenericWithCustomConstraint(m1))
	fmt.Println(getKeysWithGenericWithCustomConstraint(m2))
	fmt.Println(getKeysWithGenericWithCustomConstraint2(m1))
	fmt.Println(getKeysWithGenericWithCustomConstraint2(m2))
}

func getKeys(data map[string]int) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	return keys
}

// Как возвращать список ключей для нескольких типов словарей?
// До появления дженериков у разработчиков было несколько вариантов:
// * использование генерации кода
// * отражение
// * дублирование кода.

func getKeysWithGeneric[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Можно создать свой ограничить типы данных для дженерика
//Использование int ограничивает тип только этим типом,
//тогда как ~int ограничивает все типы, базовым типом которых является int

type customConstraint interface {
	~int | ~string
}

func getKeysWithGenericWithCustomConstraint[K customConstraint, V any](m map[K]V) []K {
	return getKeysWithGeneric(m)
}

func getKeysWithGenericWithCustomConstraint2[K ~int | ~string, V any](m map[K]V) []K {
	return getKeysWithGeneric(m)
}

// Для определения интерфейса ограничения можно использовать не только типы данных, но и методы

type customConstraintWithMethod interface {
	~int64
	String() string
}
type customInt int64

func (i customInt) String() string {
	return strconv.Itoa(int(i))
}

// Дженерики могут использоваться так же в структурах

type Node[T any] struct {
	Val  T
	next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}
