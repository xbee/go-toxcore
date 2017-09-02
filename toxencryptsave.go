package tox

/*
#include <tox/toxencryptsave.h>
*/
import "C"

const PASS_KEY_LENGTH = int(C.TOX_PASS_KEY_LENGTH)
const PASS_ENCRYPTION_EXTRA_LENGTH = int(C.TOX_PASS_ENCRYPTION_EXTRA_LENGTH)

type ToxPassKey struct {
	cpk *C.Tox_Pass_Key
}

func NewToxPassKey() *ToxPassKey {
	this := &ToxPassKey{}
	this.cpk = C.tox_pass_key_new()
	return this
}

func (this *ToxPassKey) Free() {
	C.tox_pass_key_free(this.cpk)
}

func (this *ToxPassKey) Derive(passphrase []byte) (bool, error) {
	passphrase_ := (*C.uint8_t)(&passphrase[0])

	var cerr C.TOX_ERR_KEY_DERIVATION
	ok := C.tox_pass_key_derive(this.cpk, passphrase_, C.size_t(len(passphrase)), &cerr)
	var err error
	if !bool(ok) {
		err = toxerr(cerr)
	}
	return bool(ok), err
}

func (this *ToxPassKey) DeriveWithSalt(passphrase []byte, salt []byte) (bool, error) {
	passphrase_ := (*C.uint8_t)(&passphrase[0])
	salt_ := (*C.uint8_t)(&salt[0])

	var cerr C.TOX_ERR_KEY_DERIVATION
	ok := C.tox_pass_key_derive_with_salt(this.cpk, passphrase_, C.size_t(len(passphrase)), salt_, &cerr)
	var err error
	if !bool(ok) {
		err = toxerr(cerr)
	}
	return bool(ok), err
}

func (this *ToxPassKey) Encrypt(plaintext []byte) (bool, error, []byte) {
	ciphertext := make([]byte, len(plaintext)+PASS_ENCRYPTION_EXTRA_LENGTH)
	ciphertext_ := (*C.uint8_t)(&ciphertext[0])
	plaintext_ := (*C.uint8_t)(&plaintext[0])

	var cerr C.TOX_ERR_ENCRYPTION
	ok := C.tox_pass_key_encrypt(this.cpk, plaintext_, C.size_t(len(plaintext)), ciphertext_, &cerr)

	var err error
	if !bool(ok) {
		err = toxerr(err)
	}
	return bool(ok), err, ciphertext
}

func (this *ToxPassKey) Decrypt(ciphertext []byte) (bool, error, []byte) {
	ciphertext_ := (*C.uint8_t)(&ciphertext[0])
	plaintext := make([]byte, len(ciphertext)-PASS_ENCRYPTION_EXTRA_LENGTH)
	plaintext_ := (*C.uint8_t)(&plaintext[0])

	var cerr C.TOX_ERR_DECRYPTION
	ok := C.tox_pass_key_decrypt(this.cpk, ciphertext_, C.size_t(len(ciphertext)), plaintext_, &cerr)
	var err error
	if !bool(ok) {
		err = toxerr(cerr)
	}
	return bool(ok), err, plaintext
}

func GetSalt(ciphertext []byte) (bool, error, []byte) {
	ciphertext_ := (*C.uint8_t)(&ciphertext[0])
	salt := make([]byte, int(C.TOX_PASS_SALT_LENGTH))
	salt_ := (*C.uint8_t)(&salt[0])

	var cerr C.TOX_ERR_GET_SALT
	ok := C.tox_get_salt(ciphertext_, salt_, &cerr)
	var err error
	if !bool(ok) {
		err = toxerr(cerr)
	}
	return bool(ok), err, salt
}

func IsDataEncrypted(data []byte) bool {
	data_ := (*C.uint8_t)(&data[0])
	ok := C.tox_is_data_encrypted(data_)
	if ok == C._Bool(false) {
		return false
	}
	return true
}

func PassEncrypt(plaintext []byte, passphrase []byte) (ciphertext []byte, err error) {
	ciphertext = make([]byte, len(plaintext)+PASS_ENCRYPTION_EXTRA_LENGTH)
	ciphertext_ := (*C.uint8_t)(&ciphertext[0])
	plaintext_ := (*C.uint8_t)(&plaintext[0])
	passphrase_ := (*C.uint8_t)(&passphrase[0])

	var cerr C.TOX_ERR_ENCRYPTION
	ok := C.tox_pass_encrypt(plaintext_, C.size_t(len(plaintext)), passphrase_, C.size_t(len(passphrase)), ciphertext_, &cerr)

	if !bool(ok) {
		err = toxerr(cerr)
	}
	return
}

func PassDecrypt(ciphertext []byte, passphrase []byte) (plaintext []byte, err error) {
	ciphertext_ := (*C.uint8_t)(&ciphertext[0])
	plaintext = make([]byte, len(ciphertext)-PASS_ENCRYPTION_EXTRA_LENGTH)
	plaintext_ := (*C.uint8_t)(&plaintext[0])
	passphrase_ := (*C.uint8_t)(&plaintext[0])

	var cerr C.TOX_ERR_DECRYPTION
	ok := C.tox_pass_decrypt(ciphertext_, C.size_t(len(ciphertext)), passphrase_, C.size_t(len(passphrase)), plaintext_, &cerr)

	if !bool(ok) {
		err = toxerr(cerr)
	}
	return
}
