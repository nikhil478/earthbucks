package earthbucks

// OpcodeName represents all possible opcode names.
type OpcodeName string

const (
    Opcode0                 OpcodeName = "0"
    OpcodePUSHDATA1         OpcodeName = "PUSHDATA1"
    OpcodePUSHDATA2         OpcodeName = "PUSHDATA2"
    OpcodePUSHDATA4         OpcodeName = "PUSHDATA4"
    Opcode1NEGATE           OpcodeName = "1NEGATE"
    Opcode1                 OpcodeName = "1"
    Opcode2                 OpcodeName = "2"
    Opcode3                 OpcodeName = "3"
    Opcode4                 OpcodeName = "4"
    Opcode5                 OpcodeName = "5"
    Opcode6                 OpcodeName = "6"
    Opcode7                 OpcodeName = "7"
    Opcode8                 OpcodeName = "8"
    Opcode9                 OpcodeName = "9"
    Opcode10                OpcodeName = "10"
    Opcode11                OpcodeName = "11"
    Opcode12                OpcodeName = "12"
    Opcode13                OpcodeName = "13"
    Opcode14                OpcodeName = "14"
    Opcode15                OpcodeName = "15"
    Opcode16                OpcodeName = "16"
    OpcodeIF                OpcodeName = "IF"
    OpcodeNOTIF             OpcodeName = "NOTIF"
    OpcodeELSE              OpcodeName = "ELSE"
    OpcodeENDIF             OpcodeName = "ENDIF"
    OpcodeVERIFY            OpcodeName = "VERIFY"
    OpcodeRETURN            OpcodeName = "RETURN"
    OpcodeTOALTSTACK        OpcodeName = "TOALTSTACK"
    OpcodeFROMALTSTACK      OpcodeName = "FROMALTSTACK"
    Opcode2DROP             OpcodeName = "2DROP"
    Opcode2DUP              OpcodeName = "2DUP"
    Opcode3DUP              OpcodeName = "3DUP"
    Opcode2OVER             OpcodeName = "2OVER"
    Opcode2ROT              OpcodeName = "2ROT"
    Opcode2SWAP             OpcodeName = "2SWAP"
    OpcodeIFDUP             OpcodeName = "IFDUP"
    OpcodeDEPTH             OpcodeName = "DEPTH"
    OpcodeDROP              OpcodeName = "DROP"
    OpcodeDUP               OpcodeName = "DUP"
    OpcodeNIP               OpcodeName = "NIP"
    OpcodeOVER              OpcodeName = "OVER"
    OpcodePICK              OpcodeName = "PICK"
    OpcodeROLL              OpcodeName = "ROLL"
    OpcodeROT               OpcodeName = "ROT"
    OpcodeSWAP              OpcodeName = "SWAP"
    OpcodeTUCK              OpcodeName = "TUCK"
    OpcodeCAT               OpcodeName = "CAT"
    OpcodeSUBSTR            OpcodeName = "SUBSTR"
    OpcodeLEFT              OpcodeName = "LEFT"
    OpcodeRIGHT             OpcodeName = "RIGHT"
    OpcodeSIZE              OpcodeName = "SIZE"
    OpcodeINVERT            OpcodeName = "INVERT"
    OpcodeAND               OpcodeName = "AND"
    OpcodeOR                OpcodeName = "OR"
    OpcodeXOR               OpcodeName = "XOR"
    OpcodeEQUAL             OpcodeName = "EQUAL"
    OpcodeEQUALVERIFY       OpcodeName = "EQUALVERIFY"
    Opcode1ADD              OpcodeName = "1ADD"
    Opcode1SUB              OpcodeName = "1SUB"
    Opcode2MUL              OpcodeName = "2MUL"
    Opcode2DIV              OpcodeName = "2DIV"
    OpcodeNEGATE            OpcodeName = "NEGATE"
    OpcodeABS               OpcodeName = "ABS"
    OpcodeNOT               OpcodeName = "NOT"
    Opcode0NOTEQUAL         OpcodeName = "0NOTEQUAL"
    OpcodeADD               OpcodeName = "ADD"
    OpcodeSUB               OpcodeName = "SUB"
    OpcodeMUL               OpcodeName = "MUL"
    OpcodeDIV               OpcodeName = "DIV"
    OpcodeMOD               OpcodeName = "MOD"
    OpcodeLSHIFT            OpcodeName = "LSHIFT"
    OpcodeRSHIFT            OpcodeName = "RSHIFT"
    OpcodeBOOLAND           OpcodeName = "BOOLAND"
    OpcodeBOOLOR           OpcodeName = "BOOLOR"
    OpcodeNUMEQUAL          OpcodeName = "NUMEQUAL"
    OpcodeNUMEQUALVERIFY    OpcodeName = "NUMEQUALVERIFY"
    OpcodeNUMNOTEQUAL       OpcodeName = "NUMNOTEQUAL"
    OpcodeLESSTHAN          OpcodeName = "LESSTHAN"
    OpcodeGREATERTHAN       OpcodeName = "GREATERTHAN"
    OpcodeLESSTHANOREQUAL   OpcodeName = "LESSTHANOREQUAL"
    OpcodeGREATERTHANOREQUAL OpcodeName = "GREATERTHANOREQUAL"
    OpcodeMIN               OpcodeName = "MIN"
    OpcodeMAX               OpcodeName = "MAX"
    OpcodeWITHIN            OpcodeName = "WITHIN"
    OpcodeBLAKE3            OpcodeName = "BLAKE3"
    OpcodeDOUBLEBLAKE3     OpcodeName = "DOUBLEBLAKE3"
    OpcodeCHECKSIG          OpcodeName = "CHECKSIG"
    OpcodeCHECKSIGVERIFY    OpcodeName = "CHECKSIGVERIFY"
    OpcodeCHECKMULTISIG     OpcodeName = "CHECKMULTISIG"
    OpcodeCHECKMULTISIGVERIFY OpcodeName = "CHECKMULTISIGVERIFY"
    OpcodeCHECKLOCKABSVERIFY OpcodeName = "CHECKLOCKABSVERIFY"
    OpcodeCHECKLOCKRELVERIFY OpcodeName = "CHECKLOCKRELVERIFY"
)

