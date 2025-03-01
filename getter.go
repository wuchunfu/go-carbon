package carbon

import (
	"time"
)

// DaysInYear gets total days in year like 365.
// 获取本年的总天数
func (c Carbon) DaysInYear() int {
	if c.IsInvalid() {
		return 0
	}
	if c.IsLeapYear() {
		return DaysPerLeapYear
	}
	return DaysPerNormalYear
}

// DaysInMonth gets total days in month like 30.
// 获取本月的总天数
func (c Carbon) DaysInMonth() int {
	if c.IsInvalid() {
		return 0
	}
	return c.EndOfMonth().Carbon2Time().Day()
}

// MonthOfYear gets month of year like 12.
// 获取本年的第几月
func (c Carbon) MonthOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	return int(c.Carbon2Time().Month())
}

// DayOfYear gets day of year like 365.
// 获取本年的第几天
func (c Carbon) DayOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().YearDay()
}

// DayOfMonth gets day of month like 30.
// 获取本月的第几天
func (c Carbon) DayOfMonth() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().Day()
}

// DayOfWeek gets day of week like 6.
// 获取本周的第几天
func (c Carbon) DayOfWeek() int {
	if c.IsInvalid() {
		return 0
	}
	day := int(c.Carbon2Time().Weekday())
	if day == 0 {
		return DaysPerWeek
	}
	return day
}

// WeekOfYear gets week of year like 1, see https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
// 获取本年的第几周
func (c Carbon) WeekOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	_, week := c.Carbon2Time().ISOWeek()
	return week
}

// WeekOfMonth gets week of month like 1.
// 获取本月的第几周
func (c Carbon) WeekOfMonth() int {
	if c.IsInvalid() {
		return 0
	}
	days := c.Day() + c.StartOfMonth().DayOfWeek() - 1
	if days%DaysPerWeek == 0 {
		return days / DaysPerWeek
	}
	return days/DaysPerWeek + 1
}

// DateTime gets current year, month, day, hour, minute, and second like 2020, 8, 5, 13, 14, 15.
// 获取当前年月日时分秒
func (c Carbon) DateTime() (year, month, day, hour, minute, second int) {
	if c.IsInvalid() {
		return
	}
	t := c.Carbon2Time()
	var tm time.Month
	year, tm, day = t.Date()
	hour, minute, second = t.Clock()
	return year, int(tm), day, hour, minute, second
}

// DateTimeMilli gets current year, month, day, hour, minute, second and millisecond like 2020, 8, 5, 13, 14, 15, 999.
// 获取当前年月日时分秒毫秒
func (c Carbon) DateTimeMilli() (year, month, day, hour, minute, second, millisecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day, hour, minute, second = c.DateTime()
	return year, month, day, hour, minute, second, c.Millisecond()
}

// DateTimeMicro gets current year, month, day, hour, minute, second and microsecond like 2020, 8, 5, 13, 14, 15, 999999.
// 获取当前年月日时分秒微秒
func (c Carbon) DateTimeMicro() (year, month, day, hour, minute, second, microsecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day, hour, minute, second = c.DateTime()
	return year, month, day, hour, minute, second, c.Microsecond()
}

// DateTimeNano gets current year, month, day, hour, minute, second and nanosecond like 2020, 8, 5, 13, 14, 15, 999999999.
// 获取当前年月日时分秒纳秒
func (c Carbon) DateTimeNano() (year, month, day, hour, minute, second, nanosecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day, hour, minute, second = c.DateTime()
	return year, month, day, hour, minute, second, c.Nanosecond()
}

// Date gets current year, month, and day like 2020, 8, 5.
// 获取当前年月日
func (c Carbon) Date() (year, month, day int) {
	if c.IsInvalid() {
		return
	}
	var tm time.Month
	year, tm, day = c.Carbon2Time().Date()
	return year, int(tm), day
}

// Time gets current hour, minute, and second like 13, 14, 15.
// 获取当前时分秒
func (c Carbon) Time() (hour, minute, second int) {
	if c.IsInvalid() {
		return
	}
	return c.Carbon2Time().Clock()
}

// Century gets current century like 21.
// 获取当前世纪
func (c Carbon) Century() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Year()/YearsPerCentury + 1
}

// Decade gets current decade like 20.
// 获取当前年代
func (c Carbon) Decade() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Year() % YearsPerCentury / YearsPerDecade * YearsPerDecade
}

// Year gets current year like 2020.
// 获取当前年
func (c Carbon) Year() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().Year()
}

