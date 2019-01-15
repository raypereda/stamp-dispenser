// Package stamps facilitates dispensing stamps for a postage stamp machine.
// Dispensing the minimum stamps is equivalent to the making change problem.
// https://en.wikipedia.org/wiki/Change-making_problem
package stamps

import (
	"errors"
	"math"
	"sort"
)

// Dispenser represents a stamp dispenser.
type Dispenser struct {
	Denominations []int       // the values of the stamps are in cents
	memo          map[int]int // memoizatoin of Dispense methods
	memoMax       int         // max key of memo
}

// New initializes a new dispenser
// demoninations has the value of the stamp values in cents.
func New(denominations []int) (*Dispenser, error) {
	if len(denominations) < 1 {
		return nil, errors.New("length of demominations should be at least 1 but size is zero")
	}
	// sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(denominations)))

	memo := make(map[int]int)
	return &Dispenser{Denominations: denominations, memo: memo}, nil
}

// Dispense1 returns close to minimum number of stamps needed to fill postage request.
// This uses the ordinary greedy algorithm that people use with ordinary stamp denominations.
// postage is the total value of the stamps to be dispensed.
// Computational complexity is O(d) where d is number of demominations.
func (d *Dispenser) Dispense1(postage int) (int, error) {
	if postage < 1 {
		return 0, errors.New("postage must be positive")
	}
	var stamps, quotient, remainder int
	for _, stamp := range d.Denominations {
		quotient = postage / stamp
		remainder = postage % stamp
		stamps += quotient
		postage = remainder
		if postage == 0 {
			break
		}
	}
	if postage > 0 && postage < d.Denominations[0] {
		// the postage is not exact but we can just add one more extra stamp to cover it
		return stamps + 1, nil
	}
	return stamps, nil
}

// Dispense2 returns the minimum number of stamps needed to fill postage request.
// This uses a dynamic programming, top-down algorithm.
// postage is the total value of the stamps to be dispensed.
// Computational complexity is O(n) where n is the postage.
// The worse case occurs with one denomination of value 1.
func (d *Dispenser) Dispense2(postage int) (int, error) {
	if postage < 1 {
		return 0, errors.New("postage must be positive")
	}

	p, ok := d.memo[postage] // return memoized answer
	if ok {
		return p, nil
	}

	var minStamps = math.MaxInt64 // positive infinity
	for _, demom := range d.Denominations {
		if demom > postage {
			continue
		}
		if demom == postage {
			p, _ := d.memo[postage]
			if p > 1 {
				d.memo[postage] = 1
			}
			return 1, nil
		}
		stamps, err := d.Dispense2(postage - demom)
		if err != nil {
			panic(err)
		}
		stamps++
		if stamps < minStamps {
			minStamps = stamps
		}
	}

	// there is no exact postage; dispense one stamp that is larger than the postage
	if minStamps == math.MaxInt64 {
		d.memo[postage] = 1
		return 1, nil
	}

	d.memo[postage] = minStamps
	return minStamps, nil
}

// Dispense3 returns the minimum number of stamps needed to fill postage request.
// This uses a dynamic programming, bottom-up algorithm.
// postage is the total value of the stamps to be dispensed.
// Complexity is O(n*d) where n postage and d is the number of denominations.
func (d *Dispenser) Dispense3(postage int) (int, error) {
	if postage < 1 {
		return 0, errors.New("postage must be positive")
	}
	p, ok := d.memo[postage] // return memoized answer
	if ok {
		return p, nil
	}
	for p := d.memoMax + 1; p <= postage; p++ {
		minStamps := math.MaxInt64
		for _, denom := range d.Denominations {
			if postage-denom < 0 {
				// assume can return exact postage
				continue
			}
			stamps := d.memo[postage-denom] + 1
			if stamps < minStamps {
				minStamps = stamps
			}
		}
		d.memo[p] = minStamps
	}
	if postage > d.memoMax {
		d.memoMax = postage
	}
	return d.memo[postage], nil
}
