package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinkStatuses_Set(t *testing.T) {

	ss := SinkStatuses{}

	ss.Set("bar", 1, "foo")

	if assert.Len(t, ss, 1) {
		s := ss[0]
		assert.Equal(t, "bar", s.Name)
		assert.Equal(t, "foo", s.LastMessage.Data)
		if assert.Len(t, s.Metrics, 1) {
			assert.Equal(t, uint32(1), s.Metrics[0].Replica)
			assert.Equal(t, uint64(1), s.Metrics[0].Total)
		}
	}

	ss.Set("bar", 1, "bar")

	if assert.Len(t, ss, 1) {
		s := ss[0]
		assert.Equal(t, "bar", s.LastMessage.Data)
		if assert.Len(t, s.Metrics, 1) {
			assert.Equal(t, uint32(1), s.Metrics[0].Replica)
			assert.Equal(t, uint64(2), s.Metrics[0].Total)
		}
	}

	ss.Set("bar", 0, "foo")

	if assert.Len(t, ss, 1) {
		s := ss[0]
		assert.Equal(t, "foo", s.LastMessage.Data)
		if assert.Len(t, s.Metrics, 2) {
			assert.Equal(t, uint32(1), s.Metrics[0].Replica)
			assert.Equal(t, uint64(2), s.Metrics[0].Total)
			assert.Equal(t, uint32(0), s.Metrics[1].Replica)
			assert.Equal(t, uint64(1), s.Metrics[1].Total)
		}
	}

	ss.Set("baz", 0, "foo")

	if assert.Len(t, ss, 2) {
		s := ss[1]
		assert.Equal(t, "baz", s.Name)
		assert.Equal(t, "foo", s.LastMessage.Data)
		if assert.Len(t, s.Metrics, 1) {
			assert.Equal(t, uint32(0), s.Metrics[0].Replica)
			assert.Equal(t, uint64(1), s.Metrics[0].Total)
		}
	}
}