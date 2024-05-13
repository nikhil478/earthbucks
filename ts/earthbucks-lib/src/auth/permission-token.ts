import IsoBufReader from "../iso-buf-reader";
import IsoBufWriter from "../iso-buf-writer";
import { Buffer } from "buffer";
import { Result, Ok, Err } from "../ts-results/result";

export default class PermissionToken {
  randValue: Buffer;
  timestamp: bigint; // milliseconds

  constructor(randValue: Buffer, timestamp: bigint) {
    this.randValue = randValue;
    this.timestamp = timestamp; // milliseconds
  }

  toIsoBuf(): Buffer {
    const writer = new IsoBufWriter();
    writer.writeBuffer(this.randValue);
    writer.writeUInt64BE(this.timestamp);
    return writer.toIsoBuf();
  }

  static fromIsoBuf(buf: Buffer): Result<PermissionToken, string> {
    try {
      if (buf.length !== 32 + 8) {
        return new Err("invalid buffer length");
      }
      const reader = new IsoBufReader(buf);
      const randValue = reader
        .read(32)
        .mapErr((err) => `Unable to read rand value: ${err}`)
        .unwrap();
      const timestamp = reader
        .readU64BE()
        .mapErr((err) => `Unable to read timestamp: ${err}`)
        .unwrap();
      return new Ok(new PermissionToken(randValue, timestamp));
    } catch (err) {
      return new Err(
        err?.toString() || "Unknown error parsing permission token",
      );
    }
  }

  static fromRandom(): PermissionToken {
    const randValue = crypto.getRandomValues(new Uint8Array(32));
    const timestamp = BigInt(Date.now()); // milliseconds
    return new PermissionToken(Buffer.from(randValue), timestamp);
  }

  isValid(): boolean {
    return Date.now() - Number(this.timestamp) < 15 * 60 * 1000; // 15 minutes
  }
}
