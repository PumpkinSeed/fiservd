package handler

import (
	"github.com/fluidpay/iso8583"
	"log"
	"time"
)

const (
	allTimeAnswer = `1210FA304555AAE4840600000000100000001600000000000000000920000000000000000000000000000123205001030402950123154952591221010121314C2005912950123000000020000000000020000111007640125111101111111182954212248887288158=99120010109012401      002NJ02011173420          58NJ NEWARK          123 PINE STREET      USWRIGHT AID DRUGS0030008400230202017840D00000001500077700101231110222222226`
)

func Handle(data string) string {
	log.Printf("Message arrived at %s\n", time.Now().UTC().String())
	var m = iso8583.Message{}
	m.SetEncoder(iso8583.ASCII)
	if err := m.Decode([]byte(data)); err != nil {
		log.Println(err.Error())
	}

	time.Sleep(1200*time.Millisecond)
	log.Printf("Primary account number is: %s\n", m.DE2.Value)
	// Check for fun
	if string(m.DE2.Value) == "4846811212" {
		return allTimeAnswer+"\n"
	}

	return allTimeAnswer+"\n"
}