
all:
	CGO_CFLAGS=-I$(HOME)/oss/toxcore/toxcore go install -v
	CGO_CFLAGS=-I$(HOME)/oss/toxcore/toxcore go build -o tox-bootstrap cmds/tox-bootstrap.go 

test:
	CGO_CFLAGS=-I$(HOME)/oss/toxcore/toxcore go test -v -run DHT
