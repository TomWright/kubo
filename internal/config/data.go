package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

var (
	// ErrUnexpectedType is used when you try to use a data item as a type that it is not.
	// E.g. attempting to read a string as an int.
	ErrUnexpectedType = errors.New("unexpected type")
	// ErrNotFound is used when you try to read a data item that does not exist.
	ErrNotFound = errors.New("not found")
)

// Data represents a piece of config data.
// It could be a root element, or a child element somewhere in the config tree.
type Data map[interface{}]interface{}

// Set sets the given value under the given key.
// The key supports the use of dots to identity object access.
func (d Data) Set(key string, value interface{}) error {
	return d.set(strings.Split(key, "."), value)
}

// set is the recursive func that allows Set to work as it does.
func (d Data) set(keys []string, value interface{}) error {
	switch len(keys) {
	case 0:
		return nil
	case 1:
		d[keys[0]] = value
		return nil
	}

	keyData, err := d.Data(keys[0])
	switch err {
	case ErrUnexpectedType, ErrNotFound:
		keyData = make(Data)
		d[keys[0]] = keyData
		return keyData.set(keys[1:], value)
	case nil:
		return keyData.set(keys[1:], value)
	default:
		return err
	}
}

// DataR returns a Data type that must exist at the given key.
func (d Data) DataR(key string) Data {
	result, err := d.Data(key)
	if err != nil {
		panic(err)
	}
	return result
}

// Data returns a Data type from the given key.
func (d Data) Data(key string) (Data, error) {
	val, _ := d[key]
	switch v := val.(type) {
	case Data:
		return v, nil
	case nil:
		return nil, ErrNotFound
	default:
		return nil, ErrUnexpectedType
	}
}

// StringR returns a string that must exist at the given key.
func (d Data) StringR(key string) string {
	result, err := d.String(key)
	if err != nil {
		panic(err)
	}
	return result
}

// String returns a string from the given key.
func (d Data) String(key string) (string, error) {
	val, _ := d[key]
	switch v := val.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case nil:
		return "", ErrNotFound
	default:
		return "", ErrUnexpectedType
	}
}

// BytesR returns a byte slice that must exist at the given key.
func (d Data) BytesR(key string) []byte {
	result, err := d.Bytes(key)
	if err != nil {
		panic(err)
	}
	return result
}

// Bytes returns a byte slice from the given key.
func (d Data) Bytes(key string) ([]byte, error) {
	val, _ := d[key]
	switch v := val.(type) {
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	case nil:
		return nil, ErrNotFound
	default:
		return nil, ErrUnexpectedType
	}
}

// IntR returns an int that must exist at the given key.
func (d Data) IntR(key string) int {
	result, err := d.Int(key)
	if err != nil {
		panic(err)
	}
	return result
}

// Int returns an int from the given key.
func (d Data) Int(key string) (int, error) {
	val, _ := d[key]
	switch v := val.(type) {
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case nil:
		return 0, ErrNotFound
	default:
		return 0, ErrUnexpectedType
	}
}

// FromBytes returns some Data that is represented by the given bytes.
func FromBytes(byteData []byte) (Data, error) {
	var data Data
	if err := yaml.Unmarshal(byteData, &data); err != nil {
		return data, fmt.Errorf("could not unmarshal config data: %w", err)
	}
	return data, nil
}

// LoadFromFile loads Data from the given file.
func LoadFromFile(filename string) (Data, error) {
	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %w", err)
	}
	return FromBytes(byteData)
}

// SaveToFile saves Data to the given file.
func SaveToFile(data Data, filename string) error {
	byteData, err := ToBytes(data)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, byteData, 0666); err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}
	return nil
}

// ToBytes returns the byte representation of the given Data.
func ToBytes(data Data) ([]byte, error) {
	byteData, err := yaml.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal config data: %w", err)
	}
	return byteData, nil
}
