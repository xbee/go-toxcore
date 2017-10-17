package tox

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>

void callbackConferenceInviteWrapperForC(Tox*, uint32_t, TOX_CONFERENCE_TYPE, uint8_t *, size_t, void *);
typedef void (*cb_conference_invite_ftype)(Tox *, uint32_t, TOX_CONFERENCE_TYPE, const uint8_t *, size_t, void *);
static void cb_conference_invite_wrapper_for_go(Tox *m, cb_conference_invite_ftype fn, void *userdata)
{ tox_callback_conference_invite(m, fn); }

void callbackConferenceMessageWrapperForC(Tox *, uint32_t, uint32_t, TOX_MESSAGE_TYPE, int8_t *, size_t, void *);
typedef void (*cb_conference_message_ftype)(Tox *, uint32_t, uint32_t, TOX_MESSAGE_TYPE, const uint8_t *, size_t, void *);
static void cb_conference_message_wrapper_for_go(Tox *m, cb_conference_message_ftype fn, void *userdata)
{ tox_callback_conference_message(m, fn); }

// void callbackConferenceActionWrapperForC(Tox*, uint32_t, uint32_t, uint8_t*, size_t, void*);
// typedef void (*cb_conference_action_ftype)(Tox*, uint32_t, uint32_t, const uint8_t*, size_t, void*);
// static void cb_conference_action_wrapper_for_go(Tox *m, cb_conference_action_ftype fn, void *userdata)
// { tox_callback_conference_message(m, fn); }

void callbackConferenceTitleWrapperForC(Tox*, uint32_t, uint32_t, uint8_t*, size_t, void*);
typedef void (*cb_conference_title_ftype)(Tox*, uint32_t, uint32_t, const uint8_t*, size_t, void*);
static void cb_conference_title_wrapper_for_go(Tox *m, cb_conference_title_ftype fn, void *userdata)
{ tox_callback_conference_title(m, fn); }

void callbackConferenceNameListChangeWrapperForC(Tox*, uint32_t, uint32_t, TOX_CONFERENCE_STATE_CHANGE, void*);
typedef void (*cb_conference_namelist_change_ftype)(Tox*, uint32_t, uint32_t, TOX_CONFERENCE_STATE_CHANGE, void*);
static void cb_conference_namelist_change_wrapper_for_go(Tox *m, cb_conference_namelist_change_ftype fn, void *userdata)
{ tox_callback_conference_namelist_change(m, fn); }

// fix nouse compile warning
static inline void fixnousetoxgroup() {
    cb_conference_invite_wrapper_for_go(NULL, NULL, NULL);
    cb_conference_message_wrapper_for_go(NULL, NULL, NULL);
    // cb_conference_action_wrapper_for_go(NULL, NULL, NULL);
    cb_conference_title_wrapper_for_go(NULL, NULL, NULL);
    cb_conference_namelist_change_wrapper_for_go(NULL, NULL, NULL);
}

*/
import "C"
import (
	"bytes"
	"encoding/hex"
	"errors"
	"strings"
	"unsafe"
)

// conference callback type
type cb_conference_invite_ftype func(this *Tox, friendNumber uint32, itype uint8, data []byte, userData interface{})
type cb_conference_message_ftype func(this *Tox, groupNumber uint32, peerNumber uint32, message string, userData interface{})

type cb_conference_action_ftype func(this *Tox, groupNumber uint32, peerNumber uint32, action string, userData interface{})
type cb_conference_title_ftype func(this *Tox, groupNumber uint32, peerNumber uint32, title string, userData interface{})
type cb_conference_namelist_change_ftype func(this *Tox, groupNumber uint32, peerNumber uint32, change uint8, userData interface{})

// tox_callback_conference_***

