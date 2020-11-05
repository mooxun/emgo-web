BINARY="emgo"

.PHONY: all
all: clean build install

.PHONY: linux
linux: clean build-linux install

.PHONY: build
build:
	@echo "build emgo start >>>"
	mkdir -p ./emgo-web
	go build -o ./emgo-web/emgo ./main.go
	@echo ">>> build syncd complete"

.PHONY: install
install:
	@echo "install emgo start >>>"
	mkdir -p ./emgo-web
	mkdir -p ./emgo-web/config
	cp ./config/app.yaml ./emgo-web/config/
	@echo ">>> install emgo complete"

.PHONY: clean
clean:
	@echo "clean start >>>"
	rm -fr ./emgo-web
	@echo ">>> clean complete"

.PHONY: build-linux
build-linux:
	@echo "build-linux start >>>"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./emgo-web/emgo ./main.go
	@echo ">>> build-linux complete"