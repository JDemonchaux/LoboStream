package main

import (
	"time"
	"log"
	"fmt"
)

type Node struct {
	Value string
}

func NewQueue() *Queue {
	return &Queue{
		nodes: make([]*Node, 1),
		size:  1,
	}
}
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes) + q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes) - q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func (q *Queue) Start() {
	for {
		if q.count != 0 {
			ss := q.Pop().String()
			log.Println("Queue ",ss)
			convert(ss)
		}
		time.Sleep(time.Second)
	}
}

func convert(nameFile string)  {
	log.Println("Convert", nameFile)
//	sourcePath := "tmp" + string(os .PathSeparator) + nameFile + ".mp4"
//	nameDestination := "out" + string(os .PathSeparator) + nameFile + ".mp4"
//	out, err := exec.Command("ffmpeg.exe", "-i", sourcePath, "-codec:a", "aac", "-strict", "-2", nameDestination).CombinedOutput()
//	if err != nil {
//		log.Println("some error found",err)
//	}
//
//	log.Println("out",string(out))
}