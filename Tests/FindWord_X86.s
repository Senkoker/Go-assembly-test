#include "textflag.h"
//WORDCOUNT(slice []rune) int32
TEXT ·WordCount(SB),$64-28
    MOVQ data+0(FP), AX
    MOVQ length+8(FP), BX
    MOVL $0, DX
loop:
    CMPQ BX, $0
    JE done
    SUBQ $1, BX
    MOVQ AX, 32(SP)
    MOVQ BX, 40(SP)
    MOVL DX, 48(SP)
    MOVL (AX), DI
    MOVL DI, 0(SP)
    CALL ·IsSpace(SB)
    MOVQ 32(SP), AX
    MOVQ 40(SP), BX
    MOVL 48(SP), DX
    MOVB 8(SP), CX
    ADDQ $4, AX
    CMPL CX, $1
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
