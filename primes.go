package main

import (
	"fmt"
	"math"
)

func main() {
	const max = 100
	fmt.Println("primes1")
	for _, p := range primes1(max) {
		fmt.Println(p)
	}
	fmt.Println("primes2")
	for _, p := range primes2(max) {
		fmt.Println(p)
	}
}

func primes1(max int) []int {
	primes := []int{}
	for i := 2; i <= max; i++ {
		var isComposite bool
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isComposite = true
				break
			}
		}
		if !isComposite {
			primes = append(primes, i)
		}
	}
	return primes
}

func primes2(max int) []int {
	primes := []int{}
outer:
	for i := 2; i <= max; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue outer
			}
		}
		primes = append(primes, i)
	}
	return primes
}

func primes3(max int) []int {
	primes := []int{}
outer:
	for i := 2; i <= max; i++ {
		for _, j := range primes {
			if i%j == 0 {
				continue outer
			}
		}
		primes = append(primes, i)
	}
	return primes
}

func primes4(max int) []int {
	primes := []int{}
outer:
	for i := 2; i <= max; i++ {
		sqrt := int(math.Sqrt(float64(i)))
		for _, j := range primes {
			if j > sqrt {
				break
			}
			if i%j == 0 {
				continue outer
			}
		}
		primes = append(primes, i)
	}
	return primes
}

func primes5(max int) []int {
	primes := []int{2}
outer:
	for i := 3; i <= max; i += 2 {
		sqrt := int(math.Sqrt(float64(i)))
		for _, j := range primes {
			if j > sqrt {
				break
			}
			if i%j == 0 {
				continue outer
			}
		}
		primes = append(primes, i)
	}
	return primes
}

type primeGenerator6 struct {
	primes []int
}

func (pg *primeGenerator6) Next() int {
	l := len(pg.primes)

	switch l {
	case 0:
		pg.primes = []int{2}
		return 2
	case 1:
		pg.primes = append(pg.primes, 3)
		return 3
	default:
		i := pg.primes[l-1] // fetch last prime number generated
	outer:
		i += 2 // next prime candidate is 2 numbers higher; skip even numbers
		sqrt := int(math.Sqrt(float64(i)))
		for _, j := range pg.primes {
			if j > sqrt {
				break
			}
			if i%j == 0 {
				goto outer
			}
		}
		pg.primes = append(pg.primes, i)
		return i
	}
}

type primeGenerator7 struct {
	primes []int
	next   func() int
}

func (pg *primeGenerator7) Next() int {
	if pg.next == nil {
		pg.next = func() int {
			pg.primes = []int{2}
			pg.next = func() int {
				pg.primes = append(pg.primes, 3)
				pg.next = func() int {
					i := pg.primes[len(pg.primes)-1] // fetch last prime number generated
				outer:
					i += 2 // next prime candidate is 2 numbers higher; skip even numbers
					sqrt := int(math.Sqrt(float64(i)))
					for _, j := range pg.primes {
						if j > sqrt {
							break
						}
						if i%j == 0 {
							goto outer
						}
					}
					pg.primes = append(pg.primes, i)
					return i
				}
				return 3
			}
			return 2
		}
	}
	return pg.next()
}

func primes8(max uint) []int {
	// Create a bit array to mark known composite numbers. After loop below, a 0
	// value means the corresponding natural number is prime, and a 1 value
	// means a composite number. Those defaults were chosen to take advantage of
	// the zero value for each byte in a slice, assuming all numbers are prime,
	// then mark off composite numbers.
	//
	// While it is not necessarily needed to store whether 0 and 1 are prime,
	// reserving bits for those values reduces the number of math instructions
	// required to set and check bits. Therefore, the 0th bit is for natural
	// number 0. The 1st bit is for natural number 1. And so
	// forth. Additionally, no attempt is ever made to set or check the bits
	// representing 0 or 1. While they are 0 at the end of the outer loop,
	// indicating they would be prime, they remain 0 simply because we never set
	// them at the beginning of this function.
	b, r := 1+max>>3, max&7
	if r > 0 {
		b++
	}
	composites := make([]byte, b)

	for i := uint(2); i <= max; i++ {
		for j := uint(i << 1); j <= max; j += i {
			major, minor := j>>3, j&7
			composites[major] |= 1 << minor
		}
	}

	primes := []int{2}

	for i := uint(3); i <= max; i += 2 {
		major, minor := i>>3, i&7
		if composites[major]&(1<<minor) == 0 {
			primes = append(primes, int(i))
		}
	}

	return primes
}

