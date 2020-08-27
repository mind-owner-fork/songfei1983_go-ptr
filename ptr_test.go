package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	want := ""
	assert.Equal(t, &want, FromString(want))
}

func TestFromInt(t *testing.T) {
	want := 1
	assert.Equal(t, &want, FromInt(want))
}

func TestFromInt32(t *testing.T) {
	var want int32 = 1
	assert.Equal(t, &want, FromInt32(want))
}

func TestFromInt64(t *testing.T) {
	var want int64 = 1
	assert.Equal(t, &want, FromInt64(want))
}

func TestFromUint(t *testing.T) {
	var want uint = 1
	assert.Equal(t, &want, FromUint(want))
}

func TestFromUint32(t *testing.T) {
	var want uint32 = 1
	assert.Equal(t, &want, FromUint32(want))
}

func TestFromUint64(t *testing.T) {
	var want uint64 = 1
	assert.Equal(t, &want, FromUint64(want))
}

func TestFromFloat32(t *testing.T) {
	var want float32 = 1
	assert.Equal(t, &want, FromFloat32(want))
}

func TestFromFloat64(t *testing.T) {
	var want float64 = 1
	assert.Equal(t, &want, FromFloat64(want))
}
