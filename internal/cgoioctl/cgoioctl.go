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

// Package cgoioctl contains wrapper functions around the _IOC C macro.
package cgoioctl // import "acln.ro/ioctl/internal/cgoioctl"

// #include <linux/types.h>
// #include <linux/ioctl.h>
//
// unsigned long io(__u16 type, __u16 nr) {
//         return _IOC(_IOC_NONE, type, nr, 0);
// }
//
// unsigned long ior(__u16 type, __u16 nr, __u16 size) {
//         return _IOC(_IOC_READ, type, nr, size);
// }
//
// unsigned long iow(__u16 type, __u16 nr, __u16 size) {
//         return _IOC(_IOC_WRITE, type, nr, size);
// }
//
// unsigned long iowr(__u16 type, __u16 nr, __u16 size) {
//         return _IOC(_IOC_READ | _IOC_WRITE, type, nr, size);
// }
import "C"

// IO wraps the _IO C macro.
func IO(typ, nr uint16) uint32 {
	return uint32(C.io(C.__u16(typ), C.__u16(nr)))
}

// IOR wraps the _IOR C macro.
func IOR(typ, nr, size uint16) uint32 {
	return uint32(C.ior(C.__u16(typ), C.__u16(nr), C.__u16(size)))
}

// IOW wraps the _IOW C macro.
func IOW(typ, nr, size uint16) uint32 {
	return uint32(C.iow(C.__u16(typ), C.__u16(nr), C.__u16(size)))
}

// IOWR wraps the _IOWR C macro.
func IOWR(typ, nr, size uint16) uint32 {
	return uint32(C.iowr(C.__u16(typ), C.__u16(nr), C.__u16(size)))
}
