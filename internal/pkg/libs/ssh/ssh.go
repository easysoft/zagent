package _sshUtils

import (
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

func Connect(host, user string) (client *ssh.Client, err error) {
	privateKeyFile := "xdoc/id_rsa_test"

	privateKeyBytes, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		_logUtils.Errorf(err.Error())
	}

	key, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		_logUtils.Errorf(err.Error())
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err = ssh.Dial("tcp", host+":22", config)
	if err != nil {
		_logUtils.Errorf("unable connect to ssh %s, err %s", host, err.Error())
	}

	return
}
