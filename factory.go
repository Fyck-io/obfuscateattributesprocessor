package obfuscateattributesprocessor

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
)

const (
	// The value of "type" key in configuration.
	typeStr = "obfuscateattributes"
)

// NewFactory creates a factory for the routing processor.
func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
	)
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
	}
}
