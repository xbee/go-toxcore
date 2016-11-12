package tox

/*
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>

void callbackGroupInviteWrapperForC(Tox*, int32_t, uint8_t, uint8_t *, uint16_t, void *);
typedef void (*cb_group_invite_ftype)(Tox *, int32_t, uint8_t, const uint8_t *, uint16_t, void *);
static void cb_group_invite_wrapper_for_go(Tox *m, cb_group_invite_ftype fn, void *userdata)
{ tox_callback_group_invite(m, fn, userdata); }

void callbackGroupMessageWrapperForC(Tox *, int, int, int8_t *, uint16_t, void *);
typedef void (*cb_group_message_ftype)(Tox *, int, int, const uint8_t *, uint16_t, void *);
static void cb_group_message_wrapper_for_go(Tox *m, cb_group_message_ftype fn, void *userdata)
{ tox_callback_group_message(m, fn, userdata); }

void callbackGroupActionWrapperForC(Tox*, int, int, uint8_t*, uint16_t, void*);
typedef void (*cb_group_action_ftype)(Tox*, int, int, const uint8_t*, uint16_t, void*);
static void cb_group_action_wrapper_for_go(Tox *m, cb_group_action_ftype fn, void *userdata)
{ tox_callback_group_action(m, fn, userdata); }

void callbackGroupTitleWrapperForC(Tox*, int, int, uint8_t*, uint8_t, void*);
typedef void (*cb_group_title_ftype)(Tox*, int, int, const uint8_t*, uint8_t, void*);
static void cb_group_title_wrapper_for_go(Tox *m, cb_group_title_ftype fn, void *userdata)
{ tox_callback_group_title(m, fn, userdata); }

void callbackGroupNameListChangeWrapperForC(Tox*, int, int, uint8_t, void*);
typedef void (*cb_group_namelist_change_ftype)(Tox*, int, int, uint8_t, void*);
static void cb_group_namelist_change_wrapper_for_go(Tox *m, cb_group_namelist_change_ftype fn, void *userdata)
{ tox_callback_group_namelist_change(m, fn, userdata); }

// fix nouse compile warning
static inline void fixnousetoxgroup() {
    cb_group_invite_wrapper_for_go(NULL, NULL, NULL);
    cb_group_message_wrapper_for_go(NULL, NULL, NULL);
    cb_group_action_wrapper_for_go(NULL, NULL, NULL);
    cb_group_title_wrapper_for_go(NULL, NULL, NULL);
    cb_group_namelist_change_wrapper_for_go(NULL, NULL, NULL);
}

*/
import "C"
import "unsafe"

import (
	"encoding/hex"
	"errors"
	"strings"
)

// group callback type
type cb_group_invite_ftype func(this *Tox, friendNumber uint32, itype uint8, data *uint8, length uint16, userData interface{})
type cb_group_message_ftype func(this *Tox, groupNumber int, peerNumber int, message string, userData interface{})
type cb_group_action_ftype func(this *Tox, groupNumber int, peerNumber int, action *uint8, length uint16, userData interface{})
type cb_group_title_ftype func(this *Tox, groupNumber int, peerNumber int, title *uint8, length uint8, userData interface{})
type cb_group_namelist_change_ftype func(this *Tox, groupNumber int, peerNumber int, change uint8, userData interface{})

// tox_callback_group_***

//export callbackGroupInviteWrapperForC
func callbackGroupInviteWrapperForC(m *C.Tox, a0 C.int32_t, a1 C.uint8_t, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.Get(m)
	if this.cb_group_invite != nil {
		this.cb_group_invite(this, uint32(a0), uint8(a1), (*uint8)(a2), uint16(a3), this.cb_group_invite_user_data)
	}
}

