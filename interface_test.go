package err

import "testing"

func demo(err IErr) IErr {
	return err
}

func TestDemo(t *testing.T) {
	err := NewErr("demo", 1)
	demo(err)
}