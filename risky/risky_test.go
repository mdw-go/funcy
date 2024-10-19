package risky

import (
	"testing"

	"github.com/mdw-go/funcy"
	"github.com/mdw-go/funcy/internal/should"
)

func TestField(t *testing.T) {
	type Thing struct{ ID int }
	ID := Field[Thing, int]("ID")
	sorted := funcy.Map(ID, funcy.SortAscending(ID, []Thing{{ID: 42}, {ID: 41}, {ID: 43}}))
	should.So(t, sorted, should.Equal, []int{41, 42, 43})
}
