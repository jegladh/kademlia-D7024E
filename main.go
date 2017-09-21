package main

import (
	"fmt"
	kademlia "kademlia-D7024e/d7024e"
	"strconv"
	/*"net"
	"github.com/tatsushid/go-fastping"
	"os"
	"time"
	*/)

var port int

func main() {

	port := 6969
	portstr := strconv.Itoa(port)
	//kademlia.Listen("130.240.109.93", port)
	//kademlia.Ping()
	randomID := kademlia.NewRandomKademliaID()
	newRT := kademlia.NewRoutingTable(kademlia.NewContact(randomID, "localhost:"+portstr))

	//contact := kademlia.NewContact(randomID, "localhost:"+portstr)

	for n := 0; n < 10; n++ {
		portstr := strconv.Itoa(port)
		randomID := kademlia.NewRandomKademliaID()
		newC := kademlia.NewContact(randomID, "localhost:"+portstr)
		fmt.Println(newC)
		newRT.AddContact(newC)

		port++

	}
	fmt.Println(newRT.GetBucketIndex(kademlia.NewKademliaID("0f")))
	fmt.Println(newRT.Buckets[2])

}
