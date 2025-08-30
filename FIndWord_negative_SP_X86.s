/*
#include "textflag.h"

TEXT ·WordCountTest(SB),$28-28
    MOVQ data+0(FP), AX
    MOVQ length+8(FP), BX
    ADDQ $32, SP
    MOVL $0, DX
loop:
    CMPQ BX, $0                         //24 bytes
    JE done //                  //0(SP) 4(SP) 8(SP) 16(SP)   // -0(SP) -8(SP) -16(SP) -20(SP)
    SUBQ $1, BX
    MOVQ AX, 0(SP)
    MOVQ BX, 8(SP)
    MOVL DX, 16(SP)
    MOVL (AX), DI
    MOVL DI, -24(SP)
    CALL ·IsSpace(SB)
    MOVQ -0(SP), AX
    MOVQ -8(SP), BX
    MOVL -16(SP), DX
    MOVB -40(SP), CX
    ADDQ $4, AX
    CMPB CX, $1
    JE plusWord
    JNE lastWord
plusWord:
    ADDL $1, DX
    JMP loop
lastWord:
    CMPQ BX, $0
    JE plusWord
    JMP loop
done:
    MOVL DX, ret+24(FP)
    RET
*/