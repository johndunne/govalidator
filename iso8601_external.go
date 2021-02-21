package govalidator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	BadFormatErr         = errors.New("not iso8601 format")
	utc                  = regexp.MustCompile(`^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(\\.[0-9]+)?(Z)?$`)
	iso8601_date_or_time = regexp.MustCompile(`^(P(?P<date>[\dYM]*\d+[YMD]))?(T+(?P<time>[\dHM]*\d[HMS])+)?$`)
	iso8601_duration     = regexp.MustCompile(`^(P((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<day>\d+)D)?)?(T+((?P<hour>[01]?\d|2[0-3])H)?((?P<minute>[0-5]?\d)M)?((?P<second>[0-5]?\d)S)?)?$`)
)

type Duration struct {
	Years   uint
	Months  uint
	Weeks   uint
	Days    uint
	Hours   uint
	Minutes uint
	Seconds uint
}

func ParseDuration(dur string) (*Duration, error) {
	var (
		match []string
	)

	if !iso8601_date_or_time.MatchString(dur) {
		return nil, BadFormatErr
	}
	if iso8601_duration.MatchString(dur) {
		match = iso8601_duration.FindStringSubmatch(dur)
	} else {
		return nil, BadFormatErr
	}

	d := &Duration{}

	for i, name := range iso8601_duration.SubexpNames() {
		part := match[i]
		if i == 0 || name == "" || part == "" {
			continue
		}

		parse_val, err := strconv.ParseUint(part, 10, 16)
		if err != nil {
			return nil, err
		}

		val := uint(parse_val)
		switch name {
		case "year":
			d.Years = val
		case "month":
			d.Months = val
		case "week":
			d.Weeks = val
		case "day":
			d.Days = val
		case "hour":
			d.Hours = val
		case "minute":
			d.Minutes = val
		case "second":
			d.Seconds = val
		default:
			return nil, errors.New(fmt.Sprintf("unknown field %s", name))
		}
	}

	return d, nil
}

//func (d *Duration) HasTimePart() bool {
//	return d.Hours != 0 || d.Minutes != 0 || d.Seconds != 0
//}

//func (d *Duration) ToDuration() time.Duration {
//	day := time.Hour * 24
//	year := day * 365
//
//	tot := time.Duration(0)
//
//	tot += year * time.Duration(d.Years)
//	tot += day * 7 * time.Duration(d.Weeks)
//	tot += day * time.Duration(d.Days)
//	tot += time.Hour * time.Duration(d.Hours)
//	tot += time.Minute * time.Duration(d.Minutes)
//	tot += time.Second * time.Duration(d.Seconds)
//
//	return tot
//}
