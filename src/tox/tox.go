package tox

import (
	"fmt"
	"log"
	"reflect"
	// "runtime"
	"unsafe"
)

/*
#cgo CFLAGS: -g -O2 -Wall
#cgo LDFLAGS: -ltoxcore -ltoxdns -ltoxav -ltoxencryptsave
#include <stdlib.h>
#include <tox/tox.h>

//////
// 下面的extern行不是必须的，除非这个对应的go函数在其他的文件中。
typedef void (*cb_file_send_request_ftype)(Tox *m, int32_t, uint8_t, uint64_t, uint8_t*, uint16_t, void*);
void callbackFileSendRequestWrapperForC(Tox *m, int32_t, uint8_t, uint64_t, uint8_t*, uint16_t, void*);
int CalledByCGO();
int FortytwoAbc();
static void cb_file_send_request_wrapper_for_go(Tox *tox, cb_file_send_request_ftype fn, void *userdata)
{
    tox_callback_file_send_request(tox, fn, userdata);
}


static void test_c_call_go()
{
    Tox *m = "12345";
    int32_t a = 0;
    uint8_t b = 0;
    uint64_t c = 0;
    const uint8_t *d = NULL;
    uint16_t e = 0;
    void *f = 0;

    callbackFileSendRequestWrapperForC(m, a, b, c, d, e, f);
    // CallbackFileSendRequestWrapperForC(m, a, b, c, d, e, f);
    // calledbycgo();
    FortytwoAbc();
}

*/
import "C"

type IToxer interface {
	//a int32
}

type CTox C.Tox
type CToxOptions C.Tox_Options

type Tox struct {
	// C.Tox
	i int32
	ix interface{}  // save C.Tox, dup
	x *C.Tox // save C.Tox
	iopts interface{} // C.Tox_Options
	opts *C.Tox_Options // C.Tox_Options

	// some callbacks
	cb_file_send_request func(this *Tox)
}

type Options struct {
	ix interface{}
	x *C.Tox_Options
	
	ipv6 bool
	udp_disabled bool
	proxy_type int32
	proxy_address string
	proxy_port uint16
}

// fuck,原来这个"//export"是有含义的吗，不是注释。
//export FortytwoAbc
func FortytwoAbc() C.int {
    return C.int(42)
}

func CalledByCGO() int {
	return 12
}

// 包内部函数
//export callbackFileSendRequestWrapperForC
func callbackFileSendRequestWrapperForC(m *C.Tox, a C.int32_t, b C.uint8_t, c C.uint64_t,
	d *C.uint8_t, e C.uint16_t, f unsafe.Pointer) {
	var this = (*Tox)(f)
	log.Println("called from c code", this)
	log.Println(m, a, b, c, d, e, f)
	if this.cb_file_send_request != nil {
		this.cb_file_send_request(this)
	}
}

func (this *Tox) CallbackFileSendRequest(cbfun unsafe.Pointer, userData interface{}) {
	var _userData = unsafe.Pointer(&userData)
	// var _cbfun int = callbackFileSendRequestWrapperForC
	
	var _bfun = (C.cb_file_send_request_ftype)(unsafe.Pointer(C.callbackFileSendRequestWrapperForC))
	
	C.cb_file_send_request_wrapper_for_go(this.x, _bfun, _userData);
}

func TestCCallGo() {
	log.Println("calling C...")
	C.test_c_call_go()
}

func NewTox() *Tox {
	var opts = new(C.Tox_Options)
	opts.ipv6enabled = C.uint8_t(1)
	opts.udp_disabled = C.uint8_t(0)
	fmt.Println(opts, unsafe.Pointer(opts))

	var gt = new(Tox)
	gt.iopts = opts
	gt.opts = opts

	var ct = C.tox_new(opts)
	gt.ix = ct
	gt.x = ct
	fmt.Println(reflect.TypeOf(opts), gt)

	// fmt.Println(x)
	return gt
	// return &Tox{1, gt.ix, gt.x, gt.iopts, gt.opts}
}

func (this *Tox) Kill() {
	log.Println("hoho")
	C.tox_kill(this.x)
	this.x = nil
}

