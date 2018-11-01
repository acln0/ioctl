// Copyright 2018 Andrei Tudor Călin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	nrMask   = (1 << nrBits) - 1
	typeMask = (1 << typeBits) - 1
	sizeMask = (1 << sizeBits) - 1
	dirMask  = (1 << dirBits) - 1
)

const (
	nrShift   = 0
	typeShift = nrShift + nrBits
	sizeShift = typeShift + typeBits
	dirShift  = sizeShift + sizeBits
)

// N specifies an ioctl that does not exchange data with the kernel.
type N struct {
	Name string
	Type uint16
	Nr   uint16
}

// Number returns the associated ioctl number.
func (n N) Number() uint32 {
	return number(dirNone, n.Type, n.Nr, 0)
}

// Exec executes n against fd.
func (n N) Exec(fd int) (int, error) {
	res, err := ioctlInt(fd, n.Number(), 0)
	return res, wrapError(err, n.Name, n.Number())
}

// R specifies a read ioctl: userland is writing, and the kernel is reading.
type R struct {
	Name string
	Type uint16
	Nr   uint16
	Size uint16
}

// Number returns the associated ioctl number.
func (r R) Number() uint32 {
	return number(dirRead, r.Type, r.Nr, r.Size)
}

// WriteInt executes r against fd with the integer argument val.
func (r R) WriteInt(fd int, val uintptr) error {
	_, err := ioctlInt(fd, r.Number(), val)
	return wrapError(err, r.Name, r.Number())
}

// WritePointer executes r against fd with the pointer argument ptr.
func (r R) WritePointer(fd int, ptr unsafe.Pointer) error {
	_, err := ioctlPointer(fd, r.Number(), ptr)
	return wrapError(err, r.Name, r.Number())
}

// W specifies a write ioctl: userland is reading, and the kernel is writing.
type W struct {
	Name string
	Type uint16
	Nr   uint16
	Size uint16
}

// Number returns the associated ioctl number.
func (w W) Number() uint32 {
	return number(dirWrite, w.Type, w.Nr, w.Size)
}

// ReadInt executes w against fd and returns the integer result.
func (w W) ReadInt(fd int) (int, error) {
	var res int
	_, err := ioctlPointer(fd, w.Number(), unsafe.Pointer(&res))
	return res, wrapError(err, w.Name, w.Number())
}

// ReadPointer executes w against fd and stores the result in ptr.
func (w W) ReadPointer(fd int, ptr unsafe.Pointer) error {
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

// Number returns the associated ioctl number.
func (wr WR) Number() uint32 {
	return number(dirWriteRead, wr.Type, wr.Nr, wr.Size)
}

// Exec executes wr against fd. ptr is the input / output argument.
func (wr WR) Exec(fd int, ptr unsafe.Pointer) error {
	_, err := ioctlPointer(fd, wr.Number(), ptr)
	return wrapError(err, wr.Name, wr.Number())
}

// WrapError wraps err in an Error. If err is nil, wrapError returns nil.
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
