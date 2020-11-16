package randx

import (
	"errors"
	"github.com/hsiafan/glow/intx"
	"math"
	"math/rand"
	"time"
)

// Rand is a rand with more useful methods
type Rand struct {
	rand.Rand
}

// New return a new Rand using timestamp as seed
func New() *Rand {
	return NewWithSeed(time.Now().Unix())
}

// NewWithSeed return a new Rand using seed
func NewWithSeed(seed int64) *Rand {
	return &Rand{
		Rand: *rand.New(rand.NewSource(seed)),
	}
}

var boundError = errors.New("bound less than or equals zero")
var boundRangeError = errors.New("low bound larger than/equals high bound")
var boundOverFlowError = errors.New("bound range overflows int")

// IntWithin return a random value within range [0, bound) if bound larger than 0.
// panics bound is less than or equals 0.
func (r *Rand) IntWithin(bound int) int {
	if bound <= 0 {
		panic(boundError)
	}
	if bound <= math.MaxInt32 {
		v := r.Int32Within(int32(bound))
		return int(v)
	}
	v := r.Int64Within(int64(bound))
	return int(v)
}

// IntBetween return a random value within range [low, high) if low less than high,
// The func panics an error if low is larger than or equals high, or high-low overflows int.
func (r *Rand) IntBetween(low int, high int) int {
	if low >= high {
		panic(boundRangeError)
	}
	if low < 0 && (intx.MaxInt+low) < high {
		panic(boundOverFlowError)
	}
	v := r.IntWithin(high - low)
	return low + v
}

// Int32Within return a random int32 value within range [0, bound) if bound larger than 0,
// Panics with an error if bound is less than or equals 0.
func (r *Rand) Int32Within(bound int32) int32 {
	if bound <= 0 {
		panic(boundError)
	}
	v := r.Int31()
	m := bound - 1
	if bound&m == 0 {
		// i.e., bound is a power of 2
		// returns the of high-order bits from the underlying pseudo-random number generator.
		// Linear congruential pseudo-random number generators are known to have short periods in the sequence of values of their low-order bits.
		v = int32((int64(bound) * int64(v)) >> 31)
	} else {
		// throws away numbers at the "top" of the range so that the random number is evenly distributed.
		for u := v; ; u = r.Int31() {
			v = u % bound
			if u-v+m >= 0 {
				break
			}
		}
	}
	return v
}

// Int64Within return a random int64 value within range [0, bound).
// If bound is less than or equals with 0, panics with an error
func (r *Rand) Int64Within(bound int64) int64 {
	if bound <= 0 {
		panic(boundError)
	}
	v := r.Int63()
	m := bound - 1
	if bound&m == 0 {
		// i.e., bound is a power of 2
		// em.. just use the lower bits
		v = v & m
	} else {
		// throws away numbers at the "top" of the range so that the random number is evenly distributed.
		for u := v; ; u = r.Int63() {
			v = u % bound
			if u-v+m >= 0 {
				break
			}
		}
	}
	return v
}
