package util

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	var list = []string{"a", "q", "b", "a", "c", "f", "c", "d", "e", "s"}
	list = RemoveAttentionUsersDup(list)
	t.Logf("%s", list)
}

func TestString(t *testing.T) {
	var list = []string{"a", "q", "b", "a", "c", "f", "c", "d", "e", "s"}
	fmt.Println(SliceString(list))
}
