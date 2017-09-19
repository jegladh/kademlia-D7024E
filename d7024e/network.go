package d7024e

import (
	"net"
	"fmt"
	"strconv"
	"os"
	"github.com/golang/protobuf/proto"
)

type Network struct {
}

func ErrorHandler(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func Listen(ip string, port int) {
	portstr := strconv.Itoa(port)
	serveraddr, err := net.ResolveUDPAddr("udp", ip + ":" + portstr)
	ErrorHandler(err)

	listener, err := net.ListenUDP("udp", serveraddr)
	ErrorHandler(err)

	defer listener.Close()
	fmt.Println("Listening on " + ip + ":" + portstr)

	buf := make([]byte, 1024)
	for
{		//n, conn, err := listener.ReadFromUdp(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", conn)
		ErrorHandler(err)
	}

}

func (network *Network) SendPingMessage(contact *Contact) {
	/*conn, err := net.DialUDP("udp", contact.Address, "localhost")
	ErrorHandler(err)
	n, err := conn.Write([]byte("Ping")) 
	ErrorHandler(err)
	fmt.Println("Sent " + Itoa(n) + " bytes")*/

	data, err := proto.Marshal(test)
	ErrorHandler(err)
	/*if err != nil {
		log.Fatal("marshaling error: ", err)
	}*/

	conn, err := net.DialUDP("udp", contact.Address, "localhost")
	ErrorHandler(err)
	defer conn.Close()

	_, err := conn.Write([]byte("Ping"))
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

