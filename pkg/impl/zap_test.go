package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mavolin/logstract/pkg/logstract"
)

func TestFieldsToSlice(t *testing.T) {
	testFields := logstract.Fields{
		"a": "b",
	}

	expect := []interface{}{"a", "b"}
	actual := fieldsToSlice(testFields)

	assert.Equal(t, expect, actual)
}
