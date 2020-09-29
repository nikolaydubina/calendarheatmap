package charts

import (
	"time"
)

var localeConfig = map[string]LabelsProvider{
	"en_US": LabelsProvider{
		months: map[time.Month]string{
			time.January:   "Jan",
			time.February:  "Feb",
			time.March:     "Mar",
			time.April:     "Apr",
			time.May:       "May",
			time.June:      "Jun",
			time.July:      "Jul",
			time.August:    "Aug",
			time.September: "Sep",
			time.October:   "Oct",
			time.November:  "Nov",
			time.December:  "Dec",
		},
		weekdays: map[time.Weekday]string{
			time.Monday:    "Mon",
			time.Tuesday:   "Tue",
			time.Wednesday: "Wed",
			time.Thursday:  "Thu",
			time.Friday:    "Fri",
			time.Saturday:  "Sat",
			time.Sunday:    "Sun",
		},
	},
	"ko_KR": LabelsProvider{
		months: map[time.Month]string{
			time.January:   "1월",
			time.February:  "2월",
			time.March:     "3월",
			time.April:     "4월",
			time.May:       "5월",
			time.June:      "6월",
			time.July:      "7월",
			time.August:    "8월",
			time.September: "9월",
			time.October:   "10월",
			time.November:  "11월",
			time.December:  "12월",
		},
		weekdays: map[time.Weekday]string{
			time.Monday:    "월",
			time.Tuesday:   "화",
			time.Wednesday: "수",
			time.Thursday:  "목",
			time.Friday:    "금",
			time.Saturday:  "토",
			time.Sunday:    "일",
		},
	},
}

// LabelsProvider provides labels for locale
type LabelsProvider struct {
	months   map[time.Month]string
	weekdays map[time.Weekday]string
}

// NewLabelsProvider initializes labels provider for locale
func NewLabelsProvider(locale string) LabelsProvider {
	return localeConfig[locale]
}

// GetMonth returns month label
func (p LabelsProvider) GetMonth(month time.Month) string {
	return p.months[month]
}

// GetWeekday returns weekday label
func (p LabelsProvider) GetWeekday(weekday time.Weekday) string {
	return p.weekdays[weekday]
}
