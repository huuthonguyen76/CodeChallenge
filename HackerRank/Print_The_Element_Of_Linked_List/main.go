package main

import (
	"fmt"
)

type linkList struct {
	pNext *linkList
	value int
}

func (this *linkList) AddNode(value int) {
	var _this = this
	for {
		if _this.pNext == nil {
			_this.pNext = &linkList{value: value}
			break
		} else {
			_this = _this.pNext
		}
	}
}

func (this linkList) FindNode(value int) *linkList {
	_this := &this
	for _this != nil {
		if value != _this.value {
			_this = _this.pNext
		} else {
			return _this
		}
	}
	return nil
}

func (this linkList) PrintNodes() {
	_this := &this
	for _this != nil {
		fmt.Print(_this.value, " ")
		_this = _this.pNext
	}
}

func (this *linkList) LastNode() *linkList {
	var _this = this

	for _this.pNext != nil {
		_this = _this.pNext
	}
	return _this
}

// 1 2 3 4 5
// multi define variable
func (this *linkList) Reverse() {
	var _this = this
	var headNode = new(linkList)
	var tailNode = new(linkList)
	tailNode = nil

	for _this.pNext != nil {
		headNode = _this.pNext
		_this.pNext = tailNode
		tailNode = _this
		_this = headNode
	}
	_this.pNext = tailNode
}

func main() {
	var list = &linkList{value: 123}
	list.AddNode(45)
	list.AddNode(47)
	list.AddNode(50)
	list.Reverse(list)
	list.PrintNodes()

	// var node = list.FindNode(45)

	// if node != nil {
	// 	fmt.Println(node.value)
	// }
}
