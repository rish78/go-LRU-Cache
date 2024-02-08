package main

import "fmt"

const SIZE = 5

type Node struct {
	Left  *Node
	Val   string
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{
		Head:   head,
		Tail:   tail,
		Length: 0,
	}
}

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}

}

func (c *Cache) Check(s string) {
	node := &Node{}

	if val, ok := c.Hash[s]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: s}
	}
	c.Add(node)
	c.Hash[s] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Println("REMOVING", n.Val)

	n.Left.Right = n.Right
	n.Right.Left = n.Left
	c.Queue.Length--
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(node *Node) {
	fmt.Println("ADDING", node.Val)

	tmp := c.Queue.Head.Right
	node.Right = tmp
	node.Left = c.Queue.Head
	c.Queue.Head.Right = node
	tmp.Left = node

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove((c.Queue.Tail.Left))
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)

	for i := 0; i < q.Length; i++ {
		if node != nil {
			fmt.Printf("{%s}", node.Val)
			if i < q.Length-1 {
				fmt.Printf("<-->")
			}
			node = node.Right
		} else {
			fmt.Println("Error: Node is nil")
			break
		}
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("STARTING CACHE")
	cache := NewCache()

	for _, word := range []string{"paris", "london", "amsterdam", "delhi", "london"} {
		cache.Check(word)
		cache.Display()
	}
}
