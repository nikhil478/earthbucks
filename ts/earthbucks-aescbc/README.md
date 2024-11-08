# AES+CBC Encryption/Decryption Library

This TypeScript library provides a simple interface for encrypting and decrypting data using AES (Advanced Encryption Standard) in Cipher Block Chaining (CBC) mode. It primarily exposes two methods: `encrypt` and `decrypt`, allowing users to securely handle arbitrary amounts of data.

⚠️ **Important Security Note**: This library **does not verify the authenticity of the encrypted data**. If the data is altered, it may still decrypt but result in corrupted or incorrect information. To ensure the integrity of your encrypted data, it is highly recommended that you apply a hash function (such as HMAC) to verify that the data has not been tampered with.

## Installation

```bash
npm install @earthbucks/aescbc
```

## Usage

### `encrypt` Method

Encrypt data using AES with CBC mode. This method supports 128, 192, or 256-bit AES keys and a 128-bit initialization vector (IV). You can choose to prepend the IV to the encrypted data for easier storage and transmission.

```typescript
import { WebBuf } from "webbuf";
import { encrypt } from '@earthbucks/aescbc';

/**
 * Encrypt data with AES + CBC mode.
 *
 * @param {WebBuf} messageBuf - The data to encrypt. Can be any size.
 * @param {WebBuf} aesKeyBuf - The AES key (128, 192, or 256 bits).
 * @param {WebBuf} ivBuf - The initialization vector (IV, 128 bits).
 * @param {boolean} concatIvBuf - If true, prepends the IV to the encrypted data.
 * @returns {WebBuf} - The encrypted data.
 */
const aesKey = WebBuf.from('your-32-byte-key-here', 'hex');
const iv = WebBuf.from('your-16-byte-iv-here', 'hex');
const message = WebBuf.from('Hello, World!');

const encrypted = encrypt(message, aesKey, iv);
console.log(encrypted);
```

#### Parameters
- `messageBuf`: The data to encrypt. Can be any size.
- `aesKeyBuf`: The AES key. Must be 128, 192, or 256 bits.
- `ivBuf`: The initialization vector (IV). Must be 128 bits. If not provided, a random IV is generated.
- `concatIvBuf`: If set to `true`, the IV will be prepended to the encrypted data (default: `true`).

#### Returns
- A `WebBuf` containing the encrypted data. If `concatIvBuf` is true, the IV will be prepended to the output.

### `decrypt` Method

Decrypt data that was encrypted with AES in CBC mode. If the IV was prepended to the encrypted data, you can omit the `ivBuf` parameter.

```typescript
import { WebBuf } from "webbuf";
import { decrypt } from '@earthbucks/aescbc';

/**
 * Decrypt AES-encrypted data.
 *
 * @param {WebBuf} encBuf - The encrypted data.
 * @param {WebBuf} aesKeyBuf - The AES key (128, 192, or 256 bits).
 * @param {WebBuf} ivBuf - The initialization vector (optional if IV is included in `encBuf`).
 * @returns {WebBuf} - The decrypted data.
 */
const decrypted = decrypt(encrypted, aesKey, iv);
console.log(decrypted.toString());
```

#### Parameters
- `encBuf`: The encrypted data. If the IV is prepended, pass the entire buffer here.
- `aesKeyBuf`: The AES key. Must be 128, 192, or 256 bits.
- `ivBuf`: The initialization vector (IV). Must be 128 bits. If omitted, the IV is assumed to be prepended to the encrypted data.

#### Returns
- A `WebBuf` containing the decrypted data.

## Security Notice
While this library implements AES encryption, it **does not** protect against data tampering. If an attacker modifies the encrypted data, it may decrypt incorrectly without raising an error. To ensure data authenticity, you should compute and verify a cryptographic hash (such as HMAC) alongside encryption to detect unauthorized modifications.

## Example: Encrypting and Decrypting with a Prepend IV

```typescript
const aesKey = WebBuf.from('your-32-byte-key-here', 'hex');
const message = WebBuf.from('Secret message here');
const encrypted = encrypt(message, aesKey);
console.log('Encrypted:', encrypted);

// Decrypt the data
const decrypted = decrypt(encrypted, aesKey);
console.log('Decrypted:', decrypted.toString());
```

## License

This library is open-source and licensed under the MIT License.

---

Feel free to modify this README as needed, including the installation instructions, package name, and any additional sections you may want to include.
