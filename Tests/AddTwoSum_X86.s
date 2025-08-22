#include "textflag.h"

//func SumInt(a,b int) itn
TEXT ·SumInt(SB),NOSPLIT,$0
    MOVQ first_int+0(FP), AX
    MOVQ second_int+8(FP), BX
    ADDQ AX, BX
    MOVQ BX, result+16(FP)
    RET
