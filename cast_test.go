package castlearn

import (
	"html/template"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	var eight interface{} = 8
	assert.Equal(t, 8, ToInt(8))
	assert.Equal(t, 8, ToInt(8.31))
	assert.Equal(t, 8, ToInt(eight))
	assert.Equal(t, 8, ToInt("8"))
	assert.Equal(t, 1, ToInt(true))
	assert.Equal(t, 0, ToInt(false))
	assert.Equal(t, 8, ToInt(eight))
}

func TestToInt64(t *testing.T) {
	var eight interface{} = 8
	assert.Equal(t, int64(8), ToInt64(int64(8)))
	assert.Equal(t, int64(8), ToInt64(8))
	assert.Equal(t, int64(8), ToInt64(8.31))
	assert.Equal(t, int64(8), ToInt64("8"))
	assert.Equal(t, int64(1), ToInt64(true))
	assert.Equal(t, int64(0), ToInt64(false))
	assert.Equal(t, int64(8), ToInt64(eight))
}

func TestToFloat64(t *testing.T) {
	var eight interface{} = 8
	assert.Equal(t, 8.00, ToFloat64(8))
	assert.Equal(t, 8.31, ToFloat64(8.31))
	assert.Equal(t, 8.31, ToFloat64("8.31"))
	assert.Equal(t, 8.0, ToFloat64(eight))
}

func TestToString(t *testing.T) {
	var foo interface{} = "one more time"
	assert.Equal(t, "8", ToString(8))
	assert.Equal(t, ToString(int64(16)), "16")
	assert.Equal(t, "8.12", ToString(8.12))
	assert.Equal(t, "one time", ToString([]byte("one time")))
	assert.Equal(t, "one time", ToString(template.HTML("one time")))
	assert.Equal(t, "http://somehost.foo", ToString(template.URL("http://somehost.foo")))
	assert.Equal(t, "(1+2)", ToString(template.JS("(1+2)")))
	assert.Equal(t, "a", ToString(template.CSS("a")))
	assert.Equal(t, "a", ToString(template.HTMLAttr("a")))
	assert.Equal(t, "one more time", ToString(foo))
	assert.Equal(t, "", ToString(nil))
	assert.Equal(t, ToString(true), "true")
	assert.Equal(t, ToString(false), "false")
}

type foo struct {
	val string
}

func (x foo) String() string {
	return x.val
}

func TestStringerToString(t *testing.T) {
	var x foo
	x.val = "bar"

	assert.Equal(t, "bar", ToString(x))
}

type fu struct {
	val string
}

func (x fu) Error() string {
	return x.val
}

func TestErrorToString(t *testing.T) {
	var x fu
	x.val = "bar"
	assert.Equal(t, "bar", ToString(x))
}

func TestMaps(t *testing.T) {
	var m = map[interface{}]interface{}{"tag": "tags", "group": "groups"}
	assert.Equal(t, map[string]interface{}{"tag": "tags", "group": "groups"}, ToStringMap(m))

	// ToStringMapString inputs/outputs
	var stringMapString = map[string]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var stringMapInterface = map[string]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapString = map[interface{}]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapInterface = map[interface{}]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}

	// ToStringMapStringSlice inputs/outputs
	var stringMapStringSlice = map[string][]string{"key 1": []string{"value 1", "value 2", "value 3"}, "key 2": []string{"value 1", "value 2", "value 3"}, "key 3": []string{"value 1", "value 2", "value 3"}}
	var stringMapInterfaceSlice = map[string][]interface{}{"key 1": []interface{}{"value 1", "value 2", "value 3"}, "key 2": []interface{}{"value 1", "value 2", "value 3"}, "key 3": []interface{}{"value 1", "value 2", "value 3"}}
	var stringMapStringSingleSliceFieldsResult = map[string][]string{"key 1": []string{"value", "1"}, "key 2": []string{"value", "2"}, "key 3": []string{"value", "3"}}
	var interfaceMapStringSlice = map[interface{}][]string{"key 1": []string{"value 1", "value 2", "value 3"}, "key 2": []string{"value 1", "value 2", "value 3"}, "key 3": []string{"value 1", "value 2", "value 3"}}
	var interfaceMapInterfaceSlice = map[interface{}][]interface{}{"key 1": []interface{}{"value 1", "value 2", "value 3"}, "key 2": []interface{}{"value 1", "value 2", "value 3"}, "key 3": []interface{}{"value 1", "value 2", "value 3"}}

	var stringMapStringSliceMultiple = map[string][]string{"key 1": []string{"value 1", "value 2", "value 3"}, "key 2": []string{"value 1", "value 2", "value 3"}, "key 3": []string{"value 1", "value 2", "value 3"}}
	var stringMapStringSliceSingle = map[string][]string{"key 1": []string{"value 1"}, "key 2": []string{"value 2"}, "key 3": []string{"value 3"}}

	var stringMapBool = map[interface{}]interface{}{"v1": true, "v2": false}
	assert.Equal(t, ToStringMapBool(stringMapBool), map[string]bool{"v1": true, "v2": false})

	var stringMapInterface1 = map[string]interface{}{"key 1": []string{"value 1"}, "key 2": []string{"value 2"}}
	var stringMapInterfaceResult1 = map[string][]string{"key 1": []string{"value 1"}, "key 2": []string{"value 2"}}

	// ToStringMapString tests
	assert.Equal(t, ToStringMapString(stringMapString), stringMapString)
	assert.Equal(t, ToStringMapString(stringMapInterface), stringMapString)
	assert.Equal(t, ToStringMapString(interfaceMapString), stringMapString)
	assert.Equal(t, ToStringMapString(interfaceMapInterface), stringMapString)

	// ToStringMapStringSlice tests
	assert.Equal(t, ToStringMapStringSlice(stringMapStringSlice), stringMapStringSlice)
	assert.Equal(t, ToStringMapStringSlice(stringMapInterfaceSlice), stringMapStringSlice)
	assert.Equal(t, ToStringMapStringSlice(stringMapStringSliceMultiple), stringMapStringSlice)
	assert.Equal(t, ToStringMapStringSlice(stringMapStringSliceMultiple), stringMapStringSlice)
	assert.Equal(t, ToStringMapStringSlice(stringMapString), stringMapStringSliceSingle)
	assert.Equal(t, ToStringMapStringSlice(stringMapInterface), stringMapStringSliceSingle)
	assert.Equal(t, ToStringMapStringSlice(interfaceMapStringSlice), stringMapStringSlice)
	assert.Equal(t, ToStringMapStringSlice(interfaceMapInterfaceSlice), stringMapStringSlice)
	assert.Equal(t, ToStringMapStringSlice(interfaceMapString), stringMapStringSingleSliceFieldsResult)
	assert.Equal(t, ToStringMapStringSlice(interfaceMapInterface), stringMapStringSingleSliceFieldsResult)
	assert.Equal(t, ToStringMapStringSlice(stringMapInterface1), stringMapInterfaceResult1)
}

