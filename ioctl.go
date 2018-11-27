// Copyright 2018 Andrei Tudor Călin
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// Package ioctl provides facilities to define ioctl numbers and perform
// ioctls against file descriptors.
//
// The Name field on all ioctl types is optional. If a name is specified,
// the name will be included in error messages returned by method calls.
//
// In case of errors, all method calls return errors of type Error.
package ioctl // import "acln.ro/ioctl"

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	nrBits   = 8
	typeBits = 8
)

const (
	nrShift   = 0
	typeShift = nrShift + nrBits
	sizeShift = typeShift + typeBits
	dirShift  = sizeShift + sizeBits
)

// N specifies an ioctl that does not exchange data with the kernel through
// a pointer. An N can nevertheless have an integer argument.
type N struct {
	Name string
	Type uint16
	Nr   uint16
}

// Number returns the ioctl number associated with n.
func (n N) Number() uint32 {
	return number(dirNone, n.Type, n.Nr, 0)
}

// Exec executes n against fd, with no argument.
func (n N) Exec(fd int) (int, error) {
	return n.ExecInt(fd, 0)
}

// ExecInt executes n against fd with the integer argument val.
func (n N) ExecInt(fd int, val uintptr) (int, error) {
	res, err := ioctlInt(fd, n.Number(), val)
	return res, wrapError(err, n.Name, n.Number())
}

// R specifies a read ioctl: information is passed from kernel to userspace.
type R struct {
	Name string
	Type uint16
	Nr   uint16
	Size uint16
}

// Number returns the ioctl number associated with r.
func (r R) Number() uint32 {
	return number(dirRead, r.Type, r.Nr, r.Size)
}

// Read executes r against fd with the pointer argument ptr.
//
// The size of the object ptr is pointing to must match r.Size.
func (r R) Read(fd int, ptr unsafe.Pointer) error {
	_, err := ioctlPointer(fd, r.Number(), ptr)
	return wrapError(err, r.Name, r.Number())
}

// W specifies a write ioctl: information is passed from userspace into
// the kernel.
type W struct {
	Name string
	Type uint16
	Nr   uint16
	Size uint16
}

// Number returns the ioctl number associated with w.
func (w W) Number() uint32 {
	return number(dirWrite, w.Type, w.Nr, w.Size)
}

// Write executes w against fd with the pointer argument ptr.
//
// The size of the object ptr is pointing to must match w.Size.
func (w W) Write(fd int, ptr unsafe.Pointer) error {
	_, err := ioctlPointer(fd, w.Number(), ptr)
	return wrapError(err, w.Name, w.Number())
}

// WR specifies a bidirectional ioctl.
type WR struct {
	Name string
	Type uint16
	Nr   uint16
	Size uint16
}

// Number returns the ioctl number associated with wr.
func (wr WR) Number() uint32 {
	return number(dirRead|dirWrite, wr.Type, wr.Nr, wr.Size)
}

// Exec executes wr against fd. ptr is the input / output argument.
//
// The size of the object ptr is pointing to must match wr.Size.
func (wr WR) Exec(fd int, ptr unsafe.Pointer) error {
	_, err := ioctlPointer(fd, wr.Number(), ptr)
	return wrapError(err, wr.Name, wr.Number())
}

// wrapError wraps err in an Error. If err is nil, wrapError returns nil.
func wrapError(err error, name string, number uint32) error {
	if err == nil {
		return nil
	}
	return &Error{Name: name, Number: number, Err: err}
}

// Error records an error from an ioctl(2) system call.
type Error struct {
	// Name is the name of the ioctl, e.g. "KVM_CREATE_VM".
	//
	// This field may be empty, in which case it is not included in the
	// error message.
	Name string

	// Number is the 32 bit ioctl number. It is rendered in hexadecimal
	// in the error message.
	Number uint32

	// Err is the underlying error, of type syscall.Errno.
	Err error
}

func (e *Error) Error() string {
	if e.Name != "" {
		return fmt.Sprintf("ioctl: %s (%#08x): %v", e.Name, e.Number, e.Err)
	}
	return fmt.Sprintf("ioctl: %#08x: %v", e.Number, e.Err)
}

func number(dir, typ, nr, size uint16) uint32 {
	d := uint32(dir) << dirShift
	t := uint32(typ) << typeShift
	n := uint32(nr) << nrShift
	s := uint32(size) << sizeShift
	return d | t | n | s
}

func ioctlInt(fd int, num uint32, arg uintptr) (int, error) {
	r, _, e := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(num), uintptr(arg))
	if e != 0 {
		return int(r), e
	}
	return int(r), nil
}

func ioctlPointer(fd int, num uint32, arg unsafe.Pointer) (int, error) {
	r, _, e := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(num), uintptr(arg))
	if e != 0 {
		return int(r), e
	}
	return int(r), nil
}
