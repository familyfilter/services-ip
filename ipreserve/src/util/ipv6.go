// Utility methods for processing IPV6 addresses
package util

import (
    "errors"
    "net"
)

// Given an IP address range convert it inot a pair of IP's.  Errs are returned
// if this is not a valid address range.
//
// QUESTION: Should the input address range be a CIDR or a two IP addresses
// seperated by a dash?
//
func ConvertRangeToIPs(addrRange string) (net.IP, net.IP, error) {
    
    var a net.IP
    var b net.IP
    var err error
    
    err = errors.New("this has not been programmed yet.")
    return a, b, err
}