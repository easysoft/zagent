package main

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentService "github.com/easysoft/zagent/internal/agent/service"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"testing"
)

func TestDocker(t *testing.T) {
	_logUtils.Init(agentConst.AppName)

	agentConf.Inst.Host = "192.168.0.56"
	agentConf.Inst.User = "aaron"
	agentConf.Init()

	service := agentService.NewDockerService()

	imageName := "easysoft/zentao:15.0.2"

	service.PullImage(imageName)
	image, _ := service.GetImage(imageName)
	_logUtils.Infof("pull image %s", image.ID)

	resp, _ := service.CreateContainer(imageName, nil)
	_logUtils.Infof("create container %s", resp.ID)

	err := service.StartContainer(resp.ID)
	if err != nil {
		_logUtils.Infof("start container, err %s", err.Error())
	}

	container, _ := service.GetContainer(resp.ID)
	_logUtils.Infof("get container %s", container.ID)

	info, _ := service.GetContainerInfo(resp.ID)
	_logUtils.Infof("get container info %s on %d", info.Name, info.HttpPort)

	log, err := service.GetContainerLog(container.ID)
	if err != nil {
		_logUtils.Infof("get container log, err %s", err.Error())
	}
	_logUtils.Infof("get container log %s", log)

	err = service.StopContainer(container.ID)
	if err != nil {
		_logUtils.Infof("stop container , err %s", err.Error())
	}
	_logUtils.Infof("stop container %s", container.ID)

	err = service.RemoveContainer(container.ID, true, false, true)
	if err != nil {
		_logUtils.Infof("remove container, err %s", err.Error())
	}
	_logUtils.Infof("remove container %s", container.ID)

}
