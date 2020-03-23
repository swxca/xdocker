package process

import (
	"os"
	"os/exec"
	"syscall"
)

func CreateProcess(i bool,command string) *exec.Cmd {
	cmd:=exec.Command("/proc/self/exe","init",command)
	cmd.SysProcAttr=&syscall.SysProcAttr{
		Cloneflags:syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if i{
		cmd.Stdin=os.Stdin
		cmd.Stdout=os.Stdout
		cmd.Stderr=os.Stderr
	}
	return cmd
}
