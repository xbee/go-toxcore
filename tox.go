package tox

/*
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>

//////

void callbackFriendRequestWrapperForC(Tox *, uint8_t *, uint8_t *, uint16_t, void*);
typedef void (*cb_friend_request_ftype)(Tox *, const uint8_t *, const uint8_t *, size_t, void*);
static void cb_friend_request_wrapper_for_go(Tox *m, cb_friend_request_ftype fn, void *userdata)
{ tox_callback_friend_request(m, fn); }

void callbackFriendMessageWrapperForC(Tox *, uint32_t, int, uint8_t*, uint32_t, void*);
typedef void (*cb_friend_message_ftype)(Tox *, uint32_t, TOX_MESSAGE_TYPE, const uint8_t*, size_t, void*);
static void cb_friend_message_wrapper_for_go(Tox *m, cb_friend_message_ftype fn, void *userdata)
{ tox_callback_friend_message(m, fn); }

void callbackFriendNameWrapperForC(Tox *, uint32_t, uint8_t*, uint32_t, void*);
typedef void (*cb_friend_name_ftype)(Tox *, uint32_t, const uint8_t*, size_t, void*);
static void cb_friend_name_wrapper_for_go(Tox *m, cb_friend_name_ftype fn, void *userdata)
{ tox_callback_friend_name(m, fn); }

void callbackFriendStatusMessageWrapperForC(Tox *, uint32_t, uint8_t*, uint32_t, void*);
typedef void (*cb_friend_status_message_ftype)(Tox *, uint32_t, const uint8_t*, size_t, void*);
static void cb_friend_status_message_wrapper_for_go(Tox *m, cb_friend_status_message_ftype fn, void *userdata)
{ tox_callback_friend_status_message(m, fn); }

void callbackFriendStatusWrapperForC(Tox *, uint32_t, int, void*);
typedef void (*cb_friend_status_ftype)(Tox *, uint32_t, TOX_USER_STATUS, void*);
static void cb_friend_status_wrapper_for_go(Tox *m, cb_friend_status_ftype fn, void *userdata)
{ tox_callback_friend_status(m, fn); }

void callbackFriendConnectionStatusWrapperForC(Tox *, uint32_t, int, void*);
typedef void (*cb_friend_connection_status_ftype)(Tox *, uint32_t, uint32_t, void*);
static void cb_friend_connection_status_wrapper_for_go(Tox *m, cb_friend_connection_status_ftype fn, void *userdata)
{ tox_callback_friend_connection_status(m, fn); }

void callbackFriendTypingWrapperForC(Tox *, uint32_t, uint8_t, void*);
typedef void (*cb_friend_typing_ftype)(Tox *, uint32_t, bool, void*);
static void cb_friend_typing_wrapper_for_go(Tox *m, cb_friend_typing_ftype fn, void *userdata)
{ tox_callback_friend_typing(m, fn); }

void callbackFriendReadReceiptWrapperForC(Tox *, uint32_t, uint32_t, void*);
typedef void (*cb_friend_read_receipt_ftype)(Tox *, uint32_t, uint32_t, void*);
static void cb_friend_read_receipt_wrapper_for_go(Tox *m, cb_friend_read_receipt_ftype fn, void *userdata)
{ tox_callback_friend_read_receipt(m, fn); }

void callbackFriendLossyPacketWrapperForC(Tox *, uint32_t, uint8_t*, size_t, void*);
typedef void (*cb_friend_lossy_packet_ftype)(Tox *, uint32_t, const uint8_t*, size_t, void*);
static void cb_friend_lossy_packet_wrapper_for_go(Tox *m, cb_friend_lossy_packet_ftype fn, void *userdata)
{ tox_callback_friend_lossy_packet(m, fn); }

void callbackFriendLosslessPacketWrapperForC(Tox *, uint32_t, uint8_t*, size_t, void*);
typedef void (*cb_friend_lossless_packet_ftype)(Tox *, uint32_t, const uint8_t*, size_t, void*);
static void cb_friend_lossless_packet_wrapper_for_go(Tox *m, cb_friend_lossless_packet_ftype fn, void *userdata)
{ tox_callback_friend_lossless_packet(m, fn); }

void callbackSelfConnectionStatusWrapperForC(Tox *, int, void*);
typedef void (*cb_self_connection_status_ftype)(Tox *, TOX_CONNECTION, void*);
static void cb_self_connection_status_wrapper_for_go(Tox *m, cb_self_connection_status_ftype fn, void *userdata)
{ tox_callback_self_connection_status(m, fn); }

void callbackFileRecvControlWrapperForC(Tox *tox, uint32_t friend_number, uint32_t file_number,
                                      TOX_FILE_CONTROL control, void *user_data);
typedef void (*cb_file_recv_control_ftype)(Tox *tox, uint32_t friend_number, uint32_t file_number,
                                      TOX_FILE_CONTROL control, void *useer_data);
static void cb_file_recv_control_wrapper_for_go(Tox *m, cb_file_recv_control_ftype fn, void *userdata)
{ tox_callback_file_recv_control(m, fn); }

void callbackFileRecvWrapperForC(Tox *tox, uint32_t friend_number, uint32_t file_number, uint32_t kind,
                               uint64_t file_size, uint8_t *filename, size_t filename_length, void *user_data);
typedef void (*cb_file_recv_ftype)(Tox *tox, uint32_t friend_number, uint32_t file_number, uint32_t kind,
              uint64_t file_size, const uint8_t *filename, size_t filename_length, void *user_data);
static void cb_file_recv_wrapper_for_go(Tox *m, cb_file_recv_ftype fn, void *userdata)
{ tox_callback_file_recv(m, fn); }

void callbackFileRecvChunkWrapperForC(Tox *tox, uint32_t friend_number, uint32_t file_number, uint64_t position,
                                    uint8_t *data, size_t length, void *user_data);
typedef void (*cb_file_recv_chunk_ftype)(Tox *tox, uint32_t friend_number, uint32_t file_number, uint64_t position,
                                    const uint8_t *data, size_t length, void *user_data);
static void cb_file_recv_chunk_wrapper_for_go(Tox *m, cb_file_recv_chunk_ftype fn, void *userdata)
{ tox_callback_file_recv_chunk(m, fn); }

void callbackFileChunkRequestWrapperForC(Tox *tox, uint32_t friend_number, uint32_t file_number, uint64_t position,
                                       size_t length, void *user_data);
typedef void (*cb_file_chunk_request_ftype)(Tox *tox, uint32_t friend_number, uint32_t file_number, uint64_t position,
                                       size_t length, void *user_data);
static void cb_file_chunk_request_wrapper_for_go(Tox *m, cb_file_chunk_request_ftype fn, void *userdata)
{ tox_callback_file_chunk_request(m, fn); }

// fix nouse compile warning
static inline void fixnousetox() {
    cb_friend_request_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_message_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_name_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_status_message_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_status_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_connection_status_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_typing_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_lossy_packet_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_lossless_packet_wrapper_for_go(NULL, NULL, NULL);
    cb_friend_read_receipt_wrapper_for_go(NULL, NULL, NULL);
    cb_self_connection_status_wrapper_for_go(NULL, NULL, NULL);
    cb_file_recv_control_wrapper_for_go(NULL, NULL, NULL);
    cb_file_recv_wrapper_for_go(NULL, NULL, NULL);
    cb_file_recv_chunk_wrapper_for_go(NULL, NULL, NULL);
    cb_file_chunk_request_wrapper_for_go(NULL, NULL, NULL);
}

*/
import "C"
import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	// "sync"
	"unsafe"

	deadlock "github.com/sasha-s/go-deadlock"
)

