package d7024e

type Network struct {
}

func Listen(ip string, port int) {
	/*listener, err := net.Listen("udp", ip + ":" + pstr)
	if err != nil{
		fmt.Println("Error listen", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
*/
}

func (network *Network) SendPingMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}

