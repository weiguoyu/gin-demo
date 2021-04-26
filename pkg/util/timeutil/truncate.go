package timeutil

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type TimeUnit int

const (
	zero = 0
	one  = 1

	Second = iota
	Minute
	Hour
	Day
	Month
	Year
)

func TruncateThirtyMinutes(currentTime time.Time) time.Time {
	minutes := currentTime.Minute()
	if minutes >= 30 {
		minutes = 30
	} else {
		minutes = 0
	}
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), minutes, zero, zero, currentTime.Location())
}

func TruncateHour(currentTime time.Time) time.Time {
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), zero, zero, zero, currentTime.Location())
}

func TruncateDay(currentTime time.Time) time.Time {
	year, month, day := currentTime.Date()
	return time.Date(year, month, day, zero, zero, zero, zero, currentTime.Location())
}

func TruncateMonth(currentTime time.Time) time.Time {
	year, month, _ := currentTime.Date()
	return time.Date(year, month, one, zero, zero, zero, zero, currentTime.Location())
}

func FloorTimeByMinutes(currentTime time.Time, interval int) (time.Time, error) {
	switch {
	case interval%60 == 0 && interval/60 == 1:
		return currentTime.Truncate(time.Hour), nil
	case interval%1440 == 0 && interval/1440 == 1:
		return TruncateDay(currentTime), nil
	case interval%43200 == 0 && interval/43200 == 1:
		return TruncateMonth(currentTime), nil
	default:
		return currentTime, fmt.Errorf("不支持的值: %v", interval)
	}
}

func FloorTime(currentTime time.Time, interval int, unit TimeUnit) (time.Time, error) {
	if interval <= 0 {
		return currentTime, errors.New("value应该大于0")
	}

	err := errors.New("value值超出unit范围")

	year, month, day := currentTime.Date()
	location := currentTime.Location()

	switch unit {
	case Second:
		if interval > 59 {
			return currentTime, err
		}
		if 60%interval != 0 {
			return currentTime, fmt.Errorf("秒钟不能按%v整分", interval)
		}
		startSecond := floorInt(float64(currentTime.Second()), float64(interval))
		return time.Date(year, month, day, currentTime.Hour(), currentTime.Minute(), startSecond, zero, location), nil
	case Minute:
		if interval > 59 {
			return currentTime, err
		}
		if 60%interval != 0 {
			return currentTime, fmt.Errorf("分钟不能按%v整分", interval)
		}
		startMinute := floorInt(float64(currentTime.Minute()), float64(interval))
		return time.Date(year, month, day, currentTime.Hour(), startMinute, zero, zero, location), nil
	case Hour:
		if interval > 23 {
			return currentTime, err
		}
		if 24%interval != 0 {
			return currentTime, fmt.Errorf("时钟不能按%v整分", interval)
		}
		startHour := floorInt(float64(currentTime.Hour()), float64(interval))
		return time.Date(year, month, day, startHour, zero, zero, zero, location), nil
	case Day:
		if interval > 15 {
			return currentTime, err
		}
		if 30%interval != 0 {
			return currentTime, fmt.Errorf("日期不能按%v整分", interval)
		}
		startDay := floorInt(float64(day-1), float64(interval))
		return time.Date(year, month, startDay+1, zero, zero, zero, zero, location), nil
	case Month:
		if interval > 6 {
			return currentTime, err
		}
		if 12%interval != 0 {
			return currentTime, fmt.Errorf("月份不能按%v整分", interval)
		}
		startMonth := floorInt(float64(month-1), float64(interval))
		return time.Date(year, time.Month(startMonth)+1, one, zero, zero, zero, zero, location), nil
	case Year:
		if interval > 3 {
			return currentTime, err
		}
		startYear := floorInt(float64(year-1), float64(interval))
		return time.Date(startYear, time.Month(one), one, zero, zero, zero, zero, location), nil
	default:
		return currentTime, errors.New("不支持的时间单位")
	}
}

func floorInt(src float64, interval float64) int {
	return int(math.Floor(src/interval) * interval)
}
