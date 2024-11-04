package multipartform

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/beaconsoftwarellc/gadget/v2/errors"
	"github.com/beaconsoftwarellc/gadget/v2/log"
	"github.com/beaconsoftwarellc/gadget/v2/stringutil"
	"github.com/spf13/cast"
)

const (
	// MaxMemoryBytes for form parsing
	// Note: choosing 30 MB here to account for parsing an email sent to
	// a webhook.
	MaxMemoryBytes = 30 << 20
)

type castE func([]string) (any, error)

// UnmarshalRaw the passed multi-part form encoded data into the passed
// target interface using the passed boundary for reading the form
func UnmarshalRaw(boundary string, data []byte, target any) error {
	var (
		reader          = bytes.NewReader(data)
		multiPartReader = multipart.NewReader(reader, boundary)
		form, err       = multiPartReader.ReadForm(MaxMemoryBytes)
	)
	if nil != err && !errors.Is(err, io.EOF) {
		return err
	}
	return Unmarshal(form, target)
}

// Unmarshal the passed form into the passed
// target interface using the passed boundary for reading the form.
// Expects complex types to be JSON encoded strings
func Unmarshal(form *multipart.Form, target any) error {
	if nil == form {
		return errors.New("multipart.Form cannot be nil")
	}
	var (
		targetValue = reflect.Indirect(reflect.ValueOf(target))
		targetType  = targetValue.Type()
		formValues  = clean(form.Value)
	)
	for i := 0; i < targetValue.NumField(); i++ {
		field := targetType.Field(i)
		fieldName := stringutil.Underscore(field.Name)

		formValue, ok := formValues[fieldName]
		if !ok {
			formValue, ok = formValues[field.Name]
		}
		if !ok {
			// just leave it zero valued since we could not find a
			// value for it
			continue
		}
		setField(field, targetValue.Field(i), formValue)
	}
	return nil
}

func clean(values map[string][]string) map[string][]string {
	clean := make(map[string][]string)
	for key, value := range values {
		if idx := strings.Index(key, "["); idx > 0 {
			key = key[0:idx]
		}
		clean[key] = value
	}
	return clean
}

func setField(
	field reflect.StructField,
	fieldValue reflect.Value,
	formValue []string,
) {
	fieldType := field.Type
	cast, err := getCastForType(fieldType)
	if nil != err {
		log.Warnf("error encountered getting cast for '%s': %s",
			field.Name, err)
		return
	}
	value, err := cast(formValue)
	if nil != err {
		log.Warnf("error encountered casting value for '%s': %s",
			field.Name, err)
		return
	}
	if !reflect.TypeOf(value).AssignableTo(fieldType) {
		log.Warnf("casted value %v is not assignable to %s",
			value, field.Name)
	} else {
		fieldValue.Set(reflect.ValueOf(value))
	}
}

func getCastForType(reflectType reflect.Type) (castE, error) {
	var (
		f         castE
		err       error
		errFormat = "reflect.Type='%s' is unhandled"
	)
	switch reflectType.Kind() {
	case reflect.Invalid:
		err = errors.Newf(errFormat, reflectType)
	case reflect.Bool:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToBoolE(v[0])
			}
			return r, err
		}
	case reflect.Int:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToIntE(v[0])
			}
			return r, err
		}
	case reflect.Int8:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToInt8E(v[0])
			}
			return r, err
		}
	case reflect.Int16:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToInt16E(v[0])
			}
			return r, err
		}
	case reflect.Int32:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToInt32E(v[0])
			}
			return r, err
		}
	case reflect.Int64:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToInt64E(v[0])
			}
			return r, err
		}
	case reflect.Uint:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToUintE(v[0])
			}
			return r, err
		}
	case reflect.Uint8:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToUint8E(v[0])
			}
			return r, err
		}
	case reflect.Uint16:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToUint16E(v[0])
			}
			return r, err
		}
	case reflect.Uint32:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToUint32E(v[0])
			}
			return r, err
		}
	case reflect.Uint64:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToUint64E(v[0])
			}
			return r, err
		}
	case reflect.Uintptr:
		err = errors.Newf(errFormat, reflectType)
	case reflect.Float32:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToFloat32E(v[0])
			}
			return r, err
		}
	case reflect.Float64:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r, err = cast.ToFloat64E(v[0])
			}
			return r, err
		}
	case reflect.Complex64:
		err = errors.Newf(errFormat, reflectType)
	case reflect.Complex128:
		err = errors.Newf(errFormat, reflectType)
	case reflect.Array, reflect.Slice:
		f, err = getCastToSlice(reflectType)
	case reflect.Chan:
		err = errors.Newf(errFormat, reflectType)
	case reflect.Func:
		err = errors.Newf(errFormat, reflectType)
	case reflect.Pointer:
		f = getCastToPtr(reflectType)
	case reflect.String:
		f = func(v []string) (any, error) {
			var r any
			var err error
			if len(v) > 0 {
				r = v[0]
			}
			return r, err
		}
	case reflect.Struct:
		f = getCastToStruct(reflectType)
	case reflect.Map:
		f = getCastToMap(reflectType)
	case reflect.UnsafePointer:
		err = errors.Newf(errFormat, reflectType)
	default:
		err = errors.Newf(errFormat, reflectType)
	}
	return f, err
}

func getCastToSlice(reflectType reflect.Type) (castE, error) {
	var (
		elementCastE castE
		err          error
	)
	elementCastE, err = getCastForType(reflectType.Elem())
	if nil != err {
		return nil, err
	}
	return func(sa []string) (any, error) {
		var (
			value = reflect.MakeSlice(reflectType, 0, len(sa))
		)
		for _, s := range sa {
			obj, err := elementCastE([]string{s})
			if nil != err {
				return nil, err
			}
			value = reflect.Append(value, reflect.ValueOf(obj))
		}
		return value.Interface(), nil
	}, nil
}

func getCastToPtr(reflectType reflect.Type) castE {
	return func(s []string) (any, error) {
		var (
			value  = reflect.New(reflectType.Elem())
			target = value.Interface()
			err    error
		)

		if len(s) > 0 {
			err = json.Unmarshal([]byte(s[0]), target)
		}
		return target, err
	}
}

func getCastToStruct(reflectType reflect.Type) castE {
	return func(s []string) (any, error) {
		var (
			// get a pointer to the struct
			value  = reflect.New(reflectType)
			target = value.Interface()
			err    error
		)
		if len(s) > 0 {
			// pass the pointer in since json will not
			// accept a struct
			err = json.Unmarshal([]byte(s[0]), target)
		}
		// return the struct the pointer is pointing to
		return reflect.ValueOf(target).Elem().Interface(), err
	}
}

func getCastToMap(reflectType reflect.Type) castE {
	return func(s []string) (any, error) {
		var (
			value  = reflect.MakeMap(reflectType)
			target = value.Interface()
			err    error
		)
		if len(s) > 0 {
			// KNOWN ISSUE: this is going to convert
			// 	the map from map[Y]T to map[Y]interface{}
			// 	which is really weird
			err = json.Unmarshal([]byte(s[0]), &target)
		}
		return target, err
	}
}
