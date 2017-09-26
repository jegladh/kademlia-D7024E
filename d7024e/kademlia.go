package d7024e

import (
	"fmt"
	"sync"
)

const alpha = 3

//http://blog.notdot.net/tag/kademlia struct and newkademlia found online
type Kademlia struct {
	RT *RoutingTable
}

func NewKademlia(me *Contact) *Kademlia {
	kademlia := new(Kademlia)
	kademlia.RT = NewRoutingTable(*me)
	return kademlia
}
func (kademlia *Kademlia) LookupContactThreads(target *Contact, tempContacts []*Kademlia) {
	fmt.Println("Hello world")
}

func (kademlia *Kademlia) LookupContact(target *Contact) *Contact {
	closestContacts := kademlia.RT.FindClosestContacts(target.ID, alpha)
	for i, foundContact := range closestContacts {
		if foundContact.ID == target.ID {
			return &closestContacts[i]
		}
	}
	var wg sync.WaitGroup
	tempContacts := make([]*Kademlia, alpha)
	wg.Add(alpha)
	for i := 0; i < alpha; i++ {
		go kademlia.LookupContactThreads(target, tempContacts)
	}
	wg.Wait()

	return &closestContacts[0]

}

func (kademlia *Kademlia) LookupData(hash string) {
	// for i := 0; i > (someHash); i++{
	//	if someHash == target.hash{ return someHash }
	//}

}

func (kademlia *Kademlia) Store(data []byte) {

}
