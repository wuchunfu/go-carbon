package carbon

import (
	"bytes"
	"strconv"
	"strings"
)

// String outputs a string in date and time format, implement Stringer interface.
// 实现 Stringer 接口
func (c Carbon) String() string {
	return c.ToDateTimeString()
}

// ToString outputs a string in "2006-01-02 15:04:05.999999999 -0700 MST" layout.
// 输出 "2006-01-02 15:04:05.999999999 -0700 MST" 格式字符串
func (c Carbon) ToString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().String()
}

// ToMonthString outputs a string in month layout like "January", i18n is supported.
// 输出完整月份字符串，支持i18n
func (c Carbon) ToMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["months"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 12 {
			return slice[c.Month()-1]
		}
	}
	return ""
}

// ToShortMonthString outputs a string in short month layout like "Jan", i18n is supported.
// 输出缩写月份字符串，支持i18n
func (c Carbon) ToShortMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["short_months"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 12 {
			return slice[c.Month()-1]
		}
	}
	return ""
}

// ToWeekString outputs a string in week layout like "Sunday", i18n is supported.
// 输出完整星期字符串，支持i18n
func (c Carbon) ToWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["weeks"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 7 {
			return slice[c.Week()]
		}
	}
	return ""
}

// ToShortWeekString outputs a string in short week layout like "Sun", i18n is supported.
// 输出缩写星期字符串，支持i18n
func (c Carbon) ToShortWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["short_weeks"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 7 {
			return slice[c.Week()]
		}
	}
	return ""
}

// ToDayDateTimeString outputs a string in "Mon, Jan 2, 2006 3:04 PM" layout.
// 输出 "Mon, Jan 2, 2006 3:04 PM" 格式字符串
func (c Carbon) ToDayDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DayDateTimeLayout)
}

// ToDateTimeString outputs a string in "2006-01-02 15:04:05" layout.
// 输出 "2006-01-02 15:04:05" 格式字符串
func (c Carbon) ToDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateTimeLayout)
}

// ToDateTimeMilliString outputs a string in "2006-01-02 15:04:05.999" layout.
// 输出 "2006-01-02 15:04:05.999" 格式字符串
func (c Carbon) ToDateTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateTimeMilliLayout)
}

// ToDateTimeMicroString outputs a string in "2006-01-02 15:04:05.999999" layout.
// 输出 "2006-01-02 15:04:05.999999" 格式字符串
func (c Carbon) ToDateTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateTimeMicroLayout)
}

// ToDateTimeNanoString outputs a string in "2006-01-02 15:04:05.999999999" layout.
// 输出 "2006-01-02 15:04:05.999999999" 格式字符串
func (c Carbon) ToDateTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateTimeNanoLayout)
}

// ToShortDateTimeString outputs a string in "20060102150405" layout.
// 输出 "20060102150405" 格式字符串
func (c Carbon) ToShortDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateTimeLayout)
}

// ToShortDateTimeMilliString outputs a string in "20060102150405.999" layout.
// 输出 "20060102150405.999" 格式字符串
func (c Carbon) ToShortDateTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateTimeMilliLayout)
}

// ToShortDateTimeMicroString outputs a string in "20060102150405.999999" layout.
// 输出 "20060102150405.999999" 格式字符串
func (c Carbon) ToShortDateTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateTimeMicroLayout)
}

// ToShortDateTimeNanoString outputs a string in "20060102150405.999999999" layout.
// 输出 "20060102150405.999999999" 格式字符串
func (c Carbon) ToShortDateTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateTimeNanoLayout)
}

// ToDateString outputs a string in "2006-01-02" layout.
// 输出 "2006-01-02" 格式字符串
func (c Carbon) ToDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateLayout)
}

// ToShortDateString outputs a string in "20060102" layout.
// 输出 "20060102" 格式字符串
func (c Carbon) ToShortDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateLayout)
}

// ToTimeString outputs a string in "15:04:05" layout.
// 输出 "15:04:05" 格式字符串
func (c Carbon) ToTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(TimeLayout)
}

// ToShortTimeString outputs a string in "150405" layout.
// 输出 "150405" 格式字符串
func (c Carbon) ToShortTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortTimeLayout)
}

// ToAtomString outputs a string in "2006-01-02T15:04:05Z07:00" layout.
// 输出 "2006-01-02T15:04:05Z07:00" 格式字符串
func (c Carbon) ToAtomString(timezone ...string) string {
	return c.ToRfc3339String(timezone...)
}

