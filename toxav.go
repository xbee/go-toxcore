package tox

/*
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>
#include <tox/toxav.h>

void callbackCallWrapperForC(ToxAV *toxAV, uint32_t friend_number, bool audio_enabled,
                           bool video_enabled, void *user_data);
typedef void (*cb_call_ftype)(ToxAV *toxAV, uint32_t friend_number, bool audio_enabled,
                           bool video_enabled, void *user_data);
static void cb_call_wrapper_for_go(ToxAV *m, cb_call_ftype fn, void *userdata)
{ toxav_callback_call(m, fn, userdata); }

void callbackCallStateWrapperForC(ToxAV *toxAV, uint32_t friendNumber, uint32_t state, void* user_data);
typedef void (*cb_call_state_ftype)(ToxAV *toxAV, uint32_t friendNumber, uint32_t state, void *user_data);
static void cb_call_state_wrapper_for_go(ToxAV *m, cb_call_state_ftype fn, void *userdata)
{ toxav_callback_call_state(m, fn, userdata); }

void callbackBitRateStatusWrapperForC(ToxAV *toxAV, uint32_t friendNumber, uint32_t audioBitRate, uint32_t videoBitRate, void* user_data);
typedef void (*cb_bit_rate_status_ftype)(ToxAV *toxAV, uint32_t friendNumber, uint32_t audioBitRate, uint32_t videoBitRate, void *user_data);
static void cb_bit_rate_status_wrapper_for_go(ToxAV *m, cb_call_state_ftype fn, void *userdata)
{ toxav_callback_bit_rate_status(m, fn, userdata); }


// fix nouse compile warning
static inline void fixnousetoxav() {
    cb_call_wrapper_for_go(NULL, NULL, NULL);
    cb_call_state_wrapper_for_go(NULL, NULL, NULL);
    cb_bit_rate_status_wrapper_for_go(NULL, NULL, NULL);
}

*/
import "C"
import "unsafe"

type cb_call_ftype func(this *ToxAV, groupNumber uint32, audioEnabled bool, videoEnabled bool, userData unsafe.Pointer)
type cb_call_state_ftype func(this *ToxAV, groupNumber uint32, state uint32, userData unsafe.Pointer)
type cb_bit_rate_status_ftype func(this *ToxAV, groupNumber uint32, audioBitRate uint32, videoBitRate uint32, userData unsafe.Pointer)

type ToxAV struct {
	tox   *Tox
	toxav *C.ToxAV

	cb_call                      cb_call_ftype
	cb_call_user_data            unsafe.Pointer
	cb_call_state                cb_call_state_ftype
	cb_call_state_user_data      unsafe.Pointer
	cb_bit_rate_status           cb_bit_rate_status_ftype
	cb_bit_rate_status_user_data unsafe.Pointer
}

func NewToxAV(tox *Tox) *ToxAV {
	tav := new(ToxAV)
	tav.tox = tox

	var cerr C.TOXAV_ERR_NEW
	tav.toxav = C.toxav_new(tox.toxcore, &cerr)
	if cerr != 0 {
	}

	cbAVUserDatas[tav.toxav] = tav
	return tav
}

func (this *ToxAV) Kill() {
	C.toxav_kill(this.toxav)
}

func (this *ToxAV) GetTox() *Tox {
	return this.tox
}

func (this *ToxAV) IterationInterval() uint32 {
	return uint32(C.toxav_iteration_interval(this.toxav))
}

func (this *ToxAV) Iterate() {
	C.toxav_iterate(this.toxav)
}

func (this *ToxAV) Call(friendNumber uint32, audioBitRate uint32, videoBitRate uint32) (bool, error) {
	var cerr C.TOXAV_ERR_CALL
	r := C.toxav_call(this.toxav, C.uint32_t(friendNumber), C.uint32_t(audioBitRate), C.uint32_t(videoBitRate), &cerr)
	if cerr != 0 {

	}
	return bool(r), nil
}

var cbAVUserDatas map[*C.ToxAV]*ToxAV = make(map[*C.ToxAV]*ToxAV, 0)

