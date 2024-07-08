package out

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InstallPackages(packages []string) error {
	installer := ""
	installCmd := ""

	switch runtime.GOOS {
	case "darwin", "linux":
		installer = "brew"
		installCmd = "install"
	case "windows":
		installer = "choco"
		installCmd = "install"
	default:
		return fmt.Errorf("unsupported platform")
	}

	for _, pkg := range packages {
		cmd := exec.Command(installer, installCmd, pkg)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install %s: %v", pkg, err)
		}
	}

	return nil
}