// ToANSICString outputs a string in "Mon Jan _2 15:04:05 2006" layout.
// 输出 "Mon Jan _2 15:04:05 2006" 格式字符串
func (c Carbon) ToANSICString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ANSICLayout)
}

// ToCookieString outputs a string in "Monday, 02-Jan-2006 15:04:05 MST" layout.
// 输出 "Monday, 02-Jan-2006 15:04:05 MST" 格式字符串
func (c Carbon) ToCookieString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(CookieLayout)
}

// ToRssString outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" format.
// 输出 "Mon, 02 Jan 2006 15:04:05 -0700" 格式字符串
func (c Carbon) ToRssString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RssLayout)
}

// ToW3cString outputs a string in "2006-01-02T15:04:05Z07:00" layout.
// 输出 "2006-01-02T15:04:05Z07:00" 格式字符串
func (c Carbon) ToW3cString(timezone ...string) string {
	return c.ToRfc3339String(timezone...)
}

// ToUnixDateString outputs a string in "Mon Jan _2 15:04:05 MST 2006" layout.
// 输出 "Mon Jan _2 15:04:05 MST 2006" 格式字符串
func (c Carbon) ToUnixDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(UnixDateLayout)
}

// ToRubyDateString outputs a string in "Mon Jan 02 15:04:05 -0700 2006" layout.
// 输出 "Mon Jan 02 15:04:05 -0700 2006" 格式字符串
func (c Carbon) ToRubyDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RubyDateLayout)
}

// ToKitchenString outputs a string in "3:04PM" layout.
// 输出 "3:04PM" 格式字符串
func (c Carbon) ToKitchenString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(KitchenLayout)
}

// ToIso8601String outputs a string in "2006-01-02T15:04:05-07:00" layout.
// 输出 "2006-01-02T15:04:05-07:00" 格式字符串
func (c Carbon) ToIso8601String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ISO8601Layout)
}

// ToIso8601MilliString outputs a string in "2006-01-02T15:04:05.999-07:00" layout.
// 输出 "2006-01-02T15:04:05.999-07:00" 格式字符串
func (c Carbon) ToIso8601MilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ISO8601MilliLayout)
}

// ToIso8601MicroString outputs a string in "2006-01-02T15:04:05.999999-07:00" layout.
// 输出 "2006-01-02T15:04:05.999999-07:00" 格式字符串
func (c Carbon) ToIso8601MicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ISO8601MicroLayout)
}

// ToIso8601NanoString outputs a string in "2006-01-02T15:04:05.999999999-07:00" layout.
// 输出 "2006-01-02T15:04:05.999999999-07:00" 格式字符串
func (c Carbon) ToIso8601NanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ISO8601NanoLayout)
}

// ToRfc822String outputs a string in "02 Jan 06 15:04 MST" layout.
// 输出 "02 Jan 06 15:04 MST" 格式字符串
func (c Carbon) ToRfc822String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC822Layout)
}

// ToRfc822zString outputs a string in "02 Jan 06 15:04 -0700" layout.
// 输出 "02 Jan 06 15:04 -0700" 格式字符串
func (c Carbon) ToRfc822zString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC822ZLayout)
}

// ToRfc850String outputs a string in "Monday, 02-Jan-06 15:04:05 MST" layout.
// 输出 "Monday, 02-Jan-06 15:04:05 MST" 格式字符串
func (c Carbon) ToRfc850String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC850Layout)
}

// ToRfc1036String outputs a string in "Mon, 02 Jan 06 15:04:05 -0700" layout.
// 输出 "Mon, 02 Jan 06 15:04:05 -0700" 格式字符串
func (c Carbon) ToRfc1036String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC1036Layout)
}

// ToRfc1123String outputs a string in "Mon, 02 Jan 2006 15:04:05 MST" layout.
// 输出 "Mon, 02 Jan 2006 15:04:05 MST" 格式字符串
func (c Carbon) ToRfc1123String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC1123Layout)
}

// ToRfc1123zString outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" layout.
// 输出 "Mon, 02 Jan 2006 15:04:05 -0700" 格式字符串
func (c Carbon) ToRfc1123zString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC1123ZLayout)
}

// ToRfc2822String outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" layout.
// 输出 "Mon, 02 Jan 2006 15:04:05 -0700" 格式字符串
func (c Carbon) ToRfc2822String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC2822Layout)
}

// ToRfc3339String outputs a string in "2006-01-02T15:04:05Z07:00" layout.
// 输出 "2006-01-02T15:04:05Z07:00" 格式字符串
func (c Carbon) ToRfc3339String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC3339Layout)
}

