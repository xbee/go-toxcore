package tox

/*
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>

void callbackGroupInviteWrapperForC(Tox*, uint32_t, TOX_CONFERENCE_TYPE, uint8_t *, size_t, void *);
typedef void (*cb_group_invite_ftype)(Tox *, uint32_t, TOX_CONFERENCE_TYPE, const uint8_t *, size_t, void *);
static void cb_group_invite_wrapper_for_go(Tox *m, cb_group_invite_ftype fn, void *userdata)
{ tox_callback_conference_invite(m, fn); }

void callbackGroupMessageWrapperForC(Tox *, uint32_t, uint32_t, TOX_MESSAGE_TYPE, int8_t *, size_t, void *);
typedef void (*cb_group_message_ftype)(Tox *, uint32_t, uint32_t, TOX_MESSAGE_TYPE, const uint8_t *, size_t, void *);
static void cb_group_message_wrapper_for_go(Tox *m, cb_group_message_ftype fn, void *userdata)
{ tox_callback_conference_message(m, fn); }

// void callbackGroupActionWrapperForC(Tox*, uint32_t, uint32_t, uint8_t*, size_t, void*);
// typedef void (*cb_group_action_ftype)(Tox*, uint32_t, uint32_t, const uint8_t*, size_t, void*);
// static void cb_group_action_wrapper_for_go(Tox *m, cb_group_action_ftype fn, void *userdata)
// { tox_callback_conference_message(m, fn); }

void callbackGroupTitleWrapperForC(Tox*, uint32_t, uint32_t, uint8_t*, size_t, void*);
typedef void (*cb_group_title_ftype)(Tox*, uint32_t, uint32_t, const uint8_t*, size_t, void*);
static void cb_group_title_wrapper_for_go(Tox *m, cb_group_title_ftype fn, void *userdata)
{ tox_callback_conference_title(m, fn); }

void callbackGroupNameListChangeWrapperForC(Tox*, uint32_t, uint32_t, TOX_CONFERENCE_STATE_CHANGE, void*);
typedef void (*cb_group_namelist_change_ftype)(Tox*, uint32_t, uint32_t, TOX_CONFERENCE_STATE_CHANGE, void*);
static void cb_group_namelist_change_wrapper_for_go(Tox *m, cb_group_namelist_change_ftype fn, void *userdata)
{ tox_callback_conference_namelist_change(m, fn); }

// fix nouse compile warning
static inline void fixnousetoxgroup() {
    cb_group_invite_wrapper_for_go(NULL, NULL, NULL);
    cb_group_message_wrapper_for_go(NULL, NULL, NULL);
    // cb_group_action_wrapper_for_go(NULL, NULL, NULL);
    cb_group_title_wrapper_for_go(NULL, NULL, NULL);
    cb_group_namelist_change_wrapper_for_go(NULL, NULL, NULL);
}

*/
import "C"
import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

// group callback type
type cb_group_invite_ftype func(this *Tox, friendNumber uint32, itype uint8, data []byte, userData interface{})
type cb_group_message_ftype func(this *Tox, groupNumber int, peerNumber int, message string, userData interface{})

type cb_group_action_ftype func(this *Tox, groupNumber int, peerNumber int, action string, userData interface{})
type cb_group_title_ftype func(this *Tox, groupNumber int, peerNumber int, title string, userData interface{})
type cb_group_namelist_change_ftype func(this *Tox, groupNumber int, peerNumber int, change uint8, userData interface{})

// tox_callback_group_***

