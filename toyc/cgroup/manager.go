package cgroup

import "github.com/yuchanns/toybox/cgroup/subsystems"

type Manager struct {
	Path     string
	Resource *subsystems.ResourceCfg
}

func NewCgroupManager(path string) *Manager {
	return &Manager{
		Path: path,
	}
}

func (c *Manager) Apply(pid int) error {
	for _, subSysIns := range subsystems.Ins {
		if err := subSysIns.Apply(c.Path, pid); err != nil {
			return err
		}
	}

	return nil
}

func (c *Manager) Set(res *subsystems.ResourceCfg) error {
	for _, subSysIns := range subsystems.Ins {
		if err := subSysIns.Set(c.Path, res); err != nil {
			return err
		}
	}

	return nil
}

func (c *Manager) Destory() error {
	for _, subSysIns := range subsystems.Ins {
		if err := subSysIns.Remove(c.Path); err != nil {
			return err
		}
	}

	return nil
}
