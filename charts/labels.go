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
			time.January:   "일월",
			time.February:  "이월",
			time.March:     "삼월",
			time.April:     "사월",
			time.May:       "오월",
			time.June:      "유월",
			time.July:      "칠월",
			time.August:    "팔월",
			time.September: "구월",
			time.October:   "시월",
			time.November:  "십일월",
			time.December:  "십이월",
		},
		weekdays: map[time.Weekday]string{
			time.Monday:    "월요일",
			time.Tuesday:   "화요일",
			time.Wednesday: "수요일",
			time.Thursday:  "목요일",
			time.Friday:    "금요일",
			time.Saturday:  "토요일",
			time.Sunday:    "일요일",
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
