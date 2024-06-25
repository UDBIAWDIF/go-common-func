package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const DATA_PLACE_HOLE = "_DATA_PLACE_HOLE_"

type RotateLog struct {
	lastDate     string   // 上个日期
	logDir       string   // 日志存放目录
	logBackupDir string   // 历史日志备份目录
	fileTpl      string   // 文件名模板, 例如 runtime-_DATA_PLACE_HOLE_.log
	fileHandle   *os.File // 日志文件句柄
	lock         *SpinLock
}

/**
 *	获取 RotateLog 新实例
 *
 * @param logDir 日志存放目录, 带不带结尾的目录根号都行, 正反根号也都行
 * @param fileTpl 文件名模板, 字符串中常量 DATA_PLACE_HOLE 为日期的位置, 会被替换成当前日期, 例如 runtime-_DATA_PLACE_HOLE_.log
 *
 **/
func NewRotateLog(logDir, fileTpl string) *RotateLog {
	logDir = strings.TrimRight(logDir, "/")
	logDir = strings.TrimRight(logDir, `\`)
	logDir = strings.ReplaceAll(logDir, `\`, "/")
	logHandel := &RotateLog{
		logDir:  logDir,
		fileTpl: fileTpl,
		lock:    &SpinLock{},
	}
	return logHandel
}

func (_l *RotateLog) SetBackupDir(backupDir string) *RotateLog {
	_l.logBackupDir = backupDir
	return _l
}

func (_l *RotateLog) Write(p []byte) (int, error) {
	_l.makeSureLogFile()
	return _l.fileHandle.Write(p)
}

// 确保日志文件正常建立
func (_l *RotateLog) makeSureLogFile() {
	today := time.Now().Format("2006-01-02")
	for {
		if !_l.lock.TryLock() {
			continue
		}

		if _l.lastDate != today {
			if _l.fileHandle != nil {
				_l.fileHandle.Close()
				go func() {
					oldFile := _l.fileHandle.Name()
					fmt.Printf("请实现备份旧日志文件的功能, 未被备份的日志文件为: %s\n", oldFile)
					// TODO: 移动到备份目录去
					// _l.fileHandle.Name()
				}()
			}

			os.MkdirAll(_l.logDir, os.ModePerm)
			logFileName := strings.Replace(_l.fileTpl, DATA_PLACE_HOLE, today, 1)
			logFilePath := fmt.Sprintf("%s/%s", _l.logDir, logFileName)
			_l.fileHandle, _ = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			_l.lastDate = today
		}

		_l.lock.Unlock()
		break
	}
}

// 确保日志文件正常建立
func (_l *RotateLog) usage() string {
	return `

	func configSetLog() {
		logDir := getWorkDir() + "/logs"
		logFileNameTpl := fmt.Sprintf("runtime-%s.log", rotatelog.DATA_PLACE_HOLE)
		rotateLog := rotatelog.NewRotateLog(logDir, logFileNameTpl)
		mw := io.MultiWriter(os.Stdout, rotateLog)
		log.SetOutput(mw)
		log.SetFlags(log.Llongfile | log.Ldate | log.Ltime | log.Lshortfile)
	}

`
}
