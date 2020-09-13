const Base32 = require('base32.js');
const crypto = require('crypto');

const key = 'JBSWY3DPEHPK3PXP';

const decoder = new Base32.Decoder();
const secret = decoder.write(key).finalize();

let counter = Math.floor(Date.now() / 30000);

const bytes = Buffer.alloc(8, 0);
bytes.writeUInt32BE(counter, 4);

const hmac = crypto.createHmac('sha1', secret);
hmac.update(bytes);

const hs = hmac.digest();
const n = hs[19] & 0xF

const result = (hs[n] << 24 | hs[n + 1] << 16 | hs[n + 2] << 8 | hs[n + 3]) & 0x7fffffff

console.log(result.toString().slice(-6).padStart(6, 0));