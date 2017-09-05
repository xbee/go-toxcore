
all:
	# go tool cgo -exportheader toxin_cgo_export.h -- -Ic-toxcore/toxcore dht_callback.go util.go
	go install -v 
	go build -o tox-bootstrap cmds/tox-bootstrap.go
	go build -o crawlerd crawler/crawler.go
	# CGO_CFLAGS=-I$(PWD)/c-toxcore/toxcore go install -v
	# CGO_CFLAGS=-I$(PWD)/c-toxcore/toxcore go build -o tox-bootstrap cmds/tox-bootstrap.go

test:
	CGO_CFLAGS=-I$(HOME)/oss/toxcore/toxcore go test -v -run NCN
	# CGO_CFLAGS=-I$(HOME)/oss/toxcore/toxcore go test -v -run DHT
