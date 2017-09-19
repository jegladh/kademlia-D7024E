package main

import (
	"log"
	"github.com/golang/protobuf/proto"
	kademlia "kademlia-D7024E/d7024e"
)

func main() {

	func sendProtoMessage (lable string) *protobuf.KademliaMessage {
		kademlia.sendPingMessages()
		message := protobuf.KademliaMessage {
			Label: proto.String("Ping"),
		}
		return message
	}

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