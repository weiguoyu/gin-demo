package timeutil

import (
	"fmt"
	"time"

	cons "gin-demo/pkg/constants"
)

var MysqlMaxTimestamp time.Time

const (
	DefaultTimeFormat = "2006-01-02T15:04:05.000Z"
	YearDateFormat    = "2006-01"
	DurationDay       = 24 * time.Hour
	DurationWeek      = 7 * DurationDay
	DurationMonth     = 30 * DurationDay
	DurationYear      = 365 * DurationDay
)

//mysql 时间戳最大值只能到 2038-01-19 03:14:07
var MaxMysqlTimestamp = time.Date(2038, time.Month(1), 19, 3, 14, 7, 0, time.Local)

func init() {
	MysqlMaxTimestamp, _ = time.Parse(cons.DefaultTimeLayout, cons.NoneTime)
}

// Todo:
//   Should return by parameters
func GetEndTimeThisPeriod(startTime time.Time, period uint32, periodUnit string) (time.Time, error) {
	periodUnitMap := map[string]time.Duration{
		"小时":     time.Hour * time.Duration(period),
		"分":      time.Minute * time.Duration(period),
		"天":      time.Hour * time.Duration(period*24),
		"minute": time.Minute * time.Duration(period),
	}
	endTime := startTime.Add(periodUnitMap[periodUnit])

	return endTime, nil
}

// Todo:
//   Should return by parameters
func GetStepThisPeriod(period uint32, periodUnit string) (string, error) {
	return "1 d", nil
}

// if the timestamp is '2038-01-19 03:14:07', it is end time in mysql.
func IsMysqlMaxTime(t time.Time) bool {
	return t.Equal(MysqlMaxTimestamp)
}

func FromPbTimestampMap(m map[string]interface{}) time.Time {
	var t time.Time
	if m == nil {
		t = time.Unix(0, 0).UTC() // treat nil like the empty Timestamp
	} else {
		t = time.Unix(int64(m["seconds"].(float64)), int64(m["nanos"].(float64))).UTC()
	}
	return t
}

func TimeToCronExpr(t time.Time) string {
	return fmt.Sprintf("%v %v %v %v %v ?", t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()))
}

func TimePtr(tt time.Time) *time.Time {
	return &tt
}

//获取当前时间，截断到秒
func Now() time.Time {
	return time.Now().Truncate(time.Second)
}
