import { OPCODE_TO_NAME, OP, OpcodeName, Opcode } from "./opcode";
import IsoBufWriter from "./iso-buf-writer";
import IsoBufReader from "./iso-buf-reader";
import { Buffer } from "buffer";
import { Result, Ok, Err } from "./ts-results/result";

export default class ScriptChunk {
  opcode: number;
  buf?: Buffer;

  constructor(opcode: number = 0, buf?: Buffer) {
    this.opcode = opcode;
    this.buf = buf;
  }

  getData(): Result<Buffer, string> {
    if (this.opcode === Opcode.OP_1NEGATE) {
      return Ok(Buffer.from([0x80]));
    } else if (this.opcode === Opcode.OP_0) {
      return Ok(Buffer.from([]));
    } else if (this.opcode >= Opcode.OP_1 && this.opcode <= Opcode.OP_16) {
      return Ok(Buffer.from([this.opcode - Opcode.OP_1 + 1]));
    }
    if (this.buf) {
      return Ok(this.buf);
    } else {
      return Err("no data");
    }
  }

  toIsoStr(): Result<string, string> {
    if (this.buf) {
      return Ok(`0x${this.buf.toString("hex")}`);
    } else {
      const name = OPCODE_TO_NAME[this.opcode];
      if (name !== undefined) {
        return Ok(name);
      } else {
        return Err("invalid opcode");
      }
    }
  }

  static fromIsoStr(str: string): Result<ScriptChunk, string> {
    const scriptChunk = new ScriptChunk();
    if (str.startsWith("0x")) {
      scriptChunk.buf = Buffer.from(str.slice(2), "hex");
      const len = scriptChunk.buf.length;
      const onebytelen = len <= 0xff;
      const twobytelen = len <= 0xffff;
      const fourbytelen = len <= 0xffffffff;
      if (onebytelen) {
        scriptChunk.opcode = OP.PUSHDATA1;
      } else if (twobytelen) {
        scriptChunk.opcode = OP.PUSHDATA2;
      } else if (fourbytelen) {
        scriptChunk.opcode = OP.PUSHDATA4;
      } else {
        return Err("too much data");
      }
    } else {
      function isOpcodeName(str: string): str is OpcodeName {
        return Object.prototype.hasOwnProperty.call(OP, str);
      }
      if (isOpcodeName(str)) {
        const opcode = OP[str];
        scriptChunk.opcode = opcode;
      } else {
        return Err("invalid opcode");
      }

      scriptChunk.buf = undefined;
    }
    return Ok(scriptChunk);
  }

  toIsoBuf(): Buffer {
    const opcode = this.opcode;
    if (opcode === OP.PUSHDATA1 && this.buf) {
      return Buffer.concat([
        Buffer.from([opcode]),
        new IsoBufWriter().writeUInt8(this.buf.length).toIsoBuf(),
        this.buf,
      ]);
    } else if (opcode === OP.PUSHDATA2 && this.buf) {
      return Buffer.concat([
        Buffer.from([opcode]),
        new IsoBufWriter().writeUInt16BE(this.buf.length).toIsoBuf(),
        this.buf,
      ]);
    } else if (opcode === OP.PUSHDATA4 && this.buf) {
      return Buffer.concat([
        Buffer.from([opcode]),
        new IsoBufWriter().writeUInt32BE(this.buf.length).toIsoBuf(),
        this.buf,
      ]);
    }
    return Buffer.from([opcode]);
  }

  static fromIsoBuf(buf: Buffer): Result<ScriptChunk, string> {
    const reader = new IsoBufReader(buf);
    return ScriptChunk.fromIsoBufReader(reader);
  }

  static fromIsoBufReader(reader: IsoBufReader): Result<ScriptChunk, string> {
    const opcodeRes = reader.readU8();
    if (opcodeRes.err) {
      return opcodeRes.mapErr(
        (err) =>
          `script_chunk::from_iso_buf_reader 1: unable to read opcode: ${err}`,
      );
    }
    const opcode = opcodeRes.val;
    const chunk = new ScriptChunk(opcode);
    if (opcode === OP.PUSHDATA1) {
      const lenRes = reader.readU8();
      if (lenRes.err) {
        return lenRes.mapErr(
          (err) =>
            `script_chunk::from_iso_buf_reader 2: unable to read 1 byte length: ${err}`,
        );
      }
      const len = lenRes.unwrap();
      const bufferRes = reader.read(len);
      if (bufferRes.err) {
        return bufferRes.mapErr(
          (err) =>
            `script_chunk::from_iso_buf_reader 3: unable to read buffer: ${err}`,
        );
      }
      const buffer = bufferRes.unwrap();
      if (len == 0 || (len === 1 && buffer[0] >= 1 && buffer[0] <= 16)) {
        return Err("script_chunk::from_iso_buf_reader 4: non-minimal pushdata");
      }
      chunk.buf = buffer;
    } else if (opcode === OP.PUSHDATA2) {
      const lenRes = reader.readU16BE();
      if (lenRes.err) {
        return lenRes.mapErr(
          (err) =>
            `script_chunk::from_iso_buf_reader 5: unable to read 2 byte length: ${err}`,
        );
      }
      const len = lenRes.unwrap();
      if (len <= 0xff) {
        return Err("script_chunk::from_iso_buf_reader 6: non-minimal pushdata");
      }
      const bufferRes = reader.read(len);
      if (bufferRes.err) {
        return bufferRes.mapErr(
          (err) =>
            `script_chunk::from_iso_buf_reader 7: unable to read buffer: ${err}`,
        );
      }
      const buffer = bufferRes.unwrap();
      chunk.buf = buffer;
    } else if (opcode === OP.PUSHDATA4) {
      const lenRes = reader.readU32BE();
      if (lenRes.err) {
        return lenRes.mapErr(
          (err) =>
            `script_chunk::from_iso_buf_reader 8: unable to read 4 byte length: ${err}`,
        );
      }
      const len = lenRes.unwrap();
      if (len <= 0xffff) {
        return Err("script_chunk::from_iso_buf_reader 9: non-minimal pushdata");
      }
      const bufferRes = reader.read(len);
      if (bufferRes.err) {
        return bufferRes.mapErr(
          (err) =>
            `script_chunk::from_iso_buf_reader 10: unable to read buffer: ${err}`,
        );
      }
      const buffer = bufferRes.unwrap();
      chunk.buf = buffer;
    }
    return Ok(chunk);
  }

  static fromData(data: Buffer): ScriptChunk {
    const len = data.length;
    if (len === 0) {
      return new ScriptChunk(Opcode.OP_0);
    } else if (len === 1 && data[0] >= 1 && data[0] <= 16) {
      return new ScriptChunk(data[0] + Opcode.OP_1 - 1);
    } else if (len <= 0xff) {
      return new ScriptChunk(Opcode.OP_PUSHDATA1, data);
    } else if (len <= 0xffff) {
      return new ScriptChunk(Opcode.OP_PUSHDATA2, data);
    } else if (len <= 0xffffffff) {
      return new ScriptChunk(Opcode.OP_PUSHDATA4, data);
    } else {
      return new ScriptChunk(0);
    }
  }

  static fromSmallNumber(n: number): ScriptChunk {
    if (n === -1 || (n >= 1 && n <= 16)) {
      return new ScriptChunk(n + Opcode.OP_1 - 1);
    } else {
      return new ScriptChunk(0);
    }
  }
}
