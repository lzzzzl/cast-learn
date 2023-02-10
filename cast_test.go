package castlearn

import (
	"html/template"
	"testing"

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
	assert.Equal(t, "8.12", ToString(8.12))
	assert.Equal(t, "one time", ToString([]byte("one time")))
	assert.Equal(t, "one time", ToString(template.HTML("one time")))
	assert.Equal(t, "one more time", ToString(foo))
	assert.Equal(t, "", ToString(nil))
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

	var stringMapBool = map[interface{}]interface{}{"v1": true, "v2": false}
	assert.Equal(t, ToStringMapBool(stringMapBool), map[string]bool{"v1": true, "v2": false})
}

func TestSlice(t *testing.T) {
	assert.Equal(t, []string{"a", "b"}, ToStringSlice([]string{"a", "b"}))
	assert.Equal(t, []string{"1", "3"}, ToStringSlice([]interface{}{1, 3}))

	assert.Equal(t, []int{1, 3}, ToIntSlice([]int{1, 3}))
	assert.Equal(t, []int{1, 3}, ToIntSlice([]interface{}{1.2, 3.2}))
	assert.Equal(t, []int{2, 3}, ToIntSlice([]string{"2", "3"}))
	assert.Equal(t, []int{2, 3}, ToIntSlice([2]string{"2", "3"}))
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
