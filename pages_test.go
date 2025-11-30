package pages

import (
	"fmt"
	"testing"
)

func TestNewPaginator(t *testing.T) {
	page := New(testIntSlice,
		WithPerPage(10),
	)
	page.Loop = true
	fmt.Printf("page %v: %#v\n", page.Page, page.Prev())
	page.PrevPage()
	fmt.Printf("page %v: %#v\n", page.Page, page.Current())
	page.NextPage()
	fmt.Printf("page %v: %#v\n", page.Page, page.Current())

}

var testIntSlice = []int{1, 2, 3, 4, 5, 6}
