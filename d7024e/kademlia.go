package d7024e

import (
	"fmt"
	"sync"
)

//number of "threads"
const alpha = 3

//maximum amount of lookups
const k = 20

//http://blog.notdot.net/tag/kademlia struct and newkademlia found online
type Kademlia struct {
	RT  *RoutingTable
	Net *Network
}

func NewKademlia(me *Contact) *Kademlia {
	kademlia := new(Kademlia)
	kademlia.RT = NewRoutingTable(*me)
	return kademlia
}

var neet neetwork = &Network{}

func (kademlia *Kademlia) LookupContactThreads(target *Contact, closestContacts *[]Contact, wg *sync.WaitGroup) {
	fmt.Println(closestContacts)
	defer wg.Done()
	for i := 0; i < 7; i++ {
		c := (*closestContacts)[0]
		*closestContacts = (*closestContacts)[1:]
		neet.SendFindContactMessage(target, &c)

	}
}

func (kademlia *Kademlia) LookupContact(target *Contact) *Contact {
	closestContacts := kademlia.RT.FindClosestContacts(target.ID, alpha)
	//fmt.Println(closestContacts)
	for i, foundContact := range closestContacts {
		fmt.Println(closestContacts)
		if foundContact.ID == target.ID {
			return &closestContacts[i]
		}
	}
	var wg sync.WaitGroup

	wg.Add(1)
	for i := 0; i < 1; i++ {
		go kademlia.LookupContactThreads(target, &closestContacts, &wg)
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
