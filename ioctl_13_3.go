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

// +build ppc ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le sparc sparc64

package ioctl // import "acln.ro/ioctl"

const (
	sizeBits = 13
	dirBits  = 3
)

const (
	dirNone = 1 << iota
	dirRead
	dirWrite
)