// "reflect"
// "runtime"

//////////
// friend callback type
type cb_friend_request_ftype func(this *Tox, pubkey string, message string, userData interface{})
type cb_friend_message_ftype func(this *Tox, friendNumber uint32, message string, userData interface{})
type cb_friend_name_ftype func(this *Tox, friendNumber uint32, newName string, userData interface{})
type cb_friend_status_message_ftype func(this *Tox, friendNumber uint32, newStatus string, userData interface{})
type cb_friend_status_ftype func(this *Tox, friendNumber uint32, status int, userData interface{})
type cb_friend_connection_status_ftype func(this *Tox, friendNumber uint32, status int, userData interface{})
type cb_friend_typing_ftype func(this *Tox, friendNumber uint32, isTyping uint8, userData interface{})
type cb_friend_read_receipt_ftype func(this *Tox, friendNumber uint32, receipt uint32, userData interface{})
type cb_friend_lossy_packet_ftype func(this *Tox, friendNumber uint32, data string, userData interface{})
type cb_friend_lossless_packet_ftype func(this *Tox, friendNumber uint32, data string, userData interface{})

// self callback type
type cb_self_connection_status_ftype func(this *Tox, status int, userData interface{})

// file callback type
type cb_file_recv_control_ftype func(this *Tox, friendNumber uint32, fileNumber uint32,
	control int, userData interface{})
type cb_file_recv_ftype func(this *Tox, friendNumber uint32, fileNumber uint32, kind uint32, fileSize uint64,
	fileName string, userData interface{})
type cb_file_recv_chunk_ftype func(this *Tox, friendNumber uint32, fileNumber uint32, position uint64,
	data []byte, userData interface{})
type cb_file_chunk_request_ftype func(this *Tox, friend_number uint32, file_number uint32, position uint64,
	length int, user_data interface{})

type Tox struct {
	opts       *ToxOptions
	toxopts    *C.struct_Tox_Options
	toxcore    *C.Tox // save C.Tox
	threadSafe bool
	mu         deadlock.RWMutex
	// mu sync.RWMutex

	// some callbacks, should be private
	cb_friend_requests           map[unsafe.Pointer]interface{}
	cb_friend_messages           map[unsafe.Pointer]interface{}
	cb_friend_names              map[unsafe.Pointer]interface{}
	cb_friend_status_messages    map[unsafe.Pointer]interface{}
	cb_friend_statuss            map[unsafe.Pointer]interface{}
	cb_friend_connection_statuss map[unsafe.Pointer]interface{}
	cb_friend_typings            map[unsafe.Pointer]interface{}
	cb_friend_read_receipts      map[unsafe.Pointer]interface{}
	cb_friend_lossy_packets      map[unsafe.Pointer]interface{}
	cb_friend_lossless_packets   map[unsafe.Pointer]interface{}
	cb_self_connection_statuss   map[unsafe.Pointer]interface{}

	cb_conference_invites          map[unsafe.Pointer]interface{}
	cb_conference_messages         map[unsafe.Pointer]interface{}
	cb_conference_actions          map[unsafe.Pointer]interface{}
	cb_conference_titles           map[unsafe.Pointer]interface{}
	cb_conference_namelist_changes map[unsafe.Pointer]interface{}

	cb_file_recv_controls  map[unsafe.Pointer]interface{}
	cb_file_recvs          map[unsafe.Pointer]interface{}
	cb_file_recv_chunks    map[unsafe.Pointer]interface{}
	cb_file_chunk_requests map[unsafe.Pointer]interface{}

	cb_iterate_data              interface{}
	cb_conference_message_setted bool
}

var cbUserDatas = newUserData()

