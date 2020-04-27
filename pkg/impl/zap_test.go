package impl

import (
	"reflect"
	"testing"

	"github.com/mavolin/logstract/pkg/logstract"
)

func TestFieldsToSlice(t *testing.T) {
	testFields := logstract.Fields{
		"a": "b",
		"c": 3,
		"d": true,
	}

	expect := []interface{}{"a", "b", "c", 3, "d", true}

	actual := fieldsToSlice(testFields)

	if reflect.DeepEqual(actual, expect) {
		t.Errorf("expected: %+v, but got: %+v", expect, actual)
	}
}
