package main

import (
	"errors"
    "flag"
    "fmt"
    "rest"
    "net/http"
    "strconv"
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
			clientIdStr := r.URL.Query().Get("clientid")
			clientId, err := authenticateClientId(clientIdStr)
			if err != nil {
		        fmt.Fprint(w, rest.ProcessingError(err))
		        return
			}
			
	        ret, err := rest.ReserveIpV6Addr(clientId)
	        if err != nil {
		        fmt.Fprint(w, rest.ProcessingError(err))
	        } else {
		        fmt.Fprint(w, ret)
	        }
    }
}

// Serves a request to reserve multiple IPV6 addresses
func serveHTTPReserveIPV6MultiAddr(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
	    case "GET":
			clientId, err := authenticateClientId(r.URL.Query().Get("clientid"))
			if err != nil {
		        fmt.Fprint(w, rest.ProcessingError(err))
		        return
			}
			
			numberIps, err := verifyNumberIps(r.URL.Query().Get("num"))
			if err != nil {
		        fmt.Fprint(w, rest.ProcessingError(err))
		        return
			}

	        ret, err := rest.ReserveIpV6MultiAddr(clientId, numberIps)
	        if err != nil{
		        fmt.Fprint(w, rest.ProcessingError(err))
	        } else {
		        fmt.Fprint(w, ret)
	        }
   }
}

// Authenticate a client id.  For now this is a placeholder, eventually this will
// probably be tied to the db package or some sort of authentication package.
// It returns either an error for an illegal ID (which right now just has to 
// be numeric but will also eventually need to be authenticated somehow) or
// a numeric client id.
func authenticateClientId(clientIdStr string) (int, error) {
	if clientIdStr == "" {
		return 0, errors.New("clientid is required")
	}
	return strconv.Atoi(clientIdStr)
}

func verifyNumberIps(numberIpsStr string) (int, error) {
	if numberIpsStr == "" {
		return 0, errors.New("num is required")
	}
	return strconv.Atoi(numberIpsStr)
}