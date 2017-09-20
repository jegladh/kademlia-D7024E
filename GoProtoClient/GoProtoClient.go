package main

import (
	"d7024e/kademlia-D7024E/ProtobufTest"
	"protobuf/proto"
	"net"
	"os"
	"encoding/csv"
	"flag"
	"io"
	"strconv"
	"fmt"
)

type Headers []string

func main() {
	filename := flag.String("f", "CSVV.csv", "Enter the filename to read from")
	dest := flag.String("d", "127.0.0.1:2110", "Enter the destnation socket address")
	flag.Parse()
	data, err := retrieveDataFromFile(filename)
	checkError(err)
	sendDataToDest(data, dest)
}
func retrieveDataFromFile(fname *string) ([]byte, error) {
	file, err := os.Open(*fname)
	checkError(err)
	defer file.Close()
	csvreader := csv.NewReader(file)
	var hdrs Headers
	hdrs, err = csvreader.Read()
	checkError(err)
	NODEIDINDEX := hdrs.getHeaderIndex("nodeid")
	NODEADDRESSINDEX := hdrs.getHeaderIndex("nodeaddress")
	NODEDISANCEINDEX := hdrs.getHeaderIndex("nodedistance")
	NODEBUCKETINDEX := hdrs.getHeaderIndex("nodebucket")
	ProtoMessage := new(ProtobufTest.TestMessage)
	//loop through the records
	for {
		record, err := csvreader.Read()
		if (err != io.EOF) {
			checkError(err)
		} else {

			break
		}
		//Populate items
		testMessageItem := new(ProtobufTest.TestMessage_MsgItem)
		//ID
		nodeid, err := strconv.Atoi(record[NODEIDINDEX])
		checkError(err)
		testMessageItem.NodeId = proto.Int32(int32(nodeid))
		//Address
		testMessageItem.NodeAddress = &record[NODEADDRESSINDEX]
		//Distance
		nodedistance, err := strconv.Atoi(record[NODEDISANCEINDEX])
		checkError(err)
		testMessageItem.NodeDistance = proto.Int32(int32(nodedistance))
		//Bucket
		nodebucket, err := strconv.Atoi(record[NODEBUCKETINDEX])
		checkError(err)
		testMessageItem.NodeBucket = proto.Int32(int32(nodebucket))
		ProtoMessage.Messageitems = append(ProtoMessage.Messageitems, testMessageItem)
		fmt.Println(record)
	}
	//fmt.Println(ProtoMessage.Messageitems)
	return proto.Marshal(ProtoMessage)
}
func sendDataToDest(data []byte, dst *string) {
	conn, err := net.Dial("udp", *dst)
	checkError(err)
	n, err := conn.Write(data)
	checkError(err)
	fmt.Println("Sent " + strconv.Itoa(n) + " bytes")
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func (h Headers) getHeaderIndex(headername string) int {
	if len(headername) >= 2 {
		for index, s := range h {
			if s == headername {
				return index
			}
		}
	}
	return -1
}