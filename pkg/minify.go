package pkg

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// Minify takes a yaml structure and transforms it into a minified structure.
// Example:
// nodeA:
//   nodeB: value
// Output:
// {nodeA: {nodeB: value}}
func Minify(orig []byte) ([]byte, error) {
	var (
		b bytes.Buffer
	)
	b.WriteByte('{')
	// The map transformation makes it loose the ordering...
	output := make(map[string]interface{})
	if err := yaml.Unmarshal(orig, output); err != nil {
		return nil, err
	}
	transform(output, &b)
	b.WriteByte('}')
	return b.Bytes(), nil
}

// TODO: return error
func transform(in map[string]interface{}, b *bytes.Buffer) {
	// probably tokenize, lookahead... etc.
	// how to check wether
	index := len(in)
	for k, v := range in {
		b.WriteString(k + ":")
		switch i := v.(type) {
		case string:
			// must consider multiple variables on the same level separated by ,
			b.WriteString(i)
		case map[string]interface{}:
			// this should be a recursive call with an added indentation and buffer being passed around.
			b.WriteString("{")
			transform(i, b)
			b.WriteString("}")
		case []interface{}:
		}
		// if next item, insert ,
		if index > 1 {
			b.WriteString(",")
			index--
		}
	}
}

// Maxify expands a minified yaml structure into a proper yaml structure.
// Example:
// {nodeA: {nodeB: value}}
// Output:
// nodeA:
//   nodeB: value
func Maxify(orig []byte) ([]byte, error) {
	return nil, nil
}
