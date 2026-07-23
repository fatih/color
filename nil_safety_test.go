package color

import (
	"bytes"
	"strings"
	"testing"
)

func TestNilColorSprint(t *testing.T) {
	var c *Color
	if got := c.Sprint("hi"); got != "hi" {
		t.Fatalf("Sprint: got %q", got)
	}
	if got := c.Sprintf("x=%d", 1); got != "x=1" {
		t.Fatalf("Sprintf: got %q", got)
	}
}

func TestNilWriterFprint(t *testing.T) {
	c := New(FgRed)
	c.DisableColor()
	if _, err := c.Fprint(nil, "hi"); err == nil || !strings.Contains(err.Error(), "nil") {
		t.Fatalf("expected nil Writer error, got %v", err)
	}
	var buf bytes.Buffer
	if _, err := c.Fprint(&buf, "hi"); err != nil {
		t.Fatal(err)
	}
	if buf.String() != "hi" {
		t.Fatalf("got %q", buf.String())
	}
}

func TestNilColorFprint(t *testing.T) {
	var c *Color
	var buf bytes.Buffer
	if _, err := c.Fprint(&buf, "ok"); err != nil {
		t.Fatal(err)
	}
	if buf.String() != "ok" {
		t.Fatalf("got %q", buf.String())
	}
}
