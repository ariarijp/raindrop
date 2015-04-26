package main

import (
	"fmt"
	"os"
	"code.google.com/p/goauth2/oauth"
	"github.com/digitalocean/godo"
)

func main() {
	pat := os.Getenv("DO_PAT")
	if pat == "" {
		fmt.Println("DO_PAT is required")
		return
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: pat},
	}

	client := godo.NewClient(t.Client())
	droplets, _, _ := client.Droplets.List(nil)

	for _, droplet := range droplets {
		fmt.Printf("%v\n", droplet.Networks.V4[0].IPAddress)
	}
}
