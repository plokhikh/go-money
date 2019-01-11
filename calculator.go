package money

type calculator struct{}

func (c *calculator) add(a, b *Amount) *Amount {
	return &Amount{a.val + b.val}
}

func (c *calculator) subtract(a, b *Amount) *Amount {
	return &Amount{a.val - b.val}
}

func (c *calculator) multiply(a *Amount, m int64) *Amount {
	return &Amount{a.val * m}
}

func (c *calculator) divide(a *Amount, d int64) *Amount {
	return &Amount{a.val / d}
}

func (c *calculator) modulus(a *Amount, d int64) *Amount {
	return &Amount{a.val % d}
}

func (c *calculator) allocate(a *Amount, r, s int) *Amount {
	return &Amount{a.val * int64(r) / int64(s)}
}

func (c *calculator) absolute(a *Amount) *Amount {
	if a.val < 0 {
		return &Amount{-a.val}
	}

	return &Amount{a.val}
}

func (c *calculator) negative(a *Amount) *Amount {
	if a.val > 0 {
		return &Amount{-a.val}
	}

	return &Amount{a.val}
}
