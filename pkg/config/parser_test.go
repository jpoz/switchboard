package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ParseTestCase struct {
	File           string
	ExpectedConfig Config
	Err            error
}

var ParseTestCases = map[string]ParseTestCase{
	"empty": {
		File:           ``,
		ExpectedConfig: Config{},
		Err:            nil,
	},
	"only local": {
		File: `
localProxies:
    postgres:
        proxyAddress: '0.0.0.0:6543'
        connectingAddress: '127.0.0.1:5432'`,
		ExpectedConfig: Config{
			LocalProxies: map[string]LocalProxy{
				"postgres": {
					ProxyAddr:      "0.0.0.0:6543",
					ConnectingAddr: "127.0.0.1:5432",
				},
			},
		},
		Err: nil,
	},
	"broken config": {
		File: `
localProxies:
	broken: '0.0.0.0:6543'
	blab: '127.0.0.1:5432'`,
		ExpectedConfig: Config{},
		Err:            fmt.Errorf("yaml: line 3: found character that cannot start any token"),
	},
}

func TestParse(t *testing.T) {
	assert := assert.New(t)

	for name, testCase := range ParseTestCases {
		fmt.Println(testCase.File)
		c, err := Parse([]byte(testCase.File))

		assert.Equal(c, testCase.ExpectedConfig, fmt.Sprintf("%s test case config doesn't match", name))
		assert.Equal(err, testCase.Err, fmt.Sprintf("%s test err doesn't match", name))
	}
}
