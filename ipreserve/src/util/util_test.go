package util

import (
    "net"
    "testing"
)

func TestConvertRangeToIPs(t *testing.T) {
    
    // Test completely erroneous
    _, _, err := ConvertRangeToIPs("somestring")
    if err == nil {
        t.Error("Passes in something - expecting error")
    }
    
    // Test legal value
    a, b, err := ConvertRangeToIPs(
        "2001:0db8:0000:0042:0000:8a2e:0370:7334-2001:0db8:0000:0042:0000:8a2e:0370:7338")
    if err != nil {
        t.Error("ConvertRangeToIps threw error", err)
    }
    if a.String() != "2001:db8:0:42:0:8a2e:370:7334" {
        t.Errorf("ConvertRangeToIps, Expected %s, received %s", 
            "2001:0db8:0:42:0:8a2e:370:7334", a.String())
    }
    if b.String() != "2001:db8:0:42:0:8a2e:370:7338" {
        t.Errorf("Expected %s, received %s", 
            "2001:db8:0:42:0:8a2e:370:7338", b.String())
    }
}

func TestIncrementIP(t *testing.T) {
    
    // Test basic case
    ip := IncrementIP(net.ParseIP("2001:db8:0:42:0:8a2e:370:7334"))
    if ip.String() != "2001:db8:0:42:0:8a2e:370:7335" {
        t.Errorf("IncrementIps, Expected %s, received %s", 
            "2001:db8:0:42:0:8a2e:370:7335", ip.String())
    }
    
    // Test basic roll in last digits
    ip = IncrementIP(net.ParseIP("2001:db8:0:42:0:8a2e:370:FFFF"))
    if ip.String() != "2001:db8:0:42:0:8a2e:371:0" {
        t.Errorf("IncrementIps, Expected %s, received %s", 
            "2001:db8:0:42:0:8a2e:371:0", ip.String())
    }

    // Test rolling first digits
    ip = IncrementIP(net.ParseIP("2001:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF"))
    if ip.String() != "2002::" {
        t.Errorf("IncrementIps, Expected %s, received %s", 
            "2002:0:0:0:0:0:0:0", ip.String())
    }
    
    // Test complete rollover
    ip = IncrementIP(net.ParseIP("FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF"))
    if ip.String() != "::" {
        t.Errorf("IncrementIps, Expected %s, received %s", 
            "::", ip.String())
    }
}