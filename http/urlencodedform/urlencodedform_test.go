package urlencodedform

import (
	"net/url"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	StringField    string
	IntField       int
	BoolField      bool
	StringSlice    []string
	IntSlice       []int
	BoolSlice      []bool
	TimeField      time.Time
	NestedStruct   NestedStruct
	IgnoredMap     map[string]string
	IgnoredFunc    func()
	IgnoredChan    chan int
	IgnoredPointer unsafe.Pointer
}

type NestedStruct struct {
	Value string
}

func TestURLValuesToObject_StringField(t *testing.T) {
	values := url.Values{
		"string_field": []string{"test value"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, "test value", target.StringField)
}

func TestURLValuesToObject_IntField(t *testing.T) {
	values := url.Values{
		"int_field": []string{"42"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, 42, target.IntField)
}

func TestURLValuesToObject_IntField_InvalidValue(t *testing.T) {
	values := url.Values{
		"int_field": []string{"not-a-number"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, 0, target.IntField)
}

func TestURLValuesToObject_BoolField_True(t *testing.T) {
	values := url.Values{
		"bool_field": []string{"true"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.True(t, target.BoolField)
}

func TestURLValuesToObject_BoolField_False(t *testing.T) {
	values := url.Values{
		"bool_field": []string{"false"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.False(t, target.BoolField)
}

func TestURLValuesToObject_StringSlice(t *testing.T) {
	values := url.Values{
		"string_slice": []string{"value1", "value2", "value3"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, []string{"value1", "value2", "value3"}, target.StringSlice)
}

func TestURLValuesToObject_StringSlice_WithBrackets(t *testing.T) {
	values := url.Values{
		"string_slice[]": []string{"value1", "value2"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, []string{"value1", "value2"}, target.StringSlice)
}

func TestURLValuesToObject_StringSlice_CommaSeparated(t *testing.T) {
	values := url.Values{
		"string_slice[]": []string{"value1,value2,value3"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, []string{"value1", "value2", "value3"}, target.StringSlice)
}

func TestURLValuesToObject_IntSlice(t *testing.T) {
	values := url.Values{
		"int_slice": []string{"1", "2", "3"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, target.IntSlice)
}

func TestURLValuesToObject_IntSlice_InvalidValue(t *testing.T) {
	values := url.Values{
		"int_slice": []string{"1", "invalid", "3"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, []int{1, 0, 3}, target.IntSlice)
}

func TestURLValuesToObject_BoolSlice(t *testing.T) {
	values := url.Values{
		"bool_slice": []string{"true", "false", "true"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, []bool{true, false, true}, target.BoolSlice)
}

func TestURLValuesToObject_TimeField(t *testing.T) {
	testTime := time.Date(2026, 8, 14, 12, 30, 0, 0, time.UTC)
	values := url.Values{
		"time_field": []string{testTime.Format(time.RFC3339)},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.True(t, testTime.Equal(target.TimeField))
}

func TestURLValuesToObject_TimeField_DateOnly(t *testing.T) {
	testTime := time.Date(2026, 8, 14, 12, 30, 0, 0, time.UTC)
	values := url.Values{
		"time_field": []string{testTime.Format(time.DateOnly)},
	}
	expectedTime := time.Date(testTime.Year(), testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.UTC)
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.True(t, expectedTime.Equal(target.TimeField))
}

func TestURLValuesToObject_TimeField_InvalidFormat(t *testing.T) {
	values := url.Values{
		"time_field": []string{"invalid-time"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.True(t, target.TimeField.IsZero())
}

func TestURLValuesToObject_MultipleFields(t *testing.T) {
	testTime := time.Date(2026, 8, 14, 12, 30, 0, 0, time.UTC)
	values := url.Values{
		"string_field": []string{"test"},
		"int_field":    []string{"100"},
		"bool_field":   []string{"true"},
		"string_slice": []string{"a", "b"},
		"time_field":   []string{testTime.Format(time.RFC3339)},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, "test", target.StringField)
	assert.Equal(t, 100, target.IntField)
	assert.True(t, target.BoolField)
	assert.Equal(t, []string{"a", "b"}, target.StringSlice)
	assert.True(t, testTime.Equal(target.TimeField))
}

func TestURLValuesToObject_EmptyValues(t *testing.T) {
	values := url.Values{}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, "", target.StringField)
	assert.Equal(t, 0, target.IntField)
	assert.False(t, target.BoolField)
	assert.Nil(t, target.StringSlice)
}

func TestURLValuesToObject_MixedBracketsAndRegular(t *testing.T) {
	values := url.Values{
		"string_slice":   []string{"value1"},
		"string_slice[]": []string{"value2,value3"},
	}
	var target TestStruct
	err := URLValuesToObject(values, &target)
	require.NoError(t, err)
	assert.Equal(t, []string{"value1", "value2", "value3"}, target.StringSlice)
}

func TestUnmarshal_ValidData(t *testing.T) {
	data := []byte("string_field=hello&int_field=42&bool_field=true")
	var target TestStruct
	err := Unmarshal(data, &target)
	require.NoError(t, err)
	assert.Equal(t, "hello", target.StringField)
	assert.Equal(t, 42, target.IntField)
	assert.True(t, target.BoolField)
}

func TestUnmarshal_InvalidData(t *testing.T) {
	data := []byte("%zzz")
	var target TestStruct
	err := Unmarshal(data, &target)
	assert.Error(t, err)
}

func TestUnmarshal_WithSlices(t *testing.T) {
	data := []byte("string_slice[]=a&string_slice[]=b,c&int_slice=1&int_slice=2")
	var target TestStruct
	err := Unmarshal(data, &target)
	require.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, target.StringSlice)
	assert.Equal(t, []int{1, 2}, target.IntSlice)
}
