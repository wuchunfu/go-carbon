package carbon

import (
	"time"
)

// SetTimezone sets timezone.
// 设置时区
func (c Carbon) SetTimezone(name string) Carbon {
	if c.Error != nil {
		return c
	}
	c.loc, c.Error = getLocationByTimezone(name)
	return c
}

// SetTimezone sets timezone.
// 设置时区
func SetTimezone(name string) Carbon {
	return NewCarbon().SetTimezone(name)
}

// SetLocale sets locale.
// 设置语言区域
func (c Carbon) SetLocale(locale string) Carbon {
	if c.Error != nil {
		return c
	}
	c.lang.SetLocale(locale)
	c.Error = c.lang.Error
	return c
}

// SetLocale sets locale.
// 设置语言区域
func SetLocale(locale string) Carbon {
	c := NewCarbon()
	c.lang.SetLocale(locale)
	c.Error = c.lang.Error
	return c
}

// SetLanguage sets language.
// 设置语言对象
func (c Carbon) SetLanguage(lang *Language) Carbon {
	if c.Error != nil {
		return c
	}
	c.lang, c.Error = lang, lang.Error
	return c
}

// SetLanguage sets language.
// 设置语言对象
func SetLanguage(lang *Language) Carbon {
	c := NewCarbon()
	lang.SetLocale(lang.locale)
	c.lang, c.Error = lang, lang.Error
	return c
}

// SetDateTime sets year, month, day, hour, minute and second.
// 设置年月日时分秒
func (c Carbon) SetDateTime(year, month, day, hour, minute, second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetDateTimeMilli sets year, month, day, hour, minute, second and millisecond.
// 设置年月日时分秒毫秒
func (c Carbon) SetDateTimeMilli(year, month, day, hour, minute, second, millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(year, month, day, hour, minute, second, millisecond*1e6)
}

// SetDateTimeMicro sets year, month, day, hour, minute, second and microsecond.
// 设置年月日时分秒微秒
func (c Carbon) SetDateTimeMicro(year, month, day, hour, minute, second, microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(year, month, day, hour, minute, second, microsecond*1e3)
}

// SetDateTimeNano sets year, month, day, hour, minute, second and nanosecond.
// 设置年月日时分秒纳秒
func (c Carbon) SetDateTimeNano(year, month, day, hour, minute, second, nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(year, month, day, hour, minute, second, nanosecond)
}

// SetDate sets year, month and day.
// 设置年月日
func (c Carbon) SetDate(year, month, day int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetTime sets hour, minute and second.
// 设置时分秒
func (c Carbon) SetTime(hour, minute, second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetYear sets year.
// 设置年份
func (c Carbon) SetYear(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	_, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetYearNoOverflow sets year without overflowing month.
// 设置年份(月份不溢出)
func (c Carbon) SetYearNoOverflow(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddYearsNoOverflow(year - c.Year())
}

// SetMonth sets month.
// 设置月份
func (c Carbon) SetMonth(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, _, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetMonthNoOverflow sets month without overflowing month.
// 设置月份(月份不溢出)
func (c Carbon) SetMonthNoOverflow(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddMonthsNoOverflow(month - c.Month())
}

// SetWeekStartsAt sets start day of the week.
// 设置一周的开始日期
func (c Carbon) SetWeekStartsAt(day string) Carbon {
	if c.IsInvalid() {
		return c
	}
	switch day {
	case Sunday:
		c.weekStartsAt = time.Sunday
	case Monday:
		c.weekStartsAt = time.Monday
	case Tuesday:
		c.weekStartsAt = time.Tuesday
	case Wednesday:
		c.weekStartsAt = time.Wednesday
	case Thursday:
		c.weekStartsAt = time.Thursday
	case Friday:
		c.weekStartsAt = time.Friday
	case Saturday:
		c.weekStartsAt = time.Saturday
	}
	return c
}

// SetDay sets day.
// 设置日期
func (c Carbon) SetDay(day int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetHour sets hour.
// 设置小时
func (c Carbon) SetHour(hour int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, _, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetMinute sets minute.
// 设置分钟
func (c Carbon) SetMinute(minute int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, _, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetSecond sets second.
// 设置秒数
func (c Carbon) SetSecond(second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetMillisecond sets millisecond.
// 设置毫秒
func (c Carbon) SetMillisecond(millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, millisecond*1e6)
}

// SetMicrosecond sets microsecond.
// 设置微秒
func (c Carbon) SetMicrosecond(microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, microsecond*1e3)
}

// SetNanosecond sets nanosecond.
// 设置纳秒
func (c Carbon) SetNanosecond(nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, nanosecond)
}
