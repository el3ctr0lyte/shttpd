package main

import(
	"net/http"
)

//Serve insecurely, no SSL/TLS
func ServeFiles(rootdir,addr string) error {
	err := http.ListenAndServe(addr, http.FileServer(http.Dir(rootdir)))
	return err
}

//Serve securely, with SSL/TLS
func ServeFilesSecure(rootdir,addr, certfile, keyfile string) error {
	err := http.ListenAndServeTLS(addr,certfile, keyfile, http.FileServer(http.Dir(rootdir)))
	return err
}