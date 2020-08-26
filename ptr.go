package ptr

// FromString convert to pointer
func FromString(s string) *string { return &s }

// FromInt convert to pointer
func FromInt(s int) *int { return &s }

// FromInt32 convert to pointer
func FromInt32(s int32) *int32 { return &s }

// FromInt64 convert to pointer
func FromInt64(s int64) *int64 { return &s }

// FromUint convert to pointer
func FromUint(s uint) *uint { return &s }

// FromUint32 convert to pointer
func FromUint32(s uint32) *uint32 { return &s }

// FromUint64 convert to pointer
func FromUint64(s uint64) *uint64 { return &s }

// FromFloat32 convert to pointer
func FromFloat32(s float32) *float32 { return &s }

// FromFloat64 convert to pointer
func FromFloat64(s float64) *float64 { return &s }
