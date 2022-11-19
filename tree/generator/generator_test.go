package generator

import (
	"reflect"
	"testing"

	"epi/tree"
)

func TestCanGenerateTree(t *testing.T) {
	testTree := GenerateTree([]int{1, 2, 3, 4, 5, 6})

	testTreeHandcrafted := &tree.Node{
		L: &tree.Node{
			L:   &tree.Node{Val: 1},
			R:   &tree.Node{Val: 3},
			Val: 2,
		},
		R: &tree.Node{
			L:   &tree.Node{Val: 5},
			Val: 6,
		},
		Val: 4,
	}

	if !reflect.DeepEqual(testTree, testTreeHandcrafted) {
		t.Errorf("did not generate expected tree")
	}
}

func TestCanGenerateEmptyTree(t *testing.T) {
	testTree := GenerateTree([]int{})
	if testTree != nil {
		t.Errorf("expected a nil tree")
	}
}
