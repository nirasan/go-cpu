package go_cpu

import "testing"

func TestNewRegister(t *testing.T) {
	r := NewRegister()
	r.Set32(2359271448)
	if r.Get32() != 2359271448 {
		t.Fatalf("invalid register: %v", r)
	}
	if r.Get16() != 40984 {
		t.Fatalf("invalid register: %v", r)
	}
	if r.GetH8() != 160 {
		t.Fatalf("invalid register: %v", r)
	}
	if r.GetL8() != 24 {
		t.Fatalf("invalid register: %v", r)
	}
}

func TestRegister_Get32(t *testing.T) {
	r := NewRegister()
	r.v = 100
	if r.Get32() != 100 {
		t.Fatalf("invalid register: %v", r)
	}
}

func TestRegister_Get16(t *testing.T) {
	r := NewRegister()
	samples := []struct {
		In  uint32
		Out uint32
	}{
		{0, 0},
		{100, 100},
		{v1100, 0},
		{v1101, v0001},
		{v1111, v0011},
	}
	for _, sample := range samples {
		r.v = sample.In
		if r.Get16() != uint16(sample.Out) {
			t.Errorf("invalid: %v", r.v)
		}
	}
}

func TestRegister_GetH8(t *testing.T) {
	r := NewRegister()
	samples := []struct {
		In  uint32
		Out uint32
	}{
		{0, 0},
		{100, 0},
		{v1100, 0},
		{v1101, 0},
		{v1111, v0001},
		{100 << 8, 100},
	}
	for _, sample := range samples {
		r.v = sample.In
		if r.GetH8() != uint8(sample.Out) {
			t.Errorf("invalid: %#v, %#v", sample, r.v)
		}
	}
}

func TestRegister_GetL8(t *testing.T) {
	r := NewRegister()
	samples := []struct {
		In  uint32
		Out uint32
	}{
		{0, 0},
		{100, 100},
		{v1100, 0},
		{v1101, v0001},
		{v1111, v0001},
		{100 << 8, 0},
	}
	for _, sample := range samples {
		r.v = sample.In
		if r.GetL8() != uint8(sample.Out) {
			t.Errorf("invalid: %#v, %#v", sample, r.v)
		}
	}
}


func TestRegister_Set16(t *testing.T) {
	r := NewRegister()
	samples := []struct {
		Before uint32
		In     uint32
		After  uint32
	}{
		{0, 0, 0},
		{0, 100, 100},
		{v1100, v0001, v1101},
		{v0011, v0001, v0001},
	}
	for _, sample := range samples {
		r.v = sample.Before
		r.Set16(uint16(sample.In))
		if r.v != sample.After {
			t.Errorf("invalid: %v", r.v)
		}
	}
}

func TestRegister_SetH8(t *testing.T) {
	r := NewRegister()
	samples := []struct {
		Before uint32
		In     uint32
		After  uint32
	}{
		{0, 0, 0},
		{0, 100, 100 << 8},
		{v1100, v0001, v1110},
		{v1111, 0, v1101},
	}
	for _, sample := range samples {
		r.v = sample.Before
		r.SetH8(uint8(sample.In))
		if r.v != sample.After {
			t.Errorf("invalid: %v", r.v)
		}
	}
}

func TestRegister_SetL8(t *testing.T) {
	r := NewRegister()
	samples := []struct {
		Before uint32
		In     uint32
		After  uint32
	}{
		{0, 0, 0},
		{0, 100, 100},
		{v1100, v0001, v1101},
		{v1111, 0, v1110},
	}
	for _, sample := range samples {
		r.v = sample.Before
		r.SetL8(uint8(sample.In))
		if r.v != sample.After {
			t.Errorf("invalid: %v", r.v)
		}
	}
}
