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

/////
//export onFriendIPResponse
func onFriendIPResponse(data unsafe.Pointer, number C.int32_t, ip_port C.IP_Port) {
	uds := cbi.getmapping(data)
	ipcbfn := uds[0].(func(interface{}, int32, string, uint16))
	data_ := uds[1]
	number_ := uds[2].(int32)

	ip := ip_ntoa(&ip_port.ip)
	port := net_to_host(uint16(ip_port.port))

	ipcbfn(data_, number_, ip, port)
	// TODO unmaping it?
}
func (this *DHT) AddFriend(pubkey string, ipcbfn func(interface{}, int32, string, uint16), data interface{}, number int32) int {
	retfn, retud := cbi.mapping(unsafe.Pointer(C.onFriendIPResponse), ipcbfn, data, number)

	pubkey_ := [C.TOX_PUBLIC_KEY_SIZE]byte{}
	str2ppkey(pubkey, (*C.uint8_t)(&pubkey_[0]))
	var lock_count C.uint16_t
	r := this.AddFriendC((*C.uint8_t)(&pubkey_[0]), retfn, retud, C.int32_t(number), &lock_count)
	return int(r)
}
func (this *DHT) AddFriendC(pubkey *C.uint8_t, ipcbfn *[0]byte, data unsafe.Pointer, number C.int32_t, lock_count *C.uint16_t) C.int {
	r, err := C.DHT_addfriend(this.dht, pubkey, ipcbfn, data, number, lock_count)
	if err != nil {
		log.Println(err, r)
	}
	return r
}

func (this *DHT) DelFriend() {}

func (this *DHT) GetFriendIP(pubkey string) (ip string, port uint16) {
	// int DHT_getfriendip(const DHT *dht, const uint8_t *public_key, IP_Port *ip_port);

	pubkey_ := [C.TOX_PUBLIC_KEY_SIZE]byte{}
	str2ppkey(pubkey, (*C.uint8_t)(&pubkey_[0]))
	var ip_port C.IP_Port

	r, err := C.DHT_getfriendip(this.dht, (*C.uint8_t)(&pubkey_[0]), &ip_port)
	if err != nil {
		log.Println(err)
	}
	// log.Println(r, net_to_host(uint16(ip_port.port)))
	// log.Println(ip_ntoa(&ip_port.ip))
	if r == 1 {
		ip = ip_ntoa(&ip_port.ip)
		port = net_to_host(uint16(ip_port.port))
	}

	if true {
		friendo := addrStep(unsafe.Pointer(this.dht.friends_list), 2*C.sizeof_DHT_Friend)
		friendo_ := (*C.DHT_Friend)(friendo)
		pk_ := ppkey2str(&friendo_.public_key[0])
		log.Println(pk_)
	}

	return
}

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
	ret, err := C.DHT_bootstrap_from_address(this.dht, addr_, ipv6enabled_, port_, pubkey__)
	if err != nil {
		log.Println(err)
	}
	if ret == 1 {
		return true
	}
	return false
}

func (this *DHT) ConnectAfterLoad() int {
	r := C.DHT_connect_after_load(this.dht)
	return int(r)
}
func (this *DHT) RoutePacket(pubkey string, packet []byte) int {
	pubkey_ := [C.TOX_PUBLIC_KEY_SIZE]byte{}
	str2ppkey(pubkey, (*C.uint8_t)(&pubkey_[0]))
	r := C.route_packet(this.dht, (*C.uint8_t)(&pubkey_[0]), (*C.uint8_t)(&packet[0]), C.uint16_t(len(packet)))
	_ = r
	return int(r)
}
func (this *DHT) RouteToFriend(pubkey string, packet []byte) int {
	pubkey_ := [C.TOX_PUBLIC_KEY_SIZE]byte{}
	str2ppkey(pubkey, (*C.uint8_t)(&pubkey_[0]))
	r := C.route_tofriend(this.dht, (*C.uint8_t)(&pubkey_[0]), (*C.uint8_t)(&packet[0]), C.uint16_t(len(packet)))
	_ = r
	return int(r)
}
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
	cbmap map[int32][]interface{}
	mu    sync.Mutex
}

func newcbimpl() *cbimpl {
	this := &cbimpl{}
	this.cbmap = make(map[int32][]interface{})
	return this
}

func (this *cbimpl) mapping(fn unsafe.Pointer, uds ...interface{}) (
	retfn *[0]byte, retud unsafe.Pointer) {
	retfn = (*[0]byte)(fn)
	id := atomic.AddInt32(&this.cbid, 1)
	retud = unsafe.Pointer(uintptr(id))
	this.cbmap[id] = uds
	return
}

func (this *cbimpl) unmaping(udc unsafe.Pointer) (ud interface{}) {
	return
}

func (this *cbimpl) getmapping(udc unsafe.Pointer) (uds []interface{}) {
	id := int32(uintptr(udc))
	if udx, ok := this.cbmap[id]; ok {
		uds = udx
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
	this := udx[0].(*DHT)
	// log.Println(this)
	if this.getnodes_response != nil {
		ip := ip_ntoa(&ip_port.ip)
		port := net_to_host(uint16(ip_port.port))
		pubkey_ := ppkey2str(pubkey)
		this.getnodes_response(ip, port, pubkey_, this.cb_user_data)
	}
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

//////////////
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

////////////////
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

///////////
