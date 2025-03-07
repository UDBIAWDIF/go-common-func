package funcs

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/slice"
)

// DateFormat format time.Time
func DateFormat(format string, t time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06", // 2 位数字表示的年份

		// 月
		"m", "01", // 数字表示的月份，有前导零
		"n", "1", // 数字表示的月份，没有前导零
		"M", "Jan", // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2", // 月份中的第几天，没有前导零

		"D", "Mon", // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3", // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)
	return t.Format(format)
}

// StrToLocalTime get time.Time from string
func StrToLocalTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, nil
	}
	zoneName, offset := time.Now().Zone()

	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value += " " + zoneName
	}
	return StrToTime(value)
}

// StrToTime get time.Time from string
func StrToTime(value string, zone ...bool) (time.Time, error) {
	if value == "" {
		return time.Time{}, nil
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		"20060102",
		"20060102150405",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			if len(zone) > 0 {
				t = time.Unix(t.Unix(), 0)
			}
			return t, nil
		}
	}
	return time.Time{}, err
}

// 获取今天的日期, 格式 1979-01-01
func TodayDate() string {
	today := time.Now()
	return DateFormat("Y-m-d", today)
}

// 获取昨天的日期, 格式 1978-12-31
func YesterdayDate() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return DateFormat("Y-m-d", yesterday)
}

// 获取日期是周几的索引, 星期天为0, 失败时返回 -1
func WeekDayIdx(date string) int {
	idx := -1
	toTime, err := StrToLocalTime(date)
	if err == nil {
		idx = int(toTime.Weekday())
	}
	return idx
}

// 获取日期是周几的索引, 以普通人的阅读方式返回索引, 星期一为1, 星期天为7, 失败时返回 -1
func WeekDayIdxForHuman(date string) int {
	idx := WeekDayIdx(date)
	idx = If(idx == 0, 7, idx)
	return idx
}

// 时间转成秒数, 用来作当天的工作安排类的计算, 比如 time 为 1:00 , 结果为 60 秒
// 支持 xx:xx:xx 和 xx:xx
func TimeToSecond(timeStr string) int {
	timeSplit := strings.Split(timeStr, ":")
	slice.Reverse(timeSplit)
	var sumSecond int
	for idx, eachTimeSplit := range timeSplit {
		thisTime, _ := StringToInt(eachTimeSplit)
		sumSecond += thisTime * int(math.Pow(60, float64(idx)))
	}

	return sumSecond
}

// 暂时不能用, 验证失败
func IsDateTime(dateTimeStr string) bool {
	// re := regexp.MustCompile(`^d{4}([-/年])d{2}([-/月])d{2}([-/日]) (0?[1-9]|1[0-2])(:[0-5]d){2}$`)
	re := regexp.MustCompile(`^d{4}-d{2}-d{2} (0?[1-9]|1[0-2])(:[0-5]d){2}$`)
	return re.MatchString(dateTimeStr)
}

// 两个字符串时期的时间差(秒)
func StrTimeDifferenceSeconds(strTime1, strTime2 string) (difference int, err error) {
	time1, time1Err := StrToTime(strTime1)
	time2, time2Err := StrToTime(strTime2)

	if err = time1Err; err == nil {
		err = time2Err
	}

	if err == nil {
		difference = int(time1.Sub(time2).Seconds())
	}

	return
}

// 两个字符串时期的时间差(秒)
func StrTimeDifferenceSecondsNoError(strTime1, strTime2 string) (difference int) {
	difference, _ = StrTimeDifferenceSeconds(strTime1, strTime2)
	return
}

// 两个字符串日期的时间差(天)
func StrDateDifferenceDays(strTime1, strTime2 string) (difference int, err error) {
	time1, time1Err := StrToTime(strTime1)
	time2, time2Err := StrToTime(strTime2)

	if err = time1Err; err == nil {
		err = time2Err
	}

	if err == nil {
		difference = int(time1.Sub(time2).Hours() / 24)
	}

	return
}

// 两个字符串日期的时间差(天)
func StrDateDifferenceDaysNoError(strTime1, strTime2 string) (difference int) {
	difference, _ = StrDateDifferenceDays(strTime1, strTime2)
	return
}

// TimeFormat format time.Time
// 0: "2006-01-02 15:04:05"
// 1: "2006-01-02"
// 2: "15:04:05"
// 3: "20060102150405"
// 4: "2006-01-02 15:04"
// 5: "2006-01"
func TimeFormat(t time.Time, f int) (timeStr string) {
	switch f {
	case 0:
		timeStr = t.Format("2006-01-02 15:04:05")
	case 1:
		timeStr = t.Format("2006-01-02")
	case 2:
		timeStr = t.Format("15:04:05")
	case 3:
		timeStr = t.Format("20060102150405")
	case 4:
		timeStr = t.Format("2006-01-02 15:04")
	case 5:
		timeStr = t.Format("2006-01")
	}

	return
}

// 秒数转换成具体的时间
func SecondsToTimeFormat(seconds, f int) (timeStr string) {
	return TimeFormat(time.Unix(int64(seconds), 0), f)
}

// Now format now
func Now(f ...int) string {
	var format int
	if len(f) > 0 {
		format = f[0]
	} else {
		format = 0
	}
	return TimeFormat(time.Now(), format)
}

func LastDaTeByMonth(year int, month int) string {
	firstTimeOfMonth, _ := StrToTime(fmt.Sprintf("%d-%d", year, month))
	lastTimeOfMonth := firstTimeOfMonth.AddDate(0, 1, -1)
	return TimeFormat(lastTimeOfMonth, 1)
}
