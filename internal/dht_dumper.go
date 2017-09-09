package toxin

/*
#include "DHT.h"
#include "tox.h"

*/
import "C"
import (
	"encoding/hex"
	"log"
	"strings"
	"time"
	"unsafe"
)

// not include native by underly library code
// for help use

func (this *DHT) Dump() {
	/*
	       Networking_Core *net;

	       Client_data    close_clientlist[LCLIENT_LIST];
	       uint64_t       close_lastgetnodes;
	       uint32_t       close_bootstrap_times;

	       // Note: this key should not be/is not used to transmit any sensitive materials
	       uint8_t      secret_symmetric_key[crypto_box_KEYBYTES];
	       // DHT keypair
	       uint8_t self_public_key[crypto_box_PUBLICKEYBYTES];
	       uint8_t self_secret_key[crypto_box_SECRETKEYBYTES];

	       DHT_Friend    *friends_list;
	       uint16_t       num_friends;

	       Node_format   *loaded_nodes_list;
	       uint32_t       loaded_num_nodes;
	       unsigned int   loaded_nodes_index;

	       Shared_Keys shared_keys_recv;
	       Shared_Keys shared_keys_sent;

	       struct PING   *ping;
	       Ping_Array    dht_ping_array;
	       Ping_Array    dht_harden_ping_array;
	   #ifdef ENABLE_ASSOC_DHT
	       struct Assoc  *assoc;
	   #endif
	       uint64_t       last_run;

	       Cryptopacket_Handles cryptopackethandlers[256];

	       Node_format to_bootstrap[MAX_CLOSE_TO_BOOTSTRAP_NODES];
	       unsigned int num_to_bootstrap;

	*/

	pubkey_bin := C.GoBytes(unsafe.Pointer(&this.dht.self_public_key[0]), C.TOX_PUBLIC_KEY_SIZE)
	pubkey := strings.ToUpper(hex.EncodeToString(pubkey_bin))
	log.Println("=====DUMP BEGIN=====")
	log.Println("pubkey:", pubkey, len(pubkey))
	seckey_bin := C.GoBytes(unsafe.Pointer(&this.dht.self_secret_key[0]), C.TOX_SECRET_KEY_SIZE)
	seckey := strings.ToUpper(hex.EncodeToString(seckey_bin))
	log.Println("seckey:", seckey, len(seckey))
	log.Println("close_lastgetnodes:", this.dht.close_lastgetnodes, time.Unix(int64(this.dht.close_lastgetnodes), 0))
	log.Println("close_bootstrap_times:", this.dht.close_bootstrap_times)
	log.Println("num_friends:", this.dht.num_friends)
	log.Println("loaded_num_nodes:", this.dht.loaded_num_nodes)
	log.Println("loaded_nodes_index:", this.dht.loaded_nodes_index)
	log.Println("last_run:", this.dht.last_run, time.Unix(int64(this.dht.last_run), 0))
	log.Println("num_to_bootstrap:", this.dht.num_to_bootstrap)
	log.Println("num_close_clientlist:", NewClientDataListFrom((**C.Client_data)((unsafe.Pointer)(&this.dht.close_clientlist)), C.LCLIENT_LIST).Count())
	log.Println("dht size:", this.Size())
	NewDHTFriendListFrom(this.dht.friends_list, this.dht.num_friends).Dump()

	//
	log.Println("dht is connected:", this.IsConnected())
}

// 查看是否有状态变化
var last_dht C.DHT
var last_connected int = 0

func NeedDump(dht *DHT) (need bool) {
	changed := []string{}

	connected := dht.IsConnected()
	if connected != last_connected {
		changed = append(changed, "connected")
		need = true
	}

	dhtc := dht.dht
	if dhtc.num_to_bootstrap != last_dht.num_to_bootstrap {
		changed = append(changed, "num_to_bootstrap")
		need = true
	}
	if dhtc.loaded_nodes_index != last_dht.loaded_nodes_index {
		changed = append(changed, "loaded_nodes_index")
		need = true
	}
	if dhtc.loaded_num_nodes != last_dht.loaded_num_nodes {
		changed = append(changed, "loaded_num_nodes")
		need = true
	}
	if dhtc.num_friends != last_dht.num_friends {
		changed = append(changed, "num_friends")
		need = true
	}
	if dhtc.close_bootstrap_times != last_dht.close_bootstrap_times {
		changed = append(changed, "close_bootstrap_times")
		need = true
	}

	last_dht = *dht.dht
	last_connected = connected

	if need {
		log.Println(strings.Join(changed, ", "))
	}
	return
}

func (this *DHT) DumpFriends() {
	NewDHTFriendListFrom(this.dht.friends_list, this.dht.num_friends)
}

/////////////
func (this *ClientDataList) Dump() {
	for i := 0; i < this.len; i++ {
		cdx := addrStep(unsafe.Pointer(this.cds), i*C.sizeof_Client_data)
		cd := (*C.Client_data)(cdx)
		if is_timeout(uint64(cd.assoc4.timestamp), uint64(C.BAD_NODE_TIMEOUT)) {
			continue
		}

		port := lendian_to_host16(uint16(cd.assoc4.ip_port.port))
		port = net_to_host(uint16(cd.assoc4.ip_port.port))
		pubkey := ppkey2str(&cd.public_key[0])
		log.Println(i, port, cd.assoc4.ip_port.ip.family,
			ip_ntoa(&cd.assoc4.ip_port.ip), pubkey)
	}
}

////////////////

func (this *DHTFriendList) Dump() {
	for i := 0; i < this.num; i++ {
		fix := addrStep(unsafe.Pointer(this.frs), C.sizeof_DHT_Friend*i)
		fi := (*C.DHT_Friend)(fix)
		pubkey := ppkey2str(&fi.public_key[0])
		if false {
			log.Println(i, fi.num_to_bootstrap, pubkey)
		}

		cds := NewClientDataListFrom((**C.Client_data)((unsafe.Pointer)(&fi.client_list)), C.MAX_FRIEND_CLIENTS)
		if false {
			log.Println(i, cds.Count())
			cds.Dump()
		}
		this.DumpBootstrapNodes()
	}
}

func (this *DHTFriendList) DumpClientData() {
}

func (this *DHTFriendList) DumpBootstrapNodes() {
	log.Println("Dumping...", this.frs.num_to_bootstrap)
	NewNodeFormatList(&this.frs.to_bootstrap[0], int(this.frs.num_to_bootstrap)).Dump()
}

//////////////////

func (this *NodeFormatList) Dump() {
	for idx := 0; idx < this.num; idx++ {
		nfx := addrStep(unsafe.Pointer(this.nfs), idx*C.sizeof_Node_format)
		nf := (*C.Node_format)(nfx)

		pubkey := ppkey2str(&nf.public_key[0])
		port := net_to_host(uint16(nf.ip_port.port))
		ipaddr := ip_ntoa(&nf.ip_port.ip)
		if false {
			log.Println(idx, port, ipaddr, pubkey)
		}
	}
}

///////////
