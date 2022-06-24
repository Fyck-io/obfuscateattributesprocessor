package obfuscateattributesprocessor

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

var _ component.TracesProcessor = (*obfuscateAttributes)(nil)

type obfuscateAttributes struct {
	next consumer.Traces
}

func newObfuscateAttributes(ctx context.Context, config *Config, logger *zap.Logger, next consumer.Traces) (*obfuscateAttributes, error) {
	return &obfuscateAttributes{next: next}, nil
}

// func (s *obfuscateAttributes) processTraces(ctx context.Context, batch ptrace.Traces) (ptrace.Traces, error) {
// 	for i := 0; i < batch.ResourceSpans().Len(); i++ {
// 		rs := batch.ResourceSpans().At(i)
// 		s.processResourceSpan(ctx, rs)
// 	}
// 	return batch, nil
// }

// // processResourceSpan processes the RS and all of its spans and then returns the last
// // view metric context. The context can be used for tests
// func (s *obfuscateAttributes) processResourceSpan(ctx context.Context, rs ptrace.ResourceSpans) {

// }

// // processAttrs redacts the attributes of a resource span or a span
// func (s *obfuscateAttributes) processAttrs(_ context.Context, attributes *pcommon.Map) {

// }

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
