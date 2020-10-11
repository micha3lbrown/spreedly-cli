package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/micha3lbrown/spreedly-cli/pkg/client"
)

// New instantiates a new Gateway, expects type
// func (g *Gateway) New(s string) *Gateway {
// 	return &Gateway{
// 		GatewayType: s,
// 	}
// }

// Create sends a POST to `/v1/gateways` for creation
func Create(client *client.Client, g Gateway) *GatewayJSON {
	url := "/gateways.json"
	gbyte, _ := json.Marshal(g)
	resp, err := client.WriteRequest(http.MethodPost, url, gbyte)

	if err != nil {
		fmt.Println(err)
	}

	var ng GatewayJSON
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &ng)
	// json.NewDecoder(resp.Body).Decode(ng)

	fmt.Println("Created Gateway: ", resp.StatusCode, ng.Token)
	return &ng
}

// Get ...
func Get(client *client.Client, token string) *GatewayJSON {
	url := "/gateways/" + token + ".json"

	resp, err := client.ReadRequest(http.MethodGet, url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get Gateway Status: ", resp.Status)
	var ng GatewayJSON
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &ng)
	// json.NewDecoder(resp.Body).Decode(&ng)

	fmt.Println("Get Gateway: ", ng.Name, ng.Token)
	return &ng

}

// List ....
func List(client *client.Client) *GatewayList {
	url := "/gateways.json"

	resp, err := client.ReadRequest(http.MethodGet, url)
	if err != nil {
		fmt.Println(err)
	}

	var ng GatewayList
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &ng)

	fmt.Println("Total Gateways: ", len(ng.Gateways))

	for _, g := range ng.Gateways {
		fmt.Println("Token:\t", g.Token, "\tGatewayType:\t", g.GatewayType)
	}

	return &ng
}
