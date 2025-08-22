#include "textflag.h"

TEXT ·SumSliceNotEquel(SB),NOSPLIT,$0
    MOVQ data+0(FP), AX
    MOVQ lenght+8(FP), BX
    MOVL $0, CX
condition:
    CMPQ BX, $0
    JNE loop
    MOVL CX,ret+24(FP)
    RET
loop:
    ADDL (AX),CX
    ADDQ $4, AX
    SUBQ $1, BX
    JMP condition