// uint32_t tox_do_interval(Tox *tox);
func (this *Tox) DoInterval() (int32, error) {
	r := C.tox_do_interval(this.x)
	return int32(r), nil
}

/* The main loop that needs to be run in intervals of tox_do_interval() ms. */
// void tox_do(Tox *tox);
func (this *Tox) Do() {
	C.tox_do(this.x)
}

func (this *Tox) Size() (int32, error) {
	r := C.tox_size(this.x)
	return int32(r), nil
}

func (this *Tox) Save(data interface{}) error {

	return nil
}

func (this *Tox) Load(data interface{}, length int32) error {

	return nil
}

func (this *Tox) BootstrapFromAddress(addr string, port uint16, public_key string) (int32, error) {
	var _addr *C.char = C.CString(addr)
	defer C.free(unsafe.Pointer(_addr))
	var _port C.uint16_t = C.uint16_t(port)
	var _cpubkey *C.char = C.CString(public_key)
	defer C.free(unsafe.Pointer(_cpubkey))
	
	r := C.tox_bootstrap_from_address(this.x, _addr, _port, char2uint8(_cpubkey))
	return int32(r), nil
}

func (this *Tox) GetAddress(addr interface{}) (error) {
	fmt.Println(this.i)
	return nil
}

// int32_t tox_add_friend(Tox *tox, const uint8_t *address, const uint8_t *data, uint16_t length);
func (this *Tox) AddFriend(addr interface{}, data interface{}, length int32) (int32, error) {

    return 1, nil
}

func (this *Tox) AddFriendNoRequest(public_key string) (int32, error) {
	var _pubkey = C.CString(public_key)
	defer C.free(unsafe.Pointer(_pubkey))
	
	r := C.tox_add_friend_norequest(this.x, char2uint8(_pubkey));
	return int32(r), nil
}

func (this *Tox) GetFriendNumber(public_key string) (int32, error) {
	var _pubkey = C.CString(public_key)
	defer C.free(unsafe.Pointer(_pubkey))

	r := C.tox_get_friend_number(this.x, char2uint8(_pubkey))
	return int32(r), nil
}

func (this *Tox) GetClientId(friendNumber int32, public_key string) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _pubkey = C.CString(public_key)
	defer C.free(unsafe.Pointer(_pubkey))

	r := C.tox_get_client_id(this.x, _fn, char2uint8(_pubkey))
	return int(r), nil
}

func (this *Tox) DelFriend(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_del_friend(this.x, _fn)
	return int(r), nil
}

func (this *Tox) GetFriendConnectionStatus(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_get_friend_connection_status(this.x, _fn)
	return int(r), nil
}

func (this *Tox) FriendExists(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_friend_exists(this.x, _fn)
	return int(r), nil
}

func (this *Tox) SendMesage(friendNumber int32, message string, length uint32) (int32, error) {
	var _fn = C.int32_t(friendNumber)
	var _message = C.CString(message)
	defer C.free(unsafe.Pointer(_message))
	var _length = C.uint32_t(length)

	r := C.tox_send_message(this.x, _fn, char2uint8(_message), _length)
	return int32(r), nil
}

func (this *Tox) SendAction(friendNumber int32, action string, length uint32) (int32, error) {
	var _fn = C.int32_t(friendNumber)
	var _action = C.CString(action)
	defer C.free(unsafe.Pointer(_action))
	var _length = C.uint32_t(length)

	r := C.tox_send_message(this.x, _fn, char2uint8(_action), _length)
	return int32(r), nil
}

func (this *Tox) SetName(name string, length uint16) (int, error) {
	var _name = C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	var _length = C.uint16_t(length)

	r := C.tox_set_name(this.x, char2uint8(_name), _length)
	return int(r), nil
}

func (this *Tox) GetSelfName(name string) (int, error) {
	var _name = C.CString(name)
	defer C.free(unsafe.Pointer(_name))

	r := C.tox_get_self_name(this.x, char2uint8(_name))
	return int(r), nil
}

func (this *Tox) GetName(friendNumber int32, name string) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _name = C.CString(name)
	defer C.free(unsafe.Pointer(_name))

	r := C.tox_get_name(this.x, _fn, char2uint8(_name))
	return int(r), nil
}

