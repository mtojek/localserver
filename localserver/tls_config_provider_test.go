package localserver

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	caPemPath := "resources/certs/server_ca.pem"
	caKeyPath := "resources/certs/server_ca.key"
	sut := newTLSConfigProvider()

	// when
	config := sut.Provide(caPemPath, caKeyPath)

	// then
	assert.Equal(config.ClientAuth, tls.NoClientCert, "Client certicate is not required")
	assert.NotNil(config.Rand, "Rand must be specified")
	assert.Len(config.Certificates, 1, "One certificate is expected")
}
