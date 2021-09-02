package main

import (
	"fmt"
	"log"
	"os"
)

func printHelp() {
	fmt.Println("USAGE: shttpd -a <host:port> -r <root directory>")
	fmt.Println("       shttpd -a <host:port> -r <root directory> -k <keyfile> -c <certfile>")
	fmt.Println("       shttpd -h")
	fmt.Printf("\nThe root directory is where all your html files are kept.\nThe 'index.html' file is always the default index.\nIf no 'index.html' file is found, then a directory listing is served.\nThe key file and cert file is for TLS/SSL, and if specified, will enable TLS/SSL\nautomatically.\n")
}

func main() {
	//Main variable used in the program and default values used if none are changed.
	var rootserve string = "./"
	var serveaddr string = "127.0.0.1:8080"
	var keyfile string
	var certfile string
	helpasked, certspecified, keyspecified := false, false, false

	//Check the parameters to change values if needed.
	if len(os.Args) > 1 {
		for i, v := range os.Args {
			switch v {
			case "-r":
				rootserve = os.Args[i+1]
			case "-a":
				serveaddr = os.Args[i+1]
			case "-c":
				certfile = os.Args[i+1]
				certspecified = true
			case "-k":
				keyfile = os.Args[i+1]
				keyspecified = true
			case "-h":
				helpasked = true
			}
		}
	}

	//Was help asked?
	if !helpasked {

		//if a cert and key was specified, make a TLS/SSL based connection. If not, make a normal connection.
		var err error
		if keyspecified && certspecified {
			fmt.Println("Securely Serving", rootserve, "on", serveaddr)
			err = ServeFilesSecure(rootserve, serveaddr, certfile, keyfile)
		} else {
			fmt.Println("Serving", rootserve, "on", serveaddr)
			err = ServeFiles(rootserve, serveaddr)
		}

		//Did something go wrong? Check for errors.
		if err != nil {
			log.Fatal(err)
		}
	} else {

		//Prints the help
		printHelp()
	}

}
