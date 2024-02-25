package snowball

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"math"
	"strconv"
	"strings"
)

const base62Digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Formats the Snowball ID into a binary string.
func (id SnowballID) ToBinary() string {
	return strconv.FormatUint(uint64(id), 2)
}

// Converts a binary string into a Snowball ID.
func FromBinary(sid string) (SnowballID, error) {
	id, err := strconv.ParseUint(sid, 2, 64)
	return SnowballID(id), err
}

// Formats the Snowball ID into a hexidecimal string.
func (id SnowballID) ToHex() string {
	return strconv.FormatUint(uint64(id), 16)
}

// Converts a hexidecimal string into a Snowball ID.
func FromHex(sid string) (SnowballID, error) {
	id, err := strconv.ParseUint(sid, 16, 64)
	return SnowballID(id), err
}

// Formats the Snowball ID into a base32 encoded string, using the standard encoding for Base32.
func (id SnowballID) ToBase32() string {
	b := make([]byte, 8)
	binary.NativeEndian.PutUint64(b, uint64(id))
	return base32.StdEncoding.EncodeToString(b)
}

// Converts a base32 string into a Snowball ID. Assumes the string is formatted using standard Base32.
func FromBase32(sid string) (SnowballID, error) {
	bytes, err := base32.StdEncoding.DecodeString(sid)
	if err != nil {
		return 0, errors.New("decode failed: invalid base32 string (did you use the right encoding?)")
	}

	return SnowballID(binary.NativeEndian.Uint64(bytes)), nil
}

// Formats the Snowball ID into a base64 encoded string, using the URL encoding for Base64.
func (id SnowballID) ToBase64() string {
	b := make([]byte, 8)
	binary.NativeEndian.PutUint64(b, uint64(id))
	return base64.URLEncoding.EncodeToString(b)
}

// Converts a base64 string into a Snowball ID. Assumes the string is formatted using the URL variant
// of Base64.
func FromBase64(sid string) (SnowballID, error) {
	bytes, err := base64.URLEncoding.DecodeString(sid)
	if err != nil {
		return 0, errors.New("decode failed: invalid base64 string (did you use the right encoding?)")
	}

	return SnowballID(binary.NativeEndian.Uint64(bytes)), nil
}

// Formats the Snowball ID into a base62 encoded string.
func (id SnowballID) ToBase62() string {
	result := ""
	for id > 0 {
		remainder := id % 62
		result = string(base62Digits[remainder]) + result
		id /= 62
	}

	return result
}

// Converts a base62 string into a Snowball ID.
func FromBase62(sid string) (SnowballID, error) {
	var result uint64
	for index, char := range sid {
		pow := len(sid) - (index + 1)
		pos := strings.IndexRune(base62Digits, char)
		if pos == -1 {
			return 0, errors.New("decode failed: invalid base62 string")
		}

		result += uint64(pos) * uint64(math.Pow(62, float64(pow)))
	}

	return SnowballID(result), nil
}
