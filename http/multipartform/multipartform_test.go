package multipartform

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/beaconsoftwarellc/gadget/v2/generator"
	"github.com/beaconsoftwarellc/quimby/v2/http/multipartform/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	Foo string
	Bar int
	Baz []string
}

type FormDataWithArrays struct {
	Headers string
	MyArray []string
}

type FormDataWithStructs struct {
	PtrTestStruct *TestStruct
	TestStruct    TestStruct
}

type ErryThang struct {
	PtrTestStruct           *TestStruct
	Array                   []string
	Bool                    bool
	Int                     int
	IntEight                int8
	IntSixteen              int16
	IntThirtyTwo            int32
	IntSixtyFour            int64
	YouInt                  uint
	YouIntEight             uint8
	YouIntSixteen           uint16
	YouIntThirstyTwo        uint32
	YouIntSixtyFour         uint64
	RootBeerFloatThirstyTwo float32
	RootBeerFloatSixtyFour  float64
	Yarn                    string
	MapNoWorky              map[string]string
	MapWorky                map[string]interface{}
}

func TestUnmarshalRaw_ErryThang(t *testing.T) {
	assert := assert.New(t)
	data := []byte(testdata.FormDataErryThang)
	boundary := "xYzZY"
	target := &ErryThang{}
	require.NoError(t, UnmarshalRaw(boundary, data, target))
	assert.NotNil(target.PtrTestStruct)
	assert.Equal(2, len(target.Array))
	assert.Equal(true, target.Bool)
	assert.Equal(12345678, target.Int)
	assert.Equal(int8(-1), target.IntEight)
	assert.Equal(int16(-512), target.IntSixteen)
	assert.Equal(int32(-1024), target.IntThirtyTwo)
	assert.Equal(int64(-2048), target.IntSixtyFour)
	assert.Equal(uint(1), target.YouInt)
	assert.Equal(uint8(2), target.YouIntEight)
	assert.Equal(uint16(4), target.YouIntSixteen)
	assert.Equal(uint32(16), target.YouIntThirstyTwo)
	assert.Equal(uint64(32), target.YouIntSixtyFour)
	assert.Equal(float32(1.1), target.RootBeerFloatThirstyTwo)
	assert.Equal(float64(1.2), target.RootBeerFloatSixtyFour)
	assert.Equal("spam", target.Yarn)
	assert.Nil(target.MapNoWorky)
	assert.Equal("hello", target.MapWorky["um"])
	assert.Equal("eggs and spam", target.MapWorky["green"])
}

func TestUnmarshalRaw_Arrays(t *testing.T) {
	assert := assert.New(t)
	data := []byte(testdata.FormDataWithArrays)
	boundary := "xYzZY"
	target := &FormDataWithArrays{}
	require.NoError(t, UnmarshalRaw(boundary, data, target))
	require.Equal(t, 2, len(target.MyArray))
	assert.Equal("I am a dimunitive dispenser of flavor infused water",
		target.MyArray[0])
	assert.Equal("short and stout", target.MyArray[1])
}

func TestUnmarshalRaw_Structs(t *testing.T) {
	assert := assert.New(t)
	data := []byte(testdata.FormDataWithStructs)
	boundary := "xYzZY"
	target := &FormDataWithStructs{}
	require.NoError(t, UnmarshalRaw(boundary, data, target))
	require.NotNil(t, target.PtrTestStruct)
	assert.Equal("qux", target.PtrTestStruct.Foo)
	assert.Equal(1, target.PtrTestStruct.Bar)
	assert.Equal([]string{"elm1", "elm2"}, target.PtrTestStruct.Baz)
	assert.Equal("quux", target.TestStruct.Foo)
	assert.Equal(2, target.TestStruct.Bar)
	assert.Equal([]string{"1mle", "2mle"}, target.TestStruct.Baz)
}

func Test_getCastToMap(t *testing.T) {
	assert := assert.New(t)
	expected := make(map[string]string)
	expected[generator.ID("key")] = generator.ID("value")
	expected[generator.ID("key")] = generator.ID("value")
	expected[generator.ID("key")] = generator.ID("value")
	expected[generator.ID("key")] = generator.ID("value")
	data, err := json.Marshal(expected)
	assert.NoError(err)
	argument := []string{string(data)}
	expectedType := reflect.TypeOf(expected)
	cast := getCastToMap(expectedType)
	actual, err := cast(argument)
	assert.NoError(err)
	actualMap := actual.(map[string]interface{})
	for key, value := range expected {
		assert.Equal(value, actualMap[key].(string))
	}
}

func Test_getCastToStruct(t *testing.T) {
	assert := assert.New(t)
	expected := TestStruct{
		Foo: "qux",
		Bar: 1337,
		Baz: []string{"quux", "corge", "grault"},
	}
	cast := getCastToStruct(reflect.TypeOf(expected))
	data, err := json.Marshal(expected)
	assert.NoError(err)
	actual, err := cast([]string{string(data)})
	assert.NoError(err)
	assert.Equal(expected, actual.(TestStruct))
}

func Test_getCastToPtr(t *testing.T) {
	assert := assert.New(t)
	expected := &TestStruct{
		Foo: "qux",
		Bar: 1337,
		Baz: []string{"quux", "corge", "grault"},
	}
	cast := getCastToPtr(reflect.TypeOf(expected))
	data, err := json.Marshal(expected)
	assert.NoError(err)
	actual, err := cast([]string{string(data)})
	assert.NoError(err)
	assert.Equal(expected, actual.(*TestStruct))
}

func Test_getCastToSlice_IntArray(t *testing.T) {
	assert := assert.New(t)

	expected := []int{1, 2, 3, 4}

	cast, err := getCastToSlice(reflect.TypeOf(expected))
	assert.NoError(err)

	actual, err := cast([]string{"1", "2", "3", "4"})
	assert.NoError(err)
	assert.ElementsMatch(expected, actual)
}

func Test_getCastToSlice_PtrArray(t *testing.T) {
	assert := assert.New(t)
	expected := []*TestStruct{
		{Foo: "1", Bar: 1},
		{Foo: "2", Bar: 2},
	}
	element, err := json.Marshal(expected[0])
	assert.NoError(err)
	element1, err := json.Marshal(expected[1])
	assert.NoError(err)
	data := []string{string(element), string(element1)}
	cast, err := getCastToSlice(reflect.TypeOf(expected))
	assert.NoError(err)
	obj, err := cast(data)
	assert.NoError(err)
	actual := obj.([]*TestStruct)
	assert.Equal(len(expected), len(actual))
	for i, element := range expected {
		assert.Equal(element.Foo, actual[i].Foo)
		assert.Equal(element.Bar, actual[i].Bar)
		assert.Equal(element.Baz, actual[i].Baz)
	}
}

func Test_setField(t *testing.T) {
	assert := assert.New(t)
	target := &TestStruct{
		Foo: "baz",
		Bar: 1234,
		Baz: []string{"qux", "quux"},
	}
	targetValue := reflect.Indirect(reflect.ValueOf(target))
	field := targetValue.Type().Field(0)
	fieldValue := targetValue.Field(0)

	// string valued
	expectedString := generator.String(32)
	setField(field, fieldValue, []string{expectedString})
	assert.Equal(target.Foo, expectedString)

	// int valued
	expectedInt := generator.Int()
	field = targetValue.Type().Field(1)
	fieldValue = targetValue.Field(1)
	setField(field, fieldValue, []string{fmt.Sprintf("%d", expectedInt)})
	assert.Equal(expectedInt, target.Bar)
}
