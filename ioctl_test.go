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

import (
	"testing"
)

func TestN(t *testing.T) {
	ns := []N{
		{Type: 0xae, Nr: 0x00},
		{Type: 0xae, Nr: 0x01},
	}
	for _, n := range ns {
		goNumber := n.Number()
		cgoNumber := cgoIO(n.Type, n.Nr)
		if goNumber != cgoNumber {
			t.Errorf("Go number %#08x != cgo number %#08x", goNumber, cgoNumber)
		}
	}
}

func TestR(t *testing.T) {
	rs := []R{
		{Type: 0xae, Nr: 0x00, Size: 8},
		{Type: 0xae, Nr: 0x01, Size: 16},
		{Type: 0xae, Nr: 0x02, Size: 128},
	}
	for _, r := range rs {
		goNumber := r.Number()
		cgoNumber := cgoIOR(r.Type, r.Nr, r.Size)
		if goNumber != cgoNumber {
			t.Errorf("Go number %#08x != cgo number %#08x", goNumber, cgoNumber)
		}
	}
}

func TestW(t *testing.T) {
	ws := []W{
		{Type: 0xae, Nr: 0x00, Size: 8},
		{Type: 0xae, Nr: 0x01, Size: 16},
		{Type: 0xae, Nr: 0x02, Size: 128},
	}
	for _, w := range ws {
		goNumber := w.Number()
		cgoNumber := cgoIOW(w.Type, w.Nr, w.Size)
		if goNumber != cgoNumber {
			t.Errorf("Go number %#08x != cgo number %#08x", goNumber, cgoNumber)
		}
	}
}

func TestWR(t *testing.T) {
	wrs := []WR{
		{Type: 0xae, Nr: 0x00, Size: 8},
		{Type: 0xae, Nr: 0x01, Size: 16},
		{Type: 0xae, Nr: 0x02, Size: 128},
	}
	for _, wr := range wrs {
		goNumber := wr.Number()
		cgoNumber := cgoIOWR(wr.Type, wr.Nr, wr.Size)
		if goNumber != cgoNumber {
			t.Errorf("Go number %#08x != cgo number %#08x", goNumber, cgoNumber)
		}
	}
}
