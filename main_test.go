package main

import (
	"testing"

	"github.com/apiarian/vscode-go-test-docs-gopathonly"
	"github.com/apiarian/vscode-go-test-docs-vendoronly"
	"github.com/apiarian/vscode-go-test-docs-vendoroverride"
)

func TestFoosBarsBaz(t *testing.T) {
	if got, want := gopathonly.Foo(), "foo"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	if got, want := gopathonly.Bar(), "bar"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// note the flipped wants
	if got, want := vendoroverride.Foo(), "bar"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	if got, want := vendoroverride.Bar(), "foo"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	if got, want := vendoronly.Baz(), "baz"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
