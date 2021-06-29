package traefik_vault_auth

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

// Routes used to request Kuzzle, can be customized
type Routes struct {
	// Login route used to log in to Kuzzle using Auth Basic user/pass.
	// The specified route must return 200 HTTP status code and a valid JWT when called by anonymous user.
	// Default is '/_login/local' (see: https://docs.kuzzle.io/core/2/api/controllers/auth/login/)
	// Login route using 'local' strtategy (see: https://docs.kuzzle.io/core/2/guides/main-concepts/authentication/#local-strategy)
	// It must accept JSON body containing 'username' and 'password' string fields, for example:
	// 	{
	// 		"username": "myUser",
	// 		"password": "myV3rys3cretP4ssw0rd"
	// 	}
	// You would like to update this route if you do not use 'local' strategy on your Kuzzle server
	Login string `yaml:"login,omitempty"`
}

// Vault info
type Vault struct {
	// URL use by the plugin to reach Kuzzle server.
	// NOTE: Only HTTP(s) protocol is supported
	// Examples:
	//  - HTTP: http://localhost:7512
	//	- HTTPS: https://localhost:7512
	URL    string `yaml:"url"`
	Routes Routes `yaml:"routes,omitempty"`
}

func (k *Vault) login(user string, password string) error {
// 	reqBody, _ := json.Marshal(map[string]string{
// 		"username": user,
// 		"password": password,
// 	})
	client := &http.Client{
	}

	url := fmt.Sprintf("%s%s", k.URL, k.Routes.Login)

    fmt.Println("url  "  + url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Vault-Token", "s.qztLjR0vjFm8jaB2RDUMGeyq")

    resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("Authentication request send to %s failed: %v", url, err)
	}


	if resp.StatusCode != 200 {
		return fmt.Errorf("Authentication request send to %s failed: status code %d", url, resp.StatusCode)
	}

	var jsonBody map[string]interface{}

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &jsonBody); err != nil {
		return err
	}

	res_pass := jsonBody["data"].(map[string]interface{})["data"].(map[string]interface{})[user].(string)

    fmt.Println("res_pass  "  + res_pass)
	if(res_pass != password){
		return fmt.Errorf("User %s do not have the right password: %v", user, password)
	}

// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", k.JWT))

	return nil
}

// func (k *Kuzzle) checkUser() error {
// 	client := &http.Client{}
// 	url := fmt.Sprintf("%s%s", k.URL, k.Routes.GetCurrentUser)
//
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", k.JWT))
// 	resp, err := client.Do(req)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	var jsonBody map[string]interface{}
// 	body, _ := ioutil.ReadAll(resp.Body)
//
// 	if err := json.Unmarshal(body, &jsonBody); err != nil {
// 		return err
// 	}
//
// 	kuid := jsonBody["result"].(map[string]interface{})["_id"].(string)
// 	for _, id := range k.AllowedUsers {
// 		if kuid == id {
// 			return nil
// 		}
// 	}
//
// 	return fmt.Errorf("User %s do not be part of allowed users: %v", kuid, k.AllowedUsers)
// }
