package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// 	if (host == nil) || (user == nil) || (pass == nil) {
// 	return "", errors.New("invalid parameters")
// }
// var jsonStr = []byte(`{"username":"` + *user + `,"password":"` + *pass + `"}`)

// req, _ := http.NewRequest("POST", *host+"/api/users/adm/authenticate", bytes.NewBuffer(jsonStr))

// req.Header.Set("Content-Type", "application/json")

// resp, err := http.DefaultClient.Do(req)
// if err != nil {
// 	return nil, err
// }
// var result Auth

// defer resp.Body.Close()
// body, err := ioutil.ReadAll(resp.Body)
// if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
// 	fmt.Println("Can not unmarshal JSON")
// }
// return result.Response.AuthToken, err

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}
	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signin", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
