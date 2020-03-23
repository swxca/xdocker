package process

import (
	"fmt"
	"os"
	"syscall"
)

func RunContainerInitProcess(command string) error {
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		return fmt.Errorf(err.Error())

	}
	return nil
}
