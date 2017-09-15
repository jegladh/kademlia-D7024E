package d7024e

import (
	"fmt"
	"testing"
)

//testing freppe branch
//test again
//testing branch
//go test -test.run TestListen


func TestListen(t *testing.T){
	nc := NewContact(NewKademliaID("0000"), "localhost:8000")
	Listen("localhost", 8000)
	network.SendPingMessage(nc)



}

func TestRoutingTable0(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("0000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("0000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("0001"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0010"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0011"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0100"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0101"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0110"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0111"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1001"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1010"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1011"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1100"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1101"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1110"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111"), "localhost:8002"))
	
	bucket := rt.getBucketIndex(NewKademliaID("1111"))
	fmt.Println(bucket)
	
}

func TestRoutingTable1(t *testing.T) { //glöm inte ändra ID lenght = 1


	rt := NewRoutingTable(NewContact(NewKademliaID("00"), "localhost:8000")) //Bygg från här

	rt.AddContact(NewContact(NewKademliaID("00"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("01"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("02"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("03"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("04"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("05"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("06"), "localhost:8002"))//test nodes
	rt.AddContact(NewContact(NewKademliaID("07"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("08"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("09"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("10"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("11"), "localhost:8002"))



	contacts := rt.FindClosestContacts(NewKademliaID("05"), 20)		//Hitta närmaste till denna nod
	bucket := rt.getBucketIndex(NewKademliaID("11"))
	for i := range contacts {
		fmt.Println(contacts[i].String())
		fmt.Println(bucket)
	}


	//originalet
	/*rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}*/
	
}

func TestRoutingTable2(t *testing.T) { //ID lenght = 2
	
	//binary
	rt := NewRoutingTable(NewContact(NewKademliaID("0000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("0000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("0001"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0010"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0011"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0100"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0101"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0110"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("0111"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1001"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1010"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1011"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1100"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1101"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1110"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("0101"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}

