// +build gc

// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h" // for NOSPLIT
#include "../../src/runtime/go_tls.h"

TEXT ·GoID(SB),NOSPLIT,$0-4
	get_tls(CX)
	MOVL g(CX), AX
	MOVL AX, goid+0(FP)
	RET
