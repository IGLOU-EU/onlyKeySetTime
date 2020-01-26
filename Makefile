GO     ?= go
GOFMT  ?= gofmt -s
GOPATH ?= $($(GO) env GOPATH)

BIN  ?= /usr/local/bin
UDEV ?= /etc/udev/rules.d

DASH   ?= $(if $(GOOS),_,)
GOOS   ?=
GOARCH ?=

EXEC    := onlykey-settime$(DASH)$(GOOS)$(DASH)$(GOARCH)
GOOUT   := $(shell pwd)/bin
GOFILES := main.go

.PHONY: help
help:
	@echo "Make Routines:"
	@echo " - \"\"                print Make routines list"
	@echo " - build             creates the binary, Cross-Compiling ex : 'make build GOOS='linux' GOARCH='arm''"
	@echo " - clean             delete build files"
	@echo " - install           install the binary and udev rules"
	@echo " - remove            remove binary and udev rules"
	@echo " - upgrade           build and update installed binary"

.PHONY: install
install: build
	sudo cp $(GOOUT)/$(EXEC) $(BIN)/$(EXEC)
	sudo cp 49-onlykey-settime.rules $(UDEV)/49-onlykey-settime.rules
	sudo udevadm control --reload
	make clean

.PHONY: remove
remove:
	sudo rm -rf $(BIN)/$(EXEC)
	sudo rm -rf $(UDEV)/49-onlykey-settime.rules
	sudo udevadm control --reload

.PHONY: upgrade
upgrade: build
	sudo cp $(GOOUT)/$(EXEC) $(BIN)/$(EXEC)
	make clean

.PHONY: fmt
fmt:
	$(GOFMT) -e -w $(GOFILES)

.PHONY: godep
godep:
	$(GO) get github.com/karalabe/usb

.PHONY: build
build: godep
	GOARCH=$(GOARCH) GOOS=$(GOOS) $(GO) build -o $(GOOUT)/$(EXEC) -a

.PHONY: clean
clean:
	$(GO) clean
	rm -rf $(GOOUT)/*