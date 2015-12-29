package main

import (
    "flag"
    "fmt"
    "rest"
    "net/http"
)

var (
    deleteRange = flag.String("deleterange", "", "Set to true to remove range from database")
    port = flag.String("port", "8080", "Overrides the default port.")
    ipRange = flag.String("range", "", "Specifies a range of IPV6 addresses to use.")
)

// Starts server on specified or default port and processes JSON request.  ETCD is
// used as a back end and must be running before this program is started.
//
// The optional range parameter must be used to specify an IPV6 range on the first time
// this is run.  The range is stored in ETCD and after the first run it will be
// retrieved.  Specifying it again in the future will cause the range to be augmented
// in ETCD.
func main() {
    
    flag.Parse()
    
	fmt.Println("Hello, IPV6Thing is starting up")
	fmt.Printf("IPV6 Range: %s\n", *ipRange)
	fmt.Printf("Server listenig on port %s\n", *port)
    
	http.HandleFunc("/reserveipv6addr", serveHTTPReserveIPV6Addr)
	http.HandleFunc("/reserveipv6multiaddr", serveHTTPReserveIPV6MultiAddr)
	http.ListenAndServe(":"+*port, nil)
}

// Serves a request to reserve a single IPV6 address
func serveHTTPReserveIPV6Addr(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
	    case "GET":
	        ret, err := rest.ReserveIpV6Addr(0)
	        if err != nil {
    	        // TODO: DEFINE ERROR BEHAVIOR - SEND ERR HTTP CODE????
		        fmt.Fprint(w, err)
	        } else {
		        fmt.Fprint(w, ret)
	        }
    }
}

// Serves a request to reserve multiple IPV6 addresses
func serveHTTPReserveIPV6MultiAddr(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
	    case "GET":
	        ret, err := rest.ReserveIpV6MultiAddr(1, 0)
	        if err != nil{
    	        // TODO: DEFINE ERROR BEHAVIOR - SEND ERR HTTP CODE????
		        fmt.Fprint(w, err)
	        } else {
		        fmt.Fprint(w, ret)
	        }
   }
}