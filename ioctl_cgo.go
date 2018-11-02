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
