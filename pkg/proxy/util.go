package proxy

import (
	"net"
)

// ResolveAddress of host.
func ResolveAddress(sockAddr string) (*net.TCPAddr, error) {
	addr, err := net.ResolveTCPAddr("tcp", sockAddr)
	if err != nil {
		return nil, err
	}
	return addr, nil
}
