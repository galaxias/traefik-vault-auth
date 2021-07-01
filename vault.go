package traefik_vault_auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Routes struct {
	Login string `yaml:"login,omitempty"`
}

// Vault info
type Vault struct {
	URL    string `yaml:"url"`
	Token  string `yaml:"token"`
	Routes Routes `yaml:"routes,omitempty"`
}

func (k *Vault) login(user string, password string) error {

	client := &http.Client{}

	url := fmt.Sprintf("%s%s", k.URL, k.Routes.Login)

    fmt.Println("url  "  + url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Vault-Token", k.Token)

    resp, err := client.Do(req)

    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

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

	return nil
}
