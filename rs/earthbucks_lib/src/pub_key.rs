use crate::ebx_error::EbxError;
use crate::hash::blake3_hash;
use crate::iso_hex;
use crate::priv_key::PrivKey;
use bs58;
use secp256k1::PublicKey;

#[derive(Debug, Clone)]
pub struct PubKey {
    pub buf: [u8; PubKey::SIZE],
}

impl PubKey {
    pub const SIZE: usize = 33;

    pub fn new(pub_key: [u8; PubKey::SIZE]) -> Self {
        PubKey { buf: pub_key }
    }

    pub fn from_iso_buf(vec: Vec<u8>) -> Result<Self, EbxError> {
        if vec.len() > PubKey::SIZE {
            return Err(EbxError::TooMuchDataError { source: None });
        }
        if vec.len() < PubKey::SIZE {
            return Err(EbxError::NotEnoughDataError { source: None });
        }
        let mut pub_key = [0u8; PubKey::SIZE];
        pub_key.copy_from_slice(&vec);
        Ok(PubKey::new(pub_key))
    }

    pub fn to_buffer(&self) -> &[u8; PubKey::SIZE] {
        &self.buf
    }

    pub fn from_priv_key(priv_key: &PrivKey) -> Result<Self, EbxError> {
        let pub_key_buf = priv_key.to_pub_key_buffer();
        if pub_key_buf.is_err() {
            return Err(EbxError::InvalidKeyError { source: None });
        }
        Ok(PubKey::new(pub_key_buf.unwrap()))
    }

    pub fn to_iso_hex(&self) -> String {
        hex::encode(self.buf)
    }

    pub fn from_iso_hex(hex: &str) -> Result<PubKey, EbxError> {
        let pub_key_buf = iso_hex::decode(hex)?;
        PubKey::from_iso_buf(pub_key_buf)
    }

    pub fn to_iso_str(&self) -> String {
        let check_hash = blake3_hash(&self.buf);
        let check_sum = &check_hash[0..4];
        let check_hex = hex::encode(check_sum);
        "ebxpub".to_string() + &check_hex + &bs58::encode(&self.buf).into_string()
    }

    pub fn from_iso_str(s: &str) -> Result<PubKey, EbxError> {
        if !s.starts_with("ebxpub") {
            return Err(EbxError::InvalidEncodingError { source: None });
        }
        let check_str = &s[6..14];
        let check_buf = iso_hex::decode(check_str)?;
        let buf = bs58::decode(&s[14..])
            .into_vec()
            .map_err(|_| EbxError::InvalidEncodingError { source: None })?;
        let check_hash = blake3_hash(&buf);
        let check_sum = &check_hash[0..4];
        if check_buf != check_sum {
            return Err(EbxError::InvalidChecksumError { source: None });
        }
        PubKey::from_iso_buf(buf)
    }

    pub fn is_valid(&self) -> bool {
        let public_key_obj = PublicKey::from_slice(&self.buf);
        public_key_obj.is_ok()
    }

    pub fn is_valid_string_fmt(s: &str) -> bool {
        let res = Self::from_iso_str(s);
        res.is_ok()
    }
}
