// Can you determine the values of the sequence number, acknowledgment number, source port and destination port?

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("tcpheader")
	if err != nil {
		fmt.Printf("Error reading tcpheader file %s\n", err)
	}

	sourcePort := int(dat[0])<<8 + int(dat[1])
	destinationPort := int(dat[2])<<8 + int(dat[3])
	sequenceNumber := int(dat[4])<<24 + int(dat[5])<<16 + int(dat[6])<<8 + int(dat[7])
	ackSet := dat[13]&0b00010000 == 0b00010000
	ackNumber := 0
	if ackSet {
		ackNumber = int(dat[8])<<24 + int(dat[9])<<16 + int(dat[10])<<8 + int(dat[11])
	}
	lengthOfHeader := dat[12] >> 4 * 4
	// Figure out the convention for string formatting in this case
	fmt.Printf("source port %d, destination port %d, sequence number %d, ack is set %t, ack number %d, length of header (bytes) %d\n", sourcePort, destinationPort, sequenceNumber, ackSet, ackNumber, lengthOfHeader)
}
