// Read stdin expecting a yaml file, and write data as JSON on stdout.

package yaml2json

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v2"
)

// ReadYAML reads the input, parses it as yaml, and returns the parsed result.
func ReadYAML(in io.Reader) (interface{}, error) {

	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, fmt.Errorf("failed to read yaml input: %w", err)
	}

	var t interface{}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return nil, fmt.Errorf("failed to parse yaml input: %w", err)
	}

	return t, nil
}

// PrintJSON prints the supplied value as formatted json.
func PrintJSON(out io.Writer, val interface{}) error {

	val = ConvertKeys(val)

	result, err := json.MarshalIndent(val, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	_, err = fmt.Fprintln(out, string(result))

	return err
}

// ConvertKeys replaces keys in map[interface{}] so that we can have a map[string].
// yaml supports keys of various types, but json only supports string keys.
func ConvertKeys(in interface{}) interface{} {

	msi, ok := in.(map[string]interface{})
	if ok {
		out := map[string]interface{}{}
		for k, v := range msi {
			out[k] = ConvertKeys(v)
		}

		return out
	}

	mii, ok := in.(map[interface{}]interface{})
	if ok {
		out := map[string]interface{}{}
		for k, v := range mii {
			out[toString(k)] = ConvertKeys(v)
		}

		return out
	}

	si, ok := in.([]interface{})
	if ok {
		out := []interface{}{}
		for _, i := range si {
			out = append(out, ConvertKeys(i))
		}

		return out
	}

	return in
}

// stringer has a String()
type stringer interface {
	String() string
}

func toString(v interface{}) string {

	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		return v.(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v)
	case reflect.Bool:
		return fmt.Sprintf("%t", v)
	}

	s, ok2 := v.(stringer)
	if ok2 {
		return s.String()
	}

	// last resort
	return fmt.Sprintf("%s", v)
}
