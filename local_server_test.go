package localserver

import (
	"testing"

	"crypto/tls"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestNewLocalServer(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "hostPort"
	scheme := "proto"

	// when
	sut := NewLocalServer(hostPort, scheme)

	// then
	assert.Equal(sut.hostPort, hostPort)
	assert.Equal(sut.scheme, scheme)
	assert.NotNil(sut.tlsConfigProvider, "TLS config provider should be set")
}

func TestStartServer(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "127.0.0.1:59001"
	scheme := "http"
	sut := NewLocalServer(hostPort, scheme)

	// when
	sut.Start()
	isRunning := checkIfLocalServerIsRunning(scheme + "://" + hostPort)

	// then
	sut.Stop()

	assert.True(isRunning, "Server should be available now")
}

func TestStartHTTPServer(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "127.0.0.1:59002"
	scheme := "http"
	sut := NewLocalServer(hostPort, scheme)

	// when
	sut.StartHTTP()
	isRunning := checkIfLocalServerIsRunning(scheme + "://" + hostPort)

	// then
	sut.Stop()

	assert.True(isRunning, "Server should be available now")
}

func TestStartHTTPSServer(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "127.0.0.1:59003"
	scheme := "https"
	sut := NewLocalServer(hostPort, scheme)

	// when
	sut.StartTLS("resources/certs/server_ca.pem", "resources/certs/server_ca.key")
	isRunning := checkIfLocalServerIsRunning(scheme + "://" + hostPort)

	// then
	sut.Stop()

	assert.True(isRunning, "Server should be available now")
}

func TestStopServer(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "127.0.0.1:59004"
	scheme := "http"
	sut := NewLocalServer(hostPort, scheme)
	sut.Start()

	// when
	sut.Stop()
	isRunning := checkIfLocalServerIsRunning(scheme + "://" + hostPort)

	// then
	assert.False(isRunning, "Server should not be available now")
}

func checkIfLocalServerIsRunning(url string) bool {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{
		Transport: tr,
	}

	_, error := client.Get(url)
	return nil == error
}
