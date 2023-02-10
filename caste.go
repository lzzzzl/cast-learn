package castlearn

import (
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"time"

	jww "github.com/spf13/jwalterweatherman"
)

func ToTimeE(i interface{}) (time.Time, error) {
	jww.DEBUG.Println("ToTimeE called on type:", reflect.TypeOf(i))

	switch s := i.(type) {
	case time.Time:
		return s, nil
	case string:
		d, e := StringToDate(s)
		if e == nil {
			return d, nil
		}
		return time.Time{}, fmt.Errorf("Unable to Cast %#v to Date/Time format: %s\n", i, e)
	default:
		return time.Time{}, fmt.Errorf("Unable to Cast %#v to Date/Time\n", i)
	}
}

func toBoolE(i interface{}) (bool, error) {
	jww.DEBUG.Println("ToBoolE called on type:", reflect.TypeOf(i))

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	default:
		return false, fmt.Errorf("Unable to Cast %#v to bool\n", i)
	}
}

func ToFloat64E(i interface{}) (float64, error) {
	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case int:
		return float64(s), nil

	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return float64(v), nil
		} else {
			return 0, fmt.Errorf("Unable to Cast %#v to float64", i)
		}

	default:
		return 0, fmt.Errorf("Unable to Cast %#v to float64", i)
	}
}

func ToIntE(i interface{}) (int, error) {
	jww.DEBUG.Println("ToIntE called on type:", reflect.TypeOf(i))

	switch s := i.(type) {
	case int:
		return s, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		} else {
			return 0, fmt.Errorf("Unable to Cast %#v to int", i)
		}
	case float64:
		return int(s), nil
	case bool:
		if bool(s) {
			return 1, nil
		} else {
			return 0, nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Unable to Cast %#v to int", i)
	}
}

func ToStringE(i interface{}) (string, error) {
	jww.DEBUG.Println("ToStringE called on type:", reflect.TypeOf(i))

	switch s := i.(type) {
	case string:
		return s, nil
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64), nil
	case int:
		return strconv.FormatInt(int64(i.(int)), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("Unable to Cast %#v to string", i)
	}
}

func ToStringMapStringE(i interface{}) (map[string]string, error) {
	jww.DEBUG.Println("ToStringMapStringE called on type:", reflect.TypeOf(i))

	var m = map[string]string{}

	switch v := i.(type) {
	case map[string]string:
		return v, nil
	case map[interface{}]interface{}:
		for k, v := range v {
			kStr, _ := ToStringE(k)
			vStr, _ := ToStringE(v)
			m[kStr] = vStr
		}
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]string", i)
	}

	return m, fmt.Errorf("Unable to Cast %#v to map[string]string", i)
}

func ToStringMapBoolE(i interface{}) (map[string]bool, error) {
	jww.DEBUG.Println("ToStringMapBoolE called on type:", reflect.TypeOf(i))

	var m = map[string]bool{}

	switch v := i.(type) {
	case map[string]bool:
		return v, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			kStr, _ := ToStringE(k)
			vBool, _ := toBoolE(val)
			m[kStr] = vBool
		}
		return m, nil
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]bool", i)
	}
}

func ToStringMapE(i interface{}) (map[string]interface{}, error) {
	jww.DEBUG.Println("ToStringMapE called on type:", reflect.TypeOf(i))

	var m = map[string]interface{}{}

	switch v := i.(type) {
	case map[string]interface{}:
		return v, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			kStr, _ := ToStringE(k)
			m[kStr] = val
		}
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]interface{}", i)
	}

	return m, fmt.Errorf("Unable to Cast %#v to map[string]interface{}", i)
}

func ToSliceE(i interface{}) ([]interface{}, error) {
	jww.DEBUG.Println("ToSliceE called on type:", reflect.TypeOf(i))

	var s []interface{}

	switch v := i.(type) {
	case []interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	case []map[string]interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	default:
		return s, fmt.Errorf("Unable to Cast %#v of type %v to []interface{}", i, reflect.TypeOf(i))
	}
}

func ToStringSliceE(i interface{}) ([]string, error) {
	jww.DEBUG.Println("ToStringSliceE called on type:", reflect.TypeOf(i))

	var a []string

	switch v := i.(type) {
	case []string:
		return v, nil
	case []interface{}:
		for _, u := range v {
			s, _ := ToStringE(u)
			a = append(a, s)
		}
	default:
		return a, fmt.Errorf("Unable to Cast %#v to []string", i)
	}

	return a, fmt.Errorf("Unable to Cast %#v to []string", i)
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
