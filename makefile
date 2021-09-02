all: build

build:
	go build -o shttpd main.go httpserve.go
install: build
	cp shttpd /usr/local/bin/shttpd
uninstall:
	rm /usr/local/bin/shttpd
