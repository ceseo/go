// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	_AT_HWCAP	  = 16
	_AT_HWCAP2	  = 26

	// Golang currently requires POWER5 as a minimum for ppc64, so we need
	// to check for ISA 2.03 and beyond.
	// Since we support only POWER server processors, we will assume that
	// each processor generation implements the full ISA — i.e. POWER7
	// implies VSX support, POWER8 implies HTM support, etc. The only
	// exception is the POWER6x (ISA 2.05 + mffgpr/mftgpr extension).
	_PPC_FEATURE_POWER5_PLUS  = 0x00020000  // ISA 2.03 (POWER5+)
	_PPC_FEATURE_ARCH_2_05	  = 0x00001000  // ISA 2.05 (POWER6)
	_PPC_FEATURE_POWER6_EXT	  = 0x00000200  // mffgpr/mftgpr extension (POWER6x)
	_PPC_FEATURE_ARCH_2_06	  = 0x00000100  // ISA 2.06 (POWER7).
	_PPC_FEATURE2_ARCH_2_07	  = 0x80000000  // ISA 2.07 (POWER8)
	_PPC_FEATURE2_ARCH_3_00	  = 0x00800000	// ISA 3.0 (POWER9)
)

var hwcap uint32
var hwcap2 uint32

// Function for runtime capability checks.
// mask = hwcap or hwcap2
// supports = feature mask (see __PPC_FEATURE* constants above)
func checkppc64supports(mask uint32, supports uint32) bool {
	if (mask & supports) == 0 {
		return false
	} else {
		return true
	}
}

func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_HWCAP:
		hwcap = uint32(val)
	case _AT_HWCAP2:
		hwcap2 = uint32(val)
	}
}