func (this *Tox) CallbackGroupInvite(cbfn cb_group_invite_ftype, userData interface{}) {
	this.cb_group_invite = cbfn
	this.cb_group_invite_user_data = userData

	var _cbfn = (C.cb_group_invite_ftype)(C.callbackGroupInviteWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_group_invite_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupMessageWrapperForC
func callbackGroupMessageWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.int8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.Get(m)
	if this.cb_group_message != nil {
		message := C.GoStringN((*C.char)((*C.int8_t)(a2)), C.int(a3))
		this.cb_group_message(this, int(a0), int(a1), message, this.cb_group_message_user_data)
	}
}

func (this *Tox) CallbackGroupMessage(cbfn cb_group_message_ftype, userData interface{}) {
	this.cb_group_message = cbfn
	this.cb_group_message_user_data = userData

	var _cbfn = (C.cb_group_message_ftype)(C.callbackGroupMessageWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_group_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupActionWrapperForC
func callbackGroupActionWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.Get(m)
	if this.cb_group_action != nil {
		this.cb_group_action(this, int(a0), int(a1), (*uint8)(a2), uint16(a3), this.cb_group_action_user_data)
	}
}

func (this *Tox) CallbackGroupAction(cbfn cb_group_action_ftype, userData interface{}) {
	this.cb_group_action = cbfn
	this.cb_group_action_user_data = userData

	var _cbfn = (C.cb_group_action_ftype)(C.callbackGroupActionWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_group_action_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupTitleWrapperForC
func callbackGroupTitleWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint8_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.Get(m)
	if this.cb_group_title != nil {
		this.cb_group_title(this, int(a0), int(a1), (*uint8)(a2), uint8(a3), this.cb_group_title_user_data)
	}
}

func (this *Tox) CallbackGroupTitle(cbfn cb_group_title_ftype, userData interface{}) {
	this.cb_group_title = cbfn
	this.cb_group_title_user_data = userData

	var _cbfn = (C.cb_group_title_ftype)(C.callbackGroupTitleWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_group_title_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupNameListChangeWrapperForC
func callbackGroupNameListChangeWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 C.uint8_t, a3 unsafe.Pointer) {
	var this = cbUserDatas.Get(m)
	if this.cb_group_namelist_change != nil {
		this.cb_group_namelist_change(this, int(a0), int(a1), uint8(a2), this.cb_group_namelist_change_user_data)
	}
}

func (this *Tox) CallbackGroupNameListChange(cbfn cb_group_namelist_change_ftype, userData interface{}) {
	this.cb_group_namelist_change = cbfn
	this.cb_group_namelist_change_user_data = userData

	var _cbfn = (C.cb_group_namelist_change_ftype)(C.callbackGroupNameListChangeWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_group_namelist_change_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

func (this *Tox) AddGroupChat() (int, error) {
	r := C.tox_add_groupchat(this.toxcore)
	if int(r) == -1 {
		return int(r), errors.New("add group chat failed")
	}
	return int(r), nil
}

func (this *Tox) DelGroupChat(groupNumber int) (int, error) {
	var _gn = C.int(groupNumber)

	r := C.tox_del_groupchat(this.toxcore, _gn)
	if int(r) == -1 {
		return int(r), errors.New("delete group chat failed")
	}
	return int(r), nil
}

func (this *Tox) GroupPeerName(groupNumber int, peerNumber int) (string, error) {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)
	var _name = (*C.char)(C.calloc(1, C.size_t(MAX_NAME_LENGTH)))
	defer C.free(unsafe.Pointer(_name))

	r := C.tox_group_peername(this.toxcore, _gn, _pn, char2uint8(_name))
	if r == -1 {
		return "", errors.New("get peer name failed")
	}
	name := C.GoString(_name)
	return name, nil
}

func (this *Tox) GroupPeerPubkey(groupNumber int, peerNumber int) (string, error) {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)
	var _pubkey *C.char = (*C.char)(C.calloc(1, C.size_t(PUBLIC_KEY_SIZE)))
	defer C.free(unsafe.Pointer(_pubkey))

	r := C.tox_group_peer_pubkey(this.toxcore, _gn, _pn, char2uint8(_pubkey))
	if r == C.int(-1) {
		return "", errors.New("get pubkey failed")
	}

	pubkey := hex.EncodeToString(C.GoBytes(unsafe.Pointer(_pubkey), C.int(PUBLIC_KEY_SIZE)))
	pubkey = strings.ToUpper(pubkey)
	return pubkey, nil
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
	if int(r) == -1 {
		return int(r), errors.New("join group chat failed")
	}
	return int(r), nil
}

func (this *Tox) GroupActionSend(groupNumber int, action string) (int, error) {
	var _gn = C.int(groupNumber)
	var _action = C.CString(action)
	defer C.free(unsafe.Pointer(_action))
	var _length = C.uint16_t(len(action))

	r := C.tox_group_action_send(this.toxcore, _gn, char2uint8(_action), _length)
	if int(r) == -1 {
		return int(r), errors.New("group action failed")
	}
	return int(r), nil
}

func (this *Tox) GroupMessageSend(groupNumber int, message string) (int, error) {
	var _gn = C.int(groupNumber)
	var _message = C.CString(message)
	defer C.free(unsafe.Pointer(_message))
	var _length = C.uint16_t(len(message))

	r := C.tox_group_message_send(this.toxcore, _gn, char2uint8(_message), _length)
	if int(r) == -1 {
		return int(r), errors.New("group send message failed")
	}
	return int(r), nil
}

func (this *Tox) GroupSetTitle(groupNumber int, title string) (int, error) {
	var _gn = C.int(groupNumber)
	var _title = C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	var _length = C.uint8_t(len(title))

	r := C.tox_group_set_title(this.toxcore, _gn, char2uint8(_title), _length)
	if int(r) == -1 {
		if len(title) > MAX_NAME_LENGTH {
			return int(r), errors.New("title too long")
		}
		return int(r), errors.New("set title failed")
	}
	return int(r), nil
}

func (this *Tox) GroupGetTitle(groupNumber int) (string, error) {
	var _gn = C.int(groupNumber)
	var _title = (*C.char)(C.calloc(1, C.size_t(MAX_NAME_LENGTH)))
	defer C.free(unsafe.Pointer(_title))
	var _maxlen = C.uint32_t(MAX_NAME_LENGTH)

	r := C.tox_group_get_title(this.toxcore, _gn, char2uint8(_title), _maxlen)
	if r == -1 {
		return "", errors.New("get title failed")
	}
	title := C.GoString(_title)
	return title, nil
}

func (this *Tox) GroupPeerNumberIsOurs(groupNumber int, peerNumber int) bool {
	var _gn = C.int(groupNumber)
	var _pn = C.int(peerNumber)

	r := C.tox_group_peernumber_is_ours(this.toxcore, _gn, _pn)
	return uint(r) == 1
}

func (this *Tox) GroupNumberPeers(groupNumber int) int {
	var _gn = C.int(groupNumber)

	r := C.tox_group_number_peers(this.toxcore, _gn)
	return int(r)
}

func (this *Tox) GroupGetNames(groupNumber int) []string {
	peerCount := this.GroupNumberPeers(groupNumber)
	vec := make([]string, peerCount)

	lengths := make([]uint16, peerCount)
	names := make([]byte, peerCount*MAX_NAME_LENGTH)
	clengths := (*C.uint16_t)(&lengths[0])
	cnames := (*[MAX_NAME_LENGTH]C.uint8_t)((unsafe.Pointer)(&names[0]))

	r := C.tox_group_get_names(this.toxcore, C.int(groupNumber),
		cnames, clengths, C.uint16_t(peerCount))
	if int(r) == -1 {
		return []string{}
	}

	for idx := 0; idx < peerCount; idx++ {
		len := int(lengths[idx])
		name := names[idx*MAX_NAME_LENGTH : (idx*MAX_NAME_LENGTH + len)]
		vec[idx] = string(name)
	}

	return vec
}

func (this *Tox) GroupGetPeerPubkeys(groupNumber int) []string {
	vec := make([]string, 0)
	peerCount := this.GroupNumberPeers(groupNumber)
	maxcnt := 65536
	for peerNumber := 0; peerNumber < maxcnt; peerNumber++ {
		pubkey, err := this.GroupPeerPubkey(groupNumber, peerNumber)
		if err != nil {
		} else {
			vec = append(vec, pubkey)
		}
		if len(vec) >= peerCount {
			break
		}
	}
	return vec
}

func (this *Tox) GroupGetPeers(groupNumber int) map[int]string {
	vec := make(map[int]string, 0)
	peerCount := this.GroupNumberPeers(groupNumber)
	maxcnt := 65536
	for peerNumber := 0; peerNumber < maxcnt; peerNumber++ {
		pubkey, err := this.GroupPeerPubkey(groupNumber, peerNumber)
		if err != nil {
		} else {
			vec[peerNumber] = pubkey
		}
		if len(vec) >= peerCount {
			break
		}
	}
	return vec
}

func (this *Tox) CountChatList() uint32 {
	r := C.tox_count_chatlist(this.toxcore)
	return uint32(r)
}

func (this *Tox) GetChatList() []int32 {
	var sz uint32 = this.CountChatList()
	vec := make([]int32, sz)
	vec_p := unsafe.Pointer(&vec[0])
	osz := C.tox_get_chatlist(this.toxcore, (*C.int32_t)(vec_p), C.uint32_t(sz))
	if osz == 0 {
		return vec[0:0]
	}
	return vec
}

func (this *Tox) GroupGetType(groupNumber int) (int, error) {
	var _gn = C.int(groupNumber)

	r := C.tox_group_get_type(this.toxcore, _gn)
	if int(r) == -1 {
		return int(r), errors.New("get type failed")
	}
	return int(r), nil
}