func primes9(max uint) []int {
	// Create a bit array to mark known composite numbers. After loop below, a 0
	// value means the corresponding natural number is prime, and a 1 value
	// means a composite number. Those defaults were chosen to take advantage of
	// the zero value for each byte in a slice, assuming all numbers are prime,
	// then mark off composite numbers.
	//
	// While it is not necessarily needed to store whether 0 and 1 are prime,
	// reserving bits for those values reduces the number of math instructions
	// required to set and check bits. Therefore, the 0th bit is for natural
	// number 0. The 1st bit is for natural number 1. And so
	// forth. Additionally, no attempt is ever made to set or check the bits
	// representing 0 or 1. While they are 0 at the end of the outer loop,
	// indicating they would be prime, they remain 0 simply because we never set
	// them at the beginning of this function.
	b, r := 1+max>>3, max&7
	if r > 0 {
		b++
	}
	composites := make([]byte, b)

	i := uint(2)
outer:
	for {
		// phase 1: find next prime, leaving it in 'i'.
		major, minor := i>>3, i&7
		for composites[major]&(1<<minor) != 0 {
			if minor++; minor == 8 {
				minor = 0
				if major++; major == b {
					break outer
				}
			}
			i++
		}

		// phase 2: set bits for all numbers which are multiples of 'i'.
		for j := uint(i << 1); j <= max; j += i {
			composites[j>>3] |= 1 << (j & 7)
		}

		if i++; i > max {
			break
		}
	}

	primes := []int{2}

	for i := uint(3); i <= max; i += 2 {
		if composites[i>>3]&(1<<(i&7)) == 0 {
			primes = append(primes, int(i))
		}
	}

	return primes
}

func primes10(max uint) []int {
	// Create a bit array to mark known composite numbers. After loop below, a 0
	// value means the corresponding natural number is prime, and a 1 value
	// means a composite number. Those defaults were chosen to take advantage of
	// the zero value for each byte in a slice, assuming all numbers are prime,
	// then mark off composite numbers.
	//
	// While it is not necessarily needed to store whether 0 and 1 are prime,
	// reserving bits for those values reduces the number of math instructions
	// required to set and check bits. Therefore, the 0th bit is for natural
	// number 0. The 1st bit is for natural number 1. And so
	// forth. Additionally, no attempt is ever made to set or check the bits
	// representing 0 or 1. While they are 0 at the end of the outer loop,
	// indicating they would be prime, they remain 0 simply because we never set
	// them at the beginning of this function.
	b, r := 1+max>>3, max&7
	if r > 0 {
		b++
	}
	composites := make([]byte, b)

	i := uint(2)
outer:
	for {
		// phase 1: find next prime, leaving it in 'i'.
		for composites[i>>3]&(1<<(i&7)) != 0 {
			if i++; i > max {
				break outer
			}
		}

		// phase 2: set bits for all numbers which are multiples of 'i'.
		for j := uint(i << 1); j <= max; j += i {
			composites[j>>3] |= 1 << (j & 7)
		}

		if i++; i > max {
			break
		}
	}

	primes := []int{2}

	for i := uint(3); i <= max; i += 2 {
		if composites[i>>3]&(1<<(i&7)) == 0 {
			primes = append(primes, int(i))
		}
	}

	return primes
}

func primes11(max uint) []int {
	// Create a bit array to mark known composite numbers. After loop below, a 0
	// value means the corresponding natural number is prime, and a 1 value
	// means a composite number. Those defaults were chosen to take advantage of
	// the zero value for each byte in a slice, assuming all numbers are prime,
	// then mark off composite numbers.
	//
	// While it is not necessarily needed to store whether 0 and 1 are prime,
	// reserving bits for those values reduces the number of math instructions
	// required to set and check bits. Therefore, the 0th bit is for natural
	// number 0. The 1st bit is for natural number 1. And so
	// forth. Additionally, no attempt is ever made to set or check the bits
	// representing 0 or 1. While they are 0 at the end of the outer loop,
	// indicating they would be prime, they remain 0 simply because we never set
	// them at the beginning of this function.
	b, r := 1+max>>3, max&7
	if r > 0 {
		b++
	}
	composites := make([]byte, b)

	for i := uint(2); i <= max; i++ {
		if composites[i>>3]&(1<<(i&7)) == 0 {
			// When i is prime; mark composites for all multiples of 'i'.
			for j := uint(i << 1); j <= max; j += i {
				composites[j>>3] |= 1 << (j & 7)
			}
		}
	}

	primes := []int{2}

	for i := uint(3); i <= max; i += 2 {
		if composites[i>>3]&(1<<(i&7)) == 0 {
			primes = append(primes, int(i))
		}
	}

	return primes
}

func primes12(max uint) []int {
	// Create a bit array to mark known composite numbers. After loop below, a 0
	// value means the corresponding natural number is prime, and a 1 value
	// means a composite number. Those defaults were chosen to take advantage of
	// the zero value for each byte in a slice, assuming all numbers are prime,
	// then mark off composite numbers.
	//
	// While it is not necessarily needed to store whether 0 and 1 are prime,
	// reserving bits for those values reduces the number of math instructions
	// required to set and check bits. Therefore, the 0th bit is for natural
	// number 0. The 1st bit is for natural number 1. And so
	// forth. Additionally, no attempt is ever made to set or check the bits
	// representing 0 or 1. While they are 0 at the end of the outer loop,
	// indicating they would be prime, they remain 0 simply because we never set
	// them at the beginning of this function.
	b, r := 1+max>>3, max&7
	if r > 0 {
		b++
	}
	composites := make([]byte, b)

	// Special case mark multiples of 2.
	for i := uint(4); i <= max; i += 2 {
		composites[i>>3] |= 1 << (i & 7)
	}

	// Starting at 3, check all odd numbers;
	for i := uint(3); i <= max; i += 2 {
		if composites[i>>3]&(1<<(i&7)) == 0 {
			// When i is prime; mark composites for all multiples of 'i'.
			for j := uint(i << 1); j <= max; j += i {
				composites[j>>3] |= 1 << (j & 7)
			}
		}
	}

	primes := []int{2}

	for i := uint(3); i <= max; i += 2 {
		if composites[i>>3]&(1<<(i&7)) == 0 {
			primes = append(primes, int(i))
		}
	}

	return primes
}

