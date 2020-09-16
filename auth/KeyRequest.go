package auth

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//key request to user-service
func getKey() (k *rsa.PublicKey) {
	temp, _ := http.Get("http://localhost:8888/key")
	b, err := ioutil.ReadAll(temp.Body)
	if err != nil {
		fmt.Printf("Problem with token occured ")
	}

	err = json.Unmarshal(b, &k)
	if err != nil {
		fmt.Printf("Could not Unmarshal key ")
		return
	}
	return
}
