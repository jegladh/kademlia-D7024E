package d7024e

import (
	"fmt"
	"testing"
)

/*func TestListen(t *testing.T) {
	Listen("localhost", 8000)
}*/

func TestRoutingTable(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}

func TestRoutingTable2(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("00000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("11111000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("11110000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("11110000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("11100000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("11100000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("11010000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("11000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("10100000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("10010000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("10000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("01110000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("01101000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("01100000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("01011000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("01010000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("01000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("00110000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("00101000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("00100000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("00000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("10110000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}
