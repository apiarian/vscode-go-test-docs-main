package main

import (
	"fmt"

	"github.com/apiarian/vscode-go-test-docs-gopathonly"
	"github.com/apiarian/vscode-go-test-docs-vendoronly"
	"github.com/apiarian/vscode-go-test-docs-vendoroverride"
)

func main() {
	fmt.Println("gopathonly Foo:", gopathonly.Foo()) // gopathonly at byte 248, Foo at byte 258
	fmt.Println("gopathonly Bar:", gopathonly.Bar()) // Bar at byte 351

	fmt.Println("vendoroverride Foo:", vendoroverride.Foo()) //vendoroverride at byte 416, Foo at byte 430
	fmt.Println("vendoroverride Bar:", vendoroverride.Bar()) // Bar at byte 533

	fmt.Println("vendoronly Baz:", vendoronly.Baz()) // vendoronly at byte 592, Baz at byte 603
}
