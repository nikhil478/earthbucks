import PrivKey from "../priv-key";
import PubKey from "../pub-key";
import StrictHex from "../strict-hex";
import PermissionToken from "./permission-token";
import SignedMessage from "./signed-message";

export default class SigninChallenge {
  signedMessage: SignedMessage;

  constructor(signedMessage: SignedMessage) {
    this.signedMessage = signedMessage;
  }

  static signinChallengeKeyString(domain: string): string {
    return `signin challenge for ${domain}`;
  }

  static fromRandom(domainPrivKey: PrivKey, domain: string): SigninChallenge {
    const signInPermissionStr =
      SigninChallenge.signinChallengeKeyString(domain);
    const permissionToken = PermissionToken.fromRandom();
    const message = permissionToken.toBuffer();
    //console.log("message", message.toString("hex"));
    const signedMessage = SignedMessage.fromSignMessage(
      domainPrivKey,
      message,
      signInPermissionStr,
    );
    //console.log("message", signedMessage.message.toString("hex"));
    return new SigninChallenge(signedMessage);
  }

  static fromBuffer(buf: Buffer, domain: string): SigninChallenge {
    const signinChallengeKeyStr =
      SigninChallenge.signinChallengeKeyString(domain);
    const signedMessage = SignedMessage.fromBuffer(buf, signinChallengeKeyStr);
    //console.log("signedMessage 1", signedMessage.toBuffer().toString("hex"));
    return new SigninChallenge(signedMessage);
  }

  static fromHex(hex: string, domain: string): SigninChallenge {
    const buf = StrictHex.decode(hex);
    //console.log("buf", buf.toString("hex"));
    return SigninChallenge.fromBuffer(buf, domain);
  }

  toBuffer(): Buffer {
    return this.signedMessage.toBuffer();
  }

  toHex(): string {
    return this.toBuffer().toString("hex");
  }

  isValid(domainPubKey: PubKey, domain: string): boolean {
    const message = this.signedMessage.message;
    //console.log("signedMessage", this.signedMessage.toBuffer().toString("hex"));
    //console.log("message", message.toString("hex"));
    const permissionToken = PermissionToken.fromBuffer(message);
    if (!permissionToken.isValid()) {
      //console.log("permission token invalid");
      return false;
    }
    const keyStr = SigninChallenge.signinChallengeKeyString(domain);
    if (!this.signedMessage.isValid(domainPubKey, keyStr)) {
      //console.log("signed message invalid");
      return false;
    }
    return true;
  }
}
