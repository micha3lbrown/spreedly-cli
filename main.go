package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/micha3lbrown/spreedly-cli/pkg/client"
	"github.com/micha3lbrown/spreedly-cli/pkg/gateway"
)

func main() {
	baseURL, err := url.Parse(client.SpreedlyAPIURL)
	if err != nil {
		fmt.Println(err)
	}
	var auth = client.Auth{
		AccessKey:    os.Getenv("SPREEDLY_ACCESS_KEY"),
		AccessSecret: os.Getenv("SPREEDLY_ACCESS_SECRET"),
	}

	var client = &client.Client{
		BaseURL: baseURL,
		Auth:    auth,
	}

	// ng := gateways.Create(client, *g)
	// fmt.Println(ng.Token)
	// gateways.Get(client, ng.Token)
	gateway.List(client)

	// var p pm.PaymentMethod
	// p := pm.Get(client, "blah blah blah")
	// p := pm.List(client)
	// fmt.Println(p)

}

// API interaction
//####################################3
//spreedly pm get $token
//		   pm add -data
// 		   gateway get
// 		   gateway add