func (this *Tox) GetNameSize(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_get_name_size(this.x, _fn)
	return int(r), nil
}

func (this *Tox) GetSelfNameSize() (int, error) {
	r := C.tox_get_self_name_size(this.x)
	return int(r), nil
}

func (this *Tox) SetStatusMessage(status string, length uint16) (int, error) {
	var _status = C.CString(status)
	defer C.free(unsafe.Pointer(_status))
	var _length = C.uint16_t(length)

	r := C.tox_set_status_message(this.x, char2uint8(_status), _length)
	return int(r), nil
}

func (this *Tox) SetUserStatus(status uint8) (int, error) {
	var _status = C.uint8_t(status)

	r := C.tox_set_user_status(this.x, _status)
	return int(r), nil
}

func (this *Tox) GetStatusMessageSize(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_get_status_message_size(this.x, _fn)
	return int(r), nil
}

func (this *Tox) GetSelfStatusMessageSize() (int, error) {
	r := C.tox_get_self_status_message_size(this.x)
	return int(r), nil
}

func (this *Tox) GetStatusMessage(friendNumber int32, buf string, maxlen uint32) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _buf = C.CString(buf)
	defer C.free(unsafe.Pointer(_buf))
	var _maxlen = C.uint32_t(maxlen)

	r := C.tox_get_status_message(this.x, _fn, char2uint8(_buf), _maxlen)
	buf = C.GoString(_buf)
	return int(r), nil
}

func (this *Tox) GetSelfStatusMessage(friendNumber int32, buf string, maxlen uint32) (int, error) {
	var _buf = C.CString(buf)
	defer C.free(unsafe.Pointer(_buf))
	var _maxlen = C.uint32_t(maxlen)

	r := C.tox_get_self_status_message(this.x, char2uint8(_buf), _maxlen)
	buf = C.GoString(_buf)
	return int(r), nil
}

func (this *Tox) GetUserStatus(friendNumber int32) (uint8, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_get_user_status(this.x, _fn)
	return uint8(r), nil
}

func (this *Tox) GetSelfUserStatus() (uint8, error) {
	r := C.tox_get_self_user_status(this.x)
	return uint8(r), nil
}

func (this *Tox) GetLastOnline(friendNumber int32) (uint64, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_get_last_online(this.x, _fn)
	return uint64(r), nil
}

func (this *Tox) SetUserIsTyping(friendNumber int32, is_typing uint8) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _is_typing = C.uint8_t(is_typing)

	r := C.tox_set_user_is_typing(this.x, _fn, _is_typing)
	return int(r), nil
}

func (this *Tox) GetIsTyping(friendNumber int32) (uint8, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_get_is_typing(this.x, _fn)
	return uint8(r), nil
}

func (this *Tox) CountFriendList() (uint32, error) {
	r := C.tox_count_friendlist(this.x)
	return uint32(r), nil
}

func (this *Tox) GetNumOnlineFriends() (uint32, error) {
	r := C.tox_get_num_online_friends(this.x)
	return uint32(r), nil
}

// tox_callback_***


func (this *Tox) GetNospam() (uint32, error) {
	r := C.tox_get_nospam(this.x)
	return uint32(r), nil
}

func (this *Tox) SetNospam(nospam uint32) {
	var _nospam = C.uint32_t(nospam)

	C.tox_set_nospam(this.x, _nospam)
}

func (this *Tox) GetKeys(public_key string, secret_key string) {
	var _pubkey = C.CString(public_key)
	defer C.free(unsafe.Pointer(_pubkey))
	var _seckey = C.CString(secret_key)
	defer C.free(unsafe.Pointer(_seckey))

	C.tox_get_keys(this.x, char2uint8(_pubkey), char2uint8(_seckey))
	public_key = C.GoString(_pubkey)
	secret_key = C.GoString(_seckey)
}

// tox_lossy_***

func (this *Tox) SendLossyPacket(friendNumber int32, data string, length uint32) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.uint32_t(length)

	r := C.tox_send_lossy_packet(this.x, _fn, char2uint8(_data), _length)
	return int(r), nil
}

