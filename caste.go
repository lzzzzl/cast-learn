package castlearn

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ToTimeE casts an interface to a time.Time type.
func ToTimeE(i interface{}) (time.Time, error) {
	i = indirect(i)

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		return StringToDate(v)
	case int:
		return time.Unix(int64(v), 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	default:
		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
	}
}

// ToDurationE casts an interface to a time.Duration type.
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
		return time.Duration(0), fmt.Errorf("unable to cast %#v of type %T to Duration", i, i)
	}
}

// ToBoolE casts an interface to a bool type.
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
		return false, fmt.Errorf("unable to Cast %#v of type %T to bool\n", i, i)
	}
}

// ToFloat64E casts an interface to a float64 type.
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
		return 0.0, fmt.Errorf("unable to Cast %#v of type %T to float\n", i, i)
	default:
		return 0, fmt.Errorf("unable to Cast %#v of type %T to float64\n", i, i)
	}
}

// ToInt64E casts an interface to an int64 type.
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
		return 0, fmt.Errorf("unable to Cast %#v of type %T to int64\n", i, i)
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
		return int64(0), fmt.Errorf("unable to Cast %#v of type %T to int64\n", i, i)
	}
}

// ToInt32E casts an interface to an int32 type.
func ToInt32E(i interface{}) (int32, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int64:
		return int32(s), nil
	case int:
		return int32(s), nil
	case int32:
		return s, nil
	case int16:
		return int32(s), nil
	case int8:
		return int32(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int32(v), nil
		}
		return 0, fmt.Errorf("unable to Cast %#v of type %T to int32\n", i, i)
	case float64:
		return int32(s), nil
	case bool:
		if bool(s) {
			return int32(1), nil
		}
		return int32(0), nil
	case nil:
		return int32(0), nil
	default:
		return int32(0), fmt.Errorf("unable to Cast %#v of type %T to int32\n", i, i)
	}
}

// ToInt16E casts an interface to an int16 type.
func ToInt16E(i interface{}) (int16, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int64:
		return int16(s), nil
	case int:
		return int16(s), nil
	case int32:
		return int16(s), nil
	case int16:
		return s, nil
	case int8:
		return int16(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int16(v), nil
		}
		return 0, fmt.Errorf("unable to Cast %#v of type %T to int16\n", i, i)
	case float64:
		return int16(s), nil
	case bool:
		if bool(s) {
			return int16(1), nil
		}
		return int16(0), nil
	case nil:
		return int16(0), nil
	default:
		return int16(0), fmt.Errorf("unable to Cast %#v of type %T to int16\n", i, i)
	}
}

// ToInt8E casts an interface to an int8 type.
func ToInt8E(i interface{}) (int8, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int64:
		return int8(s), nil
	case int:
		return int8(s), nil
	case int32:
		return int8(s), nil
	case int16:
		return int8(s), nil
	case int8:
		return s, nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int8(v), nil
		}
		return 0, fmt.Errorf("unable to Cast %#v of type %T to int8\n", i, i)
	case float64:
		return int8(s), nil
	case bool:
		if bool(s) {
			return int8(1), nil
		}
		return int8(0), nil
	case nil:
		return int8(0), nil
	default:
		return int8(0), fmt.Errorf("unable to Cast %#v of type %T to int8\n", i, i)
	}
}

// ToIntE casts an interface to an int type.
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
		return 0, fmt.Errorf("unable to Cast %#v of type %T to int\n", i, i)
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
		return 0, fmt.Errorf("unable to Cast %#v of type %T to int\n", i, i)
	}
}

// ToStringE casts an interface to a string type.
func ToStringE(i interface{}) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int:
		return strconv.Itoa(s), nil
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
		return "", fmt.Errorf("unable to Cast %#v of type %T to string\n", i, i)
	}
}

// ToStringMapStringE casts an interface to a map[string]string type.
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
		return m, fmt.Errorf("unable to Cast %#v of type %T to map[string]string\n", i, i)
	}

	return m, fmt.Errorf("unable to Cast %#v of type %T to map[string]string\n", i, i)
}

// ToStringMapStringSliceE casts an interface to a map[string][]string type.
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
				return m, fmt.Errorf("unable to Cast %#v of type %T to map[string]string\n", i, i)
			}
			value, err := ToStringSliceE(val)
			if err != nil {
				return m, fmt.Errorf("unable to Cast %#v of type %T to map[string]string\n", i, i)
			}
			m[key] = value
		}
	default:
		return m, fmt.Errorf("unable to Cast %#v of type %T to map[string]string\n", i, i)
	}
	return m, nil
}

// ToStringMapBoolE casts an interface to a map[string]bool type.
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
		return m, fmt.Errorf("unable to Cast %#v of type %T to map[string]bool\n", i, i)
	}
}

// ToStringMapE casts an interface to a map[string]interface{} type.
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
		return m, fmt.Errorf("unable to Cast %#v of type %T to map[string]interface{}\n", i, i)
	}
}

// ToSliceE casts an interface to a []interface{} type.
func ToSliceE(i interface{}) ([]interface{}, error) {
	var s []interface{}

	switch v := i.(type) {
	case []interface{}:
		return append(s, v...), nil
	case []map[string]interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	default:
		return s, fmt.Errorf("unable to Cast %#v of type %T to []interface{}\n", i, i)
	}
}

// ToBoolSliceE casts an interface to a []bool type.
func ToBoolSliceE(i interface{}) ([]bool, error) {
	if i == nil {
		return []bool{}, fmt.Errorf("unable to Cast %#v of type %T to []bool\n", i, i)
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
				return []bool{}, fmt.Errorf("unable to Cast %#v of type %T to []bool\n", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, fmt.Errorf("unable to Cast %#v of type %T to []bool\n", i, i)
	}
}

// ToStringSliceE casts an interface to a []string type.
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
			return a, fmt.Errorf("unable to Cast %#v of type %T to []string\n", i, i)
		}
		return []string{str}, nil
	default:
		return a, fmt.Errorf("unable to Cast %#v of type %T to []string\n", i, i)
	}
}

// ToIntSliceE casts an interface to a []int type.
func ToIntSliceE(i interface{}) ([]int, error) {
	if i == nil {
		return []int{}, fmt.Errorf("unable to Cast %#v of type %T to []int\n", i, i)
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
				return []int{}, fmt.Errorf("unable to Cast %#v of type %T to []int\n", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("unable to Cast %#v of type %T to []int\n", i, i)
	}
}

// StringToDate attempts to parse a string into a time.Time type using a
// predefined list of formats.  If no suitable format is found, an error is
// returned.
func StringToDate(s string) (time.Time, error) {
	return parseDate(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	})
}

func parseDate(s string, dates []string) (d time.Time, err error) {
	for _, dateType := range dates {
		if d, err = time.Parse(dateType, s); err == nil {
			return
		}
	}
	return d, fmt.Errorf("unable to parse date: %s", s)
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
