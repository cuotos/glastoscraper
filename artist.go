package glastoscraper

type Day int

const (
	UNDEFINED Day = iota
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

func (d Day) String() string {
	switch d {
	case UNDEFINED:
		return "Undefined"
	case THURSDAY:
		return "Thursday"
	case FRIDAY:
		return "Friday"
	case SATURDAY:
		return "Saturday"
	case SUNDAY:
		return "Sunday"
	}

	return "failure"
}

func ParseDay(day string) Day {
	switch day {
	case "WEDNESDAY":
		return WEDNESDAY
	case "THURSDAY":
		return THURSDAY
	case "FRIDAY":
		return FRIDAY
	case "SATURDAY":
		return SATURDAY
	case "SUNDAY":
		return SUNDAY
	}

	return UNDEFINED
}

type Artist struct {
	Title  string
	Stage  string
	Day    Day
	DayRaw string
	Time   string
}
