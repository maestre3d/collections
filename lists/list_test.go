package lists_test

import (
	"testing"

	"github.com/maestre3d/collections/lists"
)

func TestNewSliceList(t *testing.T) {
	tests := []struct {
		name string
		in   lists.List[string]
	}{
		{
			name: "slice",
			in:   lists.NewListSlice[string](0),
		},
		{
			name: "linked list",
			in:   lists.NewLinkedList[string](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.in
			ls.Add("foo")
			ls.Add("bar")
			ls.Add("baz")
			ls.Add("foobar")

			t.Log(ls.LastIndexOf("foo"))
			t.Log(ls.Get(1))
			t.Log(ls.IsEmpty())
			t.Log(ls.ToSlice())
			t.Log("remove last")
			ls.RemoveWithIndex(ls.Len() - 1)
			t.Log(ls.ToSlice())
			t.Log("filter remove")
			ls.RemoveIf(func(v string) bool {
				return v == "bar"
			})
			t.Log(ls.ToSlice())

			t.Log("clear")
			ls.Clear()
			t.Log(ls.ToSlice())
			t.Log(ls.Len())
			t.Log(ls.IsEmpty())
			t.Log(ls.Get(1))
		})
	}

}