func primes13(max uint) []int {
	// Create a bit array to mark known composite numbers. After loop below, a 0
	// value means the corresponding natural number is prime, and a 1 value
	// means a composite number. Those defaults were chosen to take advantage of
	// the zero value for each byte in a slice, assuming all numbers are prime,
	// then mark off composite numbers.
	//
	// While it is not necessarily needed to store whether 0 and 1 are prime,
	// reserving bits for those values reduces the number of math instructions
	// required to set and check bits. Therefore, the 0th bit is for natural
	// number 0. The 1st bit is for natural number 1. And so
	// forth. Additionally, no attempt is ever made to set or check the bits
	// representing 0 or 1. While they are 0 at the end of the outer loop,
	// indicating they would be prime, they remain 0 simply because we never set
	// them at the beginning of this function.
	b, r := 1+max>>3, max&7
	if r > 0 {
		b++
	}
	composites := make([]byte, b)

	// Special case mark multiples of 2.
	for i := uint(4); i <= max; i += 2 {
		composites[i>>3] |= 1 << (i & 7)
	}
	primes := []int{2}

	// Starting at 3, check all odd numbers;
	for i := uint(3); i <= max; i += 2 {
		if composites[i>>3]&(1<<(i&7)) == 0 {
			primes = append(primes, int(i))
			// When i is prime; mark composites for all multiples of 'i'.
			for j := uint(i << 1); j <= max; j += i {
				composites[j>>3] |= 1 << (j & 7)
			}
		}
	}

	return primes
}

func primes14(max uint) []int {
	// Create a bit array to mark known composite numbers. After loop below, a 0
	// value means the corresponding natural number is prime, and a 1 value
	// means a composite number. Those defaults were chosen to take advantage of
	// the zero value for each byte in a slice, assuming all numbers are prime,
	// then mark off composite numbers.
	//
	// While it is not necessarily needed to store whether 0 and 1 are prime,
	// reserving bits for those values reduces the number of math instructions
	// required to set and check bits. Therefore, the 0th bit is for natural
	// number 0. The 1st bit is for natural number 1. And so
	// forth. Additionally, no attempt is ever made to set or check the bits
	// representing 0 or 1. While they are 0 at the end of the outer loop,
	// indicating they would be prime, they remain 0 simply because we never set
	// them at the beginning of this function.
	b, r := 1+max>>3, max&7
	if r > 0 {
		b++
	}
	composites := make([]byte, b)

	// Special case mark multiples of 2.
	for i := uint(0); i < b; i++ {
		composites[i] = 0x55
	}
	primes := []int{2}

	// Starting at 3, check all odd numbers
	for i := uint(3); i <= max; i += 2 {
		if composites[i>>3]&(1<<(i&7)) == 0 {
			primes = append(primes, int(i))
			// When i is prime; mark composites for all multiples of 'i'.
			for j := uint(i << 1); j <= max; j += i {
				composites[j>>3] |= 1 << (j & 7)
			}
		}
	}

	return primes
}

func primes15(max uint) []int {
	// Create a bit array to mark known composite numbers. After loop below, a 0
	// value means the corresponding natural number is prime, and a 1 value
	// means a composite number. Those defaults were chosen to take advantage of
	// the zero value for each byte in a slice, assuming all numbers are prime,
	// then mark off composite numbers.
	//
	// While it is not necessarily needed to store whether 0 and 1 are prime,
	// reserving bits for those values reduces the number of math instructions
	// required to set and check bits. Therefore, the 0th bit is for natural
	// number 0. The 1st bit is for natural number 1. And so
	// forth. Additionally, no attempt is ever made to set or check the bits
	// representing 0 or 1. While they are 0 at the end of the outer loop,
	// indicating they would be prime, they remain 0 simply because we never set
	// them at the beginning of this function.
	b := 1 + max>>6
	composites := make([]uint64, b)

	// Special case mark multiples of 2.
	for i := uint(0); i < b; i++ {
		composites[i] = 0x5555555555555555
	}
	primes := []int{2}

	// Starting at 3, check all odd numbers
	for i := uint(3); i <= max; i += 2 {
		if composites[i>>6]&(1<<(i&63)) == 0 {
			primes = append(primes, int(i))
			// When i is prime; mark composites for all multiples of 'i'.
			for j := uint(i << 1); j <= max; j += i {
				composites[j>>6] |= 1 << (j & 63)
			}
		}
	}

	return primes
}
