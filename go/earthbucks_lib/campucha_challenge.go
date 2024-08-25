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

func (CompuchaChallenge) FromRandomNonce(challengeId *FixedBuf) (*CompuchaChallenge, error) {
    nonceBuf, err := FixedBuf.FromRandom(FixedBuf{}, 16)
	if err != nil {
		return nil, err
	}
    nonce, err := NewBufReader(nonceBuf.buf).ReadU128BE()
	if err != nil {
		return nil, err
	}
    return NewCompuchaChallenge(challengeId, nonce) , nil
}

func (c CompuchaChallenge) FromBufReader(br BufReader) (*CompuchaChallenge, error) {
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

func (c CompuchaChallenge) FromBuf(buf *FixedBuf) (*CompuchaChallenge, error) {
	return c.FromBufReader(*NewBufReader(buf.buf))
}

func (c CompuchaChallenge) FromHex(hex string) (*CompuchaChallenge, error) {
	fixedBuf, err := FixedBuf.FromHex(FixedBuf{}, SIZE, hex)
	if err != nil {
		return nil, err
	}
	return c.FromBuf(fixedBuf)
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

func (cc *CompuchaChallenge) isTargetValid(targetNonce U256) (*bool, error) {
	hashBuf, err := cc.Id()
	if err != nil {
		return nil, err
	}
	hashNum := U256.FromBEBuf(hashBuf.buf)
	return hashNum.bn < targetNonce.value
}