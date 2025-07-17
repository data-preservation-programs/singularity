// +build darwin

package wallet

// #cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup
import "C"
