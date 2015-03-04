package tox

import (
	"fmt"
	"log"
	"reflect"
	// "runtime"
	"unsafe"
)

/*
// #cgo CFLAGS: -g -O2 -Wall
#cgo CFLAGS: -g -O2
#cgo LDFLAGS: -ltoxcore -ltoxdns -ltoxav -ltoxencryptsave
#include <stdlib.h>
#include <tox/tox.h>

///////
int CalledByCGO();
int FortytwoAbc();

//////

void callbackFriendRequestWrapperForC(Tox *, uint8_t *, uint8_t *, uint16_t, void*);
typedef void (*cb_friend_request_ftype)(Tox *, uint8_t *, uint8_t *, uint16_t, void*);
static void cb_friend_request_wrapper_for_go(Tox *m, cb_friend_request_ftype fn, void *userdata)
{ tox_callback_friend_request(m, fn, userdata); }

void callbackFriendMessageWrapperForC(Tox *, int32_t, uint8_t*, uint16_t, void*);
typedef void (*cb_friend_message_ftype)(Tox *, int32_t, uint8_t*, uint16_t, void*);
static void cb_friend_message_wrapper_for_go(Tox *m, cb_friend_message_ftype fn, void *userdata)
{ tox_callback_friend_message(m, fn, userdata); }

void callbackFriendActionWrapperForC(Tox *, int32_t, uint8_t*, uint16_t, void*);
typedef void (*cb_friend_action_ftype)(Tox *, int32_t, uint8_t*, uint16_t, void*);
static void cb_friend_action_wrapper_for_go(Tox *m, cb_friend_action_ftype fn, void *userdata)
{ tox_callback_friend_action(m, fn, userdata); }

void callbackNameChangeWrapperForC(Tox *, int32_t, uint8_t*, uint16_t, void*);
typedef void (*cb_name_change_ftype)(Tox *, int32_t, uint8_t*, uint16_t, void*);
static void cb_name_change_wrapper_for_go(Tox *m, cb_name_change_ftype fn, void *userdata)
{ tox_callback_name_change(m, fn, userdata); }

void callbackStatusMessageWrapperForC(Tox *, int32_t, uint8_t*, uint16_t, void*);
typedef void (*cb_status_message_ftype)(Tox *, int32_t, uint8_t*, uint16_t, void*);
static void cb_status_message_wrapper_for_go(Tox *m, cb_status_message_ftype fn, void *userdata)
{ tox_callback_status_message(m, fn, userdata); }

void callbackUserStatusWrapperForC(Tox *, int32_t, uint8_t, void*);
typedef void (*cb_user_status_ftype)(Tox *, int32_t, uint8_t, void*);
static void cb_user_status_wrapper_for_go(Tox *m, cb_user_status_ftype fn, void *userdata)
{ tox_callback_user_status(m, fn, userdata); }

void callbackTypingChangeWrapperForC(Tox *, int32_t, uint8_t, void*);
typedef void (*cb_typing_change_ftype)(Tox *, int32_t, uint8_t, void*);
static void cb_typing_change_wrapper_for_go(Tox *m, cb_typing_change_ftype fn, void *userdata)
{ tox_callback_typing_change(m, fn, userdata); }

void callbackReadReceiptWrapperForC(Tox *, int32_t, uint32_t, void*);
typedef void (*cb_read_receipt_ftype)(Tox *, int32_t, uint32_t, void*);
static void cb_read_receipt_wrapper_for_go(Tox *m, cb_read_receipt_ftype fn, void *userdata)
{ tox_callback_read_receipt(m, fn, userdata); }

void callbackConnectionStatusWrapperForC(Tox *, int32_t, uint8_t, void*);
typedef void (*cb_connection_status_ftype)(Tox *, int32_t, uint8_t, void*);
static void cb_connection_status_wrapper_for_go(Tox *m, cb_connection_status_ftype fn, void *userdata)
{ tox_callback_connection_status(m, fn, userdata); }


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

void callbackAvatarInfoWrapperForC(Tox*, int32_t, uint8_t, uint8_t *, void *);
typedef void (*cb_avatar_info_ftype)(Tox *, int32_t, uint8_t, uint8_t *, void *);
static void cb_avatar_info_wrapper_for_go(Tox *m, cb_avatar_info_ftype fn, void *userdata)
{ tox_callback_avatar_info(m, fn, userdata); }

void callbackAvatarDataWrapperForC(Tox *, int32_t, uint8_t, uint8_t*, uint8_t*, uint32_t, void*);
typedef void (*cb_avatar_data_ftype)(Tox *, int32_t, uint8_t, uint8_t*, uint8_t*, uint32_t, void*);
static void cb_avatar_data_wrapper_for_go(Tox *m, cb_avatar_data_ftype fn, void *userdata)
{ tox_callback_avatar_data(m, fn, userdata); }

// 下面的extern行不是必须的，除非这个对应的go函数在其他的文件中，或者要在go中引用它。
// 声明go语言层的回调封装函数原型
void callbackFileSendRequestWrapperForC(Tox*, int32_t, uint8_t, uint64_t, uint8_t*, uint16_t, void*);
// 定义回调函数类型，这样定义后，能够在go语言层引用这种函数指针类型。
typedef void (*cb_file_send_request_ftype)(Tox*, int32_t, uint8_t, uint64_t, uint8_t*, uint16_t, void*);
// 定义C语言层的回调实现。
static void cb_file_send_request_wrapper_for_go(Tox *m, cb_file_send_request_ftype fn, void *userdata)
{ tox_callback_file_send_request(m, fn, userdata); }

void callbackFileControlWrapperForC(Tox*, int32_t, uint8_t, uint8_t, uint8_t, uint8_t*, uint16_t, void*);
typedef void (*cb_file_control_ftype)(Tox*, int32_t, uint8_t, uint8_t, uint8_t, uint8_t*, uint16_t, void*);
static void cb_file_control_wrapper_for_go(Tox *m, cb_file_control_ftype fn, void *userdata)
{ tox_callback_file_control(m, fn, userdata); }

void callbackFileDataWrapperForC(Tox*, int32_t, uint8_t, uint8_t*, uint16_t, void*);
typedef void (*cb_file_data_ftype)(Tox*, int32_t, uint8_t, uint8_t*, uint16_t, void*);
static void cb_file_data_wrapper_for_go(Tox *m, cb_file_data_ftype fn, void *userdata)
{ tox_callback_file_data(m, fn, userdata); }

////////////////
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

//////////
// friend cb type
type cb_friend_request_ftype func(this *Tox, publicKey *uint8, data *uint8, length uint16, userData unsafe.Pointer)
type cb_friend_message_ftype func(this *Tox, friendNumber int32, message *uint8, length uint16, userData unsafe.Pointer)
type cb_friend_action_ftype func(this *Tox, friendNumber int32, message *uint8, length uint16, userData unsafe.Pointer)
type cb_name_change_ftype func(this *Tox, friendNumber int32, newName *uint8, length uint16, userData unsafe.Pointer)
type cb_status_message_ftype func(this *Tox, friendNumber int32, newStatus *uint8, length uint16, userData unsafe.Pointer)
type cb_user_status_ftype func(this *Tox, friendNumber int32, status uint8, userData unsafe.Pointer)
type cb_typing_change_ftype func(this *Tox, friendNumber int32, isTyping uint8, userData unsafe.Pointer)
type cb_read_receipt_ftype func(this *Tox, friendNumber int32, receipt uint32, userData unsafe.Pointer)
type cb_connection_status_ftype func(this *Tox, friendNumber int32, status uint8, userData unsafe.Pointer)

// group cb type
type cb_group_invite_ftype func(this *Tox, friendNumber int32, itype uint8, data *uint8, length uint16, userData unsafe.Pointer)
type cb_group_message_ftype func(this *Tox, groupNumber int, peerNumber int, message *uint8, length uint16, userData unsafe.Pointer)
type cb_group_action_ftype func(this *Tox,  groupNumber int, peerNumber int, action *uint8, length uint16, userData unsafe.Pointer)
type cb_group_title_ftype func(this *Tox,  groupNumber int, peerNumber int, title *uint8, length uint8, userData unsafe.Pointer)
type cb_group_namelist_change_ftype func(this *Tox,  groupNumber int, peerNumber int, change uint8, userData unsafe.Pointer)

// avatar cb type
type cb_avatar_info_ftype func(this *Tox, friendNumber int32, format uint8, hash *uint8, userData unsafe.Pointer)
type cb_avatar_data_ftype func(this *Tox, friendNumber int32, format uint8, hash *uint8,
	data *uint8, datalen uint32, userData unsafe.Pointer)

// file cb type
type cb_file_send_request_ftype func(this *Tox, friendNumber int32, fileNumber uint8, fileSize uint64,
	fileName *uint8, fileNameLength uint16, userData unsafe.Pointer)
type cb_file_control_ftype func(this *Tox, friendNumber int32, recieveSend uint8, fileNumber uint8,
	controlType uint8, data *uint8, length uint16, userData unsafe.Pointer)
type cb_file_data_ftype func(this *Tox, friendNumber int32, fileNumber uint8, data *uint8,
	length uint16, userData unsafe.Pointer)

type Tox struct {
	// C.Tox
	i int32
	ix interface{}  // save C.Tox, dup
	x *C.Tox // save C.Tox
	iopts interface{} // C.Tox_Options
	opts *C.Tox_Options // C.Tox_Options

	// some callbacks, should be private
	cb_friend_request cb_friend_request_ftype
	cb_friend_request_user_data unsafe.Pointer
	cb_friend_message cb_friend_message_ftype
	cb_friend_message_user_data unsafe.Pointer
	cb_friend_action cb_friend_action_ftype
	cb_friend_action_user_data unsafe.Pointer
	cb_name_change cb_name_change_ftype
	cb_name_change_user_data unsafe.Pointer
	cb_status_message cb_status_message_ftype
	cb_status_message_user_data unsafe.Pointer
	cb_user_status cb_user_status_ftype
	cb_user_status_user_data unsafe.Pointer
	cb_typing_change cb_typing_change_ftype
	cb_typing_change_user_data unsafe.Pointer
	cb_read_receipt cb_read_receipt_ftype
	cb_read_receipt_user_data unsafe.Pointer
	cb_connection_status cb_connection_status_ftype
	cb_connection_status_user_data unsafe.Pointer
	
	cb_group_invite cb_group_invite_ftype
	cb_group_invite_user_data unsafe.Pointer
	cb_group_message cb_group_message_ftype
	cb_group_message_user_data unsafe.Pointer
	cb_group_action cb_group_action_ftype
	cb_group_action_user_data unsafe.Pointer
	cb_group_title cb_group_title_ftype
	cb_group_title_user_data unsafe.Pointer
	cb_group_namelist_change cb_group_namelist_change_ftype
	cb_group_namelist_change_user_data unsafe.Pointer
	
	cb_avatar_info cb_avatar_info_ftype
	cb_avatar_info_user_data unsafe.Pointer
	cb_avatar_data cb_avatar_data_ftype
	cb_avatar_data_user_data unsafe.Pointer
	
	cb_file_send_request cb_file_send_request_ftype
	cb_file_send_request_user_data unsafe.Pointer
	cb_file_control cb_file_control_ftype
	cb_file_control_user_data unsafe.Pointer
	cb_file_data cb_file_data_ftype
	cb_file_data_user_data unsafe.Pointer
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

//export callbackFriendRequestWrapperForC
func callbackFriendRequestWrapperForC(m *C.Tox, a0 *C.uint8_t, a1 *C.uint8_t, a2 C.uint16_t, a3 unsafe.Pointer) {
	var this = (*Tox)(a3)
	if this.cb_friend_request != nil {
		this.cb_friend_request(this, (*uint8)(a0), (*uint8)(a1), uint16(a2), this.cb_friend_request_user_data)
	}
}

func (this *Tox) CallbackFriendRequest(cbfun cb_friend_request_ftype, userData unsafe.Pointer) {
	this.cb_friend_request = cbfun
	this.cb_friend_request_user_data = userData

	var _cbfun = (C.cb_friend_request_ftype)(C.callbackFriendRequestWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_friend_request_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackFriendMessageWrapperForC
func callbackFriendMessageWrapperForC(m *C.Tox, a0 C.int32_t, a1 *C.uint8_t, a2 C.uint16_t, a3 unsafe.Pointer) {
	var this = (*Tox)(a3)
	if this.cb_friend_message != nil {
		this.cb_friend_message(this, int32(a0), (*uint8)(a1), uint16(a2), this.cb_friend_message_user_data)
	}
}

func (this *Tox) CallbackFriendMessage(cbfun cb_friend_message_ftype, userData unsafe.Pointer) {
	this.cb_friend_message = cbfun
	this.cb_friend_message_user_data = userData

	var _cbfun = (C.cb_friend_message_ftype)(C.callbackFriendMessageWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_friend_message_wrapper_for_go(this.x, _cbfun, _userData)
}


//export callbackFriendActionWrapperForC
func callbackFriendActionWrapperForC(m *C.Tox, a0 C.int32_t, a1 *C.uint8_t, a2 C.uint16_t, a3 unsafe.Pointer) {
	var this = (*Tox)(a3)
	if this.cb_friend_action != nil {
		this.cb_friend_action(this, int32(a0), (*uint8)(a1), uint16(a2), this.cb_friend_action_user_data)
	}
}

func (this *Tox) CallbackFriendAction(cbfun cb_friend_action_ftype, userData unsafe.Pointer) {
	this.cb_friend_action = cbfun
	this.cb_friend_action_user_data = userData

	var _cbfun = (C.cb_friend_action_ftype)(C.callbackFriendActionWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_friend_action_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackNameChangeWrapperForC
func callbackNameChangeWrapperForC(m *C.Tox, a0 C.int32_t, a1 *C.uint8_t, a2 C.uint16_t, a3 unsafe.Pointer) {
	var this = (*Tox)(a3)
	if this.cb_name_change != nil {
		this.cb_name_change(this, int32(a0), (*uint8)(a1), uint16(a2), this.cb_name_change_user_data)
	}
}

func (this *Tox) CallbackNameChange(cbfun cb_name_change_ftype, userData unsafe.Pointer) {
	this.cb_name_change = cbfun
	this.cb_name_change_user_data = userData

	var _cbfun = (C.cb_name_change_ftype)(C.callbackNameChangeWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_name_change_wrapper_for_go(this.x, _cbfun, _userData)
}


//export callbackStatusMessageWrapperForC
func callbackStatusMessageWrapperForC(m *C.Tox, a0 C.int32_t, a1 *C.uint8_t, a2 C.uint16_t, a3 unsafe.Pointer) {
	var this = (*Tox)(a3)
	if this.cb_status_message != nil {
		this.cb_status_message(this, int32(a0), (*uint8)(a1), uint16(a2), this.cb_status_message_user_data)
	}
}

func (this *Tox) CallbackStatusMessage(cbfun cb_status_message_ftype, userData unsafe.Pointer) {
	this.cb_status_message = cbfun
	this.cb_status_message_user_data = userData

	var _cbfun = (C.cb_status_message_ftype)(C.callbackStatusMessageWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_status_message_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackUserStatusWrapperForC
func callbackUserStatusWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 unsafe.Pointer) {
	var this = (*Tox)(a2)
	if this.cb_user_status != nil {
		this.cb_user_status(this, int32(a0), uint8(a1), this.cb_user_status_user_data)
	}
}

func (this *Tox) CallbackUserStatus(cbfun cb_user_status_ftype, userData unsafe.Pointer) {
	this.cb_user_status = cbfun
	this.cb_user_status_user_data = userData

	var _cbfun = (C.cb_user_status_ftype)(C.callbackUserStatusWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_user_status_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackTypingChangeWrapperForC
func callbackTypingChangeWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 unsafe.Pointer) {
	var this = (*Tox)(a2)
	if this.cb_typing_change != nil {
		this.cb_typing_change(this, int32(a0), uint8(a1), this.cb_typing_change_user_data)
	}
}

func (this *Tox) CallbackTypingChange(cbfun cb_typing_change_ftype, userData unsafe.Pointer) {
	this.cb_typing_change = cbfun
	this.cb_typing_change_user_data = userData

	var _cbfun = (C.cb_typing_change_ftype)(C.callbackTypingChangeWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_typing_change_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackReadReceiptWrapperForC
func callbackReadReceiptWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint32_t, a2 unsafe.Pointer) {
	var this = (*Tox)(a2)
	if this.cb_read_receipt != nil {
		this.cb_read_receipt(this, int32(a0), uint32(a1), this.cb_read_receipt_user_data)
	}
}

func (this *Tox) CallbackReadReceipt(cbfun cb_read_receipt_ftype, userData unsafe.Pointer) {
	this.cb_read_receipt = cbfun
	this.cb_read_receipt_user_data = userData

	var _cbfun = (C.cb_read_receipt_ftype)(C.callbackReadReceiptWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_read_receipt_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackConnectionStatusWrapperForC
func callbackConnectionStatusWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 unsafe.Pointer) {
	var this = (*Tox)(a2)
	if this.cb_connection_status != nil {
		this.cb_connection_status(this, int32(a0), uint8(a1), this.cb_connection_status_user_data)
	}
}

func (this *Tox) CallbackConnectionStatus(cbfun cb_connection_status_ftype, userData unsafe.Pointer) {
	this.cb_connection_status = cbfun
	this.cb_connection_status_user_data = userData

	var _cbfun = (C.cb_connection_status_ftype)(C.callbackConnectionStatusWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_connection_status_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackGroupInviteWrapperForC
func callbackGroupInviteWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_group_invite != nil {
		this.cb_group_invite(this, int32(a0), uint8(a1), (*uint8)(a2), uint16(a3), this.cb_group_invite_user_data)
	}
}

func (this *Tox) CallbackGroupInvite(cbfun cb_group_invite_ftype, userData unsafe.Pointer) {
	this.cb_group_invite = cbfun
	this.cb_group_invite_user_data = userData

	var _cbfun = (C.cb_group_invite_ftype)(C.callbackGroupInviteWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_group_invite_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackGroupMessageWrapperForC
func callbackGroupMessageWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_group_message != nil {
		this.cb_group_message(this, int(a0), int(a1), (*uint8)(a2), uint16(a3), this.cb_group_message_user_data)
	}
}

func (this *Tox) CallbackGroupMessage(cbfun cb_group_message_ftype, userData unsafe.Pointer) {
	this.cb_group_message = cbfun
	this.cb_group_message_user_data = userData

	var _cbfun = (C.cb_group_message_ftype)(C.callbackGroupMessageWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_group_message_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackGroupActionWrapperForC
func callbackGroupActionWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_group_action != nil {
		this.cb_group_action(this, int(a0), int(a1), (*uint8)(a2), uint16(a3), this.cb_group_action_user_data)
	}
}

func (this *Tox) CallbackGroupAction(cbfun cb_group_action_ftype, userData unsafe.Pointer) {
	this.cb_group_action = cbfun
	this.cb_group_action_user_data = userData

	var _cbfun = (C.cb_group_action_ftype)(C.callbackGroupActionWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_group_action_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackGroupTitleWrapperForC
func callbackGroupTitleWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint8_t, a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_group_title != nil {
		this.cb_group_title(this, int(a0), int(a1), (*uint8)(a2), uint8(a3), this.cb_group_title_user_data)
	}
}

func (this *Tox) CallbackGroupTitle(cbfun cb_group_title_ftype, userData unsafe.Pointer) {
	this.cb_group_title = cbfun
	this.cb_group_title_user_data = userData

	var _cbfun = (C.cb_group_title_ftype)(C.callbackGroupTitleWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_group_title_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackGroupNameListChangeWrapperForC
func callbackGroupNameListChangeWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 C.uint8_t, a3 unsafe.Pointer) {
	var this = (*Tox)(a3)
	if this.cb_group_namelist_change != nil {
		this.cb_group_namelist_change(this, int(a0), int(a1), uint8(a2), this.cb_group_namelist_change_user_data)
	}
}

func (this *Tox) CallbackGroupNameListChange(cbfun cb_group_namelist_change_ftype, userData unsafe.Pointer) {
	this.cb_group_namelist_change = cbfun
	this.cb_group_namelist_change_user_data = userData

	var _cbfun = (C.cb_group_namelist_change_ftype)(C.callbackGroupNameListChangeWrapperForC)
	var _userData = unsafe.Pointer(this)

	C.cb_group_namelist_change_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackAvatarInfoWrapperForC
func callbackAvatarInfoWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 *C.uint8_t, a3 unsafe.Pointer) {
	var this = (*Tox)(a3)
	
	if this.cb_avatar_info != nil {
		this.cb_avatar_info(this, int32(a0), uint8(a1), (*uint8)(a2), this.cb_avatar_info_user_data)
	}
}

func (this *Tox) CallbackAvatarInfo(cbfun cb_avatar_info_ftype, userData unsafe.Pointer) {
	this.cb_avatar_info = cbfun
	this.cb_avatar_info_user_data = userData

	var _cbfun = (C.cb_avatar_info_ftype)(C.callbackAvatarInfoWrapperForC)
	var _userData = unsafe.Pointer(this)
	
	C.cb_avatar_info_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackAvatarDataWrapperForC
func callbackAvatarDataWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 *C.uint8_t, a3 *C.uint8_t,
	a4 C.uint32_t, a5 unsafe.Pointer) {
	var this = (*Tox)(a5)
	if this.cb_avatar_data != nil {
		this.cb_avatar_data(this, int32(a0), uint8(a1), (*uint8)(a2), (*uint8)(a3), uint32(a4),
			this.cb_avatar_data_user_data)
	}
}

func (this *Tox) CallbackAvatarData(cbfun cb_avatar_data_ftype, userData unsafe.Pointer) {
	this.cb_avatar_data = cbfun
	this.cb_avatar_data_user_data = userData

	var _cbfun = (C.cb_avatar_data_ftype)(C.callbackAvatarDataWrapperForC)
	var _userData = unsafe.Pointer(this)
	C.cb_avatar_data_wrapper_for_go(this.x, _cbfun, _userData)
}


// 包内部函数
//export callbackFileSendRequestWrapperForC
func callbackFileSendRequestWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 C.uint64_t,
	a3 *C.uint8_t, a4 C.uint16_t, a5 unsafe.Pointer) {
	var this = (*Tox)(a5)
	log.Println("called from c code", this)
	log.Println(m, a0, a1, a2, a3, a4, a5)
	if this.cb_file_send_request != nil {
		this.cb_file_send_request(this, int32(a0), uint8(a1), uint64(a2), (*uint8)(a3),
			uint16(a4), this.cb_file_send_request_user_data)
	}
}

func (this *Tox) CallbackFileSendRequest(cbfun cb_file_send_request_ftype, userData unsafe.Pointer) {
	this.cb_file_send_request = cbfun
	this.cb_file_send_request_user_data = userData;
	var _cbfun = (C.cb_file_send_request_ftype)(unsafe.Pointer(C.callbackFileSendRequestWrapperForC))
	var _userData = unsafe.Pointer(this)
	
	C.cb_file_send_request_wrapper_for_go(this.x, _cbfun, _userData);
}

//export callbackFileControlWrapperForC
func callbackFileControlWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 C.uint8_t, a3 C.uint8_t,
	a4 *C.uint8_t, a5 C.uint16_t, a6 unsafe.Pointer) {
	var this = (*Tox)(a6)
	if this.cb_file_control != nil {
		this.cb_file_control(this, int32(a0), uint8(a1), uint8(a2), uint8(a3),
			(*uint8)(a4), uint16(a5), this.cb_file_control_user_data)
	}
}

func (this *Tox) CallbackFileControl(cbfun cb_file_control_ftype, userData unsafe.Pointer) {
	this.cb_file_control = cbfun
	this.cb_file_control_user_data = userData
	var _cbfun = (C.cb_file_control_ftype)(unsafe.Pointer(C.callbackFileControlWrapperForC))
	var _userData = unsafe.Pointer(this)

	C.cb_file_control_wrapper_for_go(this.x, _cbfun, _userData)
}

//export callbackFileDataWrapperForC
func callbackFileDataWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 *C.uint8_t, a3 C.uint16_t,
	a4 unsafe.Pointer) {
	var this = (*Tox)(a4)
	if this.cb_file_data != nil {
		this.cb_file_data(this, int32(a0), uint8(a1), (*uint8)(a2), uint16(a3), this.cb_file_data_user_data)
	}
}

func (this *Tox) CallbackFileData(cbfun cb_file_data_ftype, userData unsafe.Pointer) {
	this.cb_file_data = cbfun
	this.cb_file_data_user_data = userData
	var _cbfun = (C.cb_file_data_ftype)(unsafe.Pointer(C.callbackFileDataWrapperForC))
	var _userData = unsafe.Pointer(this)
	
	C.cb_file_data_wrapper_for_go(this.x, _cbfun, _userData)
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

