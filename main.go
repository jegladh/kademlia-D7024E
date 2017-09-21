package main

import (
	kademlia "github.com/kademlia-D7024E/d7024e"
	"fmt"
	"log"
	//"github.com/golang/protobuf/proto"
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

	/*test := &Test{
		Label: proto.String("Ping"),
		Type:  proto.Int32(17),
		Optionalgroup: &Test_OptionalGroup{
			RequiredField: proto.String("Tack o hej ping"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}

	log.Printf("Unmarshalled to: %+v", newTest)*/

}