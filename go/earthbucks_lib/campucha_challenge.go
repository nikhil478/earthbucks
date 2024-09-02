package earthbucks

const SIZE = 32

type CompuchaChallenge struct {
    challengeId *FixedBuf
    nonce       *U128
}

func NewCompuchaChallenge(challengeId *FixedBuf, nonce *U128) *CompuchaChallenge {
    return &CompuchaChallenge{
        challengeId: challengeId,
        nonce:       nonce,
    }
}

func CompuchaChallengeFromRandomNonce(challengeId *FixedBuf) (*CompuchaChallenge, error) {
    nonceBuf, err := FixedBufFromRandom(16)
	if err != nil {
		return nil, err
	}
    nonce, err := NewBufReader(nonceBuf.buf).ReadU128BE()
	if err != nil {
		return nil, err
	}
    return NewCompuchaChallenge(challengeId, nonce) , nil
}

func CompuchaChallengeFromBufReader(br BufReader) (*CompuchaChallenge, error) {
	challengeId,err := br.ReadFixed(16)
	if err != nil {
		return nil, err
	}
	nonce,err := NewBufReader(br.buf).ReadU128BE()
	if err != nil {
		return nil, err
	}
	return NewCompuchaChallenge(challengeId, nonce) , nil
}

func CompuchaChallengeFromBuf(buf *FixedBuf) (*CompuchaChallenge, error) {
	return CompuchaChallengeFromBufReader(*NewBufReader(buf.buf))
}

func CompuchaChallengeFromHex(hex string) (*CompuchaChallenge, error) {
	fixedBuf, err := FixedBufFromHex(SIZE, hex)
	if err != nil {
		return nil, err
	}
	return CompuchaChallengeFromBuf(fixedBuf)
}

func (cc *CompuchaChallenge) ToBuf() (*FixedBuf, error) {
	bw := NewBufWriter()
	bw.Write(cc.challengeId.buf)
	bw.WriteU128BE(cc.nonce)
	sysBuf := bw.ToBuf()
	fixedBuf,err := NewFixedBuf(SIZE, sysBuf)
	if err != nil {
		return nil, err
	}
	return fixedBuf, nil
}

func (cc *CompuchaChallenge) ToHex() (string, error) {
	fixedBuf, err := cc.ToBuf()
	if err != nil {
		return "", err
	}
	return fixedBuf.ToHex(), nil
}

func (cc *CompuchaChallenge) Id() (*FixedBuf, error) {
	fixedBuf,err := cc.ToBuf()
	if err != nil {
		return nil, err
	}
	hash,err := DoubleBlake3Hash(fixedBuf.buf)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// isTargetValid checks if the hash is less than the target nonce.
func (cc *CompuchaChallenge) isTargetValid(targetNonce *U256) (bool, error) {
	hashBuf, err := cc.Id()
	if err != nil {
		return false, err
	}

	// Convert the hash buffer to U256
	hashNum, err := FromBEBufU256(hashBuf.buf)
	if err != nil {
		return false, err
	}

	// Perform the comparison
	return hashNum.value.Cmp(targetNonce.value) < 0, nil
}