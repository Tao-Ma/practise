package main

import "fmt"

type IPAddr [4]byte

// Question: addr *IPAddr, addr IPAddr are different. What the hell of type
//  keyword?
func (addr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v(a)\n", n, a)
		fmt.Printf("%v: %v(&a)\n", n, &a)
		fmt.Printf("%v: %v(a.String())\n", n, a.String())
	}
}
