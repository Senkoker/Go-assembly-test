#include "textflag.h"

//func SumSlice(slice []int32) int32
TEXT ·SumSlice(SB),NOSPLIT,$0
    MOVQ data_slice+0(FP), AX
    MOVQ length_slice+8(FP), BX
    MOVQ $0, R10
loop:
    CMPQ BX, $0
    JE done
    MOVLQSX (AX), R9
    ADDQ R9,R10
    ADDQ $4, AX
    SUBQ $1, BX
    JMP loop
done:
    MOVQ R10, result+24(FP)
    RET
