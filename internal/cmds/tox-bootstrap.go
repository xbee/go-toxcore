// implementation tox-bootstrap in golang
package main

import (
	"log"
	"os"
	"time"

	"github.com/kitech/colog"
	"github.com/kitech/go-toxcore/internal"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	colog.Register()
}

func main() {
	var ip toxin.IP
	(&ip).Init()

	netcore := toxin.NewNetworkCore(ip, 12345)
	dht := toxin.NewDHT(netcore)

	onion := toxin.NewOnion(dht)
	onion_an := toxin.NewOnionAnnounce(dht)
	if onion == nil || onion_an == nil {
		log.Println("error: ", "Couldn't initialize Tox Onion. Exiting.")
		os.Exit(-1)
	}
	log.Println(onion, onion_an)

	secret_key := dht.SecretKey()
	tcpsrv := toxin.NewTCPServer(true, []uint16{33445}, secret_key, onion)
	log.Println(tcpsrv)

	rc := dht.BootstrapFromAddress("205.185.116.116", false, 33445, "A179B09749AC826FF01F37A9613F6B57118AE014D4196A0E1105A98F93A54702")
	log.Println(rc)
	if !rc {
		log.Println(rc)
	}
	dht.Dump()

	dht.LANdiscoveryInit()

	go func() {
		for {
			dht.Do()
			tcpsrv.Do()
			netcore.Poll()
			dht.Dump()
			log.Println()
			time.Sleep(200 * time.Millisecond)
		}
	}()

	select {}

}
