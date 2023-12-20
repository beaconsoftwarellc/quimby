package urlencodedform

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/beaconsoftwarellc/gadget/v2/log"
	"github.com/beaconsoftwarellc/gadget/v2/stringutil"
)

// Unmarshal parses the url encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
func Unmarshal(data []byte, target any) error {
	values, err := url.ParseQuery(string(data))
	if nil != err {
		return err
	}
	return URLValuesToObject(values, target)
}

func valuesToArray(fieldType reflect.Type, values []string) []interface{} {
	elemKind := fieldType.Elem().Kind()
	anon := make([]interface{}, len(values))
	coerce := func(v string) interface{} { return v }
	if elemKind == reflect.Int {
		coerce = func(v string) interface{} {
			i, err := strconv.Atoi(v)
			if nil != err {
				log.Warnf("error encountered coercing url value '%s' to int: %s", v, err)
			}
			return i
		}
	} else if elemKind == reflect.Bool {
		coerce = func(v string) interface{} {
			return v == "true"
		}
	} else if elemKind != reflect.String {
		log.Warnf("valuesToArray recieved unhandled kind %v", elemKind)
		return nil
	}
	for i, v := range values {
		anon[i] = coerce(v)
	}
	return anon
}

func breakUpArrayValuedParameter(values []string) []string {
	newValues := make([]string, 0, len(values))
	for i := 0; i < len(values); i++ {
		newValues = append(newValues, strings.Split(values[i], ",")...)
	}
	return newValues
}

func inspectModel(target interface{}) map[string]reflect.Type {
	v := reflect.Indirect(reflect.ValueOf(target))
	fieldMap := make(map[string]reflect.Type, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fieldMap[v.Type().Field(i).Name] = v.Field(i).Type()
	}
	return fieldMap
}

// URLValuesToObject parses the url values and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// URLValuesToObject returns an InvalidUnmarshalError.
func URLValuesToObject(values url.Values, target interface{}) error {
	var err error
	fieldMap := inspectModel(target)
	valueMap := make(map[string]interface{})
	for fieldName, fieldType := range fieldMap {
		urlFieldName := stringutil.Underscore(fieldName)
		queryValues := values[urlFieldName]
		queryValues = append(queryValues,
			breakUpArrayValuedParameter(values[urlFieldName+"[]"])...)

		if len(queryValues) == 0 {
			continue
		}
		switch fieldType.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Struct, reflect.UnsafePointer:
			continue
		case reflect.Slice, reflect.Array:
			if len(queryValues) == 0 {
				continue
			}
			if arrayValues := valuesToArray(fieldType, queryValues); nil != arrayValues {
				valueMap[fieldName] = arrayValues
			} else {
				log.Warnf("failed to assigned values to array for field '%s' kind '%s'",
					urlFieldName, fieldType.Kind())
			}
		case reflect.String:
			valueMap[fieldName] = queryValues[0]
		case reflect.Int:
			valueMap[fieldName], err = strconv.Atoi(queryValues[0])
			if nil != err {
				log.Warnf("error parsing int for field name '%s' and value '%s'",
					fieldName, queryValues[0])
			}
		case reflect.Bool:
			valueMap[fieldName] = queryValues[0] == "true"

		default:
			log.Warnf("unhandled kind %v will be omitted from url values", fieldType)
		}
	}
	data, err := json.Marshal(valueMap)
	if nil != err {
		return err
	}
	return json.Unmarshal(data, target)
}
