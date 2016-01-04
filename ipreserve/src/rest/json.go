package rest

import (
     "encoding/json"
)

type IpV6Addr struct {
	Address string `json:"address"`
}

type IpV6MultiAddr struct {
	Addresses []string `json:"addresses"`
}

// Rseerves a IPV6 address for a client.  
//
// ASSUMPITON: Client ID is an INTEGER - NEED TO VERIFY
func ReserveIpV6Addr(clientId int) (string, error) {
    
    address := IpV6Addr{Address: "2001:0db8:0000:0042:0000:8a2e:0370:7334"}
    b, err := json.Marshal(&address)
    return string(b), err
}

// Rserves mutliple IPV6 addresses for a client given a number of requried addresses.  
//
// ASSUMPITON: Client ID is an INTEGER - NEED TO VERIFY
func ReserveIpV6MultiAddr(clientId int, num int) (string, error) {
    
    addresses := IpV6MultiAddr{Addresses: []string{
            "2001:0db8:0000:0042:0000:8a2e:0370:7334",
            "2001:0db8:0000:0042:0000:8a2e:0370:7335",
            "2001:0db8:0000:0042:0000:8a2e:0370:7336",
            "2001:0db8:0000:0042:0000:8a2e:0370:7337",
    }}
    b, err := json.Marshal(&addresses)
    return string(b), err
}
