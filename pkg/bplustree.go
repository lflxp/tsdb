package pkg

import (
	"fmt"
)

var (
	M int = 5
)

type LeafNode struct {
	Key  int   //关键字
	Datas []int  //数据
	Next *LeafNode //链表
}

type BTNode struct {
	KeyNum  int 	//实际关键字个数 keyNum <m
	Key  []int   //关键字
	Children  []*BTNode // 非叶子节点的子节点数
	IsLeaf   bool
	Leaf  *LeafNode
}

type BTree struct {

}