package opticmp

import "testing"

func TestPtr(t *testing.T) {
	p := Ptr(42)
	if p == nil {
		t.Fatal("Ptr returned nil")
	}
	if *p != 42 {
		t.Fatalf("*Ptr(42) = %d, want 42", *p)
	}

	s := Ptr("hi")
	if s == nil || *s != "hi" {
		t.Fatalf(`Ptr("hi") = %v, want "hi"`, s)
	}
}
