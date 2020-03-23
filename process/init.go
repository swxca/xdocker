package process

import (
	"fmt"
	"os"
	"syscall"
)

func RunContainerInitProcess(commandPath string) error {
	argv := []string{commandPath}
	if err := syscall.Exec(commandPath, argv, os.Environ()); err != nil {
		return fmt.Errorf(err.Error())

	}
	return nil
}
