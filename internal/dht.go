package toxin

/*
#include "DHT.h"
#include "tox.h"
*/
import "C"
import "unsafe"

import (
	"encoding/hex"
	"log"
	"strings"
	"time"
)

func PackedNodeSize(ip_family uint8) int {
	r := C.packed_node_size(C.uint8_t(ip_family))
	return int(r)
}

func PackNodes(data []byte, nodes *C.Node_format, number uint16) int {
	data_ := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	len_ := C.uint16_t(len(data))
	r := C.pack_nodes(data_, len_, nodes, (C.uint16_t)(number))
	return int(r)
}

func UnpackNodes() {}

type DHT struct {
	dht *C.DHT
}

func GetSharedKey()                 {}
func (this *DHT) GetSharedKeyRecv() {}
func (this *DHT) GetSharedKeySent() {}
func (this *DHT) GetNodes()         {}
func (this *DHT) AddFriend()        {}
func (this *DHT) DelFriend()        {}

func IdClosest() {}
func AddToList() {}

func (this *DHT) NodeAddableToCloseList() {}
func (this *DHT) GetCloseNodes()          {}
func (this *DHT) RandFriendsNodes()       {}
func (this *DHT) CloseListNodes()         {}
func (this *DHT) Do()                     { C.do_DHT(this.dht) }

func (this *DHT) Bootstrap() {}

/*
int DHT_bootstrap_from_address(DHT *dht, const char *address, uint8_t ipv6enabled,
                               uint16_t port, const uint8_t *public_key);
"178.62.250.138", "33445", "788236D34978D1D5BD822F0A5BEBD2C53C64CC31CD3149350EE27D4D9A2F9B6B",
*/
func (this *DHT) BootstrapFromAddress(address string,
	ipv6enabled bool, port uint16, pubkey string) bool {
	addr_ := C.CString(address)
	defer C.free(unsafe.Pointer(addr_))

	var ipv6enabled_ C.uint8_t
	if ipv6enabled {
		ipv6enabled_ = 1
	}
	port_ := C.htons(C.uint16_t(port))
	log.Println("port:", port, port_)
	pubkey_, err := hex.DecodeString(pubkey)
	if err != nil {
		log.Println(err)
	}
	pubkey__ := (*C.uint8_t)((unsafe.Pointer)(&pubkey_[0]))
	ret := C.DHT_bootstrap_from_address(this.dht, addr_, ipv6enabled_, port_, pubkey__)
	if ret == 1 {
		return true
	}
	return false
}

func (this *DHT) ConnectAfterLoad()             {}
func (this *DHT) RoutePacket()                  {}
func (this *DHT) RouteToFriend()                {}
func (this *DHT) cryptopacket_registerhandler() {}
func (this *DHT) Size()                         {}
func (this *DHT) Save()                         {}
func (this *DHT) Load()                         {}
func NewDHT(net *NetworkCore) *DHT {
	this := &DHT{}
	this.dht = C.new_DHT(nil, net.net, true)

	return this
}
func (this *DHT) Kill()            {}
func (this *DHT) IsConnected() int { return int(C.DHT_isconnected(this.dht)) }
func (this *DHT) NonLanConnected() {}
func (this *DHT) AddToList()       {}

func (this *DHT) SecretKey() []byte {
	return C.GoBytes((unsafe.Pointer)(&this.dht.self_secret_key[0]), C.TOX_SECRET_KEY_SIZE)
}

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

	//
	log.Println("dht is connected:", this.IsConnected())
}
