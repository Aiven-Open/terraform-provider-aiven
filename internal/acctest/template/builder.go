package template

import (
	"fmt"
	"strings"
	"testing"
)

// CompositionBuilder helps build complex compositions of templates
type CompositionBuilder struct {
	registry     *Registry
	compositions []compositionEntry
}

// Add adds a new template and config to the composition
func (b *CompositionBuilder) Add(templateKey string, cfg map[string]any) *CompositionBuilder {
	b.compositions = append(b.compositions, compositionEntry{
		TemplateKey:  templateKey,
		Config:       cfg,
		ResourceType: templateKey, // For custom templates, use template key as resource type
	})
	return b
}

// AddResource helper methods to explicitly add resources or data sources
func (b *CompositionBuilder) AddResource(resourceType string, cfg map[string]any) *CompositionBuilder {
	b.compositions = append(b.compositions, compositionEntry{
		TemplateKey:  templateKey(resourceType, ResourceKindResource),
		Config:       cfg,
		ResourceType: resourceType,
		ResourceKind: ResourceKindResource,
	})
	return b
}

func (b *CompositionBuilder) AddDataSource(resourceType string, cfg map[string]any) *CompositionBuilder {
	b.compositions = append(b.compositions, compositionEntry{
		TemplateKey:  templateKey(resourceType, ResourceKindDataSource),
		Config:       cfg,
		ResourceType: resourceType,
		ResourceKind: ResourceKindDataSource,
	})
	return b
}

// AddWithConfig adds a new template and config to the composition using a resourceConfig
func (b *CompositionBuilder) AddWithConfig(templateKey string, cfg resourceConfig) *CompositionBuilder {
	b.compositions = append(b.compositions, compositionEntry{
		TemplateKey: templateKey,
		Config:      cfg.ToMap(),
	})
	return b
}

// AddIf conditional method to CompositionBuilder
func (b *CompositionBuilder) AddIf(condition bool, templateKey string, cfg map[string]any) *CompositionBuilder {
	if condition {
		return b.Add(templateKey, cfg)
	}

	return b
}

// todo: fix
func (b *CompositionBuilder) Remove(templateKey string) *CompositionBuilder {
	var newCompositions []compositionEntry
	for _, comp := range b.compositions {
		if comp.TemplateKey != templateKey {
			newCompositions = append(newCompositions, comp)
		}
	}
	b.compositions = newCompositions

	return b
}

// Render renders all templates in the composition and combines them
func (b *CompositionBuilder) Render(t testing.TB) (string, error) {
	t.Helper()

	var renderedParts = make([]string, 0, len(b.compositions))

	// Render each template
	for _, comp := range b.compositions {
		rendered, err := b.registry.Render(t, comp.TemplateKey, comp.Config)
		if err != nil {
			return "", fmt.Errorf("failed to render template %s: %w", comp.TemplateKey, err)
		}
		renderedParts = append(renderedParts, rendered)
	}

	// Combine all rendered parts
	combined := strings.Join(renderedParts, "\n\n")

	//TODO: add HCL validation?

	return combined, nil
}

func (b *CompositionBuilder) RemoveByResourceName(name string) *CompositionBuilder {
	var newCompositions []compositionEntry
	for _, comp := range b.compositions {
		if resName, ok := comp.Config["resource_name"]; !ok || resName != name {
			newCompositions = append(newCompositions, comp)
		}
	}
	b.compositions = newCompositions
	return b
}

// MustRender is like Render but fails the test on error
func (b *CompositionBuilder) MustRender(t testing.TB) string {
	t.Helper()

	result, err := b.Render(t)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

// compositionEntry represents a combination of template and its config
type compositionEntry struct {
	TemplateKey  string
	Config       map[string]any
	ResourceType string // Add this to track the resource type
	ResourceKind ResourceKind
}
