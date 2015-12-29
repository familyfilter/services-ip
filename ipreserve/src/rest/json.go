package rest

import (
    _ "encoding/json"
)
    
const (
    HARDCODED_JSON = `
        {
            "addr": "2001:0db8:0000:0042:0000:8a2e:0370:7334"
        }
    `

    HARDCODED_JSON_MULTI = `
        {
            "addresses": [
                "2001:0db8:0000:0042:0000:8a2e:0370:7334",
                "2001:0db8:0000:0042:0000:8a2e:0370:7335",
                "2001:0db8:0000:0042:0000:8a2e:0370:7336",
                "2001:0db8:0000:0042:0000:8a2e:0370:7337"
            ]
        }
    `
)

type IpV6Addr struct {
	address string `json:"address"`
}

type IpV6MultiAddr struct {
	addresses []string `json:"addresses"`
}

// Rseerves a IPV6 address for a client.  
//
// ASSUMPITON: Client ID is an INTEGER - NEED TO VERIFY
func ReserveIpV6Addr(clientId int) (string, error) {
    return HARDCODED_JSON, nil
    // TODO:Not working - fix
    //addr := IpV6Addr{address: "2001:0db8:0000:0042:0000:8a2e:0370:7334"}
    //b, err := json.Marshal(addr)
    //return string(b), err
}

// Rserves mutliple IPV6 addresses for a client given a number of requried addresses.  
//
// ASSUMPITON: Client ID is an INTEGER - NEED TO VERIFY
func ReserveIpV6MultiAddr(clientId int, num int) (string, error) {
    return HARDCODED_JSON_MULTI, nil
}
