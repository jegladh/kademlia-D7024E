package main

import (
	kademlia "github.com/kademlia-D7024E/d7024e"
	"fmt"

)

func main() {

	//ping with proto
	me := kademlia.NewContact(kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "127.0.0.1:8000")
	contact := kademlia.NewContact(kademlia.NewKademliaID("1111111100000000000000000000000000000000"), "127.0.0.1:8001")
	net := kademlia.NewNetwork(me)
	go net.SendPingMessage(&contact)
	fmt.Println(me)

	//Listening to this port
	kademlia.Listen(me)
}