package pandigital

type Status int

const (
	Indeterminate Status = iota
	Invalid
	Valid
)

type Pandigital struct {
	digits       []uint
	countByDigit [10]uint
}

func (p *Pandigital) Uint64() uint64 {
	num := uint64(0)
	radix := uint64(1)
	for len(p.digits) > 0 {
		num += uint64(p.digits[len(p.digits)-1]) * radix
		radix *= 10
		p.digits = p.digits[0 : len(p.digits)-1]
	}
	return num
}

func (p *Pandigital) AddDigits(x uint) {
	digits := make([]uint, 0)
	for {
		digit := x % 10
		digits = append(digits, digit)
		p.countByDigit[digit]++
		x /= 10
		if x == 0 {
			break
		}
	}
	// put back into most significant digit first order
	for i := len(digits) - 1; i >= 0; i-- {
		p.digits = append(p.digits, digits[i])
	}
}

func (p *Pandigital) Vet(min, max uint) Status {
	if max > 9 {
		panic("bad max")
	}

	// digits other than [min,max]?
	for digit, count := range p.countByDigit {
		if count > 0 && (uint(digit) < min || uint(digit) > max) {
			return Invalid
		}
	}
	// ensure min..max occur exactly once
	status := Valid
	for i := min; i <= max; i++ {
		count := p.countByDigit[i]
		if count == 0 {
			// future AddDigits() call might create valid Pandigital
			status = Indeterminate

		} else if count > 1 {
			return Invalid
		}
	}
	return status
}
