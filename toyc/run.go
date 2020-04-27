package main

import (
	"os"

	"github.com/yuchanns/toybox/cgroup"

	"github.com/yuchanns/toybox/cgroup/subsystems"

	log "github.com/sirupsen/logrus"
	"github.com/yuchanns/toybox/container"
)

func Run(tty bool, command string, res *subsystems.ResourceCfg) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	cgroupManager := cgroup.NewCgroupManager("toyc-cgroup")
	defer cgroupManager.Destory()
	if err := cgroupManager.Set(res); err != nil {
		log.Fatal(err)
	}

	if err := cgroupManager.Apply(parent.Process.Pid); err != nil {
		log.Fatal(err)
	}

	parent.Wait()
	os.Exit(0)
}
