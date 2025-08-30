#include "textflag.h"

// func FindElemSimd(slice []int32, target int32) bool
TEXT ·FindElemSimdYMM(SB), NOSPLIT, $0
	MOVQ	slice_base+0(FP), AX    // pointer to data
	MOVQ	slice_len+8(FP), BX     // length (elements)
	MOVL	target+24(FP), CX       // target value (32-bit)
	MOVD	CX, X0
	VPBROADCASTD	X0, Y0
	CMPQ	BX, $8
	JL	tail

loop:
	VMOVUPS	(AX), Y1
	VPCMPEQD	Y1, Y0, Y2
	VPMOVMSKB	Y2, DX
	TESTL	DX, DX
	JNZ	found
	ADDQ	$32, AX
	SUBQ	$8, BX
	CMPQ	BX, $8
	JGE	loop

tail:

	TESTQ	BX, BX
	JZ	not_found

tail_loop:
	MOVL	(AX), DX
	CMPL	DX, CX
	JE	found
	ADDQ	$4, AX
	DECQ	BX
	JNZ	tail_loop

not_found:
	MOVB	$0, ret+32(FP)
	RET

found:
	MOVB	$1, ret+32(FP)
	RET
	