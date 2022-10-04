// Code generated by @apexlang/codegen. DO NOT EDIT.

package model

import (
	"context"
	"encoding/json"
	"fmt"
)

type ns struct{}

func (n *ns) Namespace() string {
	return "apexlang.v1"
}

type Parser interface {
	Parse(ctx context.Context, source string) (*ParserResult, error)
}

type Resolver interface {
	Resolve(ctx context.Context, location string, from string) (string, error)
}

type ParserResult struct {
	ns
	Namespace *Namespace `json:"namespace,omitempty" yaml:"namespace,omitempty" msgpack:"namespace,omitempty"`
	Errors    []Error    `json:"errors,omitempty" yaml:"errors,omitempty" msgpack:"errors,omitempty"`
}

type Error struct {
	ns
	Message   string     `json:"message" yaml:"message" msgpack:"message"`
	Positions []uint32   `json:"positions" yaml:"positions" msgpack:"positions"`
	Locations []Location `json:"locations" yaml:"locations" msgpack:"locations"`
}

type Location struct {
	ns
	Line   uint32 `json:"line" yaml:"line" msgpack:"line"`
	Column uint32 `json:"column" yaml:"column" msgpack:"column"`
}

// Namespace encapsulates is used to identify and refer to elements contained in
// the Apex specification.
type Namespace struct {
	ns
	Name        string       `json:"name" yaml:"name" msgpack:"name"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Imports     []Import     `json:"imports,omitempty" yaml:"imports,omitempty" msgpack:"imports,omitempty"`
	Directives  []Directive  `json:"directives,omitempty" yaml:"directives,omitempty" msgpack:"directives,omitempty"`
	Aliases     []Alias      `json:"aliases,omitempty" yaml:"aliases,omitempty" msgpack:"aliases,omitempty"`
	Functions   []Operation  `json:"functions,omitempty" yaml:"functions,omitempty" msgpack:"functions,omitempty"`
	Interfaces  []Interface  `json:"interfaces,omitempty" yaml:"interfaces,omitempty" msgpack:"interfaces,omitempty"`
	Types       []Type       `json:"types,omitempty" yaml:"types,omitempty" msgpack:"types,omitempty"`
	Unions      []Union      `json:"unions,omitempty" yaml:"unions,omitempty" msgpack:"unions,omitempty"`
}

// Apex can integrate external definitions using the import keyword.
type Import struct {
	ns
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	All         bool         `json:"all" yaml:"all" msgpack:"all"`
	Names       []ImportRef  `json:"names,omitempty" yaml:"names,omitempty" msgpack:"names,omitempty"`
	From        string       `json:"from" yaml:"from" msgpack:"from"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

type ImportRef struct {
	ns
	Name string  `json:"name" yaml:"name" msgpack:"name"`
	As   *string `json:"as,omitempty" yaml:"as,omitempty" msgpack:"as,omitempty"`
}

// Types are the most basic component of an Apex specification. They represent data
// structures with fields. Types are defined in a language-agnostic way. This means
// that complex features like nested structures, inheritance, and
// generics/templates are omitted by design.
type Type struct {
	ns
	Name        string       `json:"name" yaml:"name" msgpack:"name"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Fields      []Field      `json:"fields" yaml:"fields" msgpack:"fields"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

// Interfaces are conceptual groups of operations that allow the developer to
// divide communication into multiple components. Typically, interfaces are named
// according to their purpose.
type Interface struct {
	ns
	Name        string       `json:"name" yaml:"name" msgpack:"name"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Operations  []Operation  `json:"operations" yaml:"operations" msgpack:"operations"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

// Alias types are used for cases when scalar types (like string) should be parsed
// our treated like a different data type in the generated code.
type Alias struct {
	ns
	Name        string       `json:"name" yaml:"name" msgpack:"name"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Type        TypeRef      `json:"type" yaml:"type" msgpack:"type"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

type Operation struct {
	ns
	Name        string       `json:"name" yaml:"name" msgpack:"name"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Parameters  []Parameter  `json:"parameters,omitempty" yaml:"parameters,omitempty" msgpack:"parameters,omitempty"`
	Unary       *Parameter   `json:"unary,omitempty" yaml:"unary,omitempty" msgpack:"unary,omitempty"`
	Returns     *TypeRef     `json:"returns,omitempty" yaml:"returns,omitempty" msgpack:"returns,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

type Parameter struct {
	ns
	Name         string       `json:"name" yaml:"name" msgpack:"name"`
	Description  *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Type         TypeRef      `json:"type" yaml:"type" msgpack:"type"`
	DefaultValue *Value       `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty" msgpack:"defaultValue,omitempty"`
	Annotations  []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

type Field struct {
	ns
	Name         string       `json:"name" yaml:"name" msgpack:"name"`
	Description  *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Type         TypeRef      `json:"type" yaml:"type" msgpack:"type"`
	DefaultValue *Value       `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty" msgpack:"defaultValue,omitempty"`
	Annotations  []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

// Unions types denote that a type can have one of several representations.
type Union struct {
	ns
	Name        string       `json:"name" yaml:"name" msgpack:"name"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Types       []TypeRef    `json:"types" yaml:"types" msgpack:"types"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

// Enumerations (or enums) are a type that is constrained to a finite set of
// allowed values.
type Enum struct {
	ns
	Name        string       `json:"name" yaml:"name" msgpack:"name"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Values      []EnumValue  `json:"values" yaml:"values" msgpack:"values"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

type EnumValue struct {
	ns
	Name        string       `json:"name" yaml:"name" msgpack:"name"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Index       uint64       `json:"index" yaml:"index" msgpack:"index"`
	Display     *string      `json:"display,omitempty" yaml:"display,omitempty" msgpack:"display,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty" yaml:"annotations,omitempty" msgpack:"annotations,omitempty"`
}

// Directives are used to ensure that an annotation's arguments match an expected
// format.
type Directive struct {
	ns
	Name        string              `json:"name" yaml:"name" msgpack:"name"`
	Description *string             `json:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty" yaml:"parameters,omitempty" msgpack:"parameters,omitempty"`
	Locations   []DirectiveLocation `json:"locations" yaml:"locations" msgpack:"locations"`
	Require     []DirectiveRequire  `json:"require" yaml:"require" msgpack:"require"`
}

type DirectiveRequire struct {
	ns
	Directive string              `json:"directive" yaml:"directive" msgpack:"directive"`
	Locations []DirectiveLocation `json:"locations" yaml:"locations" msgpack:"locations"`
}

// Annotations attach additional metadata to elements. These can be used in the
// code generation tool to implement custom functionality for your use case.
// Annotations have a name and zero or many arguments.
type Annotation struct {
	ns
	Name      string     `json:"name" yaml:"name" msgpack:"name"`
	Arguments []Argument `json:"arguments,omitempty" yaml:"arguments,omitempty" msgpack:"arguments,omitempty"`
}

type Argument struct {
	ns
	Name  string `json:"name" yaml:"name" msgpack:"name"`
	Value Value  `json:"value" yaml:"value" msgpack:"value"`
}

type Named struct {
	ns
	Kind Kind   `json:"kind" yaml:"kind" msgpack:"kind"`
	Name string `json:"name" yaml:"name" msgpack:"name"`
}

type List struct {
	ns
	Type TypeRef `json:"type" yaml:"type" msgpack:"type"`
}

type Map struct {
	ns
	KeyType   TypeRef `json:"keyType" yaml:"keyType" msgpack:"keyType"`
	ValueType TypeRef `json:"valueType" yaml:"valueType" msgpack:"valueType"`
}

type Stream struct {
	ns
	Type TypeRef `json:"type" yaml:"type" msgpack:"type"`
}

type Optional struct {
	ns
	Type TypeRef `json:"type" yaml:"type" msgpack:"type"`
}

type Reference struct {
	ns
	Name string `json:"name" yaml:"name" msgpack:"name"`
}

type ListValue struct {
	ns
	Values []Value `json:"values" yaml:"values" msgpack:"values"`
}

type ObjectValue struct {
	ns
	Fields []ObjectField `json:"fields" yaml:"fields" msgpack:"fields"`
}

type ObjectField struct {
	ns
	Name  string `json:"name" yaml:"name" msgpack:"name"`
	Value Value  `json:"value" yaml:"value" msgpack:"value"`
}

type TypeRef struct {
	Scalar   *Scalar   `json:"Scalar,omitempty" yaml:"Scalar,omitempty" msgpack:"Scalar,omitempty"`
	Named    *Named    `json:"Named,omitempty" yaml:"Named,omitempty" msgpack:"Named,omitempty"`
	List     *List     `json:"List,omitempty" yaml:"List,omitempty" msgpack:"List,omitempty"`
	Map      *Map      `json:"Map,omitempty" yaml:"Map,omitempty" msgpack:"Map,omitempty"`
	Stream   *Stream   `json:"Stream,omitempty" yaml:"Stream,omitempty" msgpack:"Stream,omitempty"`
	Optional *Optional `json:"Optional,omitempty" yaml:"Optional,omitempty" msgpack:"Optional,omitempty"`
}

type Value struct {
	Bool        *bool        `json:"bool,omitempty" yaml:"bool,omitempty" msgpack:"bool,omitempty"`
	String      *string      `json:"string,omitempty" yaml:"string,omitempty" msgpack:"string,omitempty"`
	I64         *int64       `json:"i64,omitempty" yaml:"i64,omitempty" msgpack:"i64,omitempty"`
	F64         *float64     `json:"f64,omitempty" yaml:"f64,omitempty" msgpack:"f64,omitempty"`
	Reference   *Reference   `json:"Reference,omitempty" yaml:"Reference,omitempty" msgpack:"Reference,omitempty"`
	ListValue   *ListValue   `json:"ListValue,omitempty" yaml:"ListValue,omitempty" msgpack:"ListValue,omitempty"`
	ObjectValue *ObjectValue `json:"ObjectValue,omitempty" yaml:"ObjectValue,omitempty" msgpack:"ObjectValue,omitempty"`
}

type DirectiveLocation int32

const (
	DirectiveLocationNamespace DirectiveLocation = 0
	DirectiveLocationAlias     DirectiveLocation = 1
	DirectiveLocationUnion     DirectiveLocation = 2
	DirectiveLocationEnum      DirectiveLocation = 3
	DirectiveLocationEnumValue DirectiveLocation = 4
	DirectiveLocationType      DirectiveLocation = 5
	DirectiveLocationField     DirectiveLocation = 6
	DirectiveLocationInterface DirectiveLocation = 7
	DirectiveLocationOperation DirectiveLocation = 8
	DirectiveLocationParameter DirectiveLocation = 9
)

var toStringDirectiveLocation = map[DirectiveLocation]string{
	DirectiveLocationNamespace: "NAMESPACE",
	DirectiveLocationAlias:     "ALIAS",
	DirectiveLocationUnion:     "UNION",
	DirectiveLocationEnum:      "ENUM",
	DirectiveLocationEnumValue: "ENUM_VALUE",
	DirectiveLocationType:      "TYPE",
	DirectiveLocationField:     "FIELD",
	DirectiveLocationInterface: "INTERFACE",
	DirectiveLocationOperation: "OPERATION",
	DirectiveLocationParameter: "PARAMETER",
}

var toIDDirectiveLocation = map[string]DirectiveLocation{
	"NAMESPACE":  DirectiveLocationNamespace,
	"ALIAS":      DirectiveLocationAlias,
	"UNION":      DirectiveLocationUnion,
	"ENUM":       DirectiveLocationEnum,
	"ENUM_VALUE": DirectiveLocationEnumValue,
	"TYPE":       DirectiveLocationType,
	"FIELD":      DirectiveLocationField,
	"INTERFACE":  DirectiveLocationInterface,
	"OPERATION":  DirectiveLocationOperation,
	"PARAMETER":  DirectiveLocationParameter,
}

func (e DirectiveLocation) Type() string {
	return "DirectiveLocation"
}

func (e DirectiveLocation) String() string {
	str, ok := toStringDirectiveLocation[e]
	if !ok {
		return "unknown"
	}
	return str
}

func (e *DirectiveLocation) FromString(str string) (ok bool) {
	*e, ok = toIDDirectiveLocation[str]
	return ok
}

// MarshalJSON marshals the enum as a quoted json string
func (e DirectiveLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (e *DirectiveLocation) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	if !e.FromString(str) {
		return fmt.Errorf("unknown value %q for DirectiveLocation", str)
	}
	return nil
}

type Scalar int32

const (
	ScalarString   Scalar = 1
	ScalarBool     Scalar = 2
	ScalarI8       Scalar = 3
	ScalarI16      Scalar = 4
	ScalarI32      Scalar = 5
	ScalarI64      Scalar = 6
	ScalarU8       Scalar = 7
	ScalarU16      Scalar = 8
	ScalarU32      Scalar = 9
	ScalarU64      Scalar = 10
	ScalarF32      Scalar = 11
	ScalarF64      Scalar = 12
	ScalarBytes    Scalar = 13
	ScalarDatetime Scalar = 14
	ScalarAny      Scalar = 15
	ScalarRaw      Scalar = 16
)

var toStringScalar = map[Scalar]string{
	ScalarString:   "STRING",
	ScalarBool:     "BOOL",
	ScalarI8:       "I8",
	ScalarI16:      "I16",
	ScalarI32:      "I32",
	ScalarI64:      "I64",
	ScalarU8:       "U8",
	ScalarU16:      "U16",
	ScalarU32:      "U32",
	ScalarU64:      "U64",
	ScalarF32:      "F32",
	ScalarF64:      "F64",
	ScalarBytes:    "BYTES",
	ScalarDatetime: "DATETIME",
	ScalarAny:      "ANY",
	ScalarRaw:      "RAW",
}

var toIDScalar = map[string]Scalar{
	"STRING":   ScalarString,
	"BOOL":     ScalarBool,
	"I8":       ScalarI8,
	"I16":      ScalarI16,
	"I32":      ScalarI32,
	"I64":      ScalarI64,
	"U8":       ScalarU8,
	"U16":      ScalarU16,
	"U32":      ScalarU32,
	"U64":      ScalarU64,
	"F32":      ScalarF32,
	"F64":      ScalarF64,
	"BYTES":    ScalarBytes,
	"DATETIME": ScalarDatetime,
	"ANY":      ScalarAny,
	"RAW":      ScalarRaw,
}

func (e Scalar) Type() string {
	return "Scalar"
}

func (e Scalar) String() string {
	str, ok := toStringScalar[e]
	if !ok {
		return "unknown"
	}
	return str
}

func (e *Scalar) FromString(str string) (ok bool) {
	*e, ok = toIDScalar[str]
	return ok
}

// MarshalJSON marshals the enum as a quoted json string
func (e Scalar) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (e *Scalar) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	if !e.FromString(str) {
		return fmt.Errorf("unknown value %q for Scalar", str)
	}
	return nil
}

type Kind int32

const (
	KindType      Kind = 1
	KindFunc      Kind = 2
	KindInterface Kind = 3
	KindAlias     Kind = 4
	KindUnion     Kind = 5
	KindEnum      Kind = 6
)

var toStringKind = map[Kind]string{
	KindType:      "TYPE",
	KindFunc:      "FUNC",
	KindInterface: "INTERFACE",
	KindAlias:     "ALIAS",
	KindUnion:     "UNION",
	KindEnum:      "ENUM",
}

var toIDKind = map[string]Kind{
	"TYPE":      KindType,
	"FUNC":      KindFunc,
	"INTERFACE": KindInterface,
	"ALIAS":     KindAlias,
	"UNION":     KindUnion,
	"ENUM":      KindEnum,
}

func (e Kind) Type() string {
	return "Kind"
}

func (e Kind) String() string {
	str, ok := toStringKind[e]
	if !ok {
		return "unknown"
	}
	return str
}

func (e *Kind) FromString(str string) (ok bool) {
	*e, ok = toIDKind[str]
	return ok
}

// MarshalJSON marshals the enum as a quoted json string
func (e Kind) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (e *Kind) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	if !e.FromString(str) {
		return fmt.Errorf("unknown value %q for Kind", str)
	}
	return nil
}
