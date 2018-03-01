package go_cpu

import "fmt"

type Cpu struct {
	ra *Register
	rb *Register
	rc *Register
	rd *Register
	q  *Queue
}

func NewCpu() *Cpu {
	return &Cpu{
		ra: NewRegister(),
		rb: NewRegister(),
		rc: NewRegister(),
		rd: NewRegister(),
		q:  NewQueue(),
	}
}

func (c *Cpu) GetEAX() uint32 {
	return c.ra.Get32()
}

func (c *Cpu) SetEAX(v uint32) {
	c.ra.Set32(v)
}

func (c *Cpu) GetAX() uint16 {
	return c.ra.Get16()
}

func (c *Cpu) SetAX(v uint16) {
	c.ra.Set16(v)
}

func (c *Cpu) GetAH() uint8 {
	return c.ra.GetH8()
}

func (c *Cpu) SetAH(v uint8) {
	c.ra.SetH8(v)
}

func (c *Cpu) GetAL() uint8 {
	return c.ra.GetL8()
}

func (c *Cpu) SetAL(v uint8) {
	c.ra.SetL8(v)
}

func (c *Cpu) GetEBX() uint32 {
	return c.rb.Get32()
}

func (c *Cpu) SetEBX(v uint32) {
	c.rb.Set32(v)
}

func (c *Cpu) GetBX() uint16 {
	return c.rb.Get16()
}

func (c *Cpu) SetBX(v uint16) {
	c.rb.Set16(v)
}

func (c *Cpu) GetBH() uint8 {
	return c.rb.GetH8()
}

func (c *Cpu) SetBH(v uint8) {
	c.rb.SetH8(v)
}

func (c *Cpu) GetBL() uint8 {
	return c.rb.GetL8()
}

func (c *Cpu) SetBL(v uint8) {
	c.rb.SetL8(v)
}

func (c *Cpu) GetECX() uint32 {
	return c.rc.Get32()
}

func (c *Cpu) SetECX(v uint32) {
	c.rc.Set32(v)
}

func (c *Cpu) GetCX() uint16 {
	return c.rc.Get16()
}

func (c *Cpu) SetCX(v uint16) {
	c.rc.Set16(v)
}

func (c *Cpu) GetCH() uint8 {
	return c.rc.GetH8()
}

func (c *Cpu) SetCH(v uint8) {
	c.rc.SetH8(v)
}

func (c *Cpu) GetCL() uint8 {
	return c.rc.GetL8()
}

func (c *Cpu) SetCL(v uint8) {
	c.rc.SetL8(v)
}

func (c *Cpu) GetEDX() uint32 {
	return c.rd.Get32()
}

func (c *Cpu) SetEDX(v uint32) {
	c.rd.Set32(v)
}

func (c *Cpu) GetDX() uint16 {
	return c.rd.Get16()
}

func (c *Cpu) SetDX(v uint16) {
	c.rd.Set16(v)
}

func (c *Cpu) GetDH() uint8 {
	return c.rd.GetH8()
}

func (c *Cpu) SetDH(v uint8) {
	c.rd.SetH8(v)
}

func (c *Cpu) GetDL() uint8 {
	return c.rd.GetL8()
}

func (c *Cpu) SetDL(v uint8) {
	c.rd.SetL8(v)
}

func (c *Cpu) String() string {
	return fmt.Sprintf("EAX: %d, AX: %d, AH: %d, AL: %d\n", c.GetEAX(), c.GetAX(), c.GetAH(), c.GetAL()) +
		fmt.Sprintf("EBX: %d, BX: %d, BH: %d, BL: %d\n", c.GetEBX(), c.GetBX(), c.GetBH(), c.GetBL()) +
		fmt.Sprintf("ECX: %d, CX: %d, CH: %d, CL: %d\n", c.GetECX(), c.GetCX(), c.GetCH(), c.GetCL()) +
		fmt.Sprintf("EDX: %d, DX: %d, DH: %d, DL: %d\n", c.GetEDX(), c.GetDX(), c.GetDH(), c.GetDL()) +
		fmt.Sprintf("Queue: %v\n", c.q)
}
