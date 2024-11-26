// Должны следовать вот этим соглашениям о наименованиях:
// Метод геттера должен называться Balance (а не GetBalance).
// Метод сеттера должен называться SetBalance.
package getter

import (
	"testing"
)

type Balance struct {
	value int64
}

func (b *Balance) Balance() int64 {
	return b.value
}

func (b *Balance) SetBalance(value int64) {
	b.value = value
}

func TestGetter(t *testing.T) {
	bal := Balance{5}
	bal.Balance()
	bal.SetBalance(13)
}
