package main

import (
	"errors"

	"github.com/yuchanns/toybox/cgroup/subsystems"

	"github.com/yuchanns/toybox/container"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroups limit",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
	},

	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return errors.New("missing container command")
		}
		cmd := ctx.Args().Get(0)
		tty := ctx.Bool("it")
		resCfg := &subsystems.ResourceCfg{
			MemoryLimit: ctx.String("m"),
		}
		Run(tty, cmd, resCfg)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container.",

	Action: func(ctx *cli.Context) error {
		log.Info("init")
		cmd := ctx.Args().Get(0)
		log.Infof("command %s", cmd)
		return container.RunContainerInitProcess(cmd, nil)
	},
}
