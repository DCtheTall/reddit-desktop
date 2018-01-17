package desktopimage

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func setOSXBackground(filename string) error {
	cmd := exec.Command(
		"/usr/bin/osascript",
		"-e",
		fmt.Sprintf(
			`tell application "System Events" to tell every desktop to set picture to "%s"`,
			filename,
		),
	)
	return cmd.Run()
}

func setLinuxBackground(filename string) error {
	return errors.New("Linux not supported yet")
}

func setWindowsBackground(filename string) error {
	return errors.New("Windows not supported yet")
}

/*
SetDesktopBackground of the users computer
*/
func SetDesktopBackground(filename string) error {
	switch runtime.GOOS {
	case "darwin":
		return setOSXBackground(filename)
	case "linux":
		return setLinuxBackground(filename)
	case "windows":
		return setWindowsBackground(filename)
	}
	return errors.New("Your operating system is not yet supported")
}
