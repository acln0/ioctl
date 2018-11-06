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

// +build cgotest

package ioctl // import "acln.ro/ioctl"

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

// These functions are used for testing, and live in this file because
// importing "C" in test files is not allowed. Bypass the more common
// _IO, _IOR, _IOW, and _IOWR macros, because they use sizeof their final
// argument, which is difficult to use from Go. Instead, pass the actual
// size to _IOC directly.

func cgoIO(typ, nr uint16) uint32 {
	return uint32(C.io(C.__u16(typ), C.__u16(nr)))
}

func cgoIOR(typ, nr, size uint16) uint32 {
	return uint32(C.ior(C.__u16(typ), C.__u16(nr), C.__u16(size)))
}

func cgoIOW(typ, nr, size uint16) uint32 {
	return uint32(C.iow(C.__u16(typ), C.__u16(nr), C.__u16(size)))
}

func cgoIOWR(typ, nr, size uint16) uint32 {
	return uint32(C.iowr(C.__u16(typ), C.__u16(nr), C.__u16(size)))
}
