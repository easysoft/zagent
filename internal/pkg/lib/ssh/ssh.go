package _sshUtils

import (
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_resUtils "github.com/easysoft/zagent/internal/pkg/lib/res"
	"golang.org/x/crypto/ssh"
	"path/filepath"
)

func Connect(host, user string) (client *ssh.Client, err error) {
	privateKeyFile := filepath.Join("res", agentConst.AppName, "key", "id_rsa_test")
	privateKeyBytes, err := _resUtils.ReadRes(privateKeyFile)
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