func (this *Tox) SendLossLessPacket(friendNumber int32, data string, length uint32) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.uint32_t(length)

	r := C.tox_send_lossless_packet(this.x, _fn, char2uint8(_data), _length)
	return int(r), nil
}

// tox_callback_group_***

func (this *Tox) AddGroupChat() (int, error) {
	r := C.tox_add_groupchat(this.x)
	return int(r), nil
}

func (this *Tox) DelGroupChat(groupNumber int) (int, error) {
	var _gn = C.int(groupNumber)

	r := C.tox_del_groupchat(this.x, _gn)
	return int(r), nil
}

func (this *Tox) GroupPeerName(groupNumber int, peerNumber int, name string) (int, error) {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)
	var _name = C.CString(name)
	defer C.free(unsafe.Pointer(_name))

	r := C.tox_group_peername(this.x, _gn, _pn, char2uint8(_name))
	name = C.GoString(_name)
	return int(r), nil
}

func (this *Tox) GroupPeerPubkey(groupNumber int, peerNumber int, public_key string) (int, error) {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)
	var _pubkey = C.CString(public_key)
	defer C.free(unsafe.Pointer(_pubkey))

	r := C.tox_group_peer_pubkey(this.x, _gn, _pn, char2uint8(_pubkey))
	public_key = C.GoString(_pubkey)
	return int(r), nil
}

func (this *Tox) InviteFriend(friendNumber int32, groupNumber int) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _gn = C.int(groupNumber)

	r := C.tox_invite_friend(this.x, _fn, _gn)
	return int(r), nil
}

func (this *Tox) JoinGroupChat(friendNumber int32, data string, length uint16) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.uint16_t(length)

	r := C.tox_join_groupchat(this.x, _fn, char2uint8(_data), _length)
	return int(r), nil
}

func (this *Tox) GroupActionSend(groupNumber int, action string, length uint16) (int, error) {
	var _gn = C.int(groupNumber)
	var _action = C.CString(action)
	defer C.free(unsafe.Pointer(_action))
	var _length = C.uint16_t(length)

	r := C.tox_group_action_send(this.x, _gn, char2uint8(_action), _length)
	return int(r), nil
}

func (this *Tox) GroupSetTitle(groupNumber int, title string, length uint8) (int, error) {
	var _gn = C.int(groupNumber)
	var _title = C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	var _length = C.uint8_t(length)

	r := C.tox_group_set_title(this.x, _gn, char2uint8(_title), _length)
	return int(r), nil
}

func (this *Tox) GroupGetTitle(groupNumber int, title string, maxlen uint32) (int, error) {
	var _gn = C.int(groupNumber)
	var _title = C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	var _maxlen = C.uint32_t(maxlen)

	r := C.tox_group_get_title(this.x, _gn, char2uint8(_title), _maxlen)
	title = C.GoString(_title)
	return int(r), nil
}

func (this *Tox) GroupPeerNumberIsOurs(groupNumber int, peerNumber int) (uint, error) {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)

	r := C.tox_group_peernumber_is_ours(this.x, _gn, _pn)
	return uint(r), nil
}

func (this *Tox) GroupNumberPeers(groupNumber int) (int, error) {
	var _gn = C.int(groupNumber)

	r := C.tox_group_number_peers(this.x, _gn)
	return int(r), nil
}

/*
int tox_group_get_names(const Tox *tox, int groupnumber, uint8_t names[][TOX_MAX_NAME_LENGTH],
	uint16_t lengths[],
	uint16_t length);
*/

func (this *Tox) CountChatList() (uint32, error) {
	r := C.tox_count_chatlist(this.x)
	return uint32(r), nil
}

// TODO...
func (this *Tox) GetChatList(outList []int32, listSize uint32) (uint32, error) {
	return uint32(0), nil
}

func (this *Tox) GroupGetType(groupNumber int) (int, error) {
	var _gn = C.int(groupNumber)

	r := C.tox_group_get_type(this.x, _gn)
	return int(r), nil
}

// tox_callback_avatar_**


