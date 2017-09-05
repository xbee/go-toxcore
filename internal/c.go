package toxin

/*
#cgo CFLAGS: -g -O2 -std=c99 -Wall
#cgo CFLAGS: -I/path/to/toxcore/toxcore

// toxcore
// #cgo CFLAGS: -I${SRCDIR}/toxcore/toxcore
// #cgo LDFLAGS: -lsodium ${SRCDIR}/toxcore/build/.libs/libtoxcore.a

// c-toxcore
#cgo CFLAGS: -I${SRCDIR}/c-toxcore/toxcore
#cgo LDFLAGS: -lsodium ${SRCDIR}/c-toxcore/libtoxcore.a ${SRCDIR}/c-toxcore/libtoxnetwork.a ${SRCDIR}/c-toxcore/libtoxnetcrypto.a ${SRCDIR}/c-toxcore/libtoxcrypto.a ${SRCDIR}/c-toxcore/libtoxdht.a

// #cgo pkg-config: libtoxcore
*/
import "C"

// install
// CGO_CFLAGS="-I/path/to/toxcore/toxcore" go install
// TODO what about Windows/MacOS?
