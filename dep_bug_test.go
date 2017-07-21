// Package main provides ...
package main

import (
	"testing"

	"github.com/onsi/gomega/format"
)

func TestSomething(t *testing.T) {
	format.MessageWithDiff("hello", "to equal", "goodbye")
}
