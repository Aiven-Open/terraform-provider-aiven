package template

import (
	"fmt"
	"html/template"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceKind represents the type of terraform configuration item
type ResourceKind int

const (
	ResourceKindResource ResourceKind = iota
	ResourceKindDataSource
)

// String returns the string representation of ResourceKind
func (k ResourceKind) String() string {
	switch k {
	case ResourceKindResource:
		return "resource"
	case ResourceKindDataSource:
		return "data"
	default:
		return "unknown"
	}
}

// SchemaTemplateGenerator generates test templates from Terraform schemas
type SchemaTemplateGenerator struct {
	funcMap template.FuncMap
}

// NewSchemaTemplateGenerator creates a new template generator for a specific resource type
func NewSchemaTemplateGenerator() *SchemaTemplateGenerator {
	return &SchemaTemplateGenerator{
		funcMap: template.FuncMap{
			"required": func(v any) any {
				if v == nil {
					panic("required field is missing")
				}
				return v
			},
			"renderValue": func(v interface{}) template.HTML {
				var result string
				switch val := v.(type) {
				case TemplateValue:
					if val.IsLiteral {
						result = fmt.Sprintf("%q", val.Value)
					} else {
						result = val.Value
					}
				case string:
					result = fmt.Sprintf("%q", val)
				case int, int64, float64:
					result = fmt.Sprintf("%v", val)
				case bool:
					result = fmt.Sprintf("%v", val)
				default:
					result = fmt.Sprintf("%q", val)
				}
				return template.HTML(result)
			},
		},
	}
}

// GenerateTemplate generates a template for either a resource or data source
func (g *SchemaTemplateGenerator) GenerateTemplate(r *schema.Resource, resourceType string, kind ResourceKind) string {
	var b strings.Builder

	// Header differs based on kind
	_, _ = fmt.Fprintf(&b, "%s %q %q {\n", kind, resourceType, "{{ required .resource_name }}")

	g.generateFields(&b, r.Schema, 1)

	b.WriteString("}")

	return b.String()
}

func (g *SchemaTemplateGenerator) generateFields(b *strings.Builder, s map[string]*schema.Schema, indent int) {
	// Create indent string once
	indentStr := strings.Repeat("  ", indent)

	// Collect fields by type
	var (
		required []string
		optional []string
		lists    []string
		maps     []string
	)

	for k, field := range s {
		if field.Computed && !field.Optional && !field.Required {
			continue
		}

		switch field.Type {
		case schema.TypeList, schema.TypeSet:
			lists = append(lists, k)
		case schema.TypeMap:
			maps = append(maps, k)
		default:
			if field.Required {
				required = append(required, k)
			} else {
				optional = append(optional, k)
			}
		}
	}

	// Sort all field groups for consistent ordering
	sort.Strings(required)
	sort.Strings(optional)
	sort.Strings(lists)
	sort.Strings(maps)

	// Process fields in specified order
	for _, field := range required {
		g.generateField(b, field, s[field], indentStr)
	}

	for _, field := range optional {
		g.generateField(b, field, s[field], indentStr)
	}

	for _, field := range lists {
		g.generateField(b, field, s[field], indentStr)
	}

	for _, field := range maps {
		g.generateField(b, field, s[field], indentStr)
	}
}

func (g *SchemaTemplateGenerator) generateField(b *strings.Builder, field string, schemaField *schema.Schema, indent string) {
	switch schemaField.Type {
	case schema.TypeString:
		if schemaField.Required {
			fmt.Fprintf(b, "%s%s = {{ renderValue (required .%s) }}\n", indent, field, field)
		} else {
			fmt.Fprintf(b, "%s{{- if .%s }}\n", indent, field)
			fmt.Fprintf(b, "%s%s = {{ renderValue .%s }}\n", indent, field, field)
			fmt.Fprintf(b, "%s{{- end }}\n", indent)
		}

	case schema.TypeInt, schema.TypeFloat:
		if schemaField.Required {
			fmt.Fprintf(b, "%s%s = {{ required .%s }}\n", indent, field, field)
		} else {
			fmt.Fprintf(b, "%s{{- if ne .%s nil }}\n", indent, field)
			fmt.Fprintf(b, "%s%s = {{ .%s }}\n", indent, field, field)
			fmt.Fprintf(b, "%s{{- end }}\n", indent)
		}

	case schema.TypeMap:
		if schemaField.Required {
			fmt.Fprintf(b, "%s%s = {\n", indent, field)
			fmt.Fprintf(b, "%s  {{- range $k, $v := (required .%s) }}\n", indent, field)
			fmt.Fprintf(b, "%s  {{ renderValue $k }} = {{ renderValue $v }}\n", indent)
			fmt.Fprintf(b, "%s  {{- end }}\n", indent)
			fmt.Fprintf(b, "%s}\n", indent)
		} else {
			fmt.Fprintf(b, "%s{{- if .%s }}\n", indent, field)
			fmt.Fprintf(b, "%s%s = {\n", indent, field)
			fmt.Fprintf(b, "%s  {{- range $k, $v := .%s }}\n", indent, field)
			fmt.Fprintf(b, "%s  {{ renderValue $k }} = {{ renderValue $v }}\n", indent)
			fmt.Fprintf(b, "%s  {{- end }}\n", indent)
			fmt.Fprintf(b, "%s}\n", indent)
			fmt.Fprintf(b, "%s{{- end }}\n", indent)
		}

	default:
		if schemaField.Required {
			fmt.Fprintf(b, "%s%s = {{ renderValue (required .%s) }}\n", indent, field, field)
		} else {
			fmt.Fprintf(b, "%s{{- if .%s }}\n", indent, field)
			fmt.Fprintf(b, "%s%s = {{ renderValue .%s }}\n", indent, field, field)
			fmt.Fprintf(b, "%s{{- end }}\n", indent)
		}
	}
}
