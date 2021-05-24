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
// Usage: HashBytes(data, md5.New()), HashBytes(data, sha256.New())
func HashBytes(data []byte, h hash.Hash) HashResult {
	_, _ = h.Write(data)
	return h.Sum(nil)
}

// HashReader calculate hash for all data from reader.
// If error occurred when read data, return nil result and a an error.
//
// Usage: HashReader(r, md5.New()), HashReader(r, sha256.New())
func HashReader(r io.Reader, h hash.Hash) (HashResult, error) {
	_, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// HashFile calculate hash for a file. The file is closed when hash calculating finish or an err occurred.
// If error occurred when read data, return nil result and a an error.
//
// Usage: HashFile(file, md5.New()), HashFile(file, sha256.New())
func HashFile(path string, h hash.Hash) (HashResult, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return HashReader(f, h)
}

// Hash32Bytes calculate 32 bit hash as uint result for byte array data.
//
// Usage: Hash32Bytes(data, md5.New())
func Hash32Bytes(data []byte, h hash.Hash32) uint32 {
	_, _ = h.Write(data)
	return h.Sum32()
}

// Hash32Reader calculate 32 bit hash as uint result for byte array data.
// If error occurred when read data, return 0 as result and a an error.
//
// Usage: Hash32Reader(r, md5.New())
func Hash32Reader(r io.Reader, h hash.Hash32) (uint32, error) {
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
// Usage: Hash32File(file, md5.New())
func Hash32File(path string, h hash.Hash32) (uint32, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return Hash32Reader(f, h)
}

// Hash64Bytes calculate 64 bit hash as uint result for byte array data.
func Hash64Bytes(data []byte, h hash.Hash64) uint64 {
	_, _ = h.Write(data)
	return h.Sum64()
}

// Hash64Reader calculate 64 bit hash as uint result for byte array data.
// If error occurred when read data, return 0 as result and a an error.
func Hash64Reader(r io.Reader, h hash.Hash64) (uint64, error) {
	_, err := io.Copy(h, r)
	if err != nil {
		return 0, err
	}
	return h.Sum64(), nil
}

// Hash64File calculate 64 bit hash as uint result for byte array data.
// The file is closed when hash calculating finish or an err occurred.
// If error occurred when read data, return 0 as result and a an error.
func Hash64File(path string, h hash.Hash64) (uint64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return Hash64Reader(f, h)
}
