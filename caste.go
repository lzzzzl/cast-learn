package castlearn

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ToTimeE casts an empty interface to time.Time.
func ToTimeE(i interface{}) (time.Time, error) {
	i = indirect(i)

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

// ToDurationE casts an empty interface to time.Duration.
func ToDurationE(i interface{}) (time.Duration, error) {
	i = indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int64, int32, int16, int8, int:
		d := time.Duration(ToInt64(i))
		return d, nil
	case float32, float64:
		d := time.Duration(ToFloat64(i))
		return d, nil
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			return time.ParseDuration(s)
		} else {
			return time.ParseDuration(s + "ns")
		}
	default:
		err := fmt.Errorf("Unable to Cast %#v to Duration\n", i)
		return time.Duration(0), err
	}
}

// ToBoolE casts an empty interface to a bool.
func ToBoolE(i interface{}) (bool, error) {
	i = indirect(i)

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

// ToFloat64E casts an empty interface to a float64.
func ToFloat64E(i interface{}) (float64, error) {
	i = indirect(i)

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
		}
		return 0.0, fmt.Errorf("Unable to Cast %#v to float", i)
	default:
		return 0, fmt.Errorf("Unable to Cast %#v to float64", i)
	}
}

func ToInt64E(i interface{}) (int64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int64:
		return s, nil
	case int:
		return int64(s), nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("Unable to Cast %#v to int64", i)
	case float64:
		return int64(s), nil
	case bool:
		if bool(s) {
			return int64(1), nil
		}
		return int64(0), nil
	case nil:
		return int64(0), nil
	default:
		return int64(0), fmt.Errorf("Unable to Cast %#v to int64", i)
	}
}

// ToIntE casts an empty interface to an int.
func ToIntE(i interface{}) (int, error) {
	i = indirect(i)

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
		}
		return 0, fmt.Errorf("Unable to Cast %#v to int", i)
	case float64:
		return int(s), nil
	case bool:
		if bool(s) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Unable to Cast %#v to int", i)
	}
}

// ToStringE casts an empty interface to a string.
func ToStringE(i interface{}) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64), nil
	case int64:
		return strconv.FormatInt(i.(int64), 10), nil
	case int:
		return strconv.FormatInt(int64(i.(int)), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("Unable to Cast %#v to string", i)
	}
}

// ToStringMapStringE casts an empty interface to a map[string]string.
func ToStringMapStringE(i interface{}) (map[string]string, error) {
	var m = map[string]string{}

	switch v := i.(type) {
	case map[string]string:
		return v, nil
	case map[string]interface{}:
		for k, val := range v {
			kStr, _ := ToStringE(k)
			vStr, _ := ToStringE(val)
			m[kStr] = vStr
		}
		return m, nil
	case map[interface{}]string:
		for k, val := range v {
			kStr, _ := ToStringE(k)
			vStr, _ := ToStringE(val)
			m[kStr] = vStr
		}
	case map[interface{}]interface{}:
		for k, val := range v {
			kStr, _ := ToStringE(k)
			vStr, _ := ToStringE(val)
			m[kStr] = vStr
		}
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]string", i)
	}

	return m, fmt.Errorf("Unable to Cast %#v to map[string]string", i)
}

// ToStringMapStringSliceE casts an empty interface to a map[string][]string.
func ToStringMapStringSliceE(i interface{}) (map[string][]string, error) {
	var m = map[string][]string{}

	switch v := i.(type) {
	case map[string][]string:
		return v, nil
	case map[string][]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[string]string:
		for k, val := range v {
			m[ToString(k)] = []string{val}
		}
	case map[string]interface{}:
		for k, val := range v {
			switch vt := val.(type) {
			case []interface{}:
				m[ToString(k)] = ToStringSlice(vt)
			case []string:
				m[ToString(k)] = vt
			default:
				m[ToString(k)] = []string{ToString(val)}
			}
		}
		return m, nil
	case map[interface{}][]string:
		for k, val := range v {
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[interface{}]string:
		for k, val := range v {
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[interface{}][]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			key, err := ToStringE(k)
			if err != nil {
				return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
			}
			value, err := ToStringSliceE(val)
			if err != nil {
				return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
			}
			m[key] = value
		}
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
	}
	return m, nil
}

// ToStringMapBoolE casts an empty interface to a map[string]bool.
func ToStringMapBoolE(i interface{}) (map[string]bool, error) {
	var m = map[string]bool{}

	switch v := i.(type) {
	case map[string]bool:
		return v, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			kStr, _ := ToStringE(k)
			vBool, _ := ToBoolE(val)
			m[kStr] = vBool
		}
		return m, nil
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]bool", i)
	}
}

// ToStringMapE casts an empty interface to a map[string]interface{}.
func ToStringMapE(i interface{}) (map[string]interface{}, error) {
	var m = map[string]interface{}{}

	switch v := i.(type) {
	case map[string]interface{}:
		return v, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			kStr, _ := ToStringE(k)
			m[kStr] = val
		}
		return m, nil
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]interface{}", i)
	}
}

// ToSliceE casts an empty interface to a []interface{}.
func ToSliceE(i interface{}) ([]interface{}, error) {
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

// ToBoolSliceE casts an empty interface to a []bool.
func ToBoolSliceE(i interface{}) ([]bool, error) {
	if i == nil {
		return []bool{}, fmt.Errorf("Unable to Cast %#v to []bool", i)
	}

	switch v := i.(type) {
	case []bool:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToBoolE(s.Index(j).Interface())
			if err != nil {
				return []bool{}, fmt.Errorf("Unable to Cast %#v to []bool", i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, fmt.Errorf("Unable to Cast %#v to []bool", i)
	}
}

// ToStringSliceE casts an empty interface to a []string.
func ToStringSliceE(i interface{}) ([]string, error) {
	var a []string

	switch v := i.(type) {
	case []string:
		return v, nil
	case string:
		return strings.Fields(v), nil
	case []interface{}:
		for _, u := range v {
			s, _ := ToStringE(u)
			a = append(a, s)
		}
		return a, nil
	case interface{}:
		str, err := ToStringE(v)
		if err != nil {
			return a, fmt.Errorf("Unable to Cast %#v to []string", i)
		}
		return []string{str}, nil
	default:
		return a, fmt.Errorf("Unable to Cast %#v to []string", i)
	}
}

// ToIntSliceE casts an empty interface to a []int.
func ToIntSliceE(i interface{}) ([]int, error) {
	if i == nil {
		return []int{}, fmt.Errorf("Unable to Cast %v to []int", i)
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToIntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("Unable to Cast %v to []int", i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("Unable to Cast %v to []int", i)
	}
}

// StringToDate casts an empty interface to a time.Time.
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
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
	})
}

func parseDate(s string, dates []string) (d time.Time, err error) {
	for _, dateType := range dates {
		if d, err = time.Parse(dateType, s); err == nil {
			return
		}
	}
	return d, fmt.Errorf("Unable to parse date: %s", s)
}

func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	v := reflect.ValueOf(a)
	for !v.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) &&
		!v.Type().Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()) &&
		v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
