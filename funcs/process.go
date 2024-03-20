package funcs

import (
	"os"
	"runtime"
	"syscall"

	"gitee.com/uid/go-common-func/utils"
	"github.com/shirou/gopsutil/process"
)

const OS_WIN = "windows"
const OS_LINUX = "linux"

var osIsProcessRunningMap map[string]func(int) bool = map[string]func(int) bool{
	OS_WIN:   IsProcessRunningWin,
	OS_LINUX: IsProcessRunningLinux,
}

// 判断进程是否存在
func IsProcessRunning(pid int) bool {
	return osIsProcessRunningMap[runtime.GOOS](pid)
}

// 判断进程是否存在
func IsProcessRunningLinux(pid int) bool {
	runner := utils.NewRunner()
	isRunning := false
	var process *os.Process
	err := error(nil)

	runner.Exec(func() bool {
		process, err = os.FindProcess(pid)
		return err == nil
	}).Exec(func() bool {
		err = process.Signal(syscall.Signal(0))
		return err == nil
	}).Success(func() {
		isRunning = true
	})

	return isRunning
}

func IsProcessRunningWin(pid int) bool {
	pids, _ := process.Pids()
	return Contains(pids, int32(pid))
}
