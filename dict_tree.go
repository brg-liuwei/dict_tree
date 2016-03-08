package dict_tree

import (
	"fmt"
	"strings"
)

type DictTree struct {
	tag    string
	isLeaf bool
	branch map[rune]*DictTree
}

func NewDictTree() *DictTree {
	return &DictTree{
		isLeaf: false,
		branch: make(map[rune]*DictTree),
	}
}

func (root *DictTree) addRune(r rune) *DictTree {
	if node, ok := root.branch[r]; ok {
		return node
	}
	node := NewDictTree()
	root.branch[r] = node
	return node
}

func (root *DictTree) Add(s string, tag string) {
	runes := []rune(s)
	node := root
	for _, r := range runes {
		node = node.addRune(r)
	}
	node.isLeaf = true
	node.tag = tag
}

func (root *DictTree) retrieval(runes []rune) (string, bool) {
	var t string // for longest match
	var ok bool
	node := root

	if node.isLeaf {
		t = node.tag
	}

	for _, r := range runes {
		if node, ok = node.branch[r]; !ok {
			break
		}
		if node.isLeaf {
			t = node.tag
		}
	}

	if len(t) != 0 {
		return t, true
	}
	return "", false
}

func (root *DictTree) Retrieval(s string) (tag string, ok bool) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if tag, ok = root.retrieval(runes[i:]); ok {
			return
		}
	}
	return
}

func (root *DictTree) display(spaces int, sep string) {
	isFirst := true
	for r, node := range root.branch {
		if !isFirst {
			fmt.Print(strings.Repeat(sep, 3*spaces))
		}
		isFirst = false
		fmt.Print(string(r), sep)
		if node.isLeaf {
			fmt.Println("#" + node.tag + "#")
			if len(node.branch) > 0 {
				fmt.Print(strings.Repeat(sep, 3*(spaces+1)))
			}
		}
		node.display(spaces+1, sep)
	}
}

func (root *DictTree) Display(sep string) {
	for r, node := range root.branch {
		fmt.Print(string(r), sep)
		node.display(1, sep)
		fmt.Println()
	}
}