//export callbackCallWrapperForC
func callbackCallWrapperForC(m *C.ToxAV, friendNumber C.uint32_t, audioEnabled C.bool, videoEnabled C.bool, a3 unsafe.Pointer) {
	var this = (*ToxAV)(cbAVUserDatas[m])
	if this.cb_call != nil {
		this.cb_call(this, uint32(friendNumber), bool(audioEnabled), bool(videoEnabled), this.cb_call_user_data)
	}
}

func (this *ToxAV) CallbackCall(cbfn cb_call_ftype, userData unsafe.Pointer) {
	this.cb_call = cbfn
	this.cb_call_user_data = userData

	var _cbfn = (C.cb_call_ftype)(C.callbackCallWrapperForC)
	var _userData = unsafe.Pointer(this)
	_userData = nil

	C.cb_call_wrapper_for_go(this.toxav, _cbfn, _userData)
}

func (this *ToxAV) answer(friendNumber uint32, audioBitRate uint32, videoBitRate uint32) (bool, error) {
	var cerr C.TOXAV_ERR_ANSWER
	r := C.toxav_answer(this.toxav, C.uint32_t(friendNumber), C.uint32_t(audioBitRate), C.uint32_t(videoBitRate), &cerr)
	if cerr != C.TOXAV_ERR_ANSWER_OK {

	}

	return bool(r), nil
}

//export callbackCallStateWrapperForC
func callbackCallStateWrapperForC(m *C.ToxAV, friendNumber C.uint32_t, state C.uint32_t, a3 unsafe.Pointer) {
	var this = (*ToxAV)(cbAVUserDatas[m])
	if this.cb_call_state != nil {
		this.cb_call_state(this, uint32(friendNumber), uint32(state), this.cb_call_state_user_data)
	}
}

func (this *ToxAV) CallbackCallState(cbfn cb_call_state_ftype, userData unsafe.Pointer) {
	this.cb_call_state = cbfn
	this.cb_call_state_user_data = userData

	var _cbfn = (C.cb_call_state_ftype)(C.callbackCallStateWrapperForC)
	var _userData = unsafe.Pointer(this)
	_userData = nil

	C.cb_call_state_wrapper_for_go(this.toxav, _cbfn, _userData)
}

func (this *ToxAV) CallControl(friendNumber uint32, control int) (bool, error) {
	var cerr C.TOXAV_ERR_CALL_CONTROL
	r := C.toxav_call_control(this.toxav, C.uint32_t(friendNumber), C.TOXAV_CALL_CONTROL(control), &cerr)
	if cerr != C.TOXAV_ERR_CALL_CONTROL_OK {
	}
	return bool(r), nil
}

func (this *ToxAV) BitRateSet(friendNumber uint32, audioBitRate int32, videoBitRate int32) (bool, error) {
	var cerr C.TOXAV_ERR_BIT_RATE_SET
	r := C.toxav_bit_rate_set(this.toxav, C.uint32_t(friendNumber), C.int32_t(audioBitRate), C.int32_t(videoBitRate), &cerr)
	if cerr != C.TOXAV_ERR_BIT_RATE_SET_OK {
	}
	return bool(r), nil
}

//export callbackBitRateStatusWrapperForC
func callbackBitRateStatusWrapperForC(m *C.ToxAV, friendNumber C.uint32_t, audioBitRate C.uint32_t, videoBitRate C.uint32_t, a3 unsafe.Pointer) {
	var this = (*ToxAV)(cbAVUserDatas[m])
	if this.cb_bit_rate_status != nil {
		this.cb_bit_rate_status(this, uint32(friendNumber), uint32(audioBitRate), uint32(videoBitRate), this.cb_call_state_user_data)
	}
}

func (this *ToxAV) CallbackBitRateStatus(cbfn cb_bit_rate_status_ftype, userData unsafe.Pointer) {
	this.cb_bit_rate_status = cbfn
	this.cb_bit_rate_status_user_data = userData

	var _cbfn = (C.cb_bit_rate_status_ftype)(C.callbackBitRateStatusWrapperForC)
	var _userData = unsafe.Pointer(this)
	_userData = nil

	C.cb_bit_rate_status_wrapper_for_go(this.toxav, _cbfn, _userData)
}
