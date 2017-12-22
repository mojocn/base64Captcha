// Copyright 2017 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base64Captcha

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// idLen is a length of captcha id string.
// (20 bytes of 62-letter alphabet give ~119 bits.)
const idLen = 20

// idChars are characters allowed in captcha id.
var idChars = []byte(TxtNumbers + TxtAlphabet)

// rngKey is a secret key used to deterministically derive seeds for
// PRNGs used in image and audio. Generated once during initialization.
var rngKey [32]byte

func init() {
	if _, err := io.ReadFull(rand.Reader, rngKey[:]); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
}

// Purposes for seed derivation. The goal is to make deterministic PRNG produce
// different outputs for images and audio by using different derived seeds.
const (
	imageSeedPurpose = 0x01
	audioSeedPurpose = 0x02
)

// deriveSeed returns a 16-byte PRNG seed from rngKey, purpose, id and digits.
// Same purpose, id and digits will result in the same derived seed for this
// instance of running application.
//
//   out = HMAC(rngKey, purpose || id || 0x00 || digits)  (cut to 16 bytes)
//
func deriveSeed(purpose byte, id string, digits []byte) (out [16]byte) {
	var buf [sha256.Size]byte
	h := hmac.New(sha256.New, rngKey[:])
	h.Write([]byte{purpose})
	io.WriteString(h, id)
	h.Write([]byte{0})
	h.Write(digits)
	sum := h.Sum(buf[:0])
	copy(out[:], sum)
	return
}

// randomDigits returns a byte slice of the given length containing
// pseudorandom numbers in range 0-9. The slice can be used as a captcha
// solution.
func randomDigits(length int) []byte {
	return randomBytesMod(length, 10)
}

// randomBytes returns a byte slice of the given length read from CSPRNG.
func randomBytes(length int) (b []byte) {
	b = make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
	return
}

// randomBytesMod returns a byte slice of the given length, where each byte is
// a random number modulo mod.
func randomBytesMod(length int, mod byte) (b []byte) {
	if length == 0 {
		return nil
	}
	if mod == 0 {
		panic("captcha: bad mod argument for randomBytesMod")
	}
	maxrb := 255 - byte(256%int(mod))
	b = make([]byte, length)
	i := 0
	for {
		r := randomBytes(length + (length / 4))
		for _, c := range r {
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = c % mod
			i++
			if i == length {
				return
			}
		}
	}
}

// randomId returns a new random id key string.
// golang服务器端产生随机的idKey
func randomId() string {
	b := randomBytesMod(idLen, byte(len(idChars)))
	for i, c := range b {
		b[i] = idChars[c]
	}
	return string(b)
}

//parseDigitsToString parse randomDigits to normal string
func parseDigitsToString(bytes []byte) string {
	ssbb := make([]byte, len(bytes))
	for idx, by := range bytes {
		ssbb[idx] = by + '0'
	}
	return string(ssbb)
}
