package structs

import (
	"bytes"
	"fmt"
)

// TData -> Node data struct
type TData struct {
	Char string
	Freq int
}

// Tree -> Generic tree node element
type Tree struct {
	Left,Right *Tree
	Data TData
}

// CreateNode return a tree struct pointer
func CreateNode(data TData) *Tree {
	return &Tree{Data: data}
}

// AddNode adds  new node to the tree
func (t *Tree) AddNode(data TData) {
	if t == nil {
		return
	} else if data.Freq < t.Data.Freq {
		if t.Left == nil {
			t.Left = CreateNode(data)
		} else {
			t.Left.AddNode(data)
		}
	} else {
		if t.Right == nil {
			t.Right = CreateNode(data)
		} else {
			t.Right.AddNode(data)
		}
	}
}

// SearchNode searches for a certain node in the tree
func (t *Tree) SearchNode(data TData) string {
	var code bytes.Buffer
	searchNode(t, data, &code)
	return fmt.Sprintf("Char: %s | Code: %s \n", data.Char, code.String())
}

func searchNode(root *Tree, data TData, s *bytes.Buffer) {
	if root != nil {
		if root.Data.Char == data.Char {
			s.WriteString("0")
		} else if data.Freq > root.Data.Freq {
			s.WriteString("1")
			searchNode(root.Right, data, s)
		} else {
			s.WriteString("0")
			searchNode(root.Left, data, s)
		}
	}
}


// DeleteNode deletes a certain node from the tree
func (t *Tree) DeleteNode(data TData) {
	if t == nil {
		return
	}  
	res :=	deleteNode(t, data)

	if res == nil {
		fmt.Println("Node not deleted")
		return
	}

	fmt.Println("Node deleted")

}

func deleteNode(root *Tree, data TData) *Tree{
	if root == nil {
		return nil
  }

	if data.Freq < root.Data.Freq {
		root.Left = deleteNode(root.Left, data)
		return root
	}

	if data.Freq > root.Data.Freq {
		root.Right = deleteNode(root.Right, data)
		return root
	}

	if root.Left == nil && root.Right == nil {
		root = nil
		return nil
	}

	if root.Left == nil {
		root = root.Right
		return root
	}

	if root.Right == nil {
		root = root.Left
		return root
	}

	LeftmostRightside := root.Right

	for {
		//find smallest value on the Right side
		if LeftmostRightside != nil && LeftmostRightside.Left != nil {
			LeftmostRightside = LeftmostRightside.Left
		} else {
			break
		}
	}

	root.Data.Freq = LeftmostRightside.Data.Freq
	root.Right = deleteNode(root.Right, root.Data)
	return root
}

// NewTree creates a new tree from root
func NewTree(root, left, right *Tree, data TData) {
	root = CreateNode(data)
	root.Left = left
	root.Right = right
}

// AddTree adds tree to a tree
func (t *Tree) AddTree(tree Tree) {
	if tree.Data.Freq < t.Data.Freq {
		if t.Left == nil {
			t.Left = &tree
		} else {
			t.Left.AddTree(tree)
		}
	} else {
		if t.Right == nil {
			t.Right = &tree
		} else {
			t.Right.AddTree(tree)
		}
	}
}

// Depth shows current tree depth
func (t *Tree) Depth() int {
	
	if t == nil {
		return 0
	}

	var lDepth, rDepth int
	lDepth = t.Left.Depth()
	rDepth = t.Right.Depth()

	if lDepth > rDepth {
		return lDepth
	} 
	return rDepth
}

// PreOrder prints root, left sub-tree, right sub-tree
func (t *Tree) PreOrder() {
	if t == nil {
		return
	}
	fmt.Println(t.Data.Freq)
	t.Left.PreOrder()
	t.Right.PreOrder()
}

// InOrder prints left, root, right
func (t *Tree) InOrder() {
	
	if t == nil {
		return
	}

	t.Left.InOrder()
	fmt.Println(t.Data.Freq)
	t.Right.InOrder()
}

// PostOrder prints left sub-tree, right sub-tree, roota
func (t* Tree) PostOrder() {
	if t == nil {
		return
	} 
	t.Left.PostOrder()
	t.Right.PostOrder()
	fmt.Println(t.Data.Freq)
}

// MinNode returns smallest node in the tree
func (t *Tree) MinNode() Tree {
	minNode := t

	for ; minNode.Left != nil; minNode = minNode.Left {
	}

	return *minNode
}

// MaxNode returns highest node in the tree
func (t *Tree) MaxNode() Tree{
	maxNode := t
	
	for ; maxNode.Right != nil; maxNode = maxNode.Right {
	}

	return *maxNode
}
