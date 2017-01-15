package toxin

import (
	"log"
	"testing"
	"time"
)

func TestDHT(t *testing.T) {
	var ip IP
	(&ip).Init()
	net := NewNetworkCore(ip, 12345)
	dht := NewDHT(net)
	t.Log(net, dht)
	if net == nil || dht == nil {
		t.Error("nil")
	}

	onion := NewOnion(dht)
	onion_an := NewOnionAnnounce(dht)
	if onion == nil || onion_an == nil {
		t.Error("nil")
	}

	// 205.185.116.116 33445	A179B09749AC826FF01F37A9613F6B57118AE014D4196A0E1105A98F93A54702
	// rc := dht.BootstrapFromAddress("178.62.250.138", false, 33445, "788236D34978D1D5BD822F0A5BEBD2C53C64CC31CD3149350EE27D4D9A2F9B6B")
	rc := dht.BootstrapFromAddress("205.185.116.116", false, 33445, "A179B09749AC826FF01F37A9613F6B57118AE014D4196A0E1105A98F93A54702")
	log.Println(rc)
	if !rc {
		t.Error(rc)
	}
	dht.Dump()

	go func() {
		for {
			dht.Do()
			net.Poll()
			dht.Dump()
			log.Println()
			time.Sleep(200 * time.Millisecond)
		}
	}()

	select {}
}
