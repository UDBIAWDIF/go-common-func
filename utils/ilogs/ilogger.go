package ilogs

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type LogDto struct {
	TraceId    string `json:"traceId"`
	HospitalId string `json:"hospitalId"`
	Phone      string `json:"phone"`
	Content    string `json:"content"`
	SmsType    string `json:"smsType"`
}

var logger *logrus.Entry

func init() {
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount 二者只能设置一个
	`WithMaxAge` 设置文件清理前的最长保存时间
	`WithRotationCount` 设置文件清理前最多保存的个数
	*/
	//下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。
	path := "/Users/opensource/test/go.log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	log.SetOutput(writer)
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	logger = log.WithFields(log.Fields{"request_id": "123444", "user_ip": "127.0.0.1"})
}

func write() {
	logger.Info("hello, logrus1....")
}
