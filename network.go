package d7024e

import (
	"net"
	"fmt"
	//"github.com/golang/protobuf/proto"
	"strconv"
	"os"
)

type Network struct {

}


func ErrorHandler(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
/*server inspired by https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go*/

func Listen(ip string, port int) {
	/*convert port to string*/
	portstr := strconv.Itoa(port)
	/**/
	serveraddr, err := net.ResolveUDPAddr("udp", ip + ":" + portstr)
	ErrorHandler(err)
	/*listen at port*/
	listener, err := net.ListenUDP("udp", serveraddr)
	ErrorHandler(err)
	defer listener.Close()
	fmt.Println("Listening on "+ ip + ":" + portstr)
	buf := make([]byte, 1024)
	for{
		n,conn, err := listener.ReadFromUDP(buf)
		fmt.Println("Received ",string(buf[0:n]), " from ",conn)
		ErrorHandler(err)
	}
}



/*type Contact struct {
	ID       *KademliaID
	Address  string
	distance *KademliaID
}*/
func (network *Network) SendPingMessage(contact *Contact) {
	// TODO
	conn,err := net.DialUDP("udp", contact.Address, "localhost")
	ErrorHandler(err)
	n,err := conn.Write([]byte("Ping"))
	ErrorHandler(err)
	fmt.Println("Sent " + Itoa(n) + " bytes")
	//Ping should get a pong reply
	//should update buckets?
	
	//bucket.AddContact
}


func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
	//if success
	//kademlia.LookupContact()
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
	//if success
	//kademlia.LookupData()
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
	//if success
	//kademlia.Store()
}