// Op represents opcode to its numeric value.
var Op = map[OpcodeName]uint8{
    Opcode0:                 0x00,
    OpcodePUSHDATA1:         0x4c,
    OpcodePUSHDATA2:         0x4d,
    OpcodePUSHDATA4:         0x4e,
    Opcode1NEGATE:           0x4f,
    Opcode1:                 0x51,
    Opcode2:                 0x52,
    Opcode3:                 0x53,
    Opcode4:                 0x54,
    Opcode5:                 0x55,
    Opcode6:                 0x56,
    Opcode7:                 0x57,
    Opcode8:                 0x58,
    Opcode9:                 0x59,
    Opcode10:                0x5a,
    Opcode11:                0x5b,
    Opcode12:                0x5c,
    Opcode13:                0x5d,
    Opcode14:                0x5e,
    Opcode15:                0x5f,
    Opcode16:                0x60,
    OpcodeIF:                0x63,
    OpcodeNOTIF:             0x64,
    OpcodeELSE:              0x67,
    OpcodeENDIF:             0x68,
    OpcodeVERIFY:            0x69,
    OpcodeRETURN:            0x6a,
    OpcodeTOALTSTACK:        0x6b,
    OpcodeFROMALTSTACK:      0x6c,
    Opcode2DROP:             0x6d,
    Opcode2DUP:              0x6e,
    Opcode3DUP:              0x6f,
    Opcode2OVER:             0x70,
    Opcode2ROT:              0x71,
    Opcode2SWAP:             0x72,
    OpcodeIFDUP:             0x73,
    OpcodeDEPTH:             0x74,
    OpcodeDROP:              0x75,
    OpcodeDUP:               0x76,
    OpcodeNIP:               0x77,
    OpcodeOVER:              0x78,
    OpcodePICK:              0x79,
    OpcodeROLL:              0x7a,
    OpcodeROT:               0x7b,
    OpcodeSWAP:              0x7c,
    OpcodeTUCK:              0x7d,
    OpcodeCAT:               0x7e,
    OpcodeSUBSTR:            0x7f,
    OpcodeLEFT:              0x80,
    OpcodeRIGHT:             0x81,
    OpcodeSIZE:              0x82,
    OpcodeINVERT:            0x83,
    OpcodeAND:               0x84,
    OpcodeOR:                0x85,
    OpcodeXOR:               0x86,
    OpcodeEQUAL:             0x87,
    OpcodeEQUALVERIFY:       0x88,
    Opcode1ADD:              0x8b,
    Opcode1SUB:              0x8c,
    Opcode2MUL:              0x8d,
    Opcode2DIV:              0x8e,
    OpcodeNEGATE:            0x8f,
    OpcodeABS:               0x90,
    OpcodeNOT:               0x91,
    Opcode0NOTEQUAL:         0x92,
    OpcodeADD:               0x93,
    OpcodeSUB:               0x94,
    OpcodeMUL:               0x95,
    OpcodeDIV:               0x96,
    OpcodeMOD:               0x97,
    OpcodeLSHIFT:            0x98,
    OpcodeRSHIFT:            0x99,
    OpcodeBOOLAND:           0x9a,
    OpcodeBOOLOR:            0x9b,
    OpcodeNUMEQUAL:          0x9c,
    OpcodeNUMEQUALVERIFY:    0x9d,
    OpcodeNUMNOTEQUAL:       0x9e,
    OpcodeLESSTHAN:          0x9f,
    OpcodeGREATERTHAN:       0xa0,
    OpcodeLESSTHANOREQUAL:   0xa1,
    OpcodeGREATERTHANOREQUAL: 0xa2,
    OpcodeMIN:               0xa3,
    OpcodeMAX:               0xa4,
    OpcodeWITHIN:            0xa5,
    OpcodeBLAKE3:            0xa6,
    OpcodeDOUBLEBLAKE3:      0xa7,
    OpcodeCHECKSIG:          0xac,
    OpcodeCHECKSIGVERIFY:    0xad,
    OpcodeCHECKMULTISIG:     0xae,
    OpcodeCHECKMULTISIGVERIFY: 0xaf,
    OpcodeCHECKLOCKABSVERIFY: 0xb0,
    OpcodeCHECKLOCKRELVERIFY: 0xb1,
}