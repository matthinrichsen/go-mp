package main

type Number struct {
	Zeros  bool
	Ones   int32
	Twos   int32
	Threes int32
	Fours  int32
	Fives  int32
	Sixes  int32
	Sevens int32
	Eights int32
	Nines  int32
}

func FromInt64(n int64) Number {
	num := Number{
		Zeros: n == 0,
	}

	for n > 0 {
		switch n % 10 {
		case 0:
			num.Zeros = true
		case 1:
			num.Ones++
		case 2:
			num.Twos++
		case 3:
			num.Threes++
		case 4:
			num.Fours++
		case 5:
			num.Fives++
		case 6:
			num.Sixes++
		case 7:
			num.Sevens++
		case 8:
			num.Eights++
		case 9:
			num.Nines++
		}

		n = n / 10
	}
	return num
}

func (n Number) ToInt() int64 {
	result := int64(0)
	for i := int32(0); i < n.Ones; i++ {
		result *= 10
		result += 1
	}

	for i := int32(0); i < n.Twos; i++ {
		result *= 10
		result += 2
	}

	for i := int32(0); i < n.Threes; i++ {
		result *= 10
		result += 3
	}

	for i := int32(0); i < n.Fours; i++ {
		result *= 10
		result += 4
	}

	for i := int32(0); i < n.Fives; i++ {
		result *= 10
		result += 5
	}

	for i := int32(0); i < n.Sixes; i++ {
		result *= 10
		result += 6
	}

	for i := int32(0); i < n.Sevens; i++ {
		result *= 10
		result += 7
	}

	for i := int32(0); i < n.Eights; i++ {
		result *= 10
		result += 8
	}

	for i := int32(0); i < n.Nines; i++ {
		result *= 10
		result += 9
	}
	return result
}

func (n Number) Add(digit int32) Number {
	switch digit {
	case 0:
		n.Zeros = true
	case 1:
		n.Ones++
	case 2:
		n.Twos++
	case 3:
		n.Threes++
	case 4:
		n.Fours++
	case 5:
		n.Fives++
	case 6:
		n.Sixes++
	case 7:
		n.Sevens++
	case 8:
		n.Eights++
	case 9:
		n.Nines++
	}
	return n
}

func (n Number) MP() int64 {
	depth := int64(0)

	for current := n; current.Length() > 1; depth++ {
		current = current.Next()
	}
	return depth
}

func (n Number) Next() Number {
	if n.Zeros {
		return FromInt64(0)
	}

	result := pow(2, n.Twos)
	result *= pow(3, n.Threes)
	result *= pow(4, n.Fours)
	result *= pow(5, n.Fives)
	result *= pow(6, n.Sixes)
	result *= pow(7, n.Sevens)
	result *= pow(8, n.Eights)
	result *= pow(9, n.Nines)
	return FromInt64(result)
}

func (n Number) Length() int64 {
	result := int64(n.Ones)
	result += int64(n.Twos)
	result += int64(n.Threes)
	result += int64(n.Fours)
	result += int64(n.Fives)
	result += int64(n.Sixes)
	result += int64(n.Sevens)
	result += int64(n.Eights)
	result += int64(n.Nines)

	if n.Zeros {
		if result > 0 {
			return 2
		}
		return 1
	}

	return result
}

func pow(n int64, count int32) int64 {
	if count == 0 || n == 1 {
		return 1
	}

	result := n
	for i := int32(1); i < count; i++ {
		result *= n
	}

	return result
}
