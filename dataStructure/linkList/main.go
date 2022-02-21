package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkList struct {
	Header *Node
	Length int
}

func CreateNode(item int) *Node {
	return &Node{
		Data: item,
		Next: nil,
	}
}

func CreateLinkList() *LinkList {
	return &LinkList{
		Header: CreateNode(0),
		Length: 0,
	}
}

// 链表头部增加节点
func (l *LinkList) Add(item int) {
	newNode := CreateNode(item)
	defer func() {
		l.Length++
	}()
	if l.Length == 0 {
		l.Header = newNode
	} else {
		newNode.Next = l.Header
		l.Header = newNode
	}
}

// 链表尾部增加节点
func (l *LinkList) Append(item int) {
	defer func() {
		l.Length++
	}()
	newNode := CreateNode(item)
	if l.Length == 0 {
		l.Header = newNode
	}
	if l.Length > 0 {
		current := l.Header
		for current.Next != nil { //循环找到最后一个节点
			current = current.Next
		}
		current.Next = newNode //把新节点地址给最后一个节点的Next
	}
}

// 往i插入一个节点，后插
func (l *LinkList) Insert(i, item int) {
	defer func() {
		l.Length++
	}()
	newNode := CreateNode(item)
	if l.Length == 0 {
		l.Header = newNode
	}
	if i >= l.Length {
		l.Append(item)
		return
	}
	// 找到第i个节点pre和i+1个after节点
	j := 1
	pre := l.Header
	for j != i {
		pre = pre.Next
		j++
	}
	after := pre.Next //获取到i+1个节点
	// 修改i节点，新节点的指针
	pre.Next = newNode
	newNode.Next = after
}

// 删除第i个节点
func (l *LinkList) Delete(i int) {
	defer func() {
		l.Length--
	}()
	if i > l.Length {
		return
	}
	// 删除第一个节点，把header指向第二个节点即可
	if i == 1 {
		l.Header = l.Header.Next
		return
	}
	// 找到第i-1个节点，找到第i+1个节点，修改i-1的节点的next即可
	pre := l.Header
	for j := 1; j < i-1; j++ {
		pre = pre.Next
	}
	after := pre.Next.Next
	pre.Next = after
}

// 遍历链表，显示出来
func (l *LinkList) Scan() {
	current := l.Header
	i := 1
	for current.Next != nil {
		fmt.Printf("第%d的节点是%d\n", i, current.Data)
		current = current.Next
		i++
	}
	fmt.Printf("第%d的节点是%d\n", i, current.Data)

}

func main() {
	linkList := CreateLinkList()
	linkList.Add(1)
	linkList.Add(2)
	linkList.Add(3)
	linkList.Add(4)
	linkList.Scan()
	//linkList.Insert(2, 4)
	//linkList.Scan()
	linkList.Delete(4)
	linkList.Scan()
}
