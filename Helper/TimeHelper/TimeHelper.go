package TimeHelper

import (
	"time"
)

const dateFormat = "2006-01-02"
const timeFormat = "2006-01-02 15:04:05"
const timeFormatWithMiliSecond = "2006-01-02 15:04:05.000"

func GetUtcTimeNowString() string {
	return time.Now().UTC().Format(timeFormat)
}

func GetUTC8TimeNowString() string {
	return time.Now().UTC().Add(time.Hour * time.Duration(8)).Format(timeFormat)
}

func GetUTC8DateNowString() string {
	return time.Now().Format(dateFormat)
}

func GetUTC8TimeNow() time.Time {
	return time.Now().UTC().Add(time.Hour * time.Duration(8))
}

func TimeToString(t time.Time) string {
	return t.Format(timeFormat)
}

func TimeToStringWithMiliSecond(t time.Time) string {
	return t.Format(timeFormatWithMiliSecond)
}

func ParseStringToTime(str string) (time.Time, error) {
	return time.Parse(timeFormat, str)
}

func GetMillisecondTimestamp() int64 {
	return time.Now().UnixMilli()
}
