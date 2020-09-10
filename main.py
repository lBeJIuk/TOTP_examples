import hashlib
import hmac
import time
import base64
key = "JBSWY3DPEHPK3PXP"

secret = base64.b32decode(key.encode("utf-8"))
counter = int(time.time() / 30)

bytes = bytearray()
for _ in range(0, 8):
  bytes.insert(0, counter & 0xff)
  counter >>= 8

hs = bytearray(hmac.new(secret, bytes, hashlib.sha1).digest())

n = hs[-1] & 0xF
result = (hs[n] << 24 | hs[n + 1] << 16 | hs[n + 2] << 8 | hs[n + 3]) & 0x7fffffff

print(str(result)[-6:])