package template

import (
	"fmt"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/provider"
)

var (
	globalTemplateSet *ResourceTemplateSet
	gtSetOnce         sync.Once
)

// ResourceTemplateSet represents a set of related resources and their templates
type ResourceTemplateSet struct {
	resources map[string]*schema.Resource
	registry  *Registry
	t         testing.TB
}

// NewTemplateSet creates a new template set
func NewTemplateSet(t testing.TB) *ResourceTemplateSet {
	return &ResourceTemplateSet{
		resources: make(map[string]*schema.Resource),
		registry:  NewTemplateRegistry(),
		t:         t,
	}
}

// GetTemplateSet returns a singleton instance of ResourceTemplateSet with all provider resources pre-registered
func GetTemplateSet(t testing.TB) *ResourceTemplateSet {
	gtSetOnce.Do(func() {
		globalTemplateSet = initTemplateSet(t)
	})

	globalTemplateSet.t = t // Update testing.TB for the current test context

	return globalTemplateSet
}

// AddTemplate adds a custom template to the template set
func (s *ResourceTemplateSet) AddTemplate(name, templateStr string) *ResourceTemplateSet {
	s.registry.MustAddTemplate(s.t, name, templateStr)

	return s
}

// NewBuilder creates a new composition builder
func (s *ResourceTemplateSet) NewBuilder() *CompositionBuilder {
	return &CompositionBuilder{
		registry:     s.registry,
		compositions: make([]compositionEntry, 0),
	}
}

// initTemplateSet initializes a new template set with all provider resources
func initTemplateSet(t testing.TB) *ResourceTemplateSet {
	p, err := provider.Provider("dev")
	if err != nil {
		t.Fatalf("failed to get provider: %v", err)
	}

	set := &ResourceTemplateSet{
		registry: NewTemplateRegistry(),
		t:        t,
	}

	// Register all resources
	for resourceType, resource := range p.ResourcesMap {
		set.registerResource(resourceType, resource, ResourceKindResource)
	}

	// Register all data sources
	for resourceType, resource := range p.DataSourcesMap {
		set.registerResource(resourceType, resource, ResourceKindDataSource)
	}

	return set
}

// registerResource handles the registration of a single resource or data source
func (s *ResourceTemplateSet) registerResource(resourceType string, r *schema.Resource, kind ResourceKind) {
	generator := NewSchemaTemplateGenerator()

	// Register template functions
	for name, fn := range generator.funcMap {
		s.registry.AddFunction(name, fn)
	}

	// Generate and register the template
	template := generator.GenerateTemplate(r, resourceType, kind)
	s.registry.MustAddTemplate(s.t, templateKey(resourceType, kind), template)
}

// templateKey generates a unique template key based on resource type and kind
func templateKey(resourceType string, kind ResourceKind) string {
	switch kind {
	case ResourceKindResource:
		return fmt.Sprintf("resource.%s", resourceType)
	case ResourceKindDataSource:
		return fmt.Sprintf("data.%s", resourceType)
	default:
		return resourceType
	}
}
