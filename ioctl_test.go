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

import (
	"testing"
)

func TestN(t *testing.T) {
	ns := []N{
		{Type: 0xae, Nr: 0x00},
		{Type: 0xae, Nr: 0x01},
	}
	for _, n := range ns {
		cmp(t, n.Number(), cgoIO(n.Type, n.Nr))
	}
}

func TestR(t *testing.T) {
	rs := []R{
		{Type: 0xae, Nr: 0x00, Size: 8},
		{Type: 0xae, Nr: 0x01, Size: 16},
		{Type: 0xae, Nr: 0x02, Size: 128},
	}
	for _, r := range rs {
		cmp(t, r.Number(), cgoIOR(r.Type, r.Nr, r.Size))
	}
}

func TestW(t *testing.T) {
	ws := []W{
		{Type: 0xae, Nr: 0x00, Size: 8},
		{Type: 0xae, Nr: 0x01, Size: 16},
		{Type: 0xae, Nr: 0x02, Size: 128},
	}
	for _, w := range ws {
		cmp(t, w.Number(), cgoIOW(w.Type, w.Nr, w.Size))
	}
}

func TestWR(t *testing.T) {
	wrs := []WR{
		{Type: 0xae, Nr: 0x00, Size: 8},
		{Type: 0xae, Nr: 0x01, Size: 16},
		{Type: 0xae, Nr: 0x02, Size: 128},
	}
	for _, wr := range wrs {
		cmp(t, wr.Number(), cgoIOWR(wr.Type, wr.Nr, wr.Size))
	}
}

func cmp(t *testing.T, got, want uint32) {
	t.Helper()
	if got != want {
		t.Errorf("Go number %#08x != cgo number %#08x", got, want)
	}
}
