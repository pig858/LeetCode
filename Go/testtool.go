package main

import "encoding/json"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func StringToSliceAny(s string) []interface{} {
	var slice []interface{}
	json.Unmarshal([]byte(s), &slice)
	return slice
}

// Generics syntax
func StringToSliceType[T interface{}](s string) []T {
	var slice []T
	json.Unmarshal([]byte(s), &slice)
	return slice
}

func StringTo2DSliceType[T interface{}](s string) [][]T {
	var slice [][]T
	json.Unmarshal([]byte(s), &slice)
	return slice
}

func StringToTreeNode(s string) *TreeNode {
	slice := StringToSliceAny(s)
	if len(slice) < 1 {
		return nil
	}
	tree := make([]*TreeNode, 0)

	for _, v := range slice {
		if v == nil {
			tree = append(tree, nil)
			continue
		}

		node := &TreeNode{
			Val:   int(v.(float64)),
			Left:  nil,
			Right: nil,
		}

		tree = append(tree, node)
	}

	for index := range tree {
		l := len(tree)
		if index*2+1 < l {
			tree[index].Left = tree[index*2+1]
		}
		if index*2+2 < l {
			tree[index].Right = tree[index*2+2]
		}
	}

	return tree[0]
}
