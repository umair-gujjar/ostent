// +build darwin

package ostent

/*
#include <sys/param.h> // for MAXCOMLEN
*/

import (
	"C"
	"os/exec"
	"path/filepath"
	"strings"
	"unsafe"

	sigar "github.com/ostrost/gosigar"
)

// Distrib returns system distribution and version string.
func Distrib() (string, error) {
	/* various cli to show the mac version
	sw_vers
	sw_vers -productVersion
	system_profiler SPSoftwareDataType
	defaults read loginwindow SystemVersionStampAsString
	defaults read /System/Library/CoreServices/SystemVersion ProductVersion
	*/
	std, err := exec.Command("sw_vers", "-productVersion").CombinedOutput()
	if err != nil {
		return "", err
	}
	return "Mac OS X " + strings.TrimRight(string(std), "\n\t "), nil
}

// ProcName returns argv[0].
// pbiComm originating from ProcState may be chopped, in which case
// sigar.ProcExe gives absolute executable path and the basename of it is returned.
func ProcName(pid int, pbiComm string) string {
	if len(pbiComm)+1 < C.MAXCOMLEN {
		return pbiComm
	}
	exe := sigar.ProcExe{}
	if err := exe.Get(pid); err != nil {
		return pbiComm
	}
	return filepath.Base(exe.Name)
}

func IfaDropsOut(unsafe.Pointer) uint { return 0 }