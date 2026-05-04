package main

// import (
// 	"fmt"

// 	"github.com/google/uuid"
// 	"github.com/sajeelwaien/consistent-hashing/bloomfilter"
// )

// func main() {
// 	values := make([]string, 0)
// 	for i := 0; i < 100; i++ {
// 		values = append(values, uuid.New().String())
// 	}

// 	hashFunc := bloomfilter.InitHashFunc(3)
// 	bf := bloomfilter.NewBloomFilter(50)

// 	for i := 0; i < len(values); i++ {
// 		bf.Add(values[i], &hashFunc)
// 	}

// 	bf.Print("test")
// 	fmt.Println(bf.Contains(uuid.NewString(), &hashFunc))
// 	fmt.Println(values[5], bf.Contains(values[5], &hashFunc))
// }

import (
	"fmt"

	"github.com/sajeelwaien/consistent-hashing/hashring"
	"github.com/sajeelwaien/consistent-hashing/node"
)

func main() {
	cacheNode1 := node.NewNode("Node1")
	cacheNode2 := node.NewNode("Node2")
	cacheNode3 := node.NewNode("Node3")

	rf := hashring.WithReplicationFactor(2)
	vnc := hashring.WithVirtualNodeCount(5)
	le := hashring.WithLoggingEnabled(true)

	hashRing := hashring.NewHashRing(rf, vnc, le)

	hashRing.AddNode(cacheNode1)
	hashRing.AddNode(cacheNode2)
	hashRing.AddNode(cacheNode3)

	ringNode, err := hashRing.GetPrimaryNode("key1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Primary node for key1: %v\n", ringNode.GetID())
	// for _, node := range ringNodes {
	// 	fmt.Printf("Node for key1: %v\n", node.GetID())
	// }

	fmt.Println(cacheNode1.GetID())

}
