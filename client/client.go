package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PumpkinSeed/fiservd/bridge"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	reqData = `1200F230040102A0000000000000040000001048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community11112341234234`
)

var (
	hc = http.DefaultClient
)

func Load(host, port string) error {
	reqJSON, err := json.Marshal(bridge.Wrapper{reqData})
	if err != nil {
		return err
	}

	for i := 0; i<10;i++ {
		go do(host, port, reqJSON)
	}

	return nil
}

func do(host, port string, reqJSON []byte) {
	req, err := http.NewRequest("POST", host+port, bytes.NewBuffer(reqJSON))
	if err != nil {
		log.Println(err.Error())
	}

	resp, err := hc.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(data))
}
