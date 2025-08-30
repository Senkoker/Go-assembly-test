#include "textflag.h"

// func FindElemSimd(slice []int32, target int32) bool
TEXT ·FindElemSimd(SB), NOSPLIT, $0
    MOVQ slice_base+0(FP), AX
    MOVQ slice_len+8(FP), BX
    MOVL target+24(FP), CX
    TESTQ BX, BX
    JZ not_found
    MOVD CX, X0
    PSHUFD $0, X0, X0
    CMPQ BX, $4
    JL tail

loop:
    MOVUPS (AX), X1
    PCMPEQL X0, X1
    PMOVMSKB X1, DX
    TESTL DX, DX
    JNZ found
    ADDQ $16, AX
    SUBQ $4, BX
    CMPQ BX, $4
    JGE loop

tail:
    TESTQ BX, BX
    JZ not_found

tail_loop:
    MOVL (AX), DX
    CMPL DX, CX
    JE found
    ADDQ $4, AX
    DECQ BX
    JNZ tail_loop

not_found:
    MOVB $0, ret+32(FP)
    RET

found:
    MOVB $1, ret+32(FP)
    RET
