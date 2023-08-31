package jsonschema

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/iancoleman/orderedmap"
)

type Schema struct {
	Definitions       Definitions            `json:"$defs,omitempty"`
	OneOf             []*Schema              `json:"oneOf,omitempty"`
	Items             *Schema                `json:"items,omitempty"`
	Properties        *orderedmap.OrderedMap `json:"properties,omitempty"`
	PatternProperties map[string]*Schema     `json:"patternProperties,omitempty"`
	Type              string                 `json:"type,omitempty"`
	Required          []string               `json:"required,omitempty"`
	Format            string                 `json:"format,omitempty"`
	ContentEncoding   string                 `json:"contentEncoding,omitempty"`
	Title             string                 `json:"title,omitempty"`
	Description       string                 `json:"description,omitempty"`
}

// Reflect reflects to Schema from a value using the default Reflector
func Reflect(v interface{}) *Schema {
	doc := ReflectFromType(reflect.TypeOf(v))
	render := ReflectFromType(reflect.TypeOf(&DefaultRender{}))
	render.Properties.Set("data", doc)
	return render
}

// ReflectFromType generates root schema using the default Reflector
func ReflectFromType(t reflect.Type) *Schema {
	r := &Reflector{}
	return r.ReflectFromType(t)
}

// A Reflector reflects values into a Schema.
type Reflector struct {

	// RequiredFromJSONSchemaTags will cause the Reflector to generate a schema
	// that requires any key tagged with `jsonschema:required`, overriding the
	// default of requiring any key *not* tagged with `json:,omitempty`.
	RequiredFromJSONSchemaTags bool

	// KeyNamer allows customizing of key names.
	// The default is to use the key's name as is, or the json tag if present.
	// If a json tag is present, KeyNamer will receive the tag's name as an argument, not the original key name.
	KeyNamer func(string) string
}

// Reflect reflects to Schema from a value.
func (r *Reflector) Reflect(v interface{}) *Schema {
	return r.ReflectFromType(reflect.TypeOf(v))
}

// ReflectFromType generates root schema
func (r *Reflector) ReflectFromType(t reflect.Type) *Schema {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return r.reflectTypeToSchema(Definitions{}, t)
}

type Definitions map[string]*Schema

var (
	timeType = reflect.TypeOf(time.Time{})
	uriType  = reflect.TypeOf(url.URL{})
)

// Byte slices will be encoded as base64
var byteSliceType = reflect.TypeOf([]byte(nil))

// Except for json.RawMessage
var rawMessageType = reflect.TypeOf(json.RawMessage{})

func (r *Reflector) reflectTypeToSchema(definitions Definitions, t reflect.Type) *Schema {
	st := &Schema{}
	switch t.Kind() {
	case reflect.Struct:
		r.reflectStruct(definitions, t, st)

	case reflect.Slice, reflect.Array:
		r.reflectSliceOrArray(definitions, t, st)

	case reflect.Map:
		r.reflectMap(definitions, t, st)

	case reflect.Interface:
		// empty

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		st.Type = "integer"

	case reflect.Float32, reflect.Float64:
		st.Type = "number"

	case reflect.Bool:
		st.Type = "boolean"

	case reflect.String:
		st.Type = "string"

	default:
		panic("unsupported type " + t.String())
	}

	return st
}

func (r *Reflector) reflectSliceOrArray(definitions Definitions, t reflect.Type, st *Schema) {
	if t == rawMessageType {
		return
	}
	if t.Kind() == reflect.Slice && t.Elem() == byteSliceType.Elem() {
		st.Type = "string"
		// NOTE: ContentMediaType is not set here
		st.ContentEncoding = "base64"
	} else {
		st.Type = "array"
		st.Items = r.reflectTypeToSchema(definitions, t.Elem())
	}
}

func (r *Reflector) reflectMap(definitions Definitions, t reflect.Type, st *Schema) {
	st.Type = "object"
	switch t.Key().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		st.PatternProperties = map[string]*Schema{
			"^[0-9]+$": r.reflectTypeToSchema(definitions, t.Elem()),
		}
		return
	}
	if t.Elem().Kind() != reflect.Interface {
		st.PatternProperties = map[string]*Schema{
			".*": r.reflectTypeToSchema(definitions, t.Elem()),
		}
	}
}

// Reflects a struct to a JSON Schema type.
func (r *Reflector) reflectStruct(definitions Definitions, t reflect.Type, s *Schema) {
	// Handle special types
	switch t {
	case timeType: // date-time RFC section 7.3.1
		s.Type = "string"
		s.Format = "date-time"
		return
	case uriType: // uri RFC section 7.3.6
		s.Type = "string"
		s.Format = "uri"
		return
	}

	s.Type = "object"
	s.Properties = orderedmap.New()
	r.reflectStructFields(s, definitions, t)
}

func (r *Reflector) reflectStructFields(st *Schema, definitions Definitions, t reflect.Type) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	handleField := func(f reflect.StructField) {
		name, shouldEmbed, required := r.reflectFieldName(f)
		// if anonymous and exported type should be processed recursively
		// current type should inherit properties of anonymous one
		if name == "" {
			if shouldEmbed {
				r.reflectStructFields(st, definitions, f.Type)
			}
			return
		}

		property := r.reflectTypeToSchema(definitions, f.Type)
		property.structKeywordsFromTags(f)

		st.Properties.Set(name, property)
		if required {
			st.Required = appendUniqueString(st.Required, name)
		}
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		handleField(f)
	}
}

func appendUniqueString(base []string, value string) []string {
	for _, v := range base {
		if v == value {
			return base
		}
	}
	return append(base, value)
}

func (t *Schema) structKeywordsFromTags(f reflect.StructField) {
	desc := f.Tag.Get("doc")
	// t.Title = desc
	t.Description = desc
}

func requiredFromJSONTags(tags []string) bool {
	if ignoredByJSONTags(tags) {
		return false
	}

	for _, tag := range tags[1:] {
		if tag == "omitempty" {
			return false
		}
	}
	return true
}

func ignoredByJSONTags(tags []string) bool {
	return tags[0] == "-"
}

func (r *Reflector) reflectFieldName(f reflect.StructField) (string, bool, bool) {
	jsonTagString, _ := f.Tag.Lookup("json")
	jsonTags := strings.Split(jsonTagString, ",")
	required := requiredFromJSONTags(jsonTags)

	if f.Anonymous && jsonTags[0] == "" {
		// As per JSON Marshal rules, anonymous structs are inherited
		if f.Type.Kind() == reflect.Struct {
			return "", true, false
		}

		// As per JSON Marshal rules, anonymous pointer to structs are inherited
		if f.Type.Kind() == reflect.Ptr && f.Type.Elem().Kind() == reflect.Struct {
			return "", true, false
		}
	}

	// Try to determine the name from the different combos
	name := f.Name
	if jsonTags[0] != "" {
		name = jsonTags[0]
	}
	if !f.Anonymous && f.PkgPath != "" {
		// field not anonymous and not export has no export name
		name = ""
	} else if r.KeyNamer != nil {
		name = r.KeyNamer(name)
	}

	return name, false, required
}
