package proxy

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

// LocalProxy is a proxy between two local ports
type LocalProxy struct {
	Name         string
	pAddr, cAddr *net.TCPAddr
	errsig       chan error
}

// NewLocal will build the TCP connections and return a local proxy
func NewLocal(name, proxyAddr, connectingAddr string, errsig chan error) (*LocalProxy, error) {
	pAddr, err := ResolveAddress(proxyAddr)
	if err != nil {
		return nil, err
	}

	cAddr, err := ResolveAddress(connectingAddr)
	if err != nil {
		return nil, err
	}

	localPorxy := &LocalProxy{
		Name:   name,
		pAddr:  pAddr,
		cAddr:  cAddr,
		errsig: errsig,
	}

	return localPorxy, nil
}

// Listen will attempt to connect to the connection socket address
func (lp *LocalProxy) Listen() {
	log.Infof("Listening on %s proxing to %s", lp.pAddr, lp.cAddr)

	// Start listening on proxy address
	listener, err := net.ListenTCP("tcp", lp.pAddr)
	if err != nil {
		lp.errsig <- err
		return
	}
	defer listener.Close()

	for {
		pConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer pConn.Close()

		proxyConn := &localProxyConn{
			proxyName: lp.Name,
			cAddr:     lp.cAddr,
			pConn:     pConn,
			errsig:    make(chan error),
		}
		go proxyConn.Dial()
	}
}

type localProxyConn struct {
	proxyName string
	cAddr     *net.TCPAddr
	pConn     *net.TCPConn
	errsig    chan error
}

func (lpc *localProxyConn) Dial() {
	defer lpc.pConn.Close()

	// Dial socket address
	cConn, err := net.DialTCP("tcp", nil, lpc.cAddr)
	if err != nil {
		log.Errorf("Dial error: %s", err)
		lpc.errsig <- err
		return
	}
	defer cConn.Close()

	// proxying data
	go lpc.pipe(lpc.pConn, cConn, "->")
	go lpc.pipe(cConn, lpc.pConn, "<-")

	<-lpc.errsig
}

func (lpc *localProxyConn) pipe(src, dst *net.TCPConn, prefix string) {
	buff := make([]byte, 0xffff)

	for {
		n, err := src.Read(buff)
		if err != nil {
			lpc.errsig <- err
			return
		}

		log.WithField("proxy", lpc.proxyName).Infof("%s %d bytes", prefix, n)

		n, err = dst.Write(buff[:n])
		if err != nil {
			lpc.errsig <- err
			return
		}
	}
}