//export callbackGroupInviteWrapperForC
func callbackGroupInviteWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.TOX_CONFERENCE_TYPE, a2 *C.uint8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_invites {
		cbfn := *(*cb_group_invite_ftype)(cbfni)
		data := C.GoBytes((unsafe.Pointer)(a2), C.int(a3))
		this.beforeCallback()
		cbfn(this, uint32(a0), uint8(a1), data, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackGroupInvite(cbfn cb_group_invite_ftype, userData interface{}) {
	this.CallbackGroupInviteAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupInviteAdd(cbfn cb_group_invite_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_invites[cbfnp]; ok {
		return
	}
	this.cb_group_invites[cbfnp] = userData

	var _cbfn = (C.cb_group_invite_ftype)(C.callbackGroupInviteWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_group_invite_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupMessageWrapperForC
func callbackGroupMessageWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, mtype C.TOX_MESSAGE_TYPE, a2 *C.int8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	if int(mtype) == MESSAGE_TYPE_NORMAL {
		for cbfni, ud := range this.cb_group_messages {
			cbfn := *(*cb_group_message_ftype)(cbfni)
			message := C.GoStringN((*C.char)((*C.int8_t)(a2)), C.int(a3))
			this.beforeCallback()
			cbfn(this, int(a0), int(a1), message, ud)
			this.afterCallback()
		}
	} else {
		for cbfni, ud := range this.cb_group_actions {
			cbfn := *(*cb_group_action_ftype)(cbfni)
			message := C.GoStringN((*C.char)((*C.int8_t)(a2)), C.int(a3))
			this.beforeCallback()
			cbfn(this, int(a0), int(a1), message, ud)
			this.afterCallback()
		}
	}
}

func (this *Tox) CallbackGroupMessage(cbfn cb_group_message_ftype, userData interface{}) {
	this.CallbackGroupMessageAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupMessageAdd(cbfn cb_group_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_messages[cbfnp]; ok {
		return
	}
	this.cb_group_messages[cbfnp] = userData

	if !this.cb_group_message_setted {
		this.cb_group_message_setted = true

		var _cbfn = (C.cb_group_message_ftype)(C.callbackGroupMessageWrapperForC)
		var _userData unsafe.Pointer = nil

		C.cb_group_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
	}
}

/*
//export callbackGroupActionWrapperForC
func callbackGroupActionWrapperForC(m *C.Tox, a0 C.int, a1 C.int, a2 *C.uint8_t, a3 C.uint16_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	if this.cb_group_action != nil {
		action := C.GoStringN((*C.char)((unsafe.Pointer)(a2)), C.int(a3))
		this.cb_group_action(this, int(a0), int(a1), action, this.cb_group_action_user_data)
	}
}
*/

func (this *Tox) CallbackGroupAction(cbfn cb_group_action_ftype, userData interface{}) {
	this.CallbackGroupActionAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupActionAdd(cbfn cb_group_action_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_actions[cbfnp]; ok {
		return
	}
	this.cb_group_actions[cbfnp] = userData

	if !this.cb_group_message_setted {
		this.cb_group_message_setted = true
		// var _cbfn = (C.cb_group_action_ftype)(C.callbackGroupActionWrapperForC)
		var _cbfn = (C.cb_group_message_ftype)(C.callbackGroupMessageWrapperForC)
		var _userData unsafe.Pointer = nil

		C.cb_group_message_wrapper_for_go(this.toxcore, _cbfn, _userData)
		// C.cb_group_action_wrapper_for_go(this.toxcore, _cbfn, _userData)
	}
}

//export callbackGroupTitleWrapperForC
func callbackGroupTitleWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, a2 *C.uint8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_titles {
		cbfn := *(*cb_group_title_ftype)(cbfni)
		title := C.GoStringN((*C.char)((unsafe.Pointer)(a2)), C.int(a3))
		this.beforeCallback()
		cbfn(this, int(a0), int(a1), title, ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackGroupTitle(cbfn cb_group_title_ftype, userData interface{}) {
	this.CallbackGroupTitleAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupTitleAdd(cbfn cb_group_title_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_titles[cbfnp]; ok {
		return
	}
	this.cb_group_titles[cbfnp] = userData

	var _cbfn = (C.cb_group_title_ftype)(C.callbackGroupTitleWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_group_title_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

//export callbackGroupNameListChangeWrapperForC
func callbackGroupNameListChangeWrapperForC(m *C.Tox, a0 C.uint32_t, a1 C.uint32_t, a2 C.TOX_CONFERENCE_STATE_CHANGE, a3 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_namelist_changes {
		cbfn := *(*cb_group_namelist_change_ftype)(cbfni)
		this.beforeCallback()
		cbfn(this, int(a0), int(a1), uint8(a2), ud)
		this.afterCallback()
	}
}

func (this *Tox) CallbackGroupNameListChange(cbfn cb_group_namelist_change_ftype, userData interface{}) {
	this.CallbackGroupNameListChangeAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupNameListChangeAdd(cbfn cb_group_namelist_change_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_namelist_changes[cbfnp]; ok {
		return
	}
	this.cb_group_namelist_changes[cbfnp] = userData

	var _cbfn = (C.cb_group_namelist_change_ftype)(C.callbackGroupNameListChangeWrapperForC)
	var _userData unsafe.Pointer = nil

	C.cb_group_namelist_change_wrapper_for_go(this.toxcore, _cbfn, _userData)
}

func (this *Tox) AddGroupChat() (int, error) {
	this.lock()
	defer this.unlock()

	r := C.tox_conference_new(this.toxcore, nil)
	if int(r) == -1 {
		return int(r), errors.New("add group chat failed")
	}
	return int(r), nil
}

func (this *Tox) DelGroupChat(groupNumber int) (int, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.uint32_t(groupNumber)
	var cerr C.TOX_ERR_CONFERENCE_DELETE
	r := C.tox_conference_delete(this.toxcore, _gn, &cerr)
	if bool(r) == false {
		return 1, errors.New(fmt.Sprintf("delete group chat failed:%d", int(cerr)))
	}
	return 0, nil
}

func (this *Tox) GroupPeerName(groupNumber int, peerNumber int) (string, error) {
	var _gn = C.uint32_t(groupNumber)
	var _pn = C.uint32_t(peerNumber)
	var _name = (*C.char)(C.calloc(1, C.size_t(MAX_NAME_LENGTH)))
	defer C.free(unsafe.Pointer(_name))

	r := C.tox_conference_peer_get_name(this.toxcore, _gn, _pn, char2uint8(_name), nil)
	if r == false {
		return "", errors.New("get peer name failed")
	}
	name := C.GoString(_name)
	return name, nil
}

func (this *Tox) GroupPeerPubkey(groupNumber int, peerNumber int) (string, error) {
	var _gn = C.uint32_t(groupNumber)
	var _pn = C.uint32_t(peerNumber)
	var _pubkey *C.char = (*C.char)(C.calloc(1, C.size_t(PUBLIC_KEY_SIZE)))
	defer C.free(unsafe.Pointer(_pubkey))

	r := C.tox_conference_peer_get_public_key(this.toxcore, _gn, _pn, char2uint8(_pubkey), nil)
	if r == false {
		return "", errors.New("get pubkey failed")
	}

	pubkey := hex.EncodeToString(C.GoBytes(unsafe.Pointer(_pubkey), C.int(PUBLIC_KEY_SIZE)))
	pubkey = strings.ToUpper(pubkey)
	return pubkey, nil
}

func (this *Tox) InviteFriend(friendNumber uint32, groupNumber int) (int, error) {
	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)
	var _gn = C.uint32_t(groupNumber)

	// if give a friendNumber which not exists,
	// the tox_invite_friend has a strange behaive: cause other tox_* call failed
	// and the call will return true, but only strange thing accurs
	// so just precheck the friendNumber and then go
	if !this.FriendExists(friendNumber) {
		return -1, errors.New("friend not exists")
	}

	r := C.tox_conference_invite(this.toxcore, _fn, _gn, nil)
	if r == false {
		return 0, toxerr("conference invite failed")
	}
	return 1, nil
}

func (this *Tox) JoinGroupChat(friendNumber uint32, data []byte) (int, error) {
	this.lock()
	defer this.unlock()

	if data == nil || len(data) < 10 {
		return -1, errors.New("invalid data")
	}
	var _fn = C.uint32_t(friendNumber)
	var _data = (*C.char)((unsafe.Pointer)(&data[0]))
	var length = len(data)
	var _length = C.size_t(length)

	r := C.tox_conference_join(this.toxcore, _fn, char2uint8(_data), _length, nil)
	if r == C.UINT32_MAX {
		return int(r), errors.New("join group chat failed")
	}
	return int(r), nil
}

func (this *Tox) GroupActionSend(groupNumber int, action string) (int, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.uint32_t(groupNumber)
	var _action = C.CString(action)
	defer C.free(unsafe.Pointer(_action))
	var _length = C.size_t(len(action))

	var cerr C.TOX_ERR_CONFERENCE_SEND_MESSAGE
	var mtype C.TOX_MESSAGE_TYPE = C.TOX_MESSAGE_TYPE_ACTION
	r := C.tox_conference_send_message(this.toxcore, _gn, mtype, char2uint8(_action), _length, &cerr)
	if r == false {
		return 0, errors.New("group action failed")
	}
	return 1, nil
}

func (this *Tox) GroupMessageSend(groupNumber int, message string) (int, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.uint32_t(groupNumber)
	var _message = C.CString(message)
	defer C.free(unsafe.Pointer(_message))
	var _length = C.size_t(len(message))

	var cerr C.TOX_ERR_CONFERENCE_SEND_MESSAGE
	var mtype C.TOX_MESSAGE_TYPE = C.TOX_MESSAGE_TYPE_NORMAL
	r := C.tox_conference_send_message(this.toxcore, _gn, mtype, char2uint8(_message), _length, &cerr)
	if r == false {
		return 0, errors.New("group send message failed")
	}
	return 1, nil
}

func (this *Tox) GroupSetTitle(groupNumber int, title string) (int, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.uint32_t(groupNumber)
	var _title = C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	var _length = C.size_t(len(title))

	r := C.tox_conference_set_title(this.toxcore, _gn, char2uint8(_title), _length, nil)
	if r == false {
		if len(title) > MAX_NAME_LENGTH {
			return 0, errors.New("title too long")
		}
		return 0, errors.New("set title failed")
	}
	return 1, nil
}

func (this *Tox) GroupGetTitle(groupNumber int) (string, error) {
	var _gn = C.uint32_t(groupNumber)
	var _title = (*C.char)(C.calloc(1, C.size_t(MAX_NAME_LENGTH)))
	defer C.free(unsafe.Pointer(_title))
	// var _maxlen = C.uint32_t(MAX_NAME_LENGTH)

	r := C.tox_conference_get_title(this.toxcore, _gn, char2uint8(_title), nil)
	if r == false {
		return "", errors.New("get title failed")
	}
	title := C.GoString(_title)
	return title, nil
}

func (this *Tox) GroupPeerNumberIsOurs(groupNumber int, peerNumber uint32) bool {
	var _gn = C.uint32_t(groupNumber)
	var _pn = C.uint32_t(peerNumber)

	r := C.tox_conference_peer_number_is_ours(this.toxcore, _gn, _pn, nil)
	return bool(r)
}

func (this *Tox) GroupNumberPeers(groupNumber int) int {
	var _gn = C.uint32_t(groupNumber)

	r := C.tox_conference_peer_count(this.toxcore, _gn, nil)
	return int(r)
}

func (this *Tox) GroupGetNames(groupNumber int) []string {
	peerCount := this.GroupNumberPeers(groupNumber)
	vec := make([]string, peerCount)
	if peerCount == 0 {
		return vec
	}

	// lengths := make([]uint16, peerCount)
	// names := make([]byte, peerCount*MAX_NAME_LENGTH)
	// clengths := (*C.uint16_t)(&lengths[0])
	// cnames := (*[MAX_NAME_LENGTH]C.uint8_t)((unsafe.Pointer)(&names[0]))

	for idx := 0; idx < peerCount; idx++ {
		pname, err := this.GroupPeerName(groupNumber, idx)
		if err != nil {
			return vec[0:0]
		}
		vec[idx] = pname
	}

	/*
		r := C.tox_group_get_names(this.toxcore, C.int(groupNumber),
			cnames, clengths, C.uint16_t(peerCount))
		if int(r) == -1 {
			return vec[0:0]
		}

		for idx := 0; idx < peerCount; idx++ {
			len := int(lengths[idx])
			name := names[idx*MAX_NAME_LENGTH : (idx*MAX_NAME_LENGTH + len)]
			vec[idx] = string(name)
		}
	*/
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
	r := C.tox_conference_get_chatlist_size(this.toxcore)
	return uint32(r)
}

func (this *Tox) GetChatList() []int32 {
	var sz uint32 = this.CountChatList()
	vec := make([]int32, sz)
	if sz == 0 {
		return vec
	}

	vec_p := unsafe.Pointer(&vec[0])
	C.tox_conference_get_chatlist(this.toxcore, (*C.uint32_t)(vec_p))
	return vec
}

func (this *Tox) GroupGetType(groupNumber uint32) (int, error) {
	var _gn = C.uint32_t(groupNumber)

	r := C.tox_conference_get_type(this.toxcore, _gn, nil)
	if int(r) == -1 {
		return int(r), errors.New("get type failed")
	}
	return int(r), nil
}
