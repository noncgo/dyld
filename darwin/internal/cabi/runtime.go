//go:build darwin && amd64
// +build darwin,amd64

package cabi

import (
	"unsafe"
)

//go:linkname runtime_libcCall runtime.libcCall
//go:linkname runtime_entersyscall runtime.entersyscall
//go:linkname runtime_exitsyscall runtime.exitsyscall

func runtime_libcCall(fn, arg unsafe.Pointer) int32 // from runtime/sys_libc.go
func runtime_entersyscall()                         // from runtime/proc.go
func runtime_exitsyscall()                          // from runtime/proc.go
