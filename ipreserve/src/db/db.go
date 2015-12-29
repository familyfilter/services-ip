// Routines to acccess the etcd database
package db

import (
    "fmt"
    "log"
    "net"
    "time"
    
    "github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
    "github.com/coreos/etcd/client"    
)

var conn client.Client

// Initialize database connection on startup.  Please note that there is no need to 
// de-initialize etcd since it is just a http connection.
func init() {
    
    cfg := client.Config{
        Endpoints:               []string{"http://0.0.0.0:2379"},
        Transport:               client.DefaultTransport,
        // set timeout per request to fail fast when the target endpoint is unavailable
        HeaderTimeoutPerRequest: time.Second,
    }
    var err error
    conn, err = client.New(cfg)
    if err != nil {
        log.Fatal(err)
    }
}

// Retrieve the main IP range.
func RetrieveRange() (net.IP, net.IP, error) {
    kapi := client.NewKeysAPI(conn)
    resp, err := kapi.Get(context.Background(), "/addrrange", nil)
    if err != nil {
        return nil, nil, err
    } else {
        fmt.Println(resp)  // PLACEHOLDER - NEED TO PARSE FOR REAL
        return nil, nil, nil
    }
}

// Insert the main IP range, return true if this replaces a previous IP range that is 
// of different value.  Return false if this is either brand new or the exact same IP
// range as was previously recorded.
func InsertRange (a net.IP, b net.IP) (bool, error) {
    
    ret := false
    
    aa, bb, err := RetrieveRange()
    if err == nil {
        if a.String() != aa.String() || b.String() != bb.String() {
            ret = true
        }    
    }
    // Assume that if this errors out the range was not found.  If its a database error
    // it will be picked up in the next call anyway
    
    addrRange := a.String() + "-" + a.String()
    kapi := client.NewKeysAPI(conn)
    _, err = kapi.Set(context.Background(), "/addrrange", addrRange, nil)
    if err != nil {
        return ret, err
    }
    return ret, nil
}

// Deleting the address range is mainly for teardown after test but could be also used 
// to clean up in the main program.
func DeleteRange() error {
    kapi := client.NewKeysAPI(conn)
    _, err := kapi.Delete(context.Background(), "/addrrange", nil)
    return err
}
