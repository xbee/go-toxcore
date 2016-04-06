package tox

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	// "reflect"
	// "runtime"
	"unsafe"
)

/*
// #cgo CFLAGS: -g -O2 -Wall
#cgo CFLAGS: -g -O2
#cgo LDFLAGS: -ltoxcore -ltoxdns -ltoxav -ltoxencryptsave -lvpx
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>

//////

void callbackFriendRequestWrapperForC(Tox *, uint8_t *, uint8_t *, uint16_t, void*);
typedef void (*cb_friend_request_ftype)(Tox *, const uint8_t *, const uint8_t *, size_t, void*);
static void cb_friend_request_wrapper_for_go(Tox *m, cb_friend_request_ftype fn, void *userdata)
{ tox_callback_friend_request(m, fn, userdata); }

void callbackFriendMessageWrapperForC(Tox *, uint32_t, int, uint8_t*, uint32_t, void*);
typedef void (*cb_friend_message_ftype)(Tox *, uint32_t, TOX_MESSAGE_TYPE, const uint8_t*, size_t, void*);
static void cb_friend_message_wrapper_for_go(Tox *m, cb_friend_message_ftype fn, void *userdata)
{ tox_callback_friend_message(m, fn, userdata); }

void callbackFriendNameWrapperForC(Tox *, uint32_t, uint8_t*, uint32_t, void*);
typedef void (*cb_friend_name_ftype)(Tox *, uint32_t, const uint8_t*, size_t, void*);
static void cb_friend_name_wrapper_for_go(Tox *m, cb_friend_name_ftype fn, void *userdata)
{ tox_callback_friend_name(m, fn, userdata); }

void callbackFriendStatusMessageWrapperForC(Tox *, uint32_t, uint8_t*, uint32_t, void*);
typedef void (*cb_friend_status_message_ftype)(Tox *, uint32_t, const uint8_t*, size_t, void*);
static void cb_friend_status_message_wrapper_for_go(Tox *m, cb_friend_status_message_ftype fn, void *userdata)
{ tox_callback_friend_status_message(m, fn, userdata); }

void callbackFriendStatusWrapperForC(Tox *, uint32_t, uint8_t, void*);
typedef void (*cb_friend_status_ftype)(Tox *, uint32_t, TOX_USER_STATUS, void*);
static void cb_friend_status_wrapper_for_go(Tox *m, cb_friend_status_ftype fn, void *userdata)
{ tox_callback_friend_status(m, fn, userdata); }

void callbackFriendConnectionStatusWrapperForC(Tox *, uint32_t, uint32_t, void*);
typedef void (*cb_friend_connection_status_ftype)(Tox *, uint32_t, uint32_t, void*);
static void cb_friend_connection_status_wrapper_for_go(Tox *m, cb_friend_connection_status_ftype fn, void *userdata)
{ tox_callback_friend_connection_status(m, fn, userdata); }

void callbackFriendTypingWrapperForC(Tox *, uint32_t, uint8_t, void*);
typedef void (*cb_friend_typing_ftype)(Tox *, uint32_t, uint8_t, void*);
static void cb_friend_typing_wrapper_for_go(Tox *m, cb_friend_typing_ftype fn, void *userdata)
{ tox_callback_friend_typing(m, fn, userdata); }

void callbackFriendReadReceiptWrapperForC(Tox *, uint32_t, uint32_t, void*);
typedef void (*cb_friend_read_receipt_ftype)(Tox *, uint32_t, uint32_t, void*);
static void cb_friend_read_receipt_wrapper_for_go(Tox *m, cb_friend_read_receipt_ftype fn, void *userdata)
{ tox_callback_friend_read_receipt(m, fn, userdata); }

void callbackFriendLossyPacketWrapperForC(Tox *, uint32_t, uint8_t*, size_t, void*);
typedef void (*cb_friend_lossy_packet_ftype)(Tox *, uint32_t, const uint8_t*, size_t, void*);
static void cb_friend_lossy_packet_wrapper_for_go(Tox *m, cb_friend_lossy_packet_ftype fn, void *userdata)
{ tox_callback_friend_lossy_packet(m, fn, userdata); }

void callbackFriendLosslessPacketWrapperForC(Tox *, uint32_t, uint8_t*, size_t, void*);
typedef void (*cb_friend_lossless_packet_ftype)(Tox *, uint32_t, const uint8_t*, size_t, void*);
static void cb_friend_lossless_packet_wrapper_for_go(Tox *m, cb_friend_lossless_packet_ftype fn, void *userdata)
{ tox_callback_friend_lossless_packet(m, fn, userdata); }

void callbackSelfConnectionStatusWrapperForC(Tox *, uint32_t, void*);
typedef void (*cb_self_connection_status_ftype)(Tox *, TOX_CONNECTION, void*);
static void cb_self_connection_status_wrapper_for_go(Tox *m, cb_self_connection_status_ftype fn, void *userdata)
{ tox_callback_self_connection_status(m, fn, userdata); }


void callbackGroupInviteWrapperForC(Tox*, int32_t, uint8_t, uint8_t *, uint16_t, void *);
typedef void (*cb_group_invite_ftype)(Tox *, int32_t, uint8_t, uint8_t *, uint16_t, void *);
static void cb_group_invite_wrapper_for_go(Tox *m, cb_group_invite_ftype fn, void *userdata)
{ tox_callback_group_invite(m, fn, userdata); }

void callbackGroupMessageWrapperForC(Tox *, int, int , uint8_t *, uint16_t, void *);
typedef void (*cb_group_message_ftype)(Tox *, int, int , uint8_t *, uint16_t, void *);
static void cb_group_message_wrapper_for_go(Tox *m, cb_group_message_ftype fn, void *userdata)
{ tox_callback_group_message(m, fn, userdata); }

void callbackGroupActionWrapperForC(Tox*, int, int, uint8_t*, uint16_t, void*);
typedef void (*cb_group_action_ftype)(Tox*, int, int, uint8_t*, uint16_t, void*);
static void cb_group_action_wrapper_for_go(Tox *m, cb_group_action_ftype fn, void *userdata)
{ tox_callback_group_action(m, fn, userdata); }

void callbackGroupTitleWrapperForC(Tox*, int, int, uint8_t*, uint8_t, void*);
typedef void (*cb_group_title_ftype)(Tox*, int, int, uint8_t*, uint8_t, void*);
static void cb_group_title_wrapper_for_go(Tox *m, cb_group_title_ftype fn, void *userdata)
{ tox_callback_group_title(m, fn, userdata); }

void callbackGroupNameListChangeWrapperForC(Tox*, int, int, uint8_t, void*);
typedef void (*cb_group_namelist_change_ftype)(Tox*, int, int, uint8_t, void*);
static void cb_group_namelist_change_wrapper_for_go(Tox *m, cb_group_namelist_change_ftype fn, void *userdata)
{ tox_callback_group_namelist_change(m, fn, userdata); }

void callbackFileRecvControlWrapperForC(Tox *tox, uint32_t friend_number, uint32_t file_number,
                                      TOX_FILE_CONTROL control, void *user_data);
typedef void (*cb_file_recv_control_ftype)(Tox *tox, uint32_t friend_number, uint32_t file_number,
                                      TOX_FILE_CONTROL control, void *useer_data);
static void cb_file_recv_control_wrapper_for_go(Tox *m, cb_file_recv_control_ftype fn, void *userdata)
{ tox_callback_file_recv_control(m, fn, userdata); }

void callbackFileRecvWrapperForC(Tox *tox, uint32_t friend_number, uint32_t file_number, uint32_t kind,
                               uint64_t file_size, uint8_t *filename, size_t filename_length, void *user_data);
typedef void (*cb_file_recv_ftype)(Tox *tox, uint32_t friend_number, uint32_t file_number, uint32_t kind,
              uint64_t file_size, const uint8_t *filename, size_t filename_length, void *user_data);
static void cb_file_recv_wrapper_for_go(Tox *m, cb_file_recv_ftype fn, void *userdata)
{ tox_callback_file_recv(m, fn, userdata); }

void callbackFileRecvChunkWrapperForC(Tox *tox, uint32_t friend_number, uint32_t file_number, uint64_t position,
                                    uint8_t *data, size_t length, void *user_data);
typedef void (*cb_file_recv_chunk_ftype)(Tox *tox, uint32_t friend_number, uint32_t file_number, uint64_t position,
                                    const uint8_t *data, size_t length, void *user_data);
static void cb_file_recv_chunk_wrapper_for_go(Tox *m, cb_file_recv_chunk_ftype fn, void *userdata)
{ tox_callback_file_recv_chunk(m, fn, userdata); }

void callbackFileChunkRequestWrapperForC(Tox *tox, uint32_t friend_number, uint32_t file_number, uint64_t position,
                                       size_t length, void *user_data);
typedef void (*cb_file_chunk_request_ftype)(Tox *tox, uint32_t friend_number, uint32_t file_number, uint64_t position,
                                       size_t length, void *user_data);
static void cb_file_chunk_request_wrapper_for_go(Tox *m, cb_file_chunk_request_ftype fn, void *userdata)
{ tox_callback_file_chunk_request(m, fn, userdata); }

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
    cb_group_invite_wrapper_for_go(NULL, NULL, NULL);
    cb_group_message_wrapper_for_go(NULL, NULL, NULL);
    cb_group_action_wrapper_for_go(NULL, NULL, NULL);
    cb_group_title_wrapper_for_go(NULL, NULL, NULL);
    cb_group_namelist_change_wrapper_for_go(NULL, NULL, NULL);
    cb_file_recv_control_wrapper_for_go(NULL, NULL, NULL);
    cb_file_recv_wrapper_for_go(NULL, NULL, NULL);
    cb_file_recv_chunk_wrapper_for_go(NULL, NULL, NULL);
    cb_file_chunk_request_wrapper_for_go(NULL, NULL, NULL);
}
*/
import "C"

