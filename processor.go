package obfuscateattributesprocessor

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var _ component.TracesProcessor = (*obfuscateAttributes)(nil)

type obfuscateAttributes struct {
	next consumer.Traces
}

func newObfuscateAttributes(ctx context.Context, config *Config, next consumer.Traces) (*obfuscateAttributes, error) {
	return &obfuscateAttributes{next: next}, nil
}

func (s *obfuscateAttributes) ConsumeTraces(ctx context.Context, batch ptrace.Traces) error {
	log.Println("Workinggggg")
	return nil
}

func (s *obfuscateAttributes) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

// Start the obfuscate attributes processor
func (s *obfuscateAttributes) Start(_ context.Context, _ component.Host) error {
	return nil
}

// Shutdown the obfuscate attributes processor
func (s *obfuscateAttributes) Shutdown(context.Context) error {
	return nil
}
