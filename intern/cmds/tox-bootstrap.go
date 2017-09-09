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
	rc = dht.BootstrapFromAddress("127.0.0.1", false, 33445, "398C8161D038FD328A573FFAA0F5FAAF7FFDE5E8B4350E7D15E6AFD0B993FC52")
	rc = dht.BootstrapFromAddress("85.172.30.117", false, 33445, "8E7D0B859922EF569298B4D261A8CCB5FEA14FB91ED412A7603A585A25698832")

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
			if toxin.NeedDump(dht) {
				dht.Dump()
				log.Println()
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	select {}

}
