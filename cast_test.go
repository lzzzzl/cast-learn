package castlearn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	var eight interface{} = 8
	assert.Equal(t, 8, ToInt(eight))
	assert.Equal(t, 8, ToInt("8"))
	assert.Equal(t, 1, ToInt(true))
	assert.Equal(t, 0, ToInt(false))
	assert.Equal(t, 8, ToInt(eight))
}

func TestMaps(t *testing.T) {
	var m = map[interface{}]interface{}{"tag": "tags", "group": "groups"}
	assert.Equal(t, map[string]interface{}{"tag": "tags", "group": "groups"}, ToStringMap(m))
}
