package pm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/micha3lbrown/spreedly-cli/pkg/client"
)

// Create sends a POST to `/v1/gateways` for creation
func Create(client *client.Client, pm PaymentMethod) *PaymentMethod {
	url := "/payment_methods.json"
	gbyte, _ := json.Marshal(pm)
	resp, err := client.WriteRequest(http.MethodPost, url, gbyte)

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &pm)

	fmt.Println("Created Gateway: ", resp.StatusCode, pm.Token)
	return &pm
}

// Get ...
func Get(client *client.Client, token string) *PaymentMethod {
	url := "/payment_methods/" + token + ".json"

	resp, err := client.ReadRequest(http.MethodGet, url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get Gateway Status: ", resp.Status)
	var pm PMJSON
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &pm)
	// json.NewDecoder(resp.Body).Decode(&pm)

	fmt.Println("Get PaymentMethod: ", pm.Token)
	return &pm.PaymentMethod

}

// List ....
func List(client *client.Client) *PMList {
	url := "/payment_methods.json"

	resp, err := client.ReadRequest(http.MethodGet, url)
	if err != nil {
		fmt.Println(err)
	}

	var pml PMList
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &pml)

	fmt.Println("Total Gateways: ", len(pml.PMs))

	for _, pm := range pml.PMs {
		fmt.Println("Token:\t", pm.Token, "\tGatewayType:\t")
	}

	return &pml
}
