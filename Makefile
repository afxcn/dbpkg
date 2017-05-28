RELEASES=bin/dbc-darwin-amd64 \
	 bin/dbc-linux-amd64 \
	 bin/dbc-linux-386 \
	 bin/dbc-linux-arm \
	 bin/dbc-windows-amd64.exe \
	 bin/dbc-windows-386.exe \
	 bin/dbc-solaris-amd64 

all: $(RELEASES)

bin/dbc-%: GOOS=$(firstword $(subst -, ,$*))
bin/dbc-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
bin/dbc-%: $(wildcard *.go)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build \
	     -ldflags "-X main.osarch=$(GOOS)/$(GOARCH) -s -w" \
	     -buildmode=exe \
	     -tags release \
	     -o $@

clean:
	rm -rf bin
