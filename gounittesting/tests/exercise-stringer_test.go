package tests

import (
	"fmt"
	"testing"
)

func TestIpAddr_String(t *testing.T) {
	h := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range h {
		fmt.Printf("%v: %v\n", name, ip)
	}
}