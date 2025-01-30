package template

import (
	"bytes"
	"fmt"
	"html/template"
	"sort"
	"testing"
)

// Registry holds templates for a specific resource type
type Registry struct {
	templates map[string]*template.Template
	funcMap   template.FuncMap
}

// NewTemplateRegistry creates a new template registry for a resource
func NewTemplateRegistry() *Registry {
	return &Registry{
		templates: make(map[string]*template.Template),
		funcMap:   make(template.FuncMap),
	}
}

// AddTemplate adds a new template to the registry
func (r *Registry) AddTemplate(t testing.TB, name, templateStr string) error {
	t.Helper()

	tmpl := template.New(name)
	if len(r.funcMap) > 0 {
		tmpl = tmpl.Funcs(r.funcMap)
	}

	parsed, err := tmpl.Parse(templateStr)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}
	r.templates[name] = parsed

	return nil
}

// MustAddTemplate is like AddTemplate but panics on error
func (r *Registry) MustAddTemplate(t testing.TB, name, templateStr string) {
	t.Helper()

	if err := r.AddTemplate(t, name, templateStr); err != nil {
		t.Fatalf("failed to add template %q: %s", templateStr, err)
	}
}

// Render renders a template with the given config
func (r *Registry) Render(t testing.TB, templateKey string, cfg map[string]any) (string, error) {
	t.Helper()

	tmpl, exists := r.templates[templateKey]
	if !exists {
		availableTemplates := r.getAvailableTemplates()
		return "", fmt.Errorf("template %q does not exist for resource. Available templates: %v",
			templateKey,
			availableTemplates,
		)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return buf.String(), nil
}

// MustRender is like Render but fails the test on error
func (r *Registry) MustRender(t testing.TB, templateKey string, cfg map[string]any) string {
	t.Helper()

	result, err := r.Render(t, templateKey, cfg)
	if err != nil {
		t.Fatal(err)
	}

	return result
}

// AddFunction adds a custom function to the template registry
func (r *Registry) AddFunction(name string, fn interface{}) {
	if r.funcMap == nil {
		r.funcMap = make(template.FuncMap)
	}
	r.funcMap[name] = fn
}

// HasTemplate checks if a template exists in the registry
func (r *Registry) HasTemplate(key string) bool {
	_, exists := r.templates[key]
	return exists
}

// RemoveTemplate removes a template from the registry
func (r *Registry) RemoveTemplate(key string) {
	delete(r.templates, key)
}

// NewCompositionBuilder creates a new composition builder
func (r *Registry) NewCompositionBuilder() *CompositionBuilder {
	return &CompositionBuilder{
		registry:     r,
		compositions: make([]compositionEntry, 0),
	}
}

// getAvailableTemplates returns a sorted list of available template keys
func (r *Registry) getAvailableTemplates() []string {
	templates := make([]string, 0, len(r.templates))
	for k := range r.templates {
		templates = append(templates, k)
	}
	sort.Strings(templates)

	return templates
}