//export callbackConferenceInviteWrapperForC
func callbackConferenceInviteWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.TOX_CONFERENCE_TYPE, a2 *C.uint8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_conference_invites {
		cbfn := *(*cb_conference_invite_ftype)(cbfni)
		data := C.GoBytes((unsafe.Pointer)(a2), C.int(a3))
		this.beforeCallback()
		cbfn(this, uint32(a0), uint8(a1), data, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackConferenceInvite(cbfn cb_conference_invite_ftype, userData interface{}) {
	this.CallbackConferenceInviteAdd(cbfn, userData)
}
func (this *Tox) CallbackConferenceInviteAdd(cbfn cb_conference_invite_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_invites[cbfnp]; ok {
		return
	}
	this.cb_conference_invites[cbfnp] = userData

	var _cbfn = (C.cb_conference_invite_ftype)(C.callbackConferenceInviteWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_conference_invite_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackConferenceMessageWrapperForC
func callbackConferenceMessageWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, mtype C.TOX_MESSAGE_TYPE, a2 *C.int8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	if int(mtype) == MESSAGE_TYPE_NORMAL {
		for cbfni, ud := range this.cb_conference_messages {
			cbfn := *(*cb_conference_message_ftype)(cbfni)
			message := C.GoStringN((*C.char)((*C.int8_t)(a2)), C.int(a3))
			this.beforeCallback()
			cbfn(this, uint32(a0), uint32(a1), message, ud)
			this.afterCallback()
		}
	} else {
		for cbfni, ud := range this.cb_conference_actions {
			cbfn := *(*cb_conference_action_ftype)(cbfni)
			message := C.GoStringN((*C.char)((*C.int8_t)(a2)), C.int(a3))
			this.beforeCallback()
			cbfn(this, uint32(a0), uint32(a1), message, ud)
			this.afterCallback()
		}
	}
}

func (this *Tox) CallbackConferenceMessage(cbfn cb_conference_message_ftype, userData interface{}) {
	this.CallbackConferenceMessageAdd(cbfn, userData)
}
func (this *Tox) CallbackConferenceMessageAdd(cbfn cb_conference_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_messages[cbfnp]; ok {
		return
	}
	this.cb_conference_messages[cbfnp] = userData

	if !this.cb_conference_message_setted {
		this.cb_conference_message_setted = true

		var _cbfn = (C.cb_conference_message_ftype)(C.callbackConferenceMessageWrapperForC)
		var _userData unsafe.Pointer = nil

		C.cb_conference_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
	}
}

func (this *Tox) CallbackConferenceAction(cbfn cb_conference_action_ftype, userData interface{}) {
	this.CallbackConferenceActionAdd(cbfn, userData)
}
func (this *Tox) CallbackConferenceActionAdd(cbfn cb_conference_action_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_actions[cbfnp]; ok {
		return
	}
	this.cb_conference_actions[cbfnp] = userData

	if !this.cb_conference_message_setted {
		this.cb_conference_message_setted = true
		var _cbfn = (C.cb_conference_message_ftype)(C.callbackConferenceMessageWrapperForC)
		var _userData unsafe.Pointer = nil

		C.cb_conference_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
	}
}

//export callbackConferenceTitleWrapperForC
func callbackConferenceTitleWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, a2 *C.uint8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_conference_titles {
		cbfn := *(*cb_conference_title_ftype)(cbfni)
		title := C.GoStringN((*C.char)((unsafe.Pointer)(a2)), C.int(a3))
		this.beforeCallback()
		cbfn(this, uint32(a0), uint32(a1), title, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackConferenceTitle(cbfn cb_conference_title_ftype, userData interface{}) {
	this.CallbackConferenceTitleAdd(cbfn, userData)
}
func (this *Tox) CallbackConferenceTitleAdd(cbfn cb_conference_title_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_titles[cbfnp]; ok {
		return
	}
	this.cb_conference_titles[cbfnp] = userData

	var _cbfn = (C.cb_conference_title_ftype)(C.callbackConferenceTitleWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_conference_title_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackConferenceNameListChangeWrapperForC
func callbackConferenceNameListChangeWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, a2 C.TOX_CONFERENCE_STATE_CHANGE, a3 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_conference_namelist_changes {
		cbfn := *(*cb_conference_namelist_change_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, uint32(a0), uint32(a1), uint8(a2), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackConferenceNameListChange(cbfn cb_conference_namelist_change_ftype, userData interface{}) {
	this.CallbackConferenceNameListChangeAdd(cbfn, userData)
}
func (this *Tox) CallbackConferenceNameListChangeAdd(cbfn cb_conference_namelist_change_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_conference_namelist_changes[cbfnp]; ok {
		return
	}
	this.cb_conference_namelist_changes[cbfnp] = userData

	var _cbfn = (C.cb_conference_namelist_change_ftype)(C.callbackConferenceNameListChangeWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_conference_namelist_change_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

// methods tox_conference_*
func (this *Tox) ConferenceNew() (uint32, error) {
	this.lock()
	defer this.unlock()

	var cerr C.TOX_ERR_CONFERENCE_NEW
	r := C.tox_conference_new(this.toxcore, &cerr)
	if r == C.UINT32_MAX {
		return uint32(r), toxerrf("add group chat failed: %d", cerr)
	}
	return uint32(r), nil
}

func (this *Tox) ConferenceDelete(groupNumber uint32) (int, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.uint32_t(groupNumber)
	var cerr C.TOX_ERR_CONFERENCE_DELETE
	r := C.tox_conference_delete(this.toxcore, _gn, &cerr)
	if bool(r) == false {
		return 1, toxerrf("delete group chat failed:%d", cerr)
	}
	return 0, nil
}

func (this *Tox) ConferencePeerGetName(groupNumber uint32, peerNumber uint32) (string, error) {
	var _gn = C.uint32_t(groupNumber)
	var _pn = C.uint32_t(peerNumber)
	var _name [MAX_NAME_LENGTH]byte

	var cerr C.TOX_ERR_CONFERENCE_PEER_QUERY
	r := C.tox_conference_peer_get_name(this.toxcore, _gn, _pn, (*C.uint8_t)(&_name[0]), &cerr)
	if r == false {
		return "", toxerrf("get peer name failed: %d", cerr)
	}

	return string(_name[:bytes.IndexRune(_name[:], 0)]), nil
}

func (this *Tox) ConferencePeerGetPublicKey(groupNumber uint32, peerNumber int) (string, error) {
	var _gn = C.uint32_t(groupNumber)
	var _pn = C.uint32_t(peerNumber)
	var _pubkey [PUBLIC_KEY_SIZE]byte

	var cerr C.TOX_ERR_CONFERENCE_PEER_QUERY
	r := C.tox_conference_peer_get_public_key(this.toxcore, _gn, _pn, (*C.uint8_t)(&_pubkey[0]), &cerr)
	if r == false {
		return "", toxerrf("get pubkey failed: %d", cerr)
	}

	pubkey := strings.ToUpper(hex.EncodeToString(_pubkey[:]))
	return pubkey, nil
}

func (this *Tox) ConferenceInvite(friendNumber uint32, groupNumber uint32) (int, error) {
	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)
	var _gn = C.uint32_t(groupNumber)

	// if give a friendNumber which not exists,
	// the tox_invite_friend has a strange behaive: cause other tox_* call failed
	// and the call will return true, but only strange thing accurs
	// so just precheck the friendNumber and then go
	if !this.FriendExists(friendNumber) {
		return -1, toxerrf("friend not exists: %d", friendNumber)
	}

	var cerr C.TOX_ERR_CONFERENCE_INVITE
	r := C.tox_conference_invite(this.toxcore, _fn, _gn, &cerr)
	if r == false {
		return 0, toxerrf("conference invite failed: %d", cerr)
	}
	return 1, nil
}

func (this *Tox) ConferenceJoin(friendNumber uint32, data []byte) (uint32, error) {
	this.lock()
	defer this.unlock()

	var length = len(data)

	if data == nil || length < 10 {
		return 0, errors.New("invalid data")
	}

	var _fn = C.uint32_t(friendNumber)
	var _length = C.size_t(length)

	var cerr C.TOX_ERR_CONFERENCE_JOIN
	r := C.tox_conference_join(this.toxcore, _fn, (*C.uint8_t)(&data[0]), _length, &cerr)
	if r == C.UINT32_MAX {
		return uint32(r), toxerrf("join group chat failed: %d", cerr)
	}
	return uint32(r), nil
}

func (this *Tox) ConferenceSendMessage(groupNumber uint32, mtype int, message string) (int, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.uint32_t(groupNumber)
	var _message = []byte(message)
	var _length = C.size_t(len(message))

	switch mtype {
	case MESSAGE_TYPE_NORMAL:
	case MESSAGE_TYPE_ACTION:
	default:
		return 0, toxerrf("Invalid message type: %d", mtype)
	}

	var cerr C.TOX_ERR_CONFERENCE_SEND_MESSAGE
	r := C.tox_conference_send_message(this.toxcore, _gn, (C.TOX_MESSAGE_TYPE)(mtype), (*C.uint8_t)(&_message[0]), _length, &cerr)
	if r == false {
		return 0, toxerrf("group send message failed: %d", cerr)
	}
	return 1, nil
}

func (this *Tox) ConferenceSetTitle(groupNumber uint32, title string) (int, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.uint32_t(groupNumber)
	var _title = []byte(title)
	var _length = C.size_t(len(title))

	var cerr C.TOX_ERR_CONFERENCE_TITLE
	r := C.tox_conference_set_title(this.toxcore, _gn, (*C.uint8_t)(&_title[0]), _length, &cerr)
	if r == false {
		if len(title) > MAX_NAME_LENGTH {
			return 0, errors.New("title too long")
		}
		return 0, toxerrf("set title failed:%d", cerr)
	}
	return 1, nil
}

func (this *Tox) ConferenceGetTitle(groupNumber uint32) (string, error) {
	var _gn = C.uint32_t(groupNumber)
	var _title [MAX_NAME_LENGTH]byte

	r := C.tox_conference_get_title(this.toxcore, _gn, (*C.uint8_t)(&_title[0]), nil)
	if r == false {
		return "", errors.New("get title failed")
	}
	return string(_title[:bytes.IndexRune(_title[:], 0)]), nil
}

func (this *Tox) ConferencePeerNumberIsOurs(groupNumber uint32, peerNumber uint32) bool {
	var _gn = C.uint32_t(groupNumber)
	var _pn = C.uint32_t(peerNumber)

	r := C.tox_conference_peer_number_is_ours(this.toxcore, _gn, _pn, nil)
	return bool(r)
}

func (this *Tox) ConferencePeerCount(groupNumber uint32) uint32 {
	var _gn = C.uint32_t(groupNumber)

	r := C.tox_conference_peer_count(this.toxcore, _gn, nil)
	return uint32(r)
}

// extra combined api
func (this *Tox) ConferenceGetNames(groupNumber uint32) []string {
	peerCount := this.ConferencePeerCount(groupNumber)
	vec := make([]string, peerCount)
	if peerCount == 0 {
		return vec
	}

	for idx := uint32(0); idx < peerCount; idx++ {
		pname, err := this.ConferencePeerGetName(groupNumber, idx)
		if err != nil {
			return vec[0:0]
		}
		vec[idx] = pname
	}

	return vec
}

func (this *Tox) ConferenceGetPeerPubkeys(groupNumber uint32) []string {
	vec := make([]string, 0)
	peerCount := this.ConferencePeerCount(groupNumber)
	for peerNumber := 0; peerNumber < C.UINT32_MAX; peerNumber++ {
		pubkey, err := this.ConferencePeerGetPublicKey(groupNumber, peerNumber)
		if err != nil {
		} else {
			vec = append(vec, pubkey)
		}
		if uint32(len(vec)) >= peerCount {
			break
		}
	}
	return vec
}

func (this *Tox) ConferenceGetPeers(groupNumber uint32) map[int]string {
	vec := make(map[int]string, 0)
	peerCount := this.ConferencePeerCount(groupNumber)
	for peerNumber := 0; peerNumber < C.UINT32_MAX; peerNumber++ {
		pubkey, err := this.ConferencePeerGetPublicKey(groupNumber, peerNumber)
		if err != nil {
		} else {
			vec[peerNumber] = pubkey
		}
		if uint32(len(vec)) >= peerCount {
			break
		}
	}
	return vec
}

func (this *Tox) ConferenceGetChatlistSize() uint32 {
	r := C.tox_conference_get_chatlist_size(this.toxcore)
	return uint32(r)
}

func (this *Tox) ConferenceGetChatlist() []int32 {
	var sz uint32 = this.ConferenceGetChatlistSize()
	vec := make([]int32, sz)
	if sz == 0 {
		return vec
	}

	vec_p := unsafe.Pointer(&vec[0])
	C.tox_conference_get_chatlist(this.toxcore, (*C.uint32_t)(vec_p))
	return vec
}

func (this *Tox) ConferenceGetType(groupNumber uint32) (int, error) {
	var _gn = C.uint32_t(groupNumber)

	r := C.tox_conference_get_type(this.toxcore, _gn, nil)
	if int(r) == -1 {
		return int(r), errors.New("get type failed")
	}
	return int(r), nil
}
