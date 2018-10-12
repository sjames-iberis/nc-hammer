package netconf

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

// Defines a factory method for instantiating netconf rpc sessions.

var (
	defaultLogger = log.New(os.Stderr, "logger:", log.Lshortfile)
)

// NewRpcSession connects to the  target using the ssh configuration, and establishes
// a netconf session with default configuration.
func NewRpcSession(sshcfg *ssh.ClientConfig, target string) (s Session, err error) {

	return NewRpcSessionWithConfig(sshcfg, target, defaultConfig)
}

// NewRpcSession connects to the  target using the ssh configuration, and establishes
// a netconf session with the client configuration.
func NewRpcSessionWithConfig(sshcfg *ssh.ClientConfig, target string, cfg *ClientConfig) (s Session, err error) {

	var t Transport
	if t, err = createTransport(sshcfg, target); err != nil {
		return
	}

	if s, err = NewSession(t, defaultLogger, defaultLogger, cfg); err != nil {
		t.Close() // nolint: gosec,errcheck
	}
	return
}

func createTransport(clientConfig *ssh.ClientConfig, target string) (t Transport, err error) {
	return NewSSHTransport(clientConfig, target, "netconf")
}