//////////
// friend callback type
type cb_friend_request_ftype func(this *Tox, pubkey string, message string, userData interface{})
type cb_friend_message_ftype func(this *Tox, friendNumber uint32, message string, userData interface{})
type cb_friend_name_ftype func(this *Tox, friendNumber uint32, newName *uint8, length uint16, userData interface{})
type cb_friend_status_message_ftype func(this *Tox, friendNumber uint32, newStatus string, userData interface{})
type cb_friend_status_ftype func(this *Tox, friendNumber uint32, status uint8, userData interface{})
type cb_friend_connection_status_ftype func(this *Tox, friendNumber uint32, status uint32, userData interface{})
type cb_friend_typing_ftype func(this *Tox, friendNumber uint32, isTyping uint8, userData interface{})
type cb_friend_read_receipt_ftype func(this *Tox, friendNumber uint32, receipt uint32, userData interface{})
type cb_friend_lossy_packet_ftype func(this *Tox, friendNumber uint32, data string, userData interface{})
type cb_friend_lossless_packet_ftype func(this *Tox, friendNumber uint32, data string, userData interface{})

// self callback type
type cb_self_connection_status_ftype func(this *Tox, status uint32, userData interface{})

// group callback type
type cb_group_invite_ftype func(this *Tox, friendNumber uint32, itype uint8, data *uint8, length uint16, userData interface{})
type cb_group_message_ftype func(this *Tox, groupNumber int, peerNumber int, message *uint8, length uint16, userData interface{})
type cb_group_action_ftype func(this *Tox, groupNumber int, peerNumber int, action *uint8, length uint16, userData interface{})
type cb_group_title_ftype func(this *Tox, groupNumber int, peerNumber int, title *uint8, length uint8, userData interface{})
type cb_group_namelist_change_ftype func(this *Tox, groupNumber int, peerNumber int, change uint8, userData interface{})

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
	opts    *ToxOptions
	toxopts *C.struct_Tox_Options
	toxcore *C.Tox // save C.Tox

	// some callbacks, should be private
	cb_friend_request                     cb_friend_request_ftype
	cb_friend_request_user_data           interface{}
	cb_friend_message                     cb_friend_message_ftype
	cb_friend_message_user_data           interface{}
	cb_friend_name                        cb_friend_name_ftype
	cb_friend_name_user_data              interface{}
	cb_friend_status_message              cb_friend_status_message_ftype
	cb_friend_status_message_user_data    interface{}
	cb_friend_status                      cb_friend_status_ftype
	cb_friend_status_user_data            interface{}
	cb_friend_connection_status           cb_friend_connection_status_ftype
	cb_friend_connection_status_user_data interface{}
	cb_friend_typing                      cb_friend_typing_ftype
	cb_friend_typing_user_data            interface{}
	cb_friend_read_receipt                cb_friend_read_receipt_ftype
	cb_friend_read_receipt_user_data      interface{}
	cb_friend_lossy_packet                cb_friend_lossy_packet_ftype
	cb_friend_lossy_packet_user_data      interface{}
	cb_friend_lossless_packet             cb_friend_lossless_packet_ftype
	cb_friend_lossless_packet_user_data   interface{}
	cb_self_connection_status             cb_self_connection_status_ftype
	cb_self_connection_status_user_data   interface{}

	cb_group_invite                    cb_group_invite_ftype
	cb_group_invite_user_data          interface{}
	cb_group_message                   cb_group_message_ftype
	cb_group_message_user_data         interface{}
	cb_group_action                    cb_group_action_ftype
	cb_group_action_user_data          interface{}
	cb_group_title                     cb_group_title_ftype
	cb_group_title_user_data           interface{}
	cb_group_namelist_change           cb_group_namelist_change_ftype
	cb_group_namelist_change_user_data interface{}

	cb_file_recv_control            cb_file_recv_control_ftype
	cb_file_recv_control_user_data  interface{}
	cb_file_recv                    cb_file_recv_ftype
	cb_file_recv_user_data          interface{}
	cb_file_recv_chunk              cb_file_recv_chunk_ftype
	cb_file_recv_chunk_user_data    interface{}
	cb_file_chunk_request           cb_file_chunk_request_ftype
	cb_file_chunk_request_user_data interface{}
}

