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

package ioctl

import (
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
	Type uint16
	Nr   uint16
	Size uint16
}

func (n N) marshal() uint32 {
	return marshal(dirNone, n.Type, n.Nr, n.Size)
}

// Exec executes n against fd.
func (n N) Exec(fd int) (int, error) {
	return ioctlInt(fd, n.marshal(), 0)
}

// R specifies a read ioctl: userland is writing, and the kernel is reading.
type R struct {
	Type uint16
	Nr   uint16
	Size uint16
}

func (r R) marshal() uint32 {
	return marshal(dirRead, r.Type, r.Nr, r.Size)
}

// SetInt executes r against fd with the integer argument val.
func (r R) SetInt(fd int, val uintptr) error {
	_, err := ioctlInt(fd, r.marshal(), val)
	return err
}

// SetPointer executes r against fd with the pointer argument ptr.
func (r R) SetPointer(fd int, ptr unsafe.Pointer) error {
	_, err := ioctlPointer(fd, r.marshal(), ptr)
	return err
}

// W specifies a write ioctl: userland is reading, and the kernel is writing.
type W struct {
	Type uint16
	Nr   uint16
	Size uint16
}

func (w W) marshal() uint32 {
	return marshal(dirWrite, w.Type, w.Nr, w.Size)
}

// GetInt executes w against fd and returns the integer result.
func (w W) GetInt(fd int) (int, error) {
	var res int
	_, err := ioctlPointer(fd, w.marshal(), unsafe.Pointer(&res))
	return res, err
}

// GetPointer executes w against fd and stores the result in ptr.
func (w W) GetPointer(fd int, ptr unsafe.Pointer) error {
	_, err := ioctlPointer(fd, w.marshal(), ptr)
	return err
}

// WR specifies a bidirectional ioctl.
type WR struct {
	Type uint16
	Nr   uint16
	Size uint16
}

func (wr WR) marshal() uint32 {
	return marshal(dirWriteRead, wr.Type, wr.Nr, wr.Size)
}

// Exec executes wr against fd. ptr is the read / write argument.
func (wr WR) Exec(fd int, ptr unsafe.Pointer) error {
	_, err := ioctlPointer(fd, wr.marshal(), ptr)
	return err
}

func marshal(dir, typ, nr, size uint16) uint32 {
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
