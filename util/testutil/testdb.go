//go:build !(windows && 386)

package testutil

var SupportedTestDialects = []string{"sqlite", "mysql", "postgres"}
