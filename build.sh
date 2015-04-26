go get code.google.com/p/goauth2/oauth
go get github.com/digitalocean/godo

go build -o bin/create-droplets create-droplets.go
go build -o bin/destroy-all-droplets destroy-all-droplets.go
go build -o bin/list-droplet-addresses list-droplet-addresses.go
go build -o bin/poweroff-all-droplets poweroff-all-droplets.go
go build -o bin/show-all-status show-all-status.go
