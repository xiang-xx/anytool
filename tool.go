package anytool

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const sep = "/"

func GetString(a any, path string) (string, error) {
	res, err := Get(a, path)
	if err != nil {
		return "", err
	}
	if s, ok := res.(string); ok {
		return s, nil
	}
	return "", errors.New("ValueError: value is not string")
}

func GetInt(a any, path string) (int, error) {
	res, err := Get(a, path)
	if err != nil {
		return 0, err
	}
	switch i := res.(type) {
	case int8:
		return int(i), nil
	case int16:
		return int(i), nil
	case int32:
		return int(i), nil
	case int64:
		return int(i), nil
	case int:
		return i, nil
	}
	return 0, errors.New("ValueError: value is not int")
}

func GetUint(a any, path string) (uint, error) {
	res, err := Get(a, path)
	if err != nil {
		return 0, err
	}
	switch i := res.(type) {
	case uint8:
		return uint(i), nil
	case uint16:
		return uint(i), nil
	case uint32:
		return uint(i), nil
	case uint64:
		return uint(i), nil
	case uint:
		return i, nil
	}
	return 0, errors.New("ValueError: value is not int")
}

func Get(a any, path string) (any, error) {
	if len(path) == 0 {
		return a, nil
	}

	keys, err := resolvePath(path)
	if err != nil {
		return nil, err
	}

	var ok bool
	for i, k := range keys {
		switch v := a.(type) {
		case map[string]any:
			a, ok = v[k]
			if !ok {
				return nil, fmt.Errorf("PathError: value has no key %s", k)
			}
		case []any:
			index, err := strconv.Atoi(k)
			if err != nil {
				return nil, fmt.Errorf("PathError: key %s is not integer, but value is array", k)
			}
			if index < 0 || index > len(v) {
				return nil, fmt.Errorf("PathError: out of index %d", index)
			}
			a = v[index]
		default:
			return getSlow(a, keys[i:])
		}
	}

	return a, nil
}

func getSlow(a any, keys []string) (any, error) {
	rv := reflect.ValueOf(a)
	for _, k := range keys {
		switch rv.Kind() {
		case reflect.Map:
			krv := reflect.ValueOf(k)
			rv = rv.MapIndex(krv)
			if !rv.IsValid() {
				return nil, fmt.Errorf("PathError: value has no key %s", k)
			}
		case reflect.Array, reflect.Slice:
			index, err := strconv.Atoi(k)
			if err != nil {
				return nil, fmt.Errorf("PathError: key %s is not integer, but value is array", k)
			}
			if index < 0 || index > rv.Len() {
				return nil, fmt.Errorf("PathError: out of index %d", index)
			}
			rv = rv.Index(index)
		default:
			return nil, fmt.Errorf("ValueError: unsupport value type %s", rv.Kind().String())
		}
	}
	return rv.Interface(), nil
}

func resolvePath(path string) ([]string, error) {
	keys := strings.Split(path, sep)
	for _, k := range keys {
		if k == "" {
			return nil, fmt.Errorf("PathError: %s has empty key", path)
		}
	}
	return keys, nil
}
