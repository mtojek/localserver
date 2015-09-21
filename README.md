# localserver
[![Build Status](https://travis-ci.org/mtojek/localserver.svg?branch=master)](https://travis-ci.org/mtojek/localserver)

Go HTTP/HTTPS server, useful in unit tests

Status: **Done**

Keywords: Go HTTP server, Go HTTPS server, TLS, unit tests

Simple implementation of HTTP/HTTPS server useful in unit tests. It supports starting nad stopping the listener. User can provide its own server certificates.

## Features

* HTTP and HTTPS support
* x509 certificates support
* Really simple usage

## Sample usage

~~~
func TestStartHTTPSServer(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "127.0.0.1:9000"
	scheme := "https"
	sut := NewLocalServer(hostPort, scheme)

	// when
	sut.StartTLS("resources/certs/server_ca.pem", "resources/certs/server_ca.key")
	isRunning := checkIfLocalServerIsRunning(scheme + "://" + hostPort)

	// then
	sut.Stop()

	assert.True(isRunning, "Server should be available now")
}
~~~