//export callbackFriendRequestWrapperForC
func callbackFriendRequestWrapperForC(m *C.Tox, a0 *C.uint8_t, a1 *C.uint8_t, a2 C.uint16_t, a3 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_requests {
		pubkey_b := C.GoBytes(unsafe.Pointer(a0), C.int(PUBLIC_KEY_SIZE))
		pubkey := hex.EncodeToString(pubkey_b)
		pubkey = strings.ToUpper(pubkey)
		message_b := C.GoBytes(unsafe.Pointer(a1), C.int(a2))
		message := string(message_b)
		cbfn := *(*cb_friend_request_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, pubkey, message, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendRequest(cbfn cb_friend_request_ftype, userData interface{}) {
	this.CallbackFriendRequestAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendRequestAdd(cbfn cb_friend_request_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_requests[cbfnp]; ok {
		return
	}
	this.cb_friend_requests[cbfnp] = userData

	var _cbfn = (C.cb_friend_request_ftype)(C.callbackFriendRequestWrapperForC)
	var _userData = unsafe.Pointer(this)
	_userData = nil

	C.cb_friend_request_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendMessageWrapperForC
func callbackFriendMessageWrapperForC(m *C.Tox, a0 C.uint32_t, mtype C.int,
	a1 *C.uint8_t, a2 C.uint32_t, a3 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_messages {
		message_ := C.GoStringN((*C.char)(unsafe.Pointer(a1)), (C.int)(a2))
		cbfn := *(*cb_friend_message_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(a0), message_, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendMessage(cbfn cb_friend_message_ftype, userData interface{}) {
	this.CallbackFriendMessageAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendMessageAdd(cbfn cb_friend_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_messages[cbfnp]; ok {
		return
	}
	this.cb_friend_messages[cbfnp] = userData

	var _cbfn = (C.cb_friend_message_ftype)(C.callbackFriendMessageWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendNameWrapperForC
func callbackFriendNameWrapperForC(m *C.Tox, a0 C.uint32_t, a1 *C.uint8_t, a2 C.uint32_t, a3 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_names {
		name := C.GoStringN((*C.char)((unsafe.Pointer)(a1)), C.int(a2))
		cbfn := *(*cb_friend_name_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(a0), name, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendName(cbfn cb_friend_name_ftype, userData interface{}) {
	this.CallbackFriendNameAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendNameAdd(cbfn cb_friend_name_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_names[cbfnp]; ok {
		return
	}
	this.cb_friend_names[cbfnp] = userData

	var _cbfn = (C.cb_friend_name_ftype)(C.callbackFriendNameWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_name_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendStatusMessageWrapperForC
func callbackFriendStatusMessageWrapperForC(m *C.Tox, a0 C.uint32_t, a1 *C.uint8_t, a2 C.uint32_t, a3 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_status_messages {
		statusText := C.GoStringN((*C.char)(unsafe.Pointer(a1)), C.int(a2))
		cbfn := *(*cb_friend_status_message_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(a0), statusText, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendStatusMessage(cbfn cb_friend_status_message_ftype, userData interface{}) {
	this.CallbackFriendStatusMessageAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendStatusMessageAdd(cbfn cb_friend_status_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_status_messages[cbfnp]; ok {
		return
	}
	this.cb_friend_status_messages[cbfnp] = userData

	var _cbfn = (C.cb_friend_status_message_ftype)(C.callbackFriendStatusMessageWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_status_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendStatusWrapperForC
func callbackFriendStatusWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.int, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_statuss {
		cbfn := *(*cb_friend_status_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(a0), int(a1), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendStatus(cbfn cb_friend_status_ftype, userData interface{}) {
	this.CallbackFriendStatusAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendStatusAdd(cbfn cb_friend_status_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_statuss[cbfnp]; ok {
		return
	}
	this.cb_friend_statuss[cbfnp] = userData

	var _cbfn = (C.cb_friend_status_ftype)(C.callbackFriendStatusWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_status_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendConnectionStatusWrapperForC
func callbackFriendConnectionStatusWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.int, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_connection_statuss {
		cbfn := *(*cb_friend_connection_status_ftype)((unsafe.Pointer)(cbfni))
		this.beforeCallback()
		cbfn(this, uint32(a0), int(a1), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendConnectionStatus(cbfn cb_friend_connection_status_ftype, userData interface{}) {
	this.CallbackFriendConnectionStatusAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendConnectionStatusAdd(cbfn cb_friend_connection_status_ftype, userData interface{}) {
	cbfnp := unsafe.Pointer(&cbfn)
	if _, ok := this.cb_friend_connection_statuss[cbfnp]; ok {
		return
	}
	this.cb_friend_connection_statuss[cbfnp] = userData

	var _cbfn = (C.cb_friend_connection_status_ftype)(C.callbackFriendConnectionStatusWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_connection_status_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendTypingWrapperForC
func callbackFriendTypingWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint8_t, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_typings {
		cbfn := *(*cb_friend_typing_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(a0), uint8(a1), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendTyping(cbfn cb_friend_typing_ftype, userData interface{}) {
	this.CallbackFriendTypingAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendTypingAdd(cbfn cb_friend_typing_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_typings[cbfnp]; ok {
		return
	}
	this.cb_friend_typings[cbfnp] = userData

	var _cbfn = (C.cb_friend_typing_ftype)(C.callbackFriendTypingWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_typing_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendReadReceiptWrapperForC
func callbackFriendReadReceiptWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_read_receipts {
		cbfn := *(*cb_friend_read_receipt_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(a0), uint32(a1), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendReadReceipt(cbfn cb_friend_read_receipt_ftype, userData interface{}) {
	this.CallbackFriendReadReceiptAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendReadReceiptAdd(cbfn cb_friend_read_receipt_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_read_receipts[cbfnp]; ok {
		return
	}
	this.cb_friend_read_receipts[cbfnp] = userData

	var _cbfn = (C.cb_friend_read_receipt_ftype)(C.callbackFriendReadReceiptWrapperForC)
	var _userData unsafe.Pointer

	C.cb_friend_read_receipt_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendLossyPacketWrapperForC
func callbackFriendLossyPacketWrapperForC(m *C.Tox, a0 C.uint32_t, a1 *C.uint8_t, len C.size_t, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_lossy_packets {
		cbfn := *(*cb_friend_lossy_packet_ftype)(cbfni)
		msg := C.GoStringN((*C.char)(unsafe.Pointer(a1)), C.int(len))
		this.beforeCallback()
		cbfn(this, uint32(a0), msg, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendLossyPacket(cbfn cb_friend_lossy_packet_ftype, userData interface{}) {
	this.CallbackFriendLossyPacketAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendLossyPacketAdd(cbfn cb_friend_lossy_packet_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_lossy_packets[cbfnp]; ok {
		return
	}
	this.cb_friend_lossy_packets[cbfnp] = userData

	var _cbfn = (C.cb_friend_lossy_packet_ftype)(C.callbackFriendLossyPacketWrapperForC)
	var _userData unsafe.Pointer

	C.cb_friend_lossy_packet_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendLosslessPacketWrapperForC
func callbackFriendLosslessPacketWrapperForC(m *C.Tox, a0 C.uint32_t, a1 *C.uint8_t, len C.size_t, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_friend_lossless_packets {
		cbfn := *(*cb_friend_lossless_packet_ftype)(cbfni)
		msg := C.GoStringN((*C.char)(unsafe.Pointer(a1)), C.int(len))
		this.beforeCallback()
		cbfn(this, uint32(a0), msg, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFriendLosslessPacket(cbfn cb_friend_lossless_packet_ftype, userData interface{}) {
	this.CallbackFriendLosslessPacketAdd(cbfn, userData)
}
func (this *Tox) CallbackFriendLosslessPacketAdd(cbfn cb_friend_lossless_packet_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_friend_lossless_packets[cbfnp]; ok {
		return
	}
	this.cb_friend_lossless_packets[cbfnp] = userData

	var _cbfn = (C.cb_friend_lossless_packet_ftype)(C.callbackFriendLosslessPacketWrapperForC)
	var _userData unsafe.Pointer

	C.cb_friend_lossless_packet_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackSelfConnectionStatusWrapperForC
func callbackSelfConnectionStatusWrapperForC(m *C.Tox, status C.int, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_self_connection_statuss {
		cbfn := *(*cb_self_connection_status_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, int(status), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackSelfConnectionStatus(cbfn cb_self_connection_status_ftype, userData interface{}) {
	this.CallbackSelfConnectionStatusAdd(cbfn, userData)
}
func (this *Tox) CallbackSelfConnectionStatusAdd(cbfn cb_self_connection_status_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_self_connection_statuss[cbfnp]; ok {
		return
	}
	this.cb_self_connection_statuss[cbfnp] = userData

	var _cbfn = (C.cb_self_connection_status_ftype)(C.callbackSelfConnectionStatusWrapperForC)
	// var _userData = unsafe.Pointer(this)

	C.cb_self_connection_status_wrapper_for_go(this.toxcore, _cbfn, nil)
}

// 包内部函数
//export callbackFileRecvControlWrapperForC
func callbackFileRecvControlWrapperForC(m *C.Tox, friendNumber C.uint32_t, fileNumber C.uint32_t,
	control C.TOX_FILE_CONTROL, userData unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_file_recv_controls {
		cbfn := *(*cb_file_recv_control_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(friendNumber), uint32(fileNumber), int(control), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFileRecvControl(cbfn cb_file_recv_control_ftype, userData interface{}) {
	this.CallbackFileRecvControlAdd(cbfn, userData)
}
func (this *Tox) CallbackFileRecvControlAdd(cbfn cb_file_recv_control_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_file_recv_controls[cbfnp]; ok {
		return
	}
	this.cb_file_recv_controls[cbfnp] = userData
	var _cbfn = (C.cb_file_recv_control_ftype)(unsafe.Pointer(C.callbackFileRecvControlWrapperForC))
	var _userData unsafe.Pointer = nil

	C.cb_file_recv_control_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFileRecvWrapperForC
func callbackFileRecvWrapperForC(m *C.Tox, friendNumber C.uint32_t, fileNumber C.uint32_t, kind C.uint32_t,
	fileSize C.uint64_t, fileName *C.uint8_t, fileNameLength C.size_t, userData unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_file_recvs {
		cbfn := *(*cb_file_recv_ftype)(cbfni)
		fileName_ := C.GoStringN((*C.char)(unsafe.Pointer(fileName)), C.int(fileNameLength))
		this.beforeCallback()
		cbfn(this, uint32(friendNumber), uint32(fileNumber), uint32(kind),
			uint64(fileSize), fileName_, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFileRecv(cbfn cb_file_recv_ftype, userData interface{}) {
	this.CallbackFileRecvAdd(cbfn, userData)
}
func (this *Tox) CallbackFileRecvAdd(cbfn cb_file_recv_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_file_recvs[cbfnp]; ok {
		return
	}
	this.cb_file_recvs[cbfnp] = userData
	var _cbfn = (C.cb_file_recv_ftype)(unsafe.Pointer(C.callbackFileRecvWrapperForC))
	var _userData unsafe.Pointer = nil

	C.cb_file_recv_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFileRecvChunkWrapperForC
func callbackFileRecvChunkWrapperForC(m *C.Tox, friendNumber C.uint32_t, fileNumber C.uint32_t,
	position C.uint64_t, data *C.uint8_t, length C.size_t, userData unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_file_recv_chunks {
		cbfn := *(*cb_file_recv_chunk_ftype)(cbfni)
		data_ := C.GoBytes((unsafe.Pointer)(data), C.int(length))
		this.beforeCallback()
		cbfn(this, uint32(friendNumber), uint32(fileNumber), uint64(position), data_, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFileRecvChunk(cbfn cb_file_recv_chunk_ftype, userData interface{}) {
	this.CallbackFileRecvChunkAdd(cbfn, userData)
}
func (this *Tox) CallbackFileRecvChunkAdd(cbfn cb_file_recv_chunk_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_file_recv_chunks[cbfnp]; ok {
		return
	}
	this.cb_file_recv_chunks[cbfnp] = userData

	var _cbfn = (C.cb_file_recv_chunk_ftype)(unsafe.Pointer(C.callbackFileRecvChunkWrapperForC))
	var _userData unsafe.Pointer = nil

	C.cb_file_recv_chunk_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFileChunkRequestWrapperForC
func callbackFileChunkRequestWrapperForC(m *C.Tox, friendNumber C.uint32_t, fileNumber C.uint32_t,
	position C.uint64_t, length C.size_t, userData unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_file_chunk_requests {
		cbfn := *(*cb_file_chunk_request_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(friendNumber), uint32(fileNumber), uint64(position), int(length), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackFileChunkRequest(cbfn cb_file_chunk_request_ftype, userData interface{}) {
	this.CallbackFileChunkRequestAdd(cbfn, userData)
}
func (this *Tox) CallbackFileChunkRequestAdd(cbfn cb_file_chunk_request_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_file_chunk_requests[cbfnp]; ok {
		return
	}
	this.cb_file_chunk_requests[cbfnp] = userData

	var _cbfn = (C.cb_file_chunk_request_ftype)(unsafe.Pointer(C.callbackFileChunkRequestWrapperForC))
	var _userData unsafe.Pointer = nil

	C.cb_file_chunk_request_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

func NewTox(opt *ToxOptions) *Tox {
	var tox = new(Tox)
	if opt != nil {
		tox.opts = opt
	} else {
		tox.opts = NewToxOptions()
	}
	tox.toxopts = tox.opts.toCToxOptions()

	var cerr C.TOX_ERR_NEW
	var toxcore = C.tox_new(tox.toxopts, &cerr)
	tox.toxcore = toxcore
	if toxcore == nil {
		log.Println(toxerr(cerr))
		return nil
	}
	cbUserDatas.set(toxcore, tox)

	// set default name and status
	tox.SelfSetName("NickName")
	tox.SelfSetStatusMessage("Offline")

	//
	tox.cb_friend_requests = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_messages = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_names = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_status_messages = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_statuss = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_connection_statuss = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_typings = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_read_receipts = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_lossy_packets = make(map[unsafe.Pointer]interface{})
	tox.cb_friend_lossless_packets = make(map[unsafe.Pointer]interface{})
	tox.cb_self_connection_statuss = make(map[unsafe.Pointer]interface{})

	tox.cb_conference_invites = make(map[unsafe.Pointer]interface{})
	tox.cb_conference_messages = make(map[unsafe.Pointer]interface{})
	tox.cb_conference_actions = make(map[unsafe.Pointer]interface{})
	tox.cb_conference_titles = make(map[unsafe.Pointer]interface{})
	tox.cb_conference_namelist_changes = make(map[unsafe.Pointer]interface{})

	tox.cb_file_recv_controls = make(map[unsafe.Pointer]interface{})
	tox.cb_file_recvs = make(map[unsafe.Pointer]interface{})
	tox.cb_file_recv_chunks = make(map[unsafe.Pointer]interface{})
	tox.cb_file_chunk_requests = make(map[unsafe.Pointer]interface{})

	return tox
}

func (this *Tox) Kill() {
	this.lock()
	defer this.unlock()

	if this == nil || this.toxcore == nil {
		return
	}

	C.tox_kill(this.toxcore)
	cbUserDatas.del(this.toxcore)
	this.toxcore = nil
}

// uint32_t tox_iteration_interval(Tox *tox);
func (this *Tox) IterationInterval() int {
	this.lock()
	defer this.unlock()

	r := C.tox_iteration_interval(this.toxcore)
	return int(r)
}

/* The main loop that needs to be run in intervals of tox_iteration_interval() ms. */
// void tox_iterate(Tox *tox);
// compatable with legacy version
func (this *Tox) Iterate() {
	this.lock()
	defer this.unlock()

	C.tox_iterate(this.toxcore, nil)
}

// for toktok new method
func (this *Tox) Iterate2(userData interface{}) {
	this.lock()
	defer this.unlock()

	this.cb_iterate_data = userData
	C.tox_iterate(this.toxcore, nil)
	this.cb_iterate_data = nil
}

func (this *Tox) lock() {
	if this.opts.ThreadSafe {
		this.mu.Lock()
	}
}
func (this *Tox) unlock() {
	if this.opts.ThreadSafe {
		this.mu.Unlock()
	}
}

func (this *Tox) GetSavedataSize() int32 {
	r := C.tox_get_savedata_size(this.toxcore)
	return int32(r)
}

func (this *Tox) GetSavedata() []byte {
	r := C.tox_get_savedata_size(this.toxcore)
	var savedata = make([]byte, int(r))

	C.tox_get_savedata(this.toxcore, (*C.uint8_t)(&savedata[0]))
	return savedata
}

/*
 * @param pubkey hex 64B length
 */
func (this *Tox) Bootstrap(addr string, port uint16, pubkey string) (bool, error) {
	this.lock()
	defer this.unlock()

	b_pubkey, err := hex.DecodeString(pubkey)
	if err != nil {
		return false, toxerr("Invalid pubkey")
	}

	var _addr = []byte(addr)
	var _port = C.uint16_t(port)
	var _cpubkey = (*C.uint8_t)(&b_pubkey[0])

	var cerr C.TOX_ERR_BOOTSTRAP
	r := C.tox_bootstrap(this.toxcore, (*C.char)(unsafe.Pointer(&_addr[0])), _port, _cpubkey, &cerr)
	if cerr > 0 {
		return false, toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) SelfGetAddress() string {
	var addr [ADDRESS_SIZE]byte
	var caddr = (*C.uint8_t)(unsafe.Pointer(&addr[0]))
	C.tox_self_get_address(this.toxcore, caddr)

	return strings.ToUpper(hex.EncodeToString(addr[:]))
}

func (this *Tox) SelfGetConnectionStatus() int {
	r := C.tox_self_get_connection_status(this.toxcore)
	return int(r)
}

func (this *Tox) FriendAdd(friendId string, message string) (uint32, error) {
	this.lock()
	defer this.unlock()

	friendId_b, err := hex.DecodeString(friendId)
	friendId_p := (*C.uint8_t)(&friendId_b[0])
	if err != nil {
		log.Panic(err)
	}

	cmessage := []byte(message)

	var cerr C.TOX_ERR_FRIEND_ADD
	r := C.tox_friend_add(this.toxcore, friendId_p,
		(*C.uint8_t)(&cmessage[0]), C.size_t(len(message)), &cerr)
	if cerr > 0 {
		return uint32(r), toxerr(cerr)
	}
	return uint32(r), nil
}

func (this *Tox) FriendAddNorequest(friendId string) (uint32, error) {
	this.lock()
	defer this.unlock()

	friendId_b, err := hex.DecodeString(friendId)
	if err != nil {
		return 0, err
	}
	friendId_p := (*C.uint8_t)(&friendId_b[0])

	var cerr C.TOX_ERR_FRIEND_ADD
	r := C.tox_friend_add_norequest(this.toxcore, friendId_p, &cerr)
	if cerr > 0 {
		return uint32(r), toxerr(cerr)
	}
	return uint32(r), nil
}

func (this *Tox) FriendByPublicKey(pubkey string) (uint32, error) {
	pubkey_b, err := hex.DecodeString(pubkey)
	if err != nil {
		return 0, err
	}
	var pubkey_p = (*C.uint8_t)(&pubkey_b[0])

	var cerr C.TOX_ERR_FRIEND_BY_PUBLIC_KEY
	r := C.tox_friend_by_public_key(this.toxcore, pubkey_p, &cerr)
	if cerr != C.TOX_ERR_FRIEND_BY_PUBLIC_KEY_OK {
		return uint32(r), toxerr(cerr)
	}
	return uint32(r), nil
}

func (this *Tox) FriendGetPublicKey(friendNumber uint32) (string, error) {
	var _fn = C.uint32_t(friendNumber)
	var pubkey_b = make([]byte, PUBLIC_KEY_SIZE)
	var pubkey_p = (*C.uint8_t)(&pubkey_b[0])

	var cerr C.TOX_ERR_FRIEND_GET_PUBLIC_KEY
	r := C.tox_friend_get_public_key(this.toxcore, _fn, pubkey_p, &cerr)
	if cerr > 0 || bool(r) == false {
		return "", toxerr(cerr)
	}
	pubkey_h := hex.EncodeToString(pubkey_b)
	pubkey_h = strings.ToUpper(pubkey_h)
	return pubkey_h, nil
}

func (this *Tox) FriendDelete(friendNumber uint32) (bool, error) {
	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_DELETE
	r := C.tox_friend_delete(this.toxcore, _fn, &cerr)
	if cerr > 0 {
		return bool(r), toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) FriendGetConnectionStatus(friendNumber uint32) (int, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_connection_status(this.toxcore, _fn, &cerr)
	if cerr > 0 {
		return int(r), toxerr(cerr)
	}
	return int(r), nil
}

func (this *Tox) FriendExists(friendNumber uint32) bool {
	var _fn = C.uint32_t(friendNumber)

	r := C.tox_friend_exists(this.toxcore, _fn)
	return bool(r)
}

func (this *Tox) FriendSendMessage(friendNumber uint32, message string) (uint32, error) {
	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)
	var _message = []byte(message)
	var _length = C.size_t(len(message))

	var mtype C.TOX_MESSAGE_TYPE = C.TOX_MESSAGE_TYPE_NORMAL
	var cerr C.TOX_ERR_FRIEND_SEND_MESSAGE
	r := C.tox_friend_send_message(this.toxcore, _fn, mtype, (*C.uint8_t)(&_message[0]), _length, &cerr)
	if cerr != C.TOX_ERR_FRIEND_SEND_MESSAGE_OK {
		return uint32(r), toxerr(cerr)
	}
	return uint32(r), nil
}

func (this *Tox) FriendSendAction(friendNumber uint32, action string) (uint32, error) {
	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)
	var _action = []byte(action)
	var _length = C.size_t(len(action))

	var mtype C.TOX_MESSAGE_TYPE = C.TOX_MESSAGE_TYPE_ACTION
	var cerr C.TOX_ERR_FRIEND_SEND_MESSAGE
	r := C.tox_friend_send_message(this.toxcore, _fn, mtype, (*C.uint8_t)(&_action[0]), _length, &cerr)
	if cerr > 0 {
		return uint32(r), toxerr(cerr)
	}
	return uint32(r), nil
}

func (this *Tox) SelfSetName(name string) error {
	this.lock()
	defer this.unlock()

	var _name = []byte(name)
	var _length = C.size_t(len(name))

	var cerr C.TOX_ERR_SET_INFO
	C.tox_self_set_name(this.toxcore, (*C.uint8_t)(&_name[0]), _length, &cerr)
	if cerr > 0 {
		return toxerr(cerr)
	}
	return nil
}

func (this *Tox) SelfGetName() string {
	nlen := C.tox_self_get_name_size(this.toxcore)
	_name := make([]byte, nlen)

	C.tox_self_get_name(this.toxcore, (*C.uint8_t)(&_name[0]))
	return string(_name)
}

func (this *Tox) FriendGetName(friendNumber uint32) (string, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	nlen := C.tox_friend_get_name_size(this.toxcore, _fn, &cerr)
	_name := make([]byte, nlen)

	r := C.tox_friend_get_name(this.toxcore, _fn, (*C.uint8_t)(&_name[0]), &cerr)
	if !bool(r) {
		return "", toxerr(cerr)
	}
	return string(_name), nil
}

func (this *Tox) FriendGetNameSize(friendNumber uint32) (int, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_name_size(this.toxcore, _fn, &cerr)
	if cerr > 0 {
		return int(r), toxerr(cerr)
	}
	return int(r), nil
}

func (this *Tox) SelfGetNameSize() int {
	r := C.tox_self_get_name_size(this.toxcore)
	return int(r)
}

func (this *Tox) SelfSetStatusMessage(status string) (bool, error) {
	this.lock()
	defer this.unlock()

	var _status = []byte(status)
	var _length = C.size_t(len(status))

	var cerr C.TOX_ERR_SET_INFO
	r := C.tox_self_set_status_message(this.toxcore, (*C.uint8_t)(&_status[0]), _length, &cerr)
	if cerr > 0 {
		return false, toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) SelfSetStatus(status uint8) {
	var _status = C.TOX_USER_STATUS(status)
	C.tox_self_set_status(this.toxcore, _status)
}

func (this *Tox) FriendGetStatusMessageSize(friendNumber uint32) (int, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_status_message_size(this.toxcore, _fn, &cerr)
	if cerr > 0 {
		return int(r), toxerr(cerr)
	}
	return int(r), nil
}

func (this *Tox) SelfGetStatusMessageSize() int {
	r := C.tox_self_get_status_message_size(this.toxcore)
	return int(r)
}

func (this *Tox) FriendGetStatusMessage(friendNumber uint32) (string, error) {
	var _fn = C.uint32_t(friendNumber)
	var cerr C.TOX_ERR_FRIEND_QUERY
	len := C.tox_friend_get_status_message_size(this.toxcore, _fn, &cerr)
	if cerr > 0 {
		return "", toxerr(cerr)
	}

	_buf := make([]byte, len)

	cerr = 0
	r := C.tox_friend_get_status_message(this.toxcore, _fn, (*C.uint8_t)(&_buf[0]), &cerr)
	if !bool(r) || cerr > 0 {
		return "", toxerr(cerr)
	}
	return string(_buf[:]), nil
}

func (this *Tox) SelfGetStatusMessage() (string, error) {
	nlen := C.tox_self_get_status_message_size(this.toxcore)
	var _buf = make([]byte, nlen)

	C.tox_self_get_status_message(this.toxcore, (*C.uint8_t)(&_buf[0]))
	return string(_buf[:]), nil
}

func (this *Tox) FriendGetStatus(friendNumber uint32) (int, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_status(this.toxcore, _fn, &cerr)
	if cerr > 0 {
		return int(r), toxerr(cerr)
	}
	return int(r), nil
}

func (this *Tox) SelfGetStatus() int {
	r := C.tox_self_get_status(this.toxcore)
	return int(r)
}

func (this *Tox) FriendGetLastOnline(friendNumber uint32) (uint64, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_GET_LAST_ONLINE
	r := C.tox_friend_get_last_online(this.toxcore, _fn, &cerr)
	if cerr > 0 {
		return uint64(r), toxerr(cerr)
	}
	return uint64(r), nil
}

func (this *Tox) SelfSetTyping(friendNumber uint32, typing bool) (bool, error) {
	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)
	var _typing = C._Bool(typing)

	var cerr C.TOX_ERR_SET_TYPING
	r := C.tox_self_set_typing(this.toxcore, _fn, _typing, &cerr)
	if cerr > 0 {
		return bool(r), toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) FriendGetTyping(friendNumber uint32) (bool, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_typing(this.toxcore, _fn, &cerr)
	if cerr > 0 {
		return bool(r), toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) SelfGetFriendListSize() uint32 {
	r := C.tox_self_get_friend_list_size(this.toxcore)
	return uint32(r)
}

func (this *Tox) SelfGetFriendList() []uint32 {
	sz := C.tox_self_get_friend_list_size(this.toxcore)
	vec := make([]uint32, sz)
	if sz == 0 {
		return vec
	}
	vec_p := unsafe.Pointer(&vec[0])
	C.tox_self_get_friend_list(this.toxcore, (*C.uint32_t)(vec_p))
	return vec
}

// tox_callback_***

func (this *Tox) SelfGetNospam() uint32 {
	r := C.tox_self_get_nospam(this.toxcore)
	return uint32(r)
}

func (this *Tox) SelfSetNospam(nospam uint32) {
	this.lock()
	defer this.unlock()

	var _nospam = C.uint32_t(nospam)

	C.tox_self_set_nospam(this.toxcore, _nospam)
}

func (this *Tox) SelfGetPublicKey() string {
	var _pubkey [PUBLIC_KEY_SIZE]byte

	C.tox_self_get_public_key(this.toxcore, (*C.uint8_t)(&_pubkey[0]))

	return strings.ToUpper(hex.EncodeToString(_pubkey[:]))
}

func (this *Tox) SelfGetSecretKey() string {
	var _seckey [SECRET_KEY_SIZE]byte

	C.tox_self_get_secret_key(this.toxcore, (*C.uint8_t)(&_seckey[0]))

	return strings.ToUpper(hex.EncodeToString(_seckey[:]))
}

// tox_lossy_***

func (this *Tox) FriendSendLossyPacket(friendNumber uint32, data string) error {
	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)
	var _data = []byte(data)
	var _length = C.size_t(len(data))

	var cerr C.TOX_ERR_FRIEND_CUSTOM_PACKET
	r := C.tox_friend_send_lossy_packet(this.toxcore, _fn, (*C.uint8_t)(&_data[0]), _length, &cerr)
	if !r || cerr != C.TOX_ERR_FRIEND_CUSTOM_PACKET_OK {
		return toxerr(cerr)
	}
	return nil
}

func (this *Tox) FriendSendLosslessPacket(friendNumber uint32, data string) error {
	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)
	var _data = []byte(data)
	var _length = C.size_t(len(data))

	var cerr C.TOX_ERR_FRIEND_CUSTOM_PACKET
	r := C.tox_friend_send_lossless_packet(this.toxcore, _fn, (*C.uint8_t)(&_data[0]), _length, &cerr)
	if !r || cerr != C.TOX_ERR_FRIEND_CUSTOM_PACKET_OK {
		return toxerr(cerr)
	}
	return nil
}

// tox_callback_avatar_**

func (this *Tox) Hash(data string, datalen uint32) (string, bool, error) {
	_data := []byte(data)
	_hash := make([]byte, C.TOX_HASH_LENGTH)
	var _datalen = C.size_t(datalen)

	r := C.tox_hash((*C.uint8_t)(&_hash[0]), (*C.uint8_t)(&_data[0]), _datalen)
	return string(_hash), bool(r), nil
}

// tox_callback_file_***
func (this *Tox) FileControl(friendNumber uint32, fileNumber uint32, control int) (bool, error) {
	var cerr C.TOX_ERR_FILE_CONTROL
	r := C.tox_file_control(this.toxcore, C.uint32_t(friendNumber), C.uint32_t(fileNumber),
		C.TOX_FILE_CONTROL(control), &cerr)
	if cerr > 0 {
		return false, toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) FileSend(friendNumber uint32, kind uint32, fileSize uint64, fileId string, fileName string) (uint32, error) {
	this.lock()
	defer this.unlock()

	if len(fileId) != FILE_ID_LENGTH*2 {
	}

	_fileName := []byte(fileName)

	var cerr C.TOX_ERR_FILE_SEND
	r := C.tox_file_send(this.toxcore, C.uint32_t(friendNumber), C.uint32_t(kind), C.uint64_t(fileSize),
		nil, (*C.uint8_t)(&_fileName[0]), C.size_t(len(fileName)), &cerr)
	if cerr > 0 {
		return uint32(r), toxerr(cerr)
	}
	return uint32(r), nil
}

func (this *Tox) FileSendChunk(friendNumber uint32, fileNumber uint32, position uint64, data []byte) (bool, error) {
	this.lock()
	defer this.unlock()

	if data == nil || len(data) == 0 {
		return false, toxerr("empty data")
	}
	var cerr C.TOX_ERR_FILE_SEND_CHUNK
	r := C.tox_file_send_chunk(this.toxcore, C.uint32_t(friendNumber), C.uint32_t(fileNumber),
		C.uint64_t(position), (*C.uint8_t)(&data[0]), C.size_t(len(data)), &cerr)
	if cerr > 0 {
		return bool(r), toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) FileSeek(friendNumber uint32, fileNumber uint32, position uint64) (bool, error) {
	this.lock()
	defer this.unlock()

	var cerr C.TOX_ERR_FILE_SEEK
	r := C.tox_file_seek(this.toxcore, C.uint32_t(friendNumber), C.uint32_t(fileNumber),
		C.uint64_t(position), &cerr)
	if cerr > 0 {
		return false, toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) FileGetFileId(friendNumber uint32, fileNumber uint32) (string, error) {
	var cerr C.TOX_ERR_FILE_GET
	var fileId_b = make([]byte, C.TOX_FILE_ID_LENGTH)

	r := C.tox_file_get_file_id(this.toxcore, C.uint32_t(fileNumber), C.uint32_t(fileNumber),
		(*C.uint8_t)(&fileId_b[0]), &cerr)
	if cerr > 0 || bool(r) == false {
		return "", toxerr(cerr)
	}

	var fileId_h = strings.ToUpper(hex.EncodeToString(fileId_b))
	return fileId_h, nil
}

// boostrap, see upper
func (this *Tox) AddTcpRelay(addr string, port uint16, pubkey string) (bool, error) {
	this.lock()
	defer this.unlock()

	var _addr = C.CString(addr)
	defer C.free(unsafe.Pointer(_addr))
	var _port = C.uint16_t(port)
	b_pubkey, err := hex.DecodeString(pubkey)
	if err != nil {
		log.Panic(err)
	}
	if strings.ToUpper(hex.EncodeToString(b_pubkey)) != pubkey {
		log.Panic("wtf, hex enc/dec err")
	}
	var _pubkey = (*C.uint8_t)(&b_pubkey[0])

	var cerr C.TOX_ERR_BOOTSTRAP
	r := C.tox_add_tcp_relay(this.toxcore, _addr, _port, _pubkey, &cerr)
	if cerr > 0 {
		return bool(r), toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) IsConnected() int {
	r := C.tox_self_get_connection_status(this.toxcore)
	return int(r)
}

func (this *Tox) beforeCallback(args ...interface{}) {
	this.unlock()
}
func (this *Tox) afterCallback(args ...interface{}) {
	this.lock()
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
