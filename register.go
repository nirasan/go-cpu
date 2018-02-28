package go_cpu

var (
	v0001 uint32 = 255
	v0010 uint32 = 65280
	v0011 uint32 = 65535
	v1111 uint32 = 4294967295
	v1100 uint32 = 4294901760
	v1101 uint32 = 4294902015
	v1110 uint32 = 4294967040
)

type Register struct {
	v uint32
}

func NewRegister() *Register {
	return &Register{}
}

func (r *Register) Get32() uint32 {
	return r.v
}

func (r *Register) Set32(v uint32) {
	r.v = v
}

func (r *Register) Get16() uint16 {
	return uint16(r.v & v0011)
}

func (r *Register) Set16(v uint16) {
	r.v = r.v & v1100 + uint32(v)
}

func (r *Register) GetH8() uint8 {
	return uint8((r.v & v0010) >> 8)
}

func (r *Register) SetH8(v uint8) {
	r.v = r.v & v1101 + uint32(v) << 8
}

func (r *Register) GetL8() uint8 {
	return uint8(r.v & v0001)
}

func (r *Register) SetL8(v uint8) {
	r.v = r.v & v1110 + uint32(v)
}
