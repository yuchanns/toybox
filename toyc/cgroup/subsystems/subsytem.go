package subsystems

type Subsystem interface {
	Name() string
	Set(path string, res *ResourceCfg) error
	Apply(path string, pid int) error
	Remove(path string) error
}

var Ins = []Subsystem{
	&MemorySubsystem{},
}