func TestSlice(t *testing.T) {
	assert.Equal(t, []string{"a", "b"}, ToStringSlice([]string{"a", "b"}))
	assert.Equal(t, []string{"1", "3"}, ToStringSlice([]interface{}{1, 3}))

	assert.Equal(t, []int{1, 3}, ToIntSlice([]int{1, 3}))
	assert.Equal(t, []int{1, 3}, ToIntSlice([]interface{}{1.2, 3.2}))
	assert.Equal(t, []int{2, 3}, ToIntSlice([]string{"2", "3"}))
	assert.Equal(t, []int{2, 3}, ToIntSlice([2]string{"2", "3"}))

	assert.Equal(t, []bool{true, false, true}, ToBoolSlice([]bool{true, false, true}))
	assert.Equal(t, []bool{true, false, true}, ToBoolSlice([]interface{}{true, false, true}))
	assert.Equal(t, []bool{true, false, true}, ToBoolSlice([]int{1, 0, 1}))
	assert.Equal(t, []bool{true, false, true}, ToBoolSlice([]string{"true", "false", "true"}))
}

func TestToBool(t *testing.T) {
	assert.Equal(t, ToBool(0), false)
	assert.Equal(t, ToBool(nil), false)
	assert.Equal(t, ToBool("false"), false)
	assert.Equal(t, ToBool("FALSE"), false)
	assert.Equal(t, ToBool("False"), false)
	assert.Equal(t, ToBool("f"), false)
	assert.Equal(t, ToBool("F"), false)
	assert.Equal(t, ToBool(false), false)
	assert.Equal(t, ToBool("foo"), false)

	assert.Equal(t, ToBool("true"), true)
	assert.Equal(t, ToBool("TRUE"), true)
	assert.Equal(t, ToBool("True"), true)
	assert.Equal(t, ToBool("t"), true)
	assert.Equal(t, ToBool("T"), true)
	assert.Equal(t, ToBool(1), true)
	assert.Equal(t, ToBool(true), true)
	assert.Equal(t, ToBool(-1), true)
}

func TestIndirectPointers(t *testing.T) {
	x := 13
	y := &x
	z := &y

	assert.Equal(t, 13, ToInt(y))
	assert.Equal(t, 13, ToInt(z))
}

func TestToDuration(t *testing.T) {
	var td time.Duration = 5
	tests := []struct {
		input    interface{}
		expected time.Duration
	}{
		{time.Duration(5), td},
		{int64(5), td},
		{int32(5), td},
		{int16(5), td},
		{int8(5), td},
		{int(5), td},
		{float64(5), td},
		{float32(5), td},
		{string("5"), td},
		{string("5ns"), td},
		{string("5us"), time.Microsecond * td},
		{string("5µs"), time.Microsecond * td},
		{string("5ms"), time.Millisecond * td},
		{string("5s"), time.Second * td},
		{string("5m"), time.Minute * td},
		{string("5h"), time.Hour * td},
	}
	for _, v := range tests {
		assert.Equal(t, v.expected, ToDuration(v.input))
	}
}
