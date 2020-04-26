package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if os.Args[0] == "/proc/self/exe" {
		// equal to `sudo mount --make-private /`
		if err := syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""); err != nil {
			log.Fatal(err)
		}
		// forbid execute other process, nor set-user-ID, set-group-ID
		if err := syscall.Mount("proc", "/proc", "proc", uintptr(syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV), ""); err != nil {
			log.Fatal(err)
		}
		// syscall.Exec will replace the pid 1 process from `/proc/self/exe` to `/bin/sh`
		if err := syscall.Exec("/bin/sh", []string{"/bin/sh"}, os.Environ()); err != nil {
			log.Fatal(err)
		}
	}
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWUTS | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0, // 0 means user root
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0, // 0 means user root
				HostID:      0,
				Size:        1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
	os.Exit(0)
}
