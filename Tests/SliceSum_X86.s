#include "textflag.h"

//func SumSlice(slice []int32) int
TEXT ·SumSlice(SB),NOSPLIT,$0
    MOVQ data_slice+0(FP), AX
    MOVQ length_slice+8(FP), BX
    MOVL $0,DX
loop:
    CMPQ BX, $0
    JE done
    MOVLQSX (AX), CX
    ADDL CX,DX
    ADDQ $4, AX
    SUBQ $1, BX
    JMP loop
done:
    MOVL DX, result+24(FP)
    RET
