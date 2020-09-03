// Code generated by 'go generate'; DO NOT EDIT.

package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procCredEnumerateW       = modadvapi32.NewProc("CredEnumerateW")
	procCredFree             = modadvapi32.NewProc("CredFree")
	procFileTimeToSystemTime = modkernel32.NewProc("FileTimeToSystemTime")
)

func CredEnumerateW(filter *uint16, flags uint32, count *uint32, credentials ***_CREDENTIALW) (err error) {
	r1, _, e1 := syscall.Syscall6(procCredEnumerateW.Addr(), 4, uintptr(unsafe.Pointer(filter)), uintptr(flags), uintptr(unsafe.Pointer(count)), uintptr(unsafe.Pointer(credentials)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CredFree(buffer unsafe.Pointer) {
	syscall.Syscall(procCredFree.Addr(), 1, uintptr(buffer), 0, 0)
	return
}

func FileTimeToSystemTime(fileTime *_FILETIME, systemTime *_SYSTEMTIME) (err error) {
	r1, _, e1 := syscall.Syscall(procFileTimeToSystemTime.Addr(), 2, uintptr(unsafe.Pointer(fileTime)), uintptr(unsafe.Pointer(systemTime)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
