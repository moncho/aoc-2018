package main

var instructionSet = []instruction{
	addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr,
}

type instruction func(o operation) func(r registers) registers

func addr(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA] + r[o.inputB]
		return r
	}
}

func addi(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA] + o.inputB
		return r
	}
}

func mulr(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA] * r[o.inputB]
		return r
	}
}

func muli(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA] * o.inputB
		return r
	}
}

func banr(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA] & r[o.inputB]
		return r
	}
}

func bani(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA] & o.inputB
		return r
	}
}

func borr(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA] | r[o.inputB]
		return r
	}
}

func bori(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA] | o.inputB
		return r
	}
}

func setr(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = r[o.inputA]
		return r
	}
}

func seti(o operation) func(r registers) registers {
	return func(r registers) registers {
		r[o.output] = o.inputA
		return r
	}
}

func gtir(o operation) func(r registers) registers {
	return func(r registers) registers {
		if o.inputA > r[o.inputB] {
			r[o.output] = 1
		} else {
			r[o.output] = 0
		}
		return r
	}
}

func gtri(o operation) func(r registers) registers {
	return func(r registers) registers {
		if r[o.inputA] > o.inputB {
			r[o.output] = 1
		} else {
			r[o.output] = 0
		}
		return r
	}
}
func gtrr(o operation) func(r registers) registers {
	return func(r registers) registers {
		if r[o.inputA] > r[o.inputB] {
			r[o.output] = 1
		} else {
			r[o.output] = 0
		}
		return r
	}
}

func eqir(o operation) func(r registers) registers {
	return func(r registers) registers {
		if o.inputA == r[o.inputB] {
			r[o.output] = 1
		} else {
			r[o.output] = 0
		}
		return r
	}
}

func eqri(o operation) func(r registers) registers {
	return func(r registers) registers {
		if r[o.inputA] == o.inputB {
			r[o.output] = 1
		} else {
			r[o.output] = 0
		}
		return r
	}
}
func eqrr(o operation) func(r registers) registers {
	return func(r registers) registers {
		if r[o.inputA] == r[o.inputB] {
			r[o.output] = 1
		} else {
			r[o.output] = 0
		}
		return r
	}
}
