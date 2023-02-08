package castlearn

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	jww "github.com/spf13/jwalterweatherman"
)

func ToTimeE(i interface{}) (time.Time, bool) {
	switch s := i.(type) {
	case time.Time:
		return s, true
	case string:
		d, e := StringToDate(s)
		if e == nil {
			return d, true
		}

		jww.ERROR.Printf("Unable to Cast %#v to Date/Time format:", i)
		return time.Time{}, false
	default:
		jww.ERROR.Printf("Unable to Cast %#v to Date/Time format:", i)
		return time.Time{}, false
	}
}

func toBoolE(i interface{}) (bool, bool) {
	switch b := i.(type) {
	case bool:
		return b, true
	case nil:
		return false, true
	case int:
		if i.(int) > 0 {
			return true, true
		}
		return false, true
	default:
		jww.ERROR.Printf("Unable to Cast %#v to bool", i)
		return false, false
	}
}

func ToFloat64E(i interface{}) (float64, bool) {
	switch s := i.(type) {
	case float64:
		return s, true
	case float32:
		return float64(s), true

	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return float64(v), true
		} else {
			jww.ERROR.Printf("Unable to Cast %#v to float64", i)
			return 0, false
		}

	default:
		jww.ERROR.Printf("Unable to Cast %#v to float64", i)
	}

	return 0.0, false
}

func ToIntE(i interface{}) (int, bool) {
	switch s := i.(type) {
	case int:
		return s, true
	case int64:
		return int(s), true
	case int32:
		return int(s), true
	case int16:
		return int(s), true
	case int8:
		return int(s), true
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), true
		} else {
			jww.ERROR.Printf("Unable to Cast %#v to int", i)
			return 0, false
		}
	case float64:
		return int(s), true
	case bool:
		if bool(s) {
			return 1, true
		} else {
			return 0, true
		}
	case nil:
		return 0, true
	default:
		jww.ERROR.Printf("Unable to Cast %#v to int", i)
	}

	return 0, false
}

func ToStringE(i interface{}) (string, bool) {
	switch s := i.(type) {
	case string:
		return s, true
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64), true
	case int:
		return strconv.FormatInt(int64(i.(int)), 10), true
	case []byte:
		return string(s), true
	case nil:
		return "", true
	default:
		jww.ERROR.Printf("Unable to Cast %#v to string", i)
	}

	return "", false
}

func ToStringMapStringE(i interface{}) (map[string]string, bool) {
	var m = map[string]string{}

	switch v := i.(type) {
	case map[string]string:
		return v, true
	case map[interface{}]interface{}:
		for k, v := range v {
			kStr, _ := ToStringE(k)
			vStr, _ := ToStringE(v)
			m[kStr] = vStr
		}
	default:
		return m, false
	}

	return m, true
}

func ToStringMapE(i interface{}) (map[string]interface{}, bool) {
	var m = map[string]interface{}{}

	switch v := i.(type) {
	case map[string]interface{}:
		return v, true
	case map[interface{}]interface{}:
		for k, val := range v {
			kStr, _ := ToStringE(k)
			m[kStr] = val
		}
	default:
		return m, false
	}

	return m, true
}

func ToStringSliceE(i interface{}) ([]string, bool) {
	var a []string

	switch v := i.(type) {
	case []string:
		return v, true
	case []interface{}:
		for _, u := range v {
			s, _ := ToStringE(u)
			a = append(a, s)
		}
	default:
		return a, false
	}

	return a, true
}

func StringToDate(s string) (time.Time, error) {
	return parseDate(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05Z07:00",
		"02 Jan 06 15:04 MST",
		"2006-01-02",
		"02 Jan 2006",
	})
}

func parseDate(s string, dates []string) (d time.Time, err error) {
	for _, dateType := range dates {
		if d, err = time.Parse(dateType, s); err == nil {
			return
		}
	}
	return d, errors.New(fmt.Sprintf("Unable to parse date: %s", s))
}
