package util

import (
    "testing"
)

func TestConvertRangeToIPs(t *testing.T) {
    _, _, err := ConvertRangeToIPs("somestring")
    if err != nil {
        t.Error("ConvertRangeToIps threw error", err)
    }
}