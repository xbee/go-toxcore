package toxin

/*
#include "DHT.h"
#include "tox.h"

#include "toxin_cgo_export.h"
*/
import "C"
import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
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

	getnodes_response func(string, uint16, string, interface{})
	cb_user_data      interface{}
}

func GetSharedKey()                 {}
func (this *DHT) GetSharedKeyRecv() {}
func (this *DHT) GetSharedKeySent() {}
func (this *DHT) GetNodes(ip_port *C.IP_Port, pubkey *C.uint8_t, clientid *C.uint8_t) {
	// var ip_port *C.IP_Port
	C.DHT_getnodes(this.dht, ip_port, pubkey, clientid)
}
func (this *DHT) AddFriend() {}
func (this *DHT) DelFriend() {}

func IdClosest()  {}
func AddToLists() {}

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
func (this *DHT) Size() uint32 {
	return uint32(C.DHT_size(this.dht))
}
func (this *DHT) Save() {}
func (this *DHT) Load() {}
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

//////
type cbimpl struct {
	cbid  int32
	cbmap map[int32]interface{}
	mu    sync.Mutex
}

func newcbimpl() *cbimpl {
	this := &cbimpl{}
	this.cbmap = make(map[int32]interface{})
	return this
}

func (this *cbimpl) mapping(fn unsafe.Pointer, ud interface{}) (
	retfn *[0]byte, retud unsafe.Pointer) {
	retfn = (*[0]byte)(fn)
	id := atomic.AddInt32(&this.cbid, 1)
	retud = unsafe.Pointer(uintptr(id))
	this.cbmap[id] = ud
	return
}

func (this *cbimpl) unmaping(udc unsafe.Pointer) (ud interface{}) {
	return
}

func (this *cbimpl) getmapping(udc unsafe.Pointer) (ud interface{}) {
	id := int32(uintptr(udc))
	if udx, ok := this.cbmap[id]; ok {
		ud = udx
		return
	}
	return
}

var cbi = newcbimpl()

////////
// extra: set get_nodes response callback
func (this *DHT) CallbackGetnodesResponse(cbfn func(string, uint16, string, interface{}), userData interface{}) {
	this.getnodes_response = cbfn
	this.cb_user_data = userData

	// cbfn_ := (*[0]byte)(C.onGetnodesResponse)
	cbfn_, cbud := cbi.mapping(C.onGetnodesResponse, this)
	C.DHT_callback_getnodes_response(this.dht, cbfn_, cbud)
}

// 直接使用 unsafe.Pointer(this)，如果go对这个内存进行了移动操作，则程序挂
//export onGetnodesResponse
func onGetnodesResponse(ip_port *C.IP_Port, pubkey *C.uint8_t, ud unsafe.Pointer) {
	udx := cbi.getmapping(ud)
	this := udx.(*DHT)
	// log.Println(this)
	if this.getnodes_response != nil {
		ip := ip_ntoa(&ip_port.ip)
		port := net_to_host(uint16(ip_port.port))
		pubkey_ := ppkey2str(pubkey)
		this.getnodes_response(ip, port, pubkey_, this.cb_user_data)
	}
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

type ClientData struct {
	cd *C.Client_data
}

type ClientDataList struct {
	cds **C.Client_data
	len int
}

func NewClientDataListFrom(cds **C.Client_data, len int) *ClientDataList {
	this := &ClientDataList{cds, len}
	return this
}

func addrStep(addr unsafe.Pointer, step int) unsafe.Pointer {
	naddr := uintptr(addr) + uintptr(step)
	return unsafe.Pointer(naddr)
}

func ppkey2str(key *C.uint8_t) string {
	pubkey_bin := C.GoBytes(unsafe.Pointer(key), C.TOX_PUBLIC_KEY_SIZE)
	pubkey := strings.ToUpper(hex.EncodeToString(pubkey_bin))
	return pubkey
}

func str2ppkey(pubkey string, key *C.uint8_t) {
	keybin, _ := hex.DecodeString(pubkey)
	if len(keybin) != C.TOX_PUBLIC_KEY_SIZE {
		log.Panicln("Wtf", len(keybin))
	}
	C.memcpy(unsafe.Pointer(key), unsafe.Pointer(&keybin[0]), C.TOX_PUBLIC_KEY_SIZE)
}

func (this *ClientDataList) Count() (cnt int) {
	for idx := 0; idx < this.len; idx++ {
		cdx := addrStep(unsafe.Pointer(this.cds), idx*C.sizeof_Client_data)
		cd := (*C.Client_data)(cdx)
		if is_timeout(uint64(cd.assoc4.timestamp), uint64(C.BAD_NODE_TIMEOUT)) {
			continue
		}
		cnt++
	}
	return
}

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

type DHTFriend struct {
	f *C.DHT_Friend
}

type DHTFriendList struct {
	frs *C.DHT_Friend
	num int
}

func NewDHTFriendListFrom(frs *C.DHT_Friend, num C.uint16_t) *DHTFriendList {
	this := &DHTFriendList{frs, int(num)}
	return this
}

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

type NodeFormat struct {
	nf *C.Node_format
}

func NewNodeFormatFrom(nf *C.Node_format) *NodeFormat {
	return &NodeFormat{nf}
}

func (this *NodeFormat) Pubkey() string {
	return ppkey2str(&this.nf.public_key[0])
}

func (this *NodeFormat) Port() uint16 {
	return net_to_host(uint16(this.nf.ip_port.port))
}

func (this *NodeFormat) IP() string {
	return ip_ntoa(&this.nf.ip_port.ip)
}

func (this *NodeFormat) IPPort() string {
	return fmt.Sprintf("%s:%d", this.IP(), this.Port())
}

func (this *NodeFormat) IPPortC() *C.IP_Port {
	return &this.nf.ip_port
}

func (this *NodeFormat) PubkeyC() *C.uint8_t {
	return &this.nf.public_key[0]
}

func (this *NodeFormat) Set(ip string, port uint16, pubkey string) {
	port_ := host_to_net(port)
	this.nf.ip_port.port = C.uint16_t(port_)
	str2ppkey(pubkey, &this.nf.public_key[0])
	addr_parse_ip(ip, &this.nf.ip_port.ip)
}

/////
type IP_Port2 struct {
	ipp *C.IP_Port
}

func NewIPPort2From(ipp *C.IP_Port) *IP_Port2 {
	return &IP_Port2{ipp}
}

///
type NodeFormatList struct {
	nfs *C.Node_format
	num int
}

func NewNodeFormatN(n int) *NodeFormatList {
	this := &NodeFormatList{}
	this.num = n

	this.nfs = (*C.Node_format)(C.malloc(C.sizeof_Node_format * C.size_t(n)))
	if this.nfs == nil {
		return nil
	}
	return this
}

func (this *NodeFormatList) Get(index int) *NodeFormat {
	nfx := addrStep(unsafe.Pointer(this.nfs), index*C.sizeof_Node_format)
	nf := (*C.Node_format)(nfx)
	return NewNodeFormatFrom(nf)
}

func (this *NodeFormatList) Expand() (int, bool) {
	nfs, err := C.realloc(unsafe.Pointer(this.nfs), C.size_t(2*this.num)*C.sizeof_Node_format)
	if nfs == nil {
		log.Println(err)
		return this.num, false
	}
	this.nfs = (*C.Node_format)(nfs)
	this.num *= 2
	return this.num, true
}

func NewNodeFormatList(nfs *C.Node_format, num int) *NodeFormatList {
	this := &NodeFormatList{nfs, num}
	return this
}

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
