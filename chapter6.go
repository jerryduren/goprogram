package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	x, y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//这个Path实际上是个折线，用每个顶点表示，他的所谓distance就是每段这线长度之和
func (p Path) Distance() float64 {
	sum := 0.0
	for i, _ := range p {
		if i > 0 {
			sum = sum + p[i].Distance(p[i-1])
		}
	}
	return sum
}

//Linked List
type IntLinkedList struct {
	value    int
	nextList *IntLinkedList
}

func (this *IntLinkedList) Sum() int {
	if this == nil {
		return 0
	}

	return this.value + this.nextList.Sum() //这里使用了递归算法，一个链表的和等于本节点的值+后面链表的和
}

//构造N个节点的整形单项链表
func NewIntLinkedList(n int) *IntLinkedList {
	if n <= 0 {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	head := &IntLinkedList{r.Intn(100), nil}
	tail := head
	for i := 1; i < n; i++ {
		tail.nextList = &IntLinkedList{r.Intn(100), nil}
		tail = tail.nextList
	}
	return head
}

func (this *IntLinkedList) String() string {
	result := "["
	for tail := this; tail != nil; tail = tail.nextList {
		result = result + fmt.Sprintf("%d", tail.value)
		if tail.nextList != nil {
			result = result + ", "
		}
	}
	result = result + "]"

	return result
}

func main() {
	ch := make(chan bool)

	triangle := Path{{0, 0}, {3, 0}, {0, 4}, {0, 0}} //求三角形的周长
	fmt.Println(triangle.Distance())

	fmt.Println(Path{{0, 0}, {10, 10}, {0, 10}, {0, 0}}.Distance())

	head := NewIntLinkedList(10)
	fmt.Println(head)
	fmt.Println(head.Sum())

	go func() {
		time.AfterFunc(10*time.Second, func() {
			fmt.Println("hello world!")
			fmt.Println(time.Now())
		})
		ch <- true
	}()

	<-ch //用channel来控制主协程等待子协程结束，但是这个例子不合适，因为前面的time.AfterFunc里面又启了协程
	time.Sleep(12 * time.Second)
	fmt.Println(time.Now())
}
