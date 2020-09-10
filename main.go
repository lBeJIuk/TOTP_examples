package main

import (
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"log"
	"time"
	"crypto/sha1"
	"crypto/hmac"
	"strconv"
	"bytes"
)

func main() {
	key := "JBSWY3DPEHPK3PXP"

	secret, err := base32.StdEncoding.DecodeString(key)
	if err != nil {
		log.Fatal("error:", err)
		log.Fatal("error:", secret)
	}
	counter := time.Now().Unix() / 30

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(counter))
	
	mac := hmac.New(sha1.New, secret)
	mac.Write(counterBytes)
	hs := mac.Sum(nil)
	
	n := (hs[len(hs) - 1] & 0xF)
	
	var header uint32
	reader := bytes.NewReader(hs[n : n + 4])
	_ = binary.Read(reader, binary.BigEndian, &header)
	h12 := (int(header) & 0x7fffffff) % 1000000
	otp := prefix0(strconv.Itoa(int(h12)))
	fmt.Println(otp)
}

func prefix0(otp string) string {
	if len(otp) == 6 {
		return otp
	}
	for i := (6 - len(otp)); i > 0; i-- {
		otp = "0" + otp
	}
	return otp
}