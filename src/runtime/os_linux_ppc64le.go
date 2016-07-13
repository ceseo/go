// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	_AT_HWCAP2	  = 26

	// Golang currently requires POWER8 as a minimum for ppc64le. Therefore,
	// there is no need to check for any capabilities that were introduced
	// before ISA 2.07 (Altivec/VMX, VSX, DFP, etc). Hence, only HWCAP2 is
	// required and we start checking from ISA 3.0 and beyond.
	_PPC_FEATURE2_ARCH_3_00	= 0x00800000  // ISA 3.0 (POWER9)
)

var hwcap2 uint32

// Function for runtime capability checks.
// supports = feature mask (see __PPC_FEATURE2 constants above)
func checkppc64lesupports(supports uint32) bool {
	if (hwcap2 & supports) == 0 {
		return false
	} else {
		return true
	}
}


func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_HWCAP2:
		hwcap2 = uint32(val)
	}
}
