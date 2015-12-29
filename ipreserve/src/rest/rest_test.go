package rest

import (
    "testing"
)

func TestReserveIpV6Addr(t *testing.T) {
    _, err := ReserveIpV6Addr(0)
    if err != nil {
        t.Error("ReserveIpV6Addr threw error, ", err)
    }
}

func TestReserveIpV6MultiAddr(t *testing.T) {
    _, err := ReserveIpV6MultiAddr(4, 0)
    if err != nil {
        t.Error("ResrveIpV6MultiAddr threw error, ", err)
    }
}