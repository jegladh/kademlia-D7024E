package d7024e

import (
	"net"
	"fmt"
	//"strconv"
	"os"
	"github.com/golang/protobuf/proto"
	//"log"
)

type Network struct {
	me Contact
}

func NewNetwork(me Contact) *Network {
	net := &Network{}
	net.me = me
	return net
}

func ErrorHandler(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func Listen(me Contact) {
	// convert port to string
	//portstr := strconv.Itoa(port)
	serveraddr, err := net.ResolveUDPAddr("udp", me.Address)
	ErrorHandler(err)

	// Listening at port
	listener, err := net.ListenUDP("udp", serveraddr)
	ErrorHandler(err)

	defer listener.Close()
	//fmt.Println("Listening on " + ip + ":" + portstr)

	buf := make([]byte, 1024)
	for{		
		_, _, err := listener.ReadFromUDP(buf)
		go protoMessageHandler(buf, me)
		//fmt.Println("Received ", string(buf[0:n]), " from ", conn)
		ErrorHandler(err)
	}

}

func (network *Network) SendPingMessage(contact *Contact) {
	// protomessage
	message := sendProtoMessage("ping")
	//Marshall proto message
	data, err := proto.Marshal(message)
	ErrorHandler(err)

	conn, err := net.Dial("udp", contact.Address)
	ErrorHandler(err)
	defer conn.Close()

	_, err = conn.Write(data)
	ErrorHandler(err) 
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO

	//kademlia.LookUpContact()
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO

	//kademlia.LookUpData()
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO

	//kademlia.Store()
}