func (this *Tox) SetAvatar(format uint8, data string, length uint32) (int, error) {
	var _format = C.uint8_t(format)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.uint32_t(length)

	r := C.tox_set_avatar(this.x, _format, char2uint8(_data), _length)
	return int(r), nil
}


func (this *Tox) UnsetAvatar() (int, error) {
	r := C.tox_unset_avatar(this.x)
	return int(r), nil
}

// TODO...
func (this *Tox) GetSelfAvatar(format []uint8, buf string,
	length []uint32, maxlen uint32, hash string) (int, error) {

	return int(0), nil
}

func (this *Tox) Hash(hash string, data string, datalen uint32) (int, error) {
	var _hash = C.CString(hash)
	defer C.free(unsafe.Pointer(_hash))
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _datalen = C.uint32_t(datalen)

	r := C.tox_hash(char2uint8(_hash), char2uint8(_data), _datalen)
	hash = C.GoString(_hash)
	return int(r), nil
}


func (this *Tox) RequestAvatarInfo(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_request_avatar_info(this.x, _fn)
	return int(r), nil
}

func (this *Tox) SendAvatarInfo(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_send_avatar_info(this.x, _fn)
	return int(r), nil
}

func (this *Tox) RequestAvatarData(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_request_avatar_data(this.x, _fn)
	return int(r), nil
}

// tox_callback_file_***

func (this *Tox) NewFileSender(friendNumber int32, fileSize uint64, fileName string, fnlen uint16) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _fileSize = C.uint64_t(fileSize)
	var _fileName = C.CString(fileName)
	defer C.free(unsafe.Pointer(_fileName))
	var _fnlen = C.uint16_t(fnlen)

	r := C.tox_new_file_sender(this.x, _fn, _fileSize, char2uint8(_fileName), _fnlen)
	return int(r), nil
}

func (this *Tox) FileSendControl(friendNumber int32, sendReceive uint8, fileNumber uint8,
	messageId uint8, data string, length uint16) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _sendReceive = C.uint8_t(sendReceive)
	var _fileNumber = C.uint8_t(fileNumber)
	var _messageId = C.uint8_t(messageId)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.uint16_t(length)

	r := C.tox_file_send_control(this.x, _fn, _sendReceive, _fileNumber,
		_messageId, char2uint8(_data), _length)
	return int(r), nil
}

func (this *Tox) FileSendData(friendNumber int32, fileNumber uint8, data string, length uint16) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _fileNumber = C.uint8_t(fileNumber)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.uint16_t(length)

	r := C.tox_file_send_data(this.x, _fn, _fileNumber, char2uint8(_data), _length)
	return int(r), nil
}

func (this *Tox) FileDataSize(friendNumber int32) (int, error) {
	var _fn = C.int32_t(friendNumber)

	r := C.tox_file_data_size(this.x, _fn)
	return int(r), nil
}

func (this *Tox) FileDataRemaining(friendNumber int32, fileNumber uint8, sendReceive uint8) (uint64, error) {
	var _fn = C.int32_t(friendNumber)
	var _fileNumber = C.uint8_t(fileNumber)
	var _sendReceive = C.uint8_t(sendReceive)

	r := C.tox_file_data_remaining(this.x, _fn, _fileNumber, _sendReceive)
	return uint64(r), nil
}

// boostrap, see upper

func (this *Tox) AddTcpRelay(addr string, port uint16, public_key string) (int, error) {
	var _addr = C.CString(addr)
	defer C.free(unsafe.Pointer(_addr))
	var _port = C.uint16_t(port)
	var _pubkey = C.CString(public_key)
	defer C.free(unsafe.Pointer(_pubkey))

	r := C.tox_add_tcp_relay(this.x, _addr, _port, char2uint8(_pubkey))
	return int(r), nil
}

func (this *Tox) IsConnected() (int, error) {
	r := C.tox_isconnected(this.x);
	return int(r), nil
}



////////////
/*
原则说明：
所有需要public_key的地方，在go空间内是实际串的16进制字符串表示。

*/

////////////////////
func KeepPkg() {
}

func _dirty_init() {
	fmt.Println("ddddddddd")
}

