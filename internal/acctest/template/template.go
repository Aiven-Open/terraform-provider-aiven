package template

import (
	"fmt"
	"html/template"
)

// ResourceConfig is the interface that all resource configs must implement
type resourceConfig interface {
	// ToMap converts the config to a map for template rendering
	ToMap() map[string]any
}

// Template represents a single Terraform configuration template
type Template struct {
	Name     string
	Template string
}

// In internal/acctest/template/value.go

// TemplateValue represents a value that can either be a literal string or a reference
type TemplateValue struct {
	Value     string
	IsLiteral bool
}

// Literal creates a new literal value
func Literal(v string) TemplateValue {
	return TemplateValue{Value: v, IsLiteral: true}
}

// Reference creates a new reference value
func Reference(v string) TemplateValue {
	return TemplateValue{Value: v, IsLiteral: false}
}

// String returns the properly formatted value based on whether it's a literal or reference
func (v TemplateValue) String() string {
	if v.IsLiteral {
		return fmt.Sprintf("%q", v.Value)
	}
	return v.Value
}

// MarshalText implements encoding.TextMarshaler
func (v TemplateValue) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// Add template functions for the generator
func templateValueFuncs() template.FuncMap {
	return template.FuncMap{
		"renderValue": func(v interface{}) string {
			switch val := v.(type) {
			case TemplateValue:
				return val.String()
			case string:
				return fmt.Sprintf("%q", val)
			case int, int64, float64, bool:
				return fmt.Sprintf("%v", val)
			default:
				return fmt.Sprintf("%v", val)
			}
		},
	}
}
