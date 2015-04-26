package main

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/digitalocean/godo"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	regions := []string{"nyc3", "nyc2", "ams3", "ams2", "fra1", "lon1", "sgp1", "sfo1"}

	pat := os.Getenv("DO_PAT")
	if pat == "" {
		fmt.Println("DO_PAT is required")
		return
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: pat},
	}
	client := godo.NewClient(t.Client())

	num, _ := strconv.ParseInt(os.Args[1], 10, 64)

	keys, _, _ := client.Keys.List(nil)

	var k godo.DropletCreateSSHKey
	for _, key := range keys {
		k = godo.DropletCreateSSHKey{
			ID:          key.ID,
			Fingerprint: key.Fingerprint,
		}
	}

	ks := []godo.DropletCreateSSHKey{k}

	var i int64
	for i = 0; i < num; i++ {
		resp, err := http.Get("http://meaningless-identifier.herokuapp.com/text")
		if err != nil {
			fmt.Printf("Something bad happened: %s\n\n", err)
			return
		}

		defer resp.Body.Close()

		byteArray, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Something bad happened: %s\n\n", err)
			return
		}

		name := string(byteArray)

		rand.Seed(time.Now().UnixNano())

		createRequest := &godo.DropletCreateRequest{
			Name:   fmt.Sprintf(name),
			Region: regions[rand.Intn(len(regions))],
			Size:   "512mb",
			Image: godo.DropletCreateImage{
				Slug: "ubuntu-14-04-x64",
			},
			SSHKeys: ks,
		}
		newDroplet, _, err := client.Droplets.Create(createRequest)

		if err != nil {
			fmt.Printf("Something bad happened: %s\n\n", err)
			return
		}

		fmt.Printf("%v\n", newDroplet)
	}
}
