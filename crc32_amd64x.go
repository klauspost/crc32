// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 amd64p32

package crc32

// This file contains the code to call the SSE 4.2 version of the Castagnoli
// CRC.

// haveSSE41/haveSSE42 is defined in crc_amd64.s and uses CPUID to test
// for SSE 4.1 and 4.2 support.
func haveSSE41() bool
func haveSSE42() bool
func haveCLMUL() bool

// castagnoliSSE42 is defined in crc_amd64.s and uses the SSE4.2 CRC32
// instruction.
func castagnoliSSE42(crc uint32, p []byte) uint32

// ieeeCLMUL is defined in crc_amd64.s and uses the PCLMULQDQ
// instruction.
func ieeeCLMUL(crc uint32, p []byte) uint32

var sse41 = haveSSE41()
var sse42 = haveSSE42()
var clmul = haveCLMUL()

func updateCastagnoli(crc uint32, p []byte) uint32 {
	if sse42 {
		return castagnoliSSE42(crc, p)
	}
	return update(crc, castagnoliTable, p)
}

func updateIEEE(crc uint32, p []byte) uint32 {
	if clmul && sse41 && len(p) >= 64 {
		left := len(p) & 15
		do := len(p) - left
		crc := ^ieeeCLMUL(^crc, p[:do])
		if left > 0 {
			crc = update(crc, IEEETable, p[do:])
		}
		return crc
	}
	return update(crc, IEEETable, p)
}
