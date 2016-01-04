// Utility methods for processing IPV6 addresses
package util

import (
    "errors"
    "net"
    "strings"
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
 
    if strings.Index(addrRange, "-") != 1 {
        // Format 1: 2001:0db8:0000:0042:0000:8a2e:0370:7334-2001:0db8:0000:0042:0000:8a2e:0370:7338
        // To check for this format a "-" must be fund in the string.
        addresses := strings.Split(addrRange, "-")
        if len(addresses) != 2 {
            return nil, nil, errors.New("Illegal format, exactly one '-' required")
        }
        a := net.ParseIP(addresses[0])
        if a == nil {
            return nil, nil, errors.New(addresses[0] + " is not a legal IP")
        }
        b := net.ParseIP(addresses[1])
        if b == nil {
            return nil, nil, errors.New(addresses[0] + " is not a legal IP")
        }
        return a, b, nil
        
    } else {
        // Format 2: Assume a CIDR.  To test for this method just run it through the ParseCIDR method 
        //
        // TODO: Is this desirable?
    } 
    err = errors.New("this has not been programmed yet.")
    return a, b, err
}

// This increments an IP by one.  If the IP is at the highest value it will rotate it to all 0's.
func IncrementIP(ip net.IP) net.IP {
    
    // ASSUMPTION: THIS IS IPV6 - ADD IPV4 if needed.
    
    for i:=15; i>=0; i-- { 
        ip[i] = ip[i] + 1
        if ip[i] != 0 {
            break
        }
    }
    return ip
}
