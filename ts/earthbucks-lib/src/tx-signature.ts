import { FixedBuf, WebBuf } from "./buf.js";
import { U8, U16, U32, U64 } from "./numbers.js";

export class TxSignature {
  static readonly SIGHASH_ALL = new U8(0x00000001);
  static readonly SIGHASH_NONE = new U8(0x00000002);
  static readonly SIGHASH_SINGLE = new U8(0x00000003);
  static readonly SIGHASH_ANYONECANPAY = new U8(0x00000080);
  static readonly SIZE = 65;

  hashType: U8;
  sigBuf: FixedBuf<64>;

  constructor(hashType: U8, sigBuf: FixedBuf<64>) {
    this.hashType = hashType;
    this.sigBuf = sigBuf;
  }

  toBuf(): WebBuf {
    const hashTypeBuf = WebBuf.alloc(1);
    hashTypeBuf.writeUInt8(this.hashType.n);
    return WebBuf.concat([hashTypeBuf, this.sigBuf.buf]);
  }

  static fromBuf(buf: WebBuf): TxSignature {
    if (buf.length !== TxSignature.SIZE) {
      throw new Error("Invalid TxSignature length");
    }
    const hashType = new U8(buf[0] as number);
    const sigBuf = buf.subarray(1);
    const sigFixedEbxBuf = FixedBuf.fromBuf(64, sigBuf);
    return new TxSignature(hashType, sigFixedEbxBuf);
  }
}
