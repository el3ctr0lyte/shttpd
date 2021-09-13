package main

import (
	"flag"
	"fmt"
)

func main() {
	//Main variable used in the program and default values used if none are changed.
	var rootserve string
	var serveaddr string
	var keyfile string
	var certfile string

	var err error

	flag.StringVar(&rootserve, "root", "./", "The root directory to serve from")
	flag.StringVar(&serveaddr, "addr", "127.0.0.1:8080", "The address to bind to, eg 127.0.0.1:8080")
	flag.StringVar(&keyfile, "key", "", "The SSL/TLS key file. SSL/TLS enabled if specified with cert")
	flag.StringVar(&certfile, "cert", "", "The SSL/TLS cert file. SSL/TLS enabled if specified with key")

	flag.StringVar(&rootserve, "r", "./", "Shorthand for root")
	flag.StringVar(&serveaddr, "a", "127.0.0.1:8080", "Shorthand for addr")
	flag.StringVar(&keyfile, "k", "", "Shorthand for key")
	flag.StringVar(&certfile, "c", "", "Shorthand for cert")

	flag.Parse()

	if (keyfile != "") || (certfile != "") {
		fmt.Println("Securely serving", rootserve, "on", serveaddr)
		err = ServeFilesSecure(rootserve, serveaddr, certfile, keyfile)
	} else {
		fmt.Println("Serving", rootserve, "on", serveaddr)
		err = ServeFiles(rootserve, serveaddr)
	}

	if err != nil {
		fmt.Println(err)
	}
}
