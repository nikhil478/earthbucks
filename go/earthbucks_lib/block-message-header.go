package earthbucks

import "encoding/hex"

type BlockMessageHeader struct {
	version                   *U8
	prevBlockMessageHeaderId  *FixedBuf
	messageId                 *FixedBuf
	messageNum                *U64
}

func NewBlockMessageHeader(version *U8, prevBlockMessageHeaderId *FixedBuf, messageId *FixedBuf, messageNum *U64) BlockMessageHeader {
	return BlockMessageHeader{
		version:                  version,
		prevBlockMessageHeaderId: prevBlockMessageHeaderId,
		messageId:                messageId,
		messageNum:               messageNum,
	}
}

func GetMessageHash(message string) *FixedBuf {
	messageBuf := []byte(message)
	messageHash , _ := Blake3Hash(messageBuf)
	return messageHash
}

func GetMessageId(message string) *FixedBuf {
	messageBuf := []byte(message)
	messageHash,_ := Blake3Hash(messageBuf)
	messageId,_ := Blake3Hash(messageHash.buf)
	return messageId
}

func FromMessage(prevBlockMessageHeaderId *FixedBuf, message string, messageNum *U64) BlockMessageHeader {
	messageId := GetMessageId(message)
	return NewBlockMessageHeader(&U8{value: 0}, prevBlockMessageHeaderId, messageId, messageNum)
}

func (bmh *BlockMessageHeader) ToBufWriter(bw *BufWriter) *BufWriter {
	bw.WriteU8BE(bmh.version)
	bw.Write(bmh.prevBlockMessageHeaderId.buf)
	bw.Write(bmh.messageId.buf)
	bw.WriteU64BE(bmh.messageNum)
	return bw
}

func FromBufReader(br *BufReader) (BlockMessageHeader, error) {
	version,_ := br.ReadU8()
	prevBlockMessageHeaderId,_ := br.ReadFixed(32)
	messageId,_ := br.ReadFixed(32)
	messageNum,_ := br.ReadU64BE()
	return NewBlockMessageHeader(version, prevBlockMessageHeaderId, messageId, messageNum), nil
}

func (bmh *BlockMessageHeader) ToBuf() []byte {
	return bmh.ToBufWriter(NewBufWriter()).ToBuf()
}

func FromBuf(buf []byte) (BlockMessageHeader, error) {
	return FromBufReader(NewBufReader(buf))
}

func (bmh *BlockMessageHeader) ToHex() string {
	return hex.EncodeToString(bmh.ToBuf())
}

func FromHex(hexStr string) (BlockMessageHeader, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return BlockMessageHeader{}, err
	}
	return FromBuf(buf)
}

func (bmh *BlockMessageHeader) Hash() ( *FixedBuf, error) {
	return Blake3Hash(bmh.ToBuf())
}

func (bmh *BlockMessageHeader) Id() ( *FixedBuf, error) {
	hash, _ := bmh.Hash()
	return Blake3Hash(hash.buf)
}

func (bmh *BlockMessageHeader) Verify(prevId FixedBuf, prevNum *U64, message string) bool {
	expectedVersion := U8{value: 0}
	if bmh.version.value != expectedVersion.value {
		return false
	}
	// todo fix this @nikhil478
	// expectedMessageId := GetMessageId(message).buf
	// fxBuf := FixedBuf{buf: expectedMessageId}
	// if !fxBuf.Equals(bmh.messageId) {
	// 	return false
	// }
	
	if len(prevId.buf) == 0 && prevNum != nil {
		return false
	}
	if len(prevId.buf) != 0 && prevNum == nil {
		return false
	}
	if len(prevId.buf) != 0 && prevNum != nil {
		if !prevId.Equals(bmh.prevBlockMessageHeaderId) {
			return false
		}
		expectedPrevNum := prevNum.value + 1
		if bmh.messageNum.value != expectedPrevNum {
			return false
		}
	}
	return true
}