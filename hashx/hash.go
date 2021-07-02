package hashx

import (
	"encoding/hex"
	"hash"
	"io"
	"os"
	"strings"
)

// HashResult hold hash results
type HashResult []byte

// ToHex return hash result data as lower case hexed string
func (r HashResult) ToHex() string {
	return hex.EncodeToString(r)
}

// ToHexUpper return hash result data as upper case hexed string
func (r HashResult) ToHexUpper() string {
	return strings.ToUpper(hex.EncodeToString(r))
}

// Data return hash result as binary data.
func (r HashResult) Data() []byte {
	return r
}

// HashBytes calculate hash for byte array data.
//
// Usage: HashBytes(data, md5.New), HashBytes(data, sha256.New)
func HashBytes(data []byte, f func() hash.Hash) HashResult {
	h := f()
	_, _ = h.Write(data)
	return h.Sum(nil)
}

// HashReader calculate hash for all data from reader.
// If error occurred when read data, return nil result and a an error.
//
// Usage: HashReader(r, md5.New), HashReader(r, sha256.New)
func HashReader(r io.Reader, f func() hash.Hash) (HashResult, error) {
	h := f()
	_, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// HashFile calculate hash for a file. The file is closed when hash calculating finish or an err occurred.
// If error occurred when read data, return nil result and a an error.
//
// Usage: HashFile(file, md5.New), HashFile(file, sha256.New)
func HashFile(path string, f func() hash.Hash) (HashResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return HashReader(file, f)
}

// Hash32Bytes calculate 32 bit hash as uint result for byte array data.
//
// Usage: Hash32Bytes(data, crc32.New)
func Hash32Bytes(data []byte, f func() hash.Hash32) uint32 {
	h := f()
	_, _ = h.Write(data)
	return h.Sum32()
}

// Hash32Reader calculate 32 bit hash as uint result for byte array data.
// If error occurred when read data, return 0 as result and a an error.
//
// Usage: Hash32Reader(r, crc32.New)
func Hash32Reader(r io.Reader, f func() hash.Hash32) (uint32, error) {
	h := f()
	_, err := io.Copy(h, r)
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

// Hash32File calculate 32 bit hash as uint result for byte array data.
// The file is closed when hash calculating finish or an err occurred.
// If error occurred when read data, return 0 as result and a an error.
//
// Usage: Hash32File(file, crc32.New)
func Hash32File(path string, f func() hash.Hash32) (uint32, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return Hash32Reader(file, f)
}

// Hash64Bytes calculate 64 bit hash as uint result for byte array data.
//
// Usage: Hash64File(file, crc64.New})
func Hash64Bytes(data []byte, f func() hash.Hash64) uint64 {
	h := f()
	_, _ = h.Write(data)
	return h.Sum64()
}

// Hash64Reader calculate 64 bit hash as uint result for byte array data.
// If error occurred when read data, return 0 as result and a an error.
//
// Usage: Hash64File(file, crc64.New})
func Hash64Reader(r io.Reader, f func() hash.Hash64) (uint64, error) {
	h := f()
	_, err := io.Copy(h, r)
	if err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

// Hash64File calculate 64 bit hash as uint result for byte array data.
// The file is closed when hash calculating finish or an err occurred.
// If error occurred when read data, return 0 as result and a an error.
//
// Usage: Hash64File(file, crc64.New})
func Hash64File(path string, f func() hash.Hash64) (uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return Hash64Reader(file, f)
}
