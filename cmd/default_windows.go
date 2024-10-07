package cmd

import (
	"runtime"

	"github.com/muesli/termenv"
)

// OsInit enable colors for windows' user.
func OsInit() {
	if runtime.GOOS == "windows" {
		mode, err := termenv.EnableWindowsANSIConsole()
		if err != nil {
			panic(err)
		}
		defer termenv.RestoreWindowsConsole(mode)
	}
}
