package helmetstyles

import "testing"

func TestHelmetStyles(t *testing.T) {
	if len(Styles()) != 33 {
		t.Fatalf("expected 33 styles, got %d", len(Styles()))
	}
	s := FindStyle("Retro Full Face")
	if s == nil || s.ID != "retro-full-face" || !s.HasChinBar {
		t.Fatalf("unexpected style: %+v", s)
	}
	if len(ByCategory("Motocross")) != 7 {
		t.Fatal("expected 7 motocross styles")
	}
	if len(PaintableZones("jet-long-shield")) != 5 {
		t.Fatal("expected 5 zones for open-face jet helmet")
	}
}