// Quarter gets current quarter like 3.
// 获取当前季度
func (c Carbon) Quarter() (quarter int) {
	if c.IsInvalid() {
		return
	}
	month := c.Month()
	switch {
	case month >= 10:
		quarter = 4
	case month >= 7:
		quarter = 3
	case month >= 4:
		quarter = 2
	case month >= 1:
		quarter = 1
	}
	return
}

// Month gets current month like 8.
// 获取当前月
func (c Carbon) Month() int {
	return c.MonthOfYear()
}

// Week gets current week like 6, start from 0.
// 获取当前周(从0开始)
func (c Carbon) Week() int {
	if c.IsInvalid() {
		return -1
	}
	return (c.DayOfWeek() + DaysPerWeek - int(c.weekStartsAt)) % DaysPerWeek
}

// Day gets current day like 5.
// 获取当前日
func (c Carbon) Day() int {
	return c.DayOfMonth()
}

// Hour gets current hour like 13.
// 获取当前小时
func (c Carbon) Hour() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().Hour()
}

// Minute gets current minute like 14.
// 获取当前分钟数
func (c Carbon) Minute() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().Minute()
}

// Second gets current second like 15.
// 获取当前秒数
func (c Carbon) Second() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().Second()
}

// Millisecond gets current millisecond like 999.
// 获取当前毫秒数，3位数字
func (c Carbon) Millisecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().Nanosecond() / 1e6
}

// Microsecond gets current microsecond like 999999.
// 获取当前微秒数，6位数字
func (c Carbon) Microsecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().Nanosecond() / 1e3
}

// Nanosecond gets current nanosecond like 999999999.
// 获取当前纳秒数，9位数字
func (c Carbon) Nanosecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Carbon2Time().Nanosecond()
}

// TimestampWithSecond gets timestamp with second like 1596604455, it will be replaced by Timestamp.
// 输出秒级时间戳
func (c Carbon) TimestampWithSecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.time.Unix()
}

// Timestamp gets timestamp with second like 1596604455, it is shorthand for TimestampWithSecond.
// 获取秒级时间戳, 是 TimestampWithSecond 的简写
func (c Carbon) Timestamp() int64 {
	return c.TimestampWithSecond()
}

// TimestampWithMillisecond gets timestamp with millisecond like 1596604455000, it will be replaced by TimestampMilli.
// 获取毫秒级时间戳
func (c Carbon) TimestampWithMillisecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.time.Unix()*1e3 + int64(c.time.Nanosecond())/1e6
}

// TimestampMilli gets timestamp with millisecond like 1596604455000, it is shorthand for TimestampWithMillisecond.
// 获取毫秒级时间戳, 是 TimestampMilli 的简写
func (c Carbon) TimestampMilli() int64 {
	return c.TimestampWithMillisecond()
}

// TimestampWithMicrosecond gets timestamp with microsecond like 1596604455000000, it will be replaced by TimestampMicro.
// 获取微秒级时间戳
func (c Carbon) TimestampWithMicrosecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.time.Unix()*1e6 + int64(c.time.Nanosecond())/1e3
}

// TimestampMicro gets timestamp with microsecond like 1596604455000000, it is shorthand for TimestampWithMicrosecond.
// 获取微秒级时间戳, 是 TimestampWithMicrosecond 的简写
func (c Carbon) TimestampMicro() int64 {
	return c.TimestampWithMicrosecond()
}

// TimestampWithNanosecond gets timestamp with nanosecond like 1596604455000000000, it will be replaced by TimestampNano.
// 获取纳秒级时间戳
func (c Carbon) TimestampWithNanosecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.time.UnixNano()
}

// TimestampNano gets timestamp with nanosecond like 1596604455000000000, it is shorthand for TimestampWithNanosecond.
// 获取纳秒级时间戳, 是 TimestampWithNanosecond 的简写
func (c Carbon) TimestampNano() int64 {
	return c.TimestampWithNanosecond()
}

// Location gets location name like "PRC".
// 获取位置
func (c Carbon) Location() string {
	return c.loc.String()
}

// Timezone gets timezone name like "CST".
// 获取时区
func (c Carbon) Timezone() string {
	name, _ := c.time.Zone()
	return name
}

// Offset gets offset seconds from the UTC timezone like 28800.
// 获取距离UTC时区的偏移量，单位秒
func (c Carbon) Offset() int {
	_, offset := c.time.Zone()
	return offset
}

// Locale gets locale name like "zh-CN".
// 获取语言区域
func (c Carbon) Locale() string {
	return c.lang.locale
}

// Age gets age like 18.
// 获取年龄
func (c Carbon) Age() int {
	if c.IsInvalid() {
		return 0
	}
	now := c.Now()
	if c.Timestamp() > now.Timestamp() {
		return 0
	}
	return int(c.DiffInYears(now))
}
