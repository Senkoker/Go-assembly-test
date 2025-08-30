#include "textflag.h"

// func WordCount(slice []rune) int32
TEXT ·WordCount(SB), NOSPLIT, $0
    MOVQ slice_base+0(FP), AX
    MOVQ slice_len+8(FP), BX
    MOVL $0, DX
    MOVL $0, SI

    TESTQ BX, BX
    JZ done

loop:
    MOVL (AX), DI
    PUSHQ AX
    PUSHQ BX
    PUSHQ DX
    PUSHQ SI

    MOVL DI, 0(SP)
    CALL ·IsSpace(SB)

    MOVB 0(SP), CL

    POPQ SI
    POPQ DX
    POPQ BX
    POPQ AX

    CMPB CL, $1
    JE spaceFound
    MOVL $1, SI
    JMP nextChar

spaceFound:

    CMPL SI, $1
    JNE nextChar
    ADDL $1, DX
    MOVL $0, SI

nextChar:
    ADDQ $4, AX
    SUBQ $1, BX
    JNZ loop


    CMPL SI, $1
    JNE done
    ADDL $1, DX

done:
    MOVL DX, ret+24(FP)
    RET