package transformer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/binding/spec"
	"github.com/cloudevents/sdk-go/v2/binding/test"
)

func TestDeleteAttribute(t *testing.T) {
	withSubjectEvent := test.MinEvent()
	withSubjectEvent.Context = withSubjectEvent.Context.AsV1()
	require.NoError(t, withSubjectEvent.Context.SetSubject("aaa"))

	withTimeEvent := withSubjectEvent.Clone()
	require.NoError(t, withTimeEvent.Context.SetTime(time.Now()))

	noSubjectEvent := withSubjectEvent.Clone()
	require.NoError(t, noSubjectEvent.Context.SetSubject(""))

	test.RunTransformerTests(t, context.Background(), []test.TransformerTestArgs{
		{
			Name:         "Remove subject from Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(withSubjectEvent.Clone()),
			WantEvent:    noSubjectEvent,
			Transformers: binding.Transformers{DeleteAttribute(spec.Subject)},
		},
		{
			Name:         "Remove subject from Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(withSubjectEvent.Clone()),
			WantEvent:    noSubjectEvent,
			Transformers: binding.Transformers{DeleteAttribute(spec.Subject)},
		},
		{
			Name:         "Remove subject from Event message",
			InputEvent:   withSubjectEvent,
			WantEvent:    noSubjectEvent,
			Transformers: binding.Transformers{DeleteAttribute(spec.Subject)},
		},
		{
			Name:         "Remove time from Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(withTimeEvent.Clone()),
			WantEvent:    withSubjectEvent.Clone(),
			Transformers: binding.Transformers{DeleteAttribute(spec.Time)},
		},
		{
			Name:         "Remove time from Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(withTimeEvent.Clone()),
			WantEvent:    withSubjectEvent.Clone(),
			Transformers: binding.Transformers{DeleteAttribute(spec.Time)},
		},
		{
			Name:         "Remove time from Event message",
			InputEvent:   withTimeEvent,
			WantEvent:    withSubjectEvent,
			Transformers: binding.Transformers{DeleteAttribute(spec.Time)},
		},
		{
			Name:         "Do nothing with Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(withSubjectEvent.Clone()),
			WantEvent:    withSubjectEvent.Clone(),
			Transformers: binding.Transformers{DeleteAttribute(spec.Time)},
		},
		{
			Name:         "Do nothing with Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(withSubjectEvent.Clone()),
			WantEvent:    withSubjectEvent.Clone(),
			Transformers: binding.Transformers{DeleteAttribute(spec.Time)},
		},
		{
			Name:         "Do nothing with Event message",
			InputEvent:   withSubjectEvent,
			WantEvent:    withSubjectEvent,
			Transformers: binding.Transformers{DeleteAttribute(spec.Time)},
		},
	})
}

func TestDeleteExtension(t *testing.T) {
	e := test.MinEvent()
	e.Context = e.Context.AsV1()

	extName := "aaa"
	extValue := "bbb"
	expectedEventWithExtension := e.Clone()
	require.NoError(t, expectedEventWithExtension.Context.SetExtension(extName, extValue))

	test.RunTransformerTests(t, context.Background(), []test.TransformerTestArgs{
		{
			Name:         "No change to Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(expectedEventWithExtension.Clone()),
			WantEvent:    expectedEventWithExtension.Clone(),
			Transformers: binding.Transformers{DeleteExtension("ccc")},
		},
		{
			Name:         "No change to Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(expectedEventWithExtension.Clone()),
			WantEvent:    expectedEventWithExtension.Clone(),
			Transformers: binding.Transformers{DeleteExtension("ccc")},
		},
		{
			Name:         "No change to Event message",
			InputEvent:   expectedEventWithExtension,
			WantEvent:    expectedEventWithExtension,
			Transformers: binding.Transformers{DeleteExtension("ccc")},
		},
		{
			Name:         "Delete extension 'aaa' from Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(expectedEventWithExtension.Clone()),
			WantEvent:    e.Clone(),
			Transformers: binding.Transformers{DeleteExtension(extName)},
		},
		{
			Name:         "Delete extension 'aaa' from Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(expectedEventWithExtension.Clone()),
			WantEvent:    e.Clone(),
			Transformers: binding.Transformers{DeleteExtension(extName)},
		},
		{
			Name:         "Delete extension 'aaa' from Event message",
			InputEvent:   expectedEventWithExtension,
			WantEvent:    e,
			Transformers: binding.Transformers{DeleteExtension(extName)},
		},
	})
}
