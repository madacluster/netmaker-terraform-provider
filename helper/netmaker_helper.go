package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	Code     int    `json:"Code"`
	Message  string `json:"Message"`
	Response struct {
		UserName  string `json:"UserName"`
		AuthToken string `json:"AuthToken"`
	} `json:"Response"`
}

func NewClient(host, user, pass *string) (string, error) {

	if (host == nil) || (user == nil) || (pass == nil) {
		return "", errors.New("invalid parameters")
	}
	var jsonStr = []byte(`{"username":"` + *user + `,"password":"` + *pass + `"}`)

	req, _ := http.NewRequest("POST", *host+"/api/users/adm/authenticate", bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	var result Auth

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return result.Response.AuthToken, err
}
