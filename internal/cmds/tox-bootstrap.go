// implementation tox-bootstrap in golang
package main

import (
	"log"

	"github.com/kitech/colog"
	toxin "github.com/kitech/go-toxcore/internal"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	colog.Register()
}

func main() {
	var ip toxin.IP
	(&ip).Init()
	log.Println(ip)

	netcore := toxin.NewNetworkCore(ip, 12345)
	log.Println(netcore)
	dht := toxin.NewDHT(netcore)
	log.Println(dht)
}
