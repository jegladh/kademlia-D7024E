package d7024e

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"sync"
)

//number of "threads"
const alpha = 3

//maximum amount of lookups
const k = 20

//http://blog.notdot.net/tag/kademlia struct and newkademlia found online
type Kademlia struct {
	RT        *RoutingTable
	mapdemlia map[string]string
	//Net *Network
}

//type for closest contacts ish
type Contacts []Contact

func NewKademlia(me *Contact, rt *RoutingTable) *Kademlia {
	kademlia := new(Kademlia)
	kademlia.RT = rt
	kademlia.mapdemlia = make(map[string]string)
	return kademlia
}

var neet neetwork = &MockNetwork{}

//Functions for sorting
func (c Contacts) Len() int {
	return len(c)
}

func (c Contacts) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Contacts) Less(i, j int) bool {
	return c[i].distance.Less(c[j].distance)
}

func (kademlia *Kademlia) LookupContactThreads(target *Contact, closestContacts *Contacts, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < k; i++ {
		if len(*closestContacts) == 0 {
			return
		}
		mutex.Lock()
		c := (*closestContacts)[0]
		*closestContacts = (*closestContacts)[1:]
		a := neet.SendFindContactMessage(target, &c)
		*closestContacts = append(*closestContacts, a...)
		sort.Sort(*closestContacts)
		mutex.Unlock()
	}
}

func (kademlia *Kademlia) LookupContact(target *Contact) *Contact {
	var mutex sync.Mutex
	var closestContacts Contacts = kademlia.RT.FindClosestContacts(target.ID, alpha)
	for i, foundContact := range closestContacts {
		if foundContact.ID == target.ID {
			return &closestContacts[i]
		}
	}
	var wg sync.WaitGroup

	wg.Add(alpha)
	for i := 0; i < alpha; i++ {

		fmt.Println("Spawned new thread")
		go kademlia.LookupContactThreads(target, &closestContacts, &wg, &mutex)
	}
	wg.Wait()
	for i := 0; i < len(closestContacts); i++ {
		fmt.Println(closestContacts[i])
	}
	for i := 0; i < len(closestContacts); i++ {
		if closestContacts[i].ID == target.ID {
			//fmt.Println(closestContacts[i])
			return &(closestContacts[i])
		}
	}

	return nil
}

func (kademlia *Kademlia) LookupDataThreads(hash string, closestContacts *Contacts, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < k; i++ {
		if len(*closestContacts) == 0 {
			return
		}
		mutex.Lock()
		c := (*closestContacts)[0]
		*closestContacts = (*closestContacts)[1:]
		neet.SendFindDataMessage(hash, &c)
	}
	mutex.Unlock()
}

func (kademlia *Kademlia) LookupData(hash string, contact *Contact) string {
	fmt.Println("Lookup data starting...")
	var closestContacts Contacts = kademlia.RT.FindClosestContacts(contact.ID, alpha)
	var mutex sync.Mutex
	i, ok := kademlia.mapdemlia[hash]
	if ok {
		return i
	}

	var wg sync.WaitGroup

	wg.Add(alpha)
	for i := 0; i < alpha; i++ {
		fmt.Println("Spawned new thread")
		go kademlia.LookupDataThreads(hash, &closestContacts, &wg, &mutex)
	}
	wg.Wait()

	for i := 0; i < len(closestContacts); i++ {

		//if closestContacts[i]. == target.ID {
		//p, ok := kademlia.mapdemlia[hash]
		//fmt.Println(closestContacts[i])
		//return &(closestContacts[i])
	}
	//}
	return ""
}

//fmt.Println(target.ID.String())

// for i := 0; i > (someHash); i++{
//	if someHash == target.hash{ return someHash }
//}väldigt lik LookupContactsl

func (kademlia *Kademlia) Store(data []byte) {
	hashedval := KademliaID(sha1.Sum(data))
	var closestContacts Contacts = kademlia.RT.FindClosestContacts(&hashedval, alpha)
	for i := 0; i < len(closestContacts); i++ {
		go neet.SendStoreMessage(data)
	}

	//hashedvalue := data //hash here?
}

func (kademlia *Kademlia) Cat(hash string) {
	fmt.Println(hash)
	//Print the content of a file with address as argument
}

func (kademlia *Kademlia) Pin(hash string) {
	//Pin data so it can't be deleted

}

func (kademlia *Kademlia) Unpin(hash string) {
	//unpin data so
}
