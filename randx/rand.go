package randx

import (
	"math"
	"math/rand"
	"strconv"
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

// IntWithin return a random value within range [0, bound)
// If bound is less than or equals with 0, this method panics
func (r *Rand) IntWithin(bound int) int {
	if bound <= 0 {
		panic("bound less or equal than zero: " + strconv.Itoa(bound))
	}
	if bound <= math.MaxInt32 {
		return int(r.Int32Within(int32(bound)))
	}
	return int(r.Int64Within(int64(bound)))
}

// IntBetween return a random value within range [low, high)
func (r *Rand) IntBetween(low int, high int) int {
	if low >= high {
		panic("high " + strconv.Itoa(high) + " less or equal than low: " + strconv.Itoa(low))
	}
	return low + r.IntWithin(high-low)
}

// Int32Within return a random value within range [0, bound)
// If bound is less than or equals with 0, this method panics
func (r *Rand) Int32Within(bound int32) int32 {
	if bound <= 0 {
		panic("bound less or equal than zero: " + strconv.Itoa(int(bound)))
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

// Int64Within return a random value within range [0, bound)
// If bound is less than or equals with 0, this method panics
func (r *Rand) Int64Within(bound int64) int64 {
	if bound <= 0 {
		panic("bound less or equal than zero: " + strconv.FormatInt(bound, 10))
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

// Int64Between return a random value within range [low, high)
func (r *Rand) Int64Between(low int64, high int64) int64 {
	if low >= high {
		panic("high " + strconv.FormatInt(high, 10) + " less or equal than low: " + strconv.FormatInt(low, 10))
	}
	return low + r.Int64Within(high-low)
}
