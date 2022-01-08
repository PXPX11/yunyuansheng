package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func respondHandler(w http.ResponseWriter, r *http.Request) {
	for i, j := range r.Header {
		w.Header().Set(i, j[0])
	}
	//fmt.Fprintln(w, w.Header())
	env := os.Getenv("VERSION")
	if env != "" {
		w.Header().Set("VERSION", env)
		fmt.Printf(env)
	} else {
		fmt.Printf("VERSION DOES NOT FOUND! \n")
	}
	//HOME := os.Getenv("HOME")
	//fmt.Println(HOME)
	hostIp, port, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		ip := net.ParseIP(hostIp)
		if ip != nil {
			fmt.Printf("We get the IP" + hostIp + "and the port is " + port)
		}
	} else {
		fmt.Printf("We did not get the required response due to the error of " + err.Error())
	}
	fmt.Printf("\n HTTP Respond Code is %v", http.StatusOK)
}
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HTTP Respond Code is %v", http.StatusOK)
}
func main() {
	http.HandleFunc("/", respondHandler)
	http.HandleFunc("/healthz", healthzHandler)
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Printf("Web server start failed!")
	}
}