var cbUserDatas map[*C.Tox]*Tox = make(map[*C.Tox]*Tox, 0)

//export callbackFriendRequestWrapperForC
func callbackFriendRequestWrapperForC(m *C.Tox, a0 *C.uint8_t, a1 *C.uint8_t, a2 C.uint16_t, a3 unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_friend_request != nil {
		pubkey_b := C.GoBytes(unsafe.Pointer(a0), 32)
		pubkey := hex.EncodeToString(pubkey_b)
		pubkey = strings.ToUpper(pubkey)
		message_b := C.GoBytes(unsafe.Pointer(a1), C.int(a2))
		message := string(message_b)
		this.cb_friend_request(this, pubkey, message, this.cb_friend_request_user_data)
	}
}

func (this *Tox) CallbackFriendRequest(cbfn cb_friend_request_ftype, userData interface{}) {
	this.cb_friend_request = cbfn
	this.cb_friend_request_user_data = userData

	var _cbfn = (C.cb_friend_request_ftype)(C.callbackFriendRequestWrapperForC)
	var _userData = unsafe.Pointer(this)
	_userData = nil

	C.cb_friend_request_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendMessageWrapperForC
func callbackFriendMessageWrapperForC(m *C.Tox, a0 C.uint32_t, mtype C.int,
	a1 *C.uint8_t, a2 C.uint32_t, a3 unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_friend_message != nil {
		message_ := C.GoStringN((*C.char)(unsafe.Pointer(a1)), (C.int)(a2))
		this.cb_friend_message(this, uint32(a0), message_, this.cb_friend_message_user_data)
	}
}

func (this *Tox) CallbackFriendMessage(cbfn cb_friend_message_ftype, userData interface{}) {
	this.cb_friend_message = cbfn
	this.cb_friend_message_user_data = userData

	var _cbfn = (C.cb_friend_message_ftype)(C.callbackFriendMessageWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendNameWrapperForC
func callbackFriendNameWrapperForC(m *C.Tox, a0 C.uint32_t, a1 *C.uint8_t, a2 C.uint32_t, a3 unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_friend_name != nil {
		this.cb_friend_name(this, uint32(a0), (*uint8)(a1), uint16(a2), this.cb_friend_name_user_data)
	}
}

func (this *Tox) CallbackFriendName(cbfn cb_friend_name_ftype, userData interface{}) {
	this.cb_friend_name = cbfn
	this.cb_friend_name_user_data = userData

	var _cbfn = (C.cb_friend_name_ftype)(C.callbackFriendNameWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_friend_name_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendStatusMessageWrapperForC
func callbackFriendStatusMessageWrapperForC(m *C.Tox, a0 C.uint32_t, a1 *C.uint8_t, a2 C.uint32_t, a3 unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_friend_status_message != nil {
		statusText := C.GoStringN((*C.char)(unsafe.Pointer(a1)), C.int(a2))
		this.cb_friend_status_message(this, uint32(a0), statusText, this.cb_friend_status_message_user_data)
	}
}

func (this *Tox) CallbackFriendStatusMessage(cbfn cb_friend_status_message_ftype, userData interface{}) {
	this.cb_friend_status_message = cbfn
	this.cb_friend_status_message_user_data = userData

	var _cbfn = (C.cb_friend_status_message_ftype)(C.callbackFriendStatusMessageWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_status_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendStatusWrapperForC
func callbackFriendStatusWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint8_t, a2 unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_friend_status != nil {
		this.cb_friend_status(this, uint32(a0), uint8(a1), this.cb_friend_status_user_data)
	}
}

func (this *Tox) CallbackFriendStatus(cbfn cb_friend_status_ftype, userData interface{}) {
	this.cb_friend_status = cbfn
	this.cb_friend_status_user_data = userData

	var _cbfn = (C.cb_friend_status_ftype)(C.callbackFriendStatusWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_status_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendConnectionStatusWrapperForC
func callbackFriendConnectionStatusWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, a2 unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_friend_connection_status != nil {
		this.cb_friend_connection_status(this, uint32(a0), uint32(a1), this.cb_friend_connection_status_user_data)
	}
}

func (this *Tox) CallbackFriendConnectionStatus(cbfn cb_friend_connection_status_ftype, userData interface{}) {
	this.cb_friend_connection_status = cbfn
	this.cb_friend_connection_status_user_data = userData

	var _cbfn = (C.cb_friend_connection_status_ftype)(C.callbackFriendConnectionStatusWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_connection_status_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendTypingWrapperForC
func callbackFriendTypingWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint8_t, a2 unsafe.Pointer) {
	var this = (*Tox)(a2)
	if this.cb_friend_typing != nil {
		this.cb_friend_typing(this, uint32(a0), uint8(a1), this.cb_friend_typing_user_data)
	}
}

func (this *Tox) CallbackFriendTyping(cbfn cb_friend_typing_ftype, userData interface{}) {
	this.cb_friend_typing = cbfn
	this.cb_friend_typing_user_data = userData

	var _cbfn = (C.cb_friend_typing_ftype)(C.callbackFriendTypingWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_friend_typing_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendReadReceiptWrapperForC
func callbackFriendReadReceiptWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, a2 unsafe.Pointer) {
	var this = cbUserDatas[m]
	if this.cb_friend_read_receipt != nil {
		this.cb_friend_read_receipt(this, uint32(a0), uint32(a1), this.cb_friend_read_receipt_user_data)
	}
}

func (this *Tox) CallbackFriendReadReceipt(cbfn cb_friend_read_receipt_ftype, userData interface{}) {
	this.cb_friend_read_receipt = cbfn
	this.cb_friend_read_receipt_user_data = userData

	var _cbfn = (C.cb_friend_read_receipt_ftype)(C.callbackFriendReadReceiptWrapperForC)
	var _userData unsafe.Pointer

	C.cb_friend_read_receipt_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendLossyPacketWrapperForC
func callbackFriendLossyPacketWrapperForC(m *C.Tox, a0 C.uint32_t, a1 *C.uint8_t, len C.size_t, a2 unsafe.Pointer) {
	var this = cbUserDatas[m]
	if this.cb_friend_lossy_packet != nil {
		msg := C.GoStringN((*C.char)(unsafe.Pointer(a1)), C.int(len))
		this.cb_friend_lossy_packet(this, uint32(a0), msg, this.cb_friend_lossy_packet_user_data)
	}
}

func (this *Tox) CallbackFriendLossyPacket(cbfn cb_friend_lossy_packet_ftype, userData interface{}) {
	this.cb_friend_lossy_packet = cbfn
	this.cb_friend_lossy_packet_user_data = userData

	var _cbfn = (C.cb_friend_lossy_packet_ftype)(C.callbackFriendLossyPacketWrapperForC)
	var _userData unsafe.Pointer

	C.cb_friend_lossy_packet_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFriendLosslessPacketWrapperForC
func callbackFriendLosslessPacketWrapperForC(m *C.Tox, a0 C.uint32_t, a1 *C.uint8_t, len C.size_t, a2 unsafe.Pointer) {
	var this = cbUserDatas[m]
	if this.cb_friend_lossless_packet != nil {
		msg := C.GoStringN((*C.char)(unsafe.Pointer(a1)), C.int(len))
		this.cb_friend_lossless_packet(this, uint32(a0), msg, this.cb_friend_lossless_packet_user_data)
	}
}

func (this *Tox) CallbackFriendLosslessPacket(cbfn cb_friend_lossless_packet_ftype, userData interface{}) {
	this.cb_friend_lossless_packet = cbfn
	this.cb_friend_lossless_packet_user_data = userData

	var _cbfn = (C.cb_friend_lossless_packet_ftype)(C.callbackFriendLosslessPacketWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_friend_lossless_packet_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackSelfConnectionStatusWrapperForC
func callbackSelfConnectionStatusWrapperForC(m *C.Tox, status C.uint32_t, a2 unsafe.Pointer) {
	var this = cbUserDatas[m]
	if this.cb_self_connection_status != nil {
		this.cb_self_connection_status(this, uint32(status), this.cb_self_connection_status_user_data)
	}
}

func (this *Tox) CallbackSelfConnectionStatus(cbfn cb_self_connection_status_ftype, userData interface{}) {
	this.cb_self_connection_status = cbfn
	this.cb_self_connection_status_user_data = userData

	var _cbfn = (C.cb_self_connection_status_ftype)(C.callbackSelfConnectionStatusWrapperForC)
	// var _userData = unsafe.Pointer(this)

	C.cb_self_connection_status_wrapper_for_go(this.toxcore, _cbfn, nil)
}

//export callbackGroupInviteWrapperForC
func callbackGroupInviteWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_group_invite != nil {
		this.cb_group_invite(this, uint32(a0), uint8(a1), (*uint8)(a2), uint16(a3), this.cb_group_invite_user_data)
	}
}

func (this *Tox) CallbackGroupInvite(cbfn cb_group_invite_ftype, userData interface{}) {
	this.cb_group_invite = cbfn
	this.cb_group_invite_user_data = userData

	var _cbfn = (C.cb_group_invite_ftype)(C.callbackGroupInviteWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_group_invite_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupMessageWrapperForC
func callbackGroupMessageWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_group_message != nil {
		this.cb_group_message(this, int(a0), int(a1), (*uint8)(a2), uint16(a3), this.cb_group_message_user_data)
	}
}

func (this *Tox) CallbackGroupMessage(cbfn cb_group_message_ftype, userData interface{}) {
	this.cb_group_message = cbfn
	this.cb_group_message_user_data = userData

	var _cbfn = (C.cb_group_message_ftype)(C.callbackGroupMessageWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_group_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupActionWrapperForC
func callbackGroupActionWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_group_action != nil {
		this.cb_group_action(this, int(a0), int(a1), (*uint8)(a2), uint16(a3), this.cb_group_action_user_data)
	}
}

func (this *Tox) CallbackGroupAction(cbfn cb_group_action_ftype, userData interface{}) {
	this.cb_group_action = cbfn
	this.cb_group_action_user_data = userData

	var _cbfn = (C.cb_group_action_ftype)(C.callbackGroupActionWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_group_action_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupTitleWrapperForC
func callbackGroupTitleWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint8_t, a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_group_title != nil {
		this.cb_group_title(this, int(a0), int(a1), (*uint8)(a2), uint8(a3), this.cb_group_title_user_data)
	}
}

func (this *Tox) CallbackGroupTitle(cbfn cb_group_title_ftype, userData interface{}) {
	this.cb_group_title = cbfn
	this.cb_group_title_user_data = userData

	var _cbfn = (C.cb_group_title_ftype)(C.callbackGroupTitleWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_group_title_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupNameListChangeWrapperForC
func callbackGroupNameListChangeWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 C.uint8_t, a3 unsafe.Pointer) {
	var this = (*Tox)(a3)
	if this.cb_group_namelist_change != nil {
		this.cb_group_namelist_change(this, int(a0), int(a1), uint8(a2), this.cb_group_namelist_change_user_data)
	}
}

func (this *Tox) CallbackGroupNameListChange(cbfn cb_group_namelist_change_ftype, userData interface{}) {
	this.cb_group_namelist_change = cbfn
	this.cb_group_namelist_change_user_data = userData

	var _cbfn = (C.cb_group_namelist_change_ftype)(C.callbackGroupNameListChangeWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_group_namelist_change_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

// 包内部函数
//export callbackFileRecvControlWrapperForC
func callbackFileRecvControlWrapperForC(m *C.Tox, friendNumber C.uint32_t, fileNumber C.uint32_t,
	control C.TOX_FILE_CONTROL, userData unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_file_recv_control != nil {
		this.cb_file_recv_control(this, uint32(friendNumber), uint32(fileNumber),
			int(control), this.cb_file_recv_control_user_data)
	}
}

func (this *Tox) CallbackFileRecvControl(cbfn cb_file_recv_control_ftype, userData interface{}) {
	this.cb_file_recv_control = cbfn
	this.cb_file_recv_control_user_data = userData
	var _cbfn = (C.cb_file_recv_control_ftype)(unsafe.Pointer(C.callbackFileRecvControlWrapperForC))
	var _userData unsafe.Pointer = nil

	C.cb_file_recv_control_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFileRecvWrapperForC
func callbackFileRecvWrapperForC(m *C.Tox, friendNumber C.uint32_t, fileNumber C.uint32_t, kind C.uint32_t,
	fileSize C.uint64_t, fileName *C.uint8_t, fileNameLength C.size_t, userData unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_file_recv != nil {
		fileName_ := C.GoStringN((*C.char)(unsafe.Pointer(fileName)), C.int(fileNameLength))
		this.cb_file_recv(this, uint32(friendNumber), uint32(fileNumber), uint32(kind),
			uint64(fileSize), fileName_, this.cb_file_recv_user_data)
	}
}

func (this *Tox) CallbackFileRecv(cbfn cb_file_recv_ftype, userData interface{}) {
	this.cb_file_recv = cbfn
	this.cb_file_recv_user_data = userData
	var _cbfn = (C.cb_file_recv_ftype)(unsafe.Pointer(C.callbackFileRecvWrapperForC))
	var _userData unsafe.Pointer = nil

	C.cb_file_recv_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFileRecvChunkWrapperForC
func callbackFileRecvChunkWrapperForC(m *C.Tox, friendNumber C.uint32_t, fileNumber C.uint32_t,
	position C.uint64_t, data *C.uint8_t, length C.size_t, userData unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_file_recv != nil {
		data_ := C.GoBytes((unsafe.Pointer)(data), C.int(length))
		this.cb_file_recv_chunk(this, uint32(friendNumber), uint32(fileNumber), uint64(position),
			data_, this.cb_file_recv_chunk_user_data)
	}
}

func (this *Tox) CallbackFileRecvChunk(cbfn cb_file_recv_chunk_ftype, userData interface{}) {
	this.cb_file_recv_chunk = cbfn
	this.cb_file_recv_chunk_user_data = userData
	var _cbfn = (C.cb_file_recv_chunk_ftype)(unsafe.Pointer(C.callbackFileRecvChunkWrapperForC))
	var _userData unsafe.Pointer = nil

	C.cb_file_recv_chunk_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackFileChunkRequestWrapperForC
func callbackFileChunkRequestWrapperForC(m *C.Tox, friendNumber C.uint32_t, fileNumber C.uint32_t,
	position C.uint64_t, length C.size_t, userData unsafe.Pointer) {
	var this = (*Tox)(cbUserDatas[m])
	if this.cb_file_recv != nil {
		this.cb_file_chunk_request(this, uint32(friendNumber), uint32(fileNumber), uint64(position),
			int(length), this.cb_file_chunk_request_user_data)
	}
}

func (this *Tox) CallbackFileChunkRequest(cbfn cb_file_chunk_request_ftype, userData interface{}) {
	this.cb_file_chunk_request = cbfn
	this.cb_file_chunk_request_user_data = userData
	var _cbfn = (C.cb_file_chunk_request_ftype)(unsafe.Pointer(C.callbackFileChunkRequestWrapperForC))
	var _userData unsafe.Pointer = nil

	C.cb_file_chunk_request_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

func NewTox(opt *ToxOptions) *Tox {
	var tox = new(Tox)
	tox.opts = NewToxOptions()
	if opt != nil {
		tox.opts = opt
	}
	tox.toxopts = tox.opts.toCToxOptions()

	var cerr C.TOX_ERR_NEW
	var toxcore = C.tox_new(tox.toxopts, &cerr)
	tox.toxcore = toxcore
	if toxcore == nil {
		log.Println(toxerr(cerr))
		return nil
	}
	cbUserDatas[toxcore] = tox

	return tox
}

func (this *Tox) Kill() {
	C.tox_kill(this.toxcore)
	this.toxcore = nil
}

// uint32_t tox_iteration_interval(Tox *tox);
func (this *Tox) IterationInterval() int {
	r := C.tox_iteration_interval(this.toxcore)
	return int(r)
}

/* The main loop that needs to be run in intervals of tox_iteration_interval() ms. */
// void tox_iterate(Tox *tox);
func (this *Tox) Iterate() {
	C.tox_iterate(this.toxcore)
}

func (this *Tox) GetSavedataSize() int32 {
	r := C.tox_get_savedata_size(this.toxcore)
	return int32(r)
}

func (this *Tox) GetSavedata() []byte {
	r := C.tox_get_savedata_size(this.toxcore)
	var savedata = make([]byte, int(r))

	C.tox_get_savedata(this.toxcore, bytes2uint8(savedata))
	return savedata
}

/*
 * @param pubkey hex 64B length
 */
func (this *Tox) Bootstrap(addr string, port uint16, pubkey string) (bool, error) {
	b_pubkey, err := hex.DecodeString(pubkey)
	if err != nil {
		log.Panic(err)
	}

	var _addr *C.char = C.CString(addr)
	defer C.free(unsafe.Pointer(_addr))
	var _port C.uint16_t = C.uint16_t(port)
	var _cpubkey *C.char = (*C.char)(unsafe.Pointer(&b_pubkey[0]))

	var cerr C.TOX_ERR_BOOTSTRAP
	r := C.tox_bootstrap(this.toxcore, _addr, _port, char2uint8(_cpubkey), &cerr)
	return bool(r), nil
}

func (this *Tox) SelfGetAddress() string {
	var addr [C.TOX_ADDRESS_SIZE]byte
	var caddr = (*C.char)(unsafe.Pointer(&addr[0]))
	C.tox_self_get_address(this.toxcore, char2uint8(caddr))

	haddr := hex.EncodeToString(addr[0:])
	return strings.ToUpper(haddr)
}

func (this *Tox) SelfGetConnectionStatus() int {
	r := C.tox_self_get_connection_status(this.toxcore)
	return int(r)
}

// int32_t tox_friend_add(Tox *tox, const uint8_t *address, const uint8_t *data, uint16_t length);
func (this *Tox) FriendAdd(friendId string, message string) (uint32, error) {
	friendId_b, err := hex.DecodeString(friendId)
	friendId_p := unsafe.Pointer(&friendId_b[0])
	if err != nil {
		log.Panic(err)
	}

	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))

	var cerr C.TOX_ERR_FRIEND_ADD
	r := C.tox_friend_add(this.toxcore, pointer2uint8(friendId_p),
		char2uint8(cmessage), C.size_t(len(message)), &cerr)
	return uint32(r), nil
}

func (this *Tox) FriendAddNorequest(friendId string) (uint32, error) {
	friendId_b, err := hex.DecodeString(friendId)
	friendId_p := unsafe.Pointer(&friendId_b[0])
	if err != nil {
		log.Panic(err)
	}

	var cerr C.TOX_ERR_FRIEND_ADD
	r := C.tox_friend_add_norequest(this.toxcore, pointer2uint8(friendId_p), &cerr)
	return uint32(r), nil
}

func (this *Tox) FriendByPublicKey(pubkey string) (uint32, error) {
	var pubkey_b, _ = hex.DecodeString(pubkey)
	var pubkey_p = unsafe.Pointer(&pubkey_b[0])

	var cerr C.TOX_ERR_FRIEND_BY_PUBLIC_KEY
	r := C.tox_friend_by_public_key(this.toxcore, pointer2uint8(pubkey_p), &cerr)
	if cerr != C.TOX_ERR_FRIEND_BY_PUBLIC_KEY_OK {
		return uint32(r), toxerr(cerr)
	}
	return uint32(r), nil
}

func (this *Tox) FriendGetPublicKey(friendNumber uint32) (string, error) {
	var _fn = C.uint32_t(friendNumber)
	var pubkey_b = make([]byte, 32)
	var pubkey_p = unsafe.Pointer(&pubkey_b[0])

	var cerr C.TOX_ERR_FRIEND_GET_PUBLIC_KEY
	r := C.tox_friend_get_public_key(this.toxcore, _fn, pointer2uint8(pubkey_p), &cerr)
	if bool(r) {
		pubkey_h := hex.EncodeToString(pubkey_b)
		pubkey_h = strings.ToUpper(pubkey_h)
		return pubkey_h, nil
	}
	return "", toxerr(cerr)
}

func (this *Tox) FriendDelete(friendNumber uint32) (bool, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_DELETE
	r := C.tox_friend_delete(this.toxcore, _fn, &cerr)
	return bool(r), nil
}

func (this *Tox) FriendGetConnectionStatus(friendNumber uint32) (int, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_connection_status(this.toxcore, _fn, &cerr)
	return int(r), nil

}

func (this *Tox) FriendExists(friendNumber uint32) (bool, error) {
	var _fn = C.uint32_t(friendNumber)

	r := C.tox_friend_exists(this.toxcore, _fn)
	return bool(r), nil
}

func (this *Tox) FriendSendMessage(friendNumber uint32, message string) (int32, error) {
	var _fn = C.uint32_t(friendNumber)
	var _message = C.CString(message)
	defer C.free(unsafe.Pointer(_message))
	var _length = C.size_t(len(message))

	var mtype C.TOX_MESSAGE_TYPE = C.TOX_MESSAGE_TYPE_NORMAL
	var cerr C.TOX_ERR_FRIEND_SEND_MESSAGE
	r := C.tox_friend_send_message(this.toxcore, _fn, mtype, char2uint8(_message), _length, &cerr)
	if cerr != C.TOX_ERR_FRIEND_SEND_MESSAGE_OK {
		return int32(r), toxerr(cerr)
	}
	return int32(r), nil
}

func (this *Tox) FriendSendAction(friendNumber uint32, action string) (int32, error) {
	var _fn = C.uint32_t(friendNumber)
	var _action = C.CString(action)
	defer C.free(unsafe.Pointer(_action))
	var _length = C.size_t(len(action))

	var mtype C.TOX_MESSAGE_TYPE = C.TOX_MESSAGE_TYPE_ACTION
	var cerr C.TOX_ERR_FRIEND_SEND_MESSAGE
	r := C.tox_friend_send_message(this.toxcore, _fn, mtype, char2uint8(_action), _length, &cerr)
	return int32(r), nil
}

func (this *Tox) SelfSetName(name string) (bool, error) {
	var _name = C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	var _length = C.size_t(len(name))

	var cerr C.TOX_ERR_SET_INFO
	r := C.tox_self_set_name(this.toxcore, char2uint8(_name), _length, &cerr)
	if cerr > 0 {
	}
	return bool(r), nil
}

func (this *Tox) SelfGetName() (string, error) {
	var _name = (*C.char)(C.malloc(C.tox_self_get_name_size(this.toxcore)))
	defer C.free(unsafe.Pointer(_name))

	C.tox_self_get_name(this.toxcore, char2uint8(_name))
	return C.GoString(_name), nil
}

func (this *Tox) FriendGetName(friendNumber uint32) (string, error) {
	var _fn = C.uint32_t(friendNumber)
	var cerr C.TOX_ERR_FRIEND_QUERY

	var _name = (*C.char)(C.malloc(C.tox_friend_get_name_size(this.toxcore, _fn, &cerr)))
	defer C.free(unsafe.Pointer(_name))

	r := C.tox_friend_get_name(this.toxcore, _fn, char2uint8(_name), &cerr)
	if !bool(r) {
		return "", toxerr(cerr)
	}
	return C.GoString(_name), nil
}

func (this *Tox) FriendGetNameSize(friendNumber uint32) (int, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_name_size(this.toxcore, _fn, &cerr)
	return int(r), nil
}

func (this *Tox) SelfGetNameSize() (int, error) {
	r := C.tox_self_get_name_size(this.toxcore)
	return int(r), nil
}

func (this *Tox) SelfSetStatusMessage(status string) (bool, error) {
	var _status = C.CString(status)
	defer C.free(unsafe.Pointer(_status))
	var _length = C.size_t(len(status))

	var cerr C.TOX_ERR_SET_INFO
	r := C.tox_self_set_status_message(this.toxcore, char2uint8(_status), _length, &cerr)
	return bool(r), nil
}

func (this *Tox) SelfSetStatus(status uint8) error {
	var _status = C.TOX_USER_STATUS(status)

	C.tox_self_set_status(this.toxcore, _status)
	return nil
}

func (this *Tox) FriendGetStatusMessageSize(friendNumber uint32) (int, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_status_message_size(this.toxcore, _fn, &cerr)
	return int(r), nil
}

func (this *Tox) SelfGetStatusMessageSize() (int, error) {
	r := C.tox_self_get_status_message_size(this.toxcore)
	return int(r), nil
}

func (this *Tox) FriendGetStatusMessage(friendNumber uint32) (string, error) {
	var _fn = C.uint32_t(friendNumber)
	var cerr C.TOX_ERR_FRIEND_QUERY
	var _buf = (*C.char)(C.malloc(C.tox_friend_get_status_message_size(this.toxcore, _fn, &cerr)))
	defer C.free(unsafe.Pointer(_buf))

	r := C.tox_friend_get_status_message(this.toxcore, _fn, char2uint8(_buf), &cerr)
	if !bool(r) {
		return "", toxerr(cerr)
	}

	return C.GoString(_buf), nil
}

func (this *Tox) SelfGetStatusMessage() (string, error) {
	var _buf = (*C.char)(C.malloc(C.tox_self_get_status_message_size(this.toxcore)))
	defer C.free(unsafe.Pointer(_buf))

	C.tox_self_get_status_message(this.toxcore, char2uint8(_buf))
	return C.GoString(_buf), nil
}

func (this *Tox) FriendGetStatus(friendNumber uint32) (uint8, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_status(this.toxcore, _fn, &cerr)
	return uint8(r), nil
}

func (this *Tox) SelfGetUserStatus() (uint8, error) {
	r := C.tox_self_get_status(this.toxcore)
	return uint8(r), nil
}

func (this *Tox) GetLastOnline(friendNumber uint32) (uint64, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_GET_LAST_ONLINE
	r := C.tox_friend_get_last_online(this.toxcore, _fn, &cerr)
	return uint64(r), nil
}

func (this *Tox) SelfSetTyping(friendNumber uint32, typing bool) (bool, error) {
	var _fn = C.uint32_t(friendNumber)
	var _typing = C._Bool(typing)

	var cerr C.TOX_ERR_SET_TYPING
	r := C.tox_self_set_typing(this.toxcore, _fn, _typing, &cerr)
	return bool(r), nil
}

func (this *Tox) FriendGetTyping(friendNumber uint32) (bool, error) {
	var _fn = C.uint32_t(friendNumber)

	var cerr C.TOX_ERR_FRIEND_QUERY
	r := C.tox_friend_get_typing(this.toxcore, _fn, &cerr)
	return bool(r), nil
}

func (this *Tox) SelfGetFriendListSize() (uint32, error) {
	r := C.tox_self_get_friend_list_size(this.toxcore)
	return uint32(r), nil
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
	var _nospam = C.uint32_t(nospam)

	C.tox_self_set_nospam(this.toxcore, _nospam)
}

func (this *Tox) SelfGetPublicKey() string {
	var _pubkey = (*C.char)(C.malloc(32))
	defer C.free(unsafe.Pointer(_pubkey))

	C.tox_self_get_public_key(this.toxcore, char2uint8(_pubkey))
	b_pubkey := C.GoBytes(unsafe.Pointer(_pubkey), 32)
	h_pubkey := hex.EncodeToString(b_pubkey)

	return strings.ToUpper(h_pubkey)
}

func (this *Tox) SelfGetSecretKey() string {
	var _seckey = (*C.char)(C.malloc(32))
	defer C.free(unsafe.Pointer(_seckey))

	C.tox_self_get_secret_key(this.toxcore, char2uint8(_seckey))
	b_seckey := C.GoBytes(unsafe.Pointer(_seckey), 32)
	h_seckey := hex.EncodeToString(b_seckey)

	return strings.ToUpper(h_seckey)
}

// tox_lossy_***

func (this *Tox) FriendSendLossyPacket(friendNumber uint32, data string) error {
	var _fn = C.uint32_t(friendNumber)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.size_t(len(data))

	var cerr C.TOX_ERR_FRIEND_CUSTOM_PACKET
	r := C.tox_friend_send_lossy_packet(this.toxcore, _fn, char2uint8(_data), _length, &cerr)
	if !r || cerr != C.TOX_ERR_FRIEND_CUSTOM_PACKET_OK {
		return toxerr(cerr)
	}
	return nil
}

func (this *Tox) FriendSendLossLessPacket(friendNumber uint32, data string) error {
	var _fn = C.uint32_t(friendNumber)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.size_t(len(data))

	var cerr C.TOX_ERR_FRIEND_CUSTOM_PACKET
	r := C.tox_friend_send_lossless_packet(this.toxcore, _fn, char2uint8(_data), _length, &cerr)
	if !r || cerr != C.TOX_ERR_FRIEND_CUSTOM_PACKET_OK {
		return toxerr(cerr)
	}
	return nil
}

// tox_callback_group_***

func (this *Tox) AddGroupChat() (int, error) {
	r := C.tox_add_groupchat(this.toxcore)
	return int(r), nil
}

func (this *Tox) DelGroupChat(groupNumber int) (int, error) {
	var _gn = C.int(groupNumber)

	r := C.tox_del_groupchat(this.toxcore, _gn)
	return int(r), nil
}

func (this *Tox) GroupPeerName(groupNumber int, peerNumber int, name string) (int, error) {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)
	var _name = C.CString(name)
	defer C.free(unsafe.Pointer(_name))

	r := C.tox_group_peername(this.toxcore, _gn, _pn, char2uint8(_name))
	name = C.GoString(_name)
	return int(r), nil
}

func (this *Tox) GroupPeerPubkey(groupNumber int, peerNumber int, public_key string) (int, error) {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)
	var _pubkey = C.CString(public_key)
	defer C.free(unsafe.Pointer(_pubkey))

	r := C.tox_group_peer_pubkey(this.toxcore, _gn, _pn, char2uint8(_pubkey))
	public_key = C.GoString(_pubkey)
	return int(r), nil
}

func (this *Tox) InviteFriend(friendNumber uint32, groupNumber int) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _gn = C.int(groupNumber)

	r := C.tox_invite_friend(this.toxcore, _fn, _gn)
	return int(r), nil
}

func (this *Tox) JoinGroupChat(friendNumber uint32, data string, length uint16) (int, error) {
	var _fn = C.int32_t(friendNumber)
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _length = C.uint16_t(length)

	r := C.tox_join_groupchat(this.toxcore, _fn, char2uint8(_data), _length)
	return int(r), nil
}

func (this *Tox) GroupActionSend(groupNumber int, action string, length uint16) (int, error) {
	var _gn = C.int(groupNumber)
	var _action = C.CString(action)
	defer C.free(unsafe.Pointer(_action))
	var _length = C.uint16_t(length)

	r := C.tox_group_action_send(this.toxcore, _gn, char2uint8(_action), _length)
	return int(r), nil
}

func (this *Tox) GroupSetTitle(groupNumber int, title string, length uint8) (int, error) {
	var _gn = C.int(groupNumber)
	var _title = C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	var _length = C.uint8_t(length)

	r := C.tox_group_set_title(this.toxcore, _gn, char2uint8(_title), _length)
	return int(r), nil
}

func (this *Tox) GroupGetTitle(groupNumber int, title string, maxlen uint32) (int, error) {
	var _gn = C.int(groupNumber)
	var _title = C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	var _maxlen = C.uint32_t(maxlen)

	r := C.tox_group_get_title(this.toxcore, _gn, char2uint8(_title), _maxlen)
	title = C.GoString(_title)
	return int(r), nil
}

func (this *Tox) GroupPeerNumberIsOurs(groupNumber int, peerNumber int) (uint, error) {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)

	r := C.tox_group_peernumber_is_ours(this.toxcore, _gn, _pn)
	return uint(r), nil
}

func (this *Tox) GroupNumberPeers(groupNumber int) (int, error) {
	var _gn = C.int(groupNumber)

	r := C.tox_group_number_peers(this.toxcore, _gn)
	return int(r), nil
}

/*
int tox_group_get_names(const Tox *tox, int groupnumber, uint8_t names[][TOX_MAX_NAME_LENGTH],
	uint16_t lengths[],
	uint16_t length);
*/

func (this *Tox) CountChatList() (uint32, error) {
	r := C.tox_count_chatlist(this.toxcore)
	return uint32(r), nil
}

// TODO...
func (this *Tox) GetChatList(outList []int32, listSize uint32) (uint32, error) {
	return uint32(0), nil
}

func (this *Tox) GroupGetType(groupNumber int) (int, error) {
	var _gn = C.int(groupNumber)

	r := C.tox_group_get_type(this.toxcore, _gn)
	return int(r), nil
}

// tox_callback_avatar_**

func (this *Tox) Hash(hash string, data string, datalen uint32) (bool, error) {
	var _hash = C.CString(hash)
	defer C.free(unsafe.Pointer(_hash))
	var _data = C.CString(data)
	defer C.free(unsafe.Pointer(_data))
	var _datalen = C.size_t(datalen)

	r := C.tox_hash(char2uint8(_hash), char2uint8(_data), _datalen)
	hash = C.GoString(_hash)
	return bool(r), nil
}

// tox_callback_file_***
func (this *Tox) FileControl(friendNumber uint32, fileNumber uint32, control int) (bool, error) {
	var cerr C.TOX_ERR_FILE_CONTROL
	r := C.tox_file_control(this.toxcore, C.uint32_t(friendNumber), C.uint32_t(fileNumber),
		C.TOX_FILE_CONTROL(control), &cerr)
	return bool(r), nil
}

func (this *Tox) FileSend(friendNumber uint32, kind uint32, fileSize uint64, fileId string, fileName string) (uint32, error) {
	fileName_ := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileName_))

	var cerr C.TOX_ERR_FILE_SEND
	r := C.tox_file_send(this.toxcore, C.uint32_t(friendNumber), C.uint32_t(kind), C.uint64_t(fileSize),
		nil, char2uint8(fileName_), C.size_t(len(fileName)), &cerr)
	return uint32(r), nil
}

func (this *Tox) FileSendChunk(friendNumber uint32, fileNumber uint32, position uint64, data []byte) (bool, error) {
	var cerr C.TOX_ERR_FILE_SEND_CHUNK
	r := C.tox_file_send_chunk(this.toxcore, C.uint32_t(friendNumber), C.uint32_t(fileNumber),
		C.uint64_t(position), pointer2uint8((unsafe.Pointer)(&data[0])), C.size_t(len(data)), &cerr)
	if !bool(r) {
		return bool(r), toxerr(cerr)
	}
	return bool(r), nil
}

func (this *Tox) FileSeek(friendNumber uint32, fileNumber uint32, position uint64) (bool, error) {
	var cerr C.TOX_ERR_FILE_SEEK
	r := C.tox_file_seek(this.toxcore, C.uint32_t(friendNumber), C.uint32_t(fileNumber),
		C.uint64_t(position), &cerr)
	return bool(r), nil
}

func (this *Tox) FileGetFileId(friendNumber uint32, fileNumber uint32) (string, error) {
	var cerr C.TOX_ERR_FILE_GET
	var fileId_b = make([]byte, C.TOX_FILE_ID_LENGTH)

	r := C.tox_file_get_file_id(this.toxcore, C.uint32_t(fileNumber), C.uint32_t(fileNumber),
		pointer2uint8((unsafe.Pointer)(&fileId_b[0])), &cerr)
	if !bool(r) {
		return "", toxerr(cerr)
	}

	var fileId_h = strings.ToUpper(hex.EncodeToString(fileId_b))
	return fileId_h, nil
}

// boostrap, see upper

func (this *Tox) AddTcpRelay(addr string, port uint16, pubkey string) (bool, error) {
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
	var _pubkey = (*C.char)(unsafe.Pointer(&b_pubkey[0]))

	var cerr C.TOX_ERR_BOOTSTRAP
	r := C.tox_add_tcp_relay(this.toxcore, _addr, _port, char2uint8(_pubkey), &cerr)
	return bool(r), nil
}

func (this *Tox) IsConnected() (int, error) {
	r := C.tox_self_get_connection_status(this.toxcore)
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
