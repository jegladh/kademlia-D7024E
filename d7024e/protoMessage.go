package d7024e

import (
	//"fmt"
	//"net"
	"github.com/golang/protobuf/proto"
	"github.com/kademlia-D7024E/protobuf"

	)

type protoMessage struct {

}

func sendProtoMessage (lable string) *protobuf.KademliaMessage {
	message := &protobuf.KademliaMessage {
		Label: proto.String("Ping"),
	}
	return message
}
/*
func protoMessageHandler(data []byte, me Contact) {
	message := &protobuf.KademliaMessage{}
	err := proto.Unmarshal(data, message)
	ErrorHandler(err)
	fmt.Println("\n", message)

}*/