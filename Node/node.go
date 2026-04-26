package node

import (
	"errors"
	"log"
	"sync"
)

type Node struct {
	id          string
	isConnected bool
	data        sync.Map
}

var (
	ErrorNodeNotConnected = errors.New("Node is not connected")
	ErrorKeyNotFound      = errors.New("Key not found in node")
)

func NewNode(id string) *Node {
	node := &Node{
		id:          id,
		isConnected: true,
		data:        sync.Map{},
	}

	return node
}

func (n *Node) GetID() string {
	return n.id
}

func (node *Node) ping() error {
	if !node.isConnected {
		return ErrorNodeNotConnected
	}
	return nil
}

func (node *Node) Get(key string) (string, error) {
	if err := node.ping(); err != nil {
		return "", err
	}

	if value, ok := node.data.Load(key); ok {
		log.Printf("[Node: %s] ✅ GET key: %v (found)\n", node.id, key)
		return value.(string), nil
	}

	log.Printf("[Node: %s] ❌ GET key: %v (not found)\n", node.id, key)
	return "", ErrorKeyNotFound
}

func (node *Node) Set(key string, value string) error {
	if err := node.ping(); err != nil {
		return err
	}
	node.data.Store(key, value)
	log.Printf("[Node: %s] ✅ SET key: %v, value: %v\n", node.id, key, value)
	return nil
}

func (node *Node) Delete(key string) error {
	if err := node.ping(); err != nil {
		return err
	}

	_, err := node.Get(key)
	if err != nil {
		return err
	}
	node.data.Delete(key)
	log.Printf("[Node: %s] ✅ DELETE key: %v\n", node.id, key)
	return nil
}
