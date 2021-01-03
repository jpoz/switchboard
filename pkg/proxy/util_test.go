package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveAddress(t *testing.T) {
	assert := assert.New(t)

	addr, err := ResolveAddress("0.0.0.0:1234")

	assert.Equal(addr.Port, 1234)
	assert.Nil(err)

	addr, err = ResolveAddress("im not an address")

	assert.Nil(addr)
	assert.EqualError(err, "address im not an address: missing port in address")
}
