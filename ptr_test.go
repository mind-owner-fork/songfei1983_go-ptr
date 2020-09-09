package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	t.Parallel()
	want := ""
	assert.Equal(t, &want, FromString(want))
}

func TestFromInt(t *testing.T) {
	t.Parallel()
	want := 1
	assert.Equal(t, &want, FromInt(want))
}

func TestFromInt32(t *testing.T) {
	t.Parallel()
	var want int32 = 1
	assert.Equal(t, &want, FromInt32(want))
}

func TestFromInt64(t *testing.T) {
	t.Parallel()
	var want int64 = 1
	assert.Equal(t, &want, FromInt64(want))
}

func TestFromUint(t *testing.T) {
	t.Parallel()
	var want uint = 1
	assert.Equal(t, &want, FromUint(want))
}

func TestFromUint32(t *testing.T) {
	t.Parallel()
	var want uint32 = 1
	assert.Equal(t, &want, FromUint32(want))
}

func TestFromUint64(t *testing.T) {
	t.Parallel()
	var want uint64 = 1
	assert.Equal(t, &want, FromUint64(want))
}

func TestFromFloat32(t *testing.T) {
	t.Parallel()
	var want float32 = 1
	assert.Equal(t, &want, FromFloat32(want))
}

func TestFromFloat64(t *testing.T) {
	t.Parallel()
	var want float64 = 1
	assert.Equal(t, &want, FromFloat64(want))
}

func BenchmarkFromInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt(i)
	}
}