// ToRfc3339MilliString outputs a string in "2006-01-02T15:04:05.999Z07:00" layout.
// 输出 "2006-01-02T15:04:05.999Z07:00" 格式字符串
func (c Carbon) ToRfc3339MilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC3339MilliLayout)
}

// ToRfc3339MicroString outputs a string in "2006-01-02T15:04:05.999999Z07:00" layout.
// 输出 "2006-01-02T15:04:05.999999Z07:00" 格式字符串
func (c Carbon) ToRfc3339MicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC3339MicroLayout)
}

// ToRfc3339NanoString outputs a string in "2006-01-02T15:04:05.999999999Z07:00" layout.
// 输出 "2006-01-02T15:04:05.999999999Z07:00" 格式字符串
func (c Carbon) ToRfc3339NanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC3339NanoLayout)
}

// ToRfc7231String outputs a string in "Mon, 02 Jan 2006 15:04:05 GMT" layout.
// 输出 "Mon, 02 Jan 2006 15:04:05 GMT" 格式字符串
func (c Carbon) ToRfc7231String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC7231Layout)
}

// ToLayoutString outputs a string by layout.
// 输出指定布局的时间字符串
func (c Carbon) ToLayoutString(layout string, timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(layout)
}

// Layout outputs a string by layout, it is shorthand for ToLayoutString.
// 输出指定布局的时间字符串, 是 ToLayoutString 的简写
func (c Carbon) Layout(layout string, timezone ...string) string {
	return c.ToLayoutString(layout, timezone...)
}

// ToFormatString outputs a string by format.
// 输出指定格式的时间字符串
func (c Carbon) ToFormatString(format string, timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(format); i++ {
		if layout, ok := formats[format[i]]; ok {
			buffer.WriteString(c.Carbon2Time().Format(layout))
		} else {
			switch format[i] {
			case '\\': // 原样输出，不解析
				buffer.WriteByte(format[i+1])
				i++
				continue
			case 'W': // ISO-8601 格式数字表示的年份中的第几周，取值范围 1-52
				buffer.WriteString(strconv.Itoa(c.WeekOfYear()))
			case 'N': // ISO-8601 格式数字表示的星期中的第几天，取值范围 1-7
				buffer.WriteString(strconv.Itoa(c.DayOfWeek()))
			case 'S': // 月份中第几天的英文缩写后缀，如st, nd, rd, th
				suffix := "th"
				switch c.Day() {
				case 1, 21, 31:
					suffix = "st"
				case 2, 22:
					suffix = "nd"
				case 3, 23:
					suffix = "rd"
				}
				buffer.WriteString(suffix)
			case 'L': // 是否为闰年，如果是闰年为 1，否则为 0
				if c.IsLeapYear() {
					buffer.WriteString("1")
				} else {
					buffer.WriteString("0")
				}
			case 'G': // 数字表示的小时，24 小时格式，没有前导零，取值范围 0-23
				buffer.WriteString(strconv.Itoa(c.Hour()))
			case 'U': // 秒级时间戳，如 1611818268
				buffer.WriteString(strconv.FormatInt(c.Timestamp(), 10))
			case 'u': // 数字表示的毫秒，如 999
				buffer.WriteString(strconv.Itoa(c.Millisecond()))
			case 'w': // 数字表示的星期中的第几天，取值范围 0-6
				buffer.WriteString(strconv.Itoa(c.DayOfWeek() - 1))
			case 't': // 指定的月份有几天，取值范围 28-31
				buffer.WriteString(strconv.Itoa(c.DaysInMonth()))
			case 'z': // 年份中的第几天，取值范围 0-365
				buffer.WriteString(strconv.Itoa(c.DayOfYear() - 1))
			case 'e': // 当前位置，如 UTC，GMT，Atlantic/Azores
				buffer.WriteString(c.Location())
			case 'Q': // 当前季度，取值范围 1-4
				buffer.WriteString(strconv.Itoa(c.Quarter()))
			case 'C': // 当前世纪数，取值范围 0-99
				buffer.WriteString(strconv.Itoa(c.Century()))
			default:
				buffer.WriteByte(format[i])
			}
		}
	}
	return buffer.String()
}

// Format outputs a string by format, it is shorthand for ToFormatString.
// 输出指定格式的时间字符串, 是 ToFormatString 的简写
func (c Carbon) Format(format string, timezone ...string) string {
	return c.ToFormatString(format, timezone...)
}
