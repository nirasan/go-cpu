package go_cpu

import "testing"

func TestCpu_String(t *testing.T) {
	c := NewCpu()
	c.SetEAX(400000)
	t.Log(c.String())
}
