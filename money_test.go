package money

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	m := New(1, "EUR")

	if m.amount.val != 100 {
		t.Errorf("Expected %d got %d", 100, m.amount.val)
	}

	if m.currency.Code != "EUR" {
		t.Errorf("Expected currency %s got %s", "EUR", m.currency.Code)
	}

	m = New(-100, "EUR")

	if m.amount.val != -10000 {
		t.Errorf("Expected %d got %d", -10000, m.amount.val)
	}
}

func TestCurrency(t *testing.T) {
	code := "MOCK"
	decimals := 5
	AddCurrency(code, "M$", "1 $", ".", ",", decimals)
	m := New(1, code)
	c := m.Currency().Code
	if c != code {
		t.Errorf("Expected %s got %s", code, c)
	}
	f := m.Currency().Fraction
	if f != decimals {
		t.Errorf("Expected %d got %d", decimals, f)
	}
}

func TestMoney_SameCurrency(t *testing.T) {
	m := New(0, "EUR")
	om := New(0, "USD")

	if m.SameCurrency(om) {
		t.Errorf("Expected %s not to be same as %s", m.currency.Code, om.currency.Code)
	}

	om = New(0, "EUR")

	if !m.SameCurrency(om) {
		t.Errorf("Expected %s to be same as %s", m.currency.Code, om.currency.Code)
	}
}

func TestMoney_Equals(t *testing.T) {
	m := New(0, "EUR")
	tcs := []struct {
		amount   float64
		expected bool
	}{
		{-1, false},
		{0, true},
		{1, false},
	}

	for _, tc := range tcs {
		om := New(tc.amount, "EUR")
		r, err := m.Equals(om)

		if err != nil || r != tc.expected {
			t.Errorf("Expected %d Equals %d == %t got %t", m.amount.val,
				om.amount.val, tc.expected, r)
		}
	}
}

func TestMoney_GreaterThan(t *testing.T) {
	m := New(0, "EUR")
	tcs := []struct {
		amount   float64
		expected bool
	}{
		{-1, true},
		{0, false},
		{1, false},
	}

	for _, tc := range tcs {
		om := New(tc.amount, "EUR")
		r, err := m.GreaterThan(om)

		if err != nil || r != tc.expected {
			t.Errorf("Expected %d Greater Than %d == %t got %t", m.amount.val,
				om.amount.val, tc.expected, r)
		}
	}
}

func TestMoney_GreaterThanOrEqual(t *testing.T) {
	m := New(0, "EUR")
	tcs := []struct {
		amount   float64
		expected bool
	}{
		{-1, true},
		{0, true},
		{1, false},
	}

	for _, tc := range tcs {
		om := New(tc.amount, "EUR")
		r, err := m.GreaterThanOrEqual(om)

		if err != nil || r != tc.expected {
			t.Errorf("Expected %d Equals Or Greater Than %d == %t got %t", m.amount.val,
				om.amount.val, tc.expected, r)
		}
	}
}

func TestMoney_LessThan(t *testing.T) {
	m := New(0, "EUR")
	tcs := []struct {
		amount   float64
		expected bool
	}{
		{-1, false},
		{0, false},
		{1, true},
	}

	for _, tc := range tcs {
		om := New(tc.amount, "EUR")
		r, err := m.LessThan(om)

		if err != nil || r != tc.expected {
			t.Errorf("Expected %d Less Than %d == %t got %t", m.amount.val,
				om.amount.val, tc.expected, r)
		}
	}
}

func TestMoney_LessThanOrEqual(t *testing.T) {
	m := New(0, "EUR")
	tcs := []struct {
		amount   float64
		expected bool
	}{
		{-1, false},
		{0, true},
		{1, true},
	}

	for _, tc := range tcs {
		om := New(tc.amount, "EUR")
		r, err := m.LessThanOrEqual(om)

		if err != nil || r != tc.expected {
			t.Errorf("Expected %d Equal Or Less Than %d == %t got %t", m.amount.val,
				om.amount.val, tc.expected, r)
		}
	}
}

func TestMoney_IsZero(t *testing.T) {
	tcs := []struct {
		amount   float64
		expected bool
	}{
		{-1, false},
		{0, true},
		{1, false},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		r := m.IsZero()

		if r != tc.expected {
			t.Errorf("Expected %d to be zero == %t got %t", m.amount.val, tc.expected, r)
		}
	}
}

func TestMoney_IsNegative(t *testing.T) {
	tcs := []struct {
		amount   float64
		expected bool
	}{
		{-1, true},
		{0, false},
		{1, false},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		r := m.IsNegative()

		if r != tc.expected {
			t.Errorf("Expected %d to be negative == %t got %t", m.amount.val,
				tc.expected, r)
		}
	}
}

func TestMoney_IsPositive(t *testing.T) {
	tcs := []struct {
		amount   float64
		expected bool
	}{
		{-1, false},
		{0, false},
		{1, true},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		r := m.IsPositive()

		if r != tc.expected {
			t.Errorf("Expected %d to be positive == %t got %t", m.amount.val,
				tc.expected, r)
		}
	}
}

func TestMoney_Absolute(t *testing.T) {
	tcs := []struct {
		amount   float64
		expected int64
	}{
		{-1, 100},
		{0, 0},
		{1, 100},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		r := m.Absolute().amount.val

		if r != tc.expected {
			t.Errorf("Expected absolute %d to be %d got %d", m.amount.val,
				tc.expected, r)
		}
	}
}

func TestMoney_Negative(t *testing.T) {
	tcs := []struct {
		amount   float64
		expected int64
	}{
		{-1, -100},
		{0, -0},
		{1, -100},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		r := m.Negative().amount.val

		if r != tc.expected {
			t.Errorf("Expected absolute %d to be %d got %d", m.amount.val,
				tc.expected, r)
		}
	}
}

func TestMoney_Add(t *testing.T) {
	tcs := []struct {
		amount1  float64
		amount2  float64
		expected float64
	}{
		{5, 5, 10},
		{10, 5, 15},
		{1, -1, 0},
	}

	for _, tc := range tcs {
		m := New(tc.amount1, "EUR")
		om := New(tc.amount2, "EUR")
		r, err := m.Add(om)

		if err != nil {
			t.Error(err)
		}

		if r.Amount() != tc.expected {
			t.Errorf("Expected %f + %f = %f got %d", tc.amount1, tc.amount2,
				tc.expected, r.amount.val)
		}
	}

}

func TestMoney_Add2(t *testing.T) {
	m := New(100, "EUR")
	dm := New(100, "GBP")
	r, err := m.Add(dm)

	if r != nil || err == nil {
		t.Error("Expected err")
	}
}

func TestMoney_Subtract(t *testing.T) {
	tcs := []struct {
		amount1  float64
		amount2  float64
		expected int64
	}{
		{5, 5, 0},
		{10, 5, 500},
		{1, -1, 200},
	}

	for _, tc := range tcs {
		m := New(tc.amount1, "EUR")
		om := New(tc.amount2, "EUR")
		r, err := m.Subtract(om)

		if err != nil {
			t.Error(err)
		}

		if r.amount.val != tc.expected {
			t.Errorf("Expected %f - %f = %d got %d", tc.amount1, tc.amount2,
				tc.expected, r.amount.val)
		}
	}
}

func TestMoney_Subtract2(t *testing.T) {
	m := New(100, "EUR")
	dm := New(100, "GBP")
	r, err := m.Subtract(dm)

	if r != nil || err == nil {
		t.Error("Expected err")
	}
}

func TestMoney_Multiply(t *testing.T) {
	tcs := []struct {
		amount     float64
		multiplier int64
		expected   int64
	}{
		{5, 5, 2500},
		{10, 5, 5000},
		{1, -1, -100},
		{1, 0, 0},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		r := m.Multiply(tc.multiplier).amount.val

		if r != tc.expected {
			t.Errorf("Expected %f * %d = %d got %d", tc.amount, tc.multiplier, tc.expected, r)
		}
	}
}

func TestMoney_Divide(t *testing.T) {
	tcs := []struct {
		amount   float64
		divisor  int64
		expected int64
	}{
		{5, 5, 100},
		{10, 5, 200},
		{1, -1, -100},
		{10, 3, 333},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		r := m.Divide(tc.divisor).amount.val

		if r != tc.expected {
			t.Errorf("Expected %f / %d = %d got %d", tc.amount, tc.divisor, tc.expected, r)
		}
	}
}

func TestMoney_Split(t *testing.T) {
	tcs := []struct {
		amount   float64
		split    int
		expected []int64
	}{
		{100, 3, []int64{3334, 3333, 3333}},
		{100, 4, []int64{2500, 2500, 2500, 2500}},
		{5, 3, []int64{167, 167, 166}},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		var rs []int64
		split, _ := m.Split(tc.split)

		for _, party := range split {
			rs = append(rs, party.amount.val)
		}

		if !reflect.DeepEqual(tc.expected, rs) {
			t.Errorf("Expected split of %f to be %v got %v", tc.amount, tc.expected, rs)
		}
	}
}

func TestMoney_Split2(t *testing.T) {
	m := New(100, "EUR")
	r, err := m.Split(-10)

	if r != nil || err == nil {
		t.Error("Expected err")
	}
}

func TestMoney_Allocate(t *testing.T) {
	tcs := []struct {
		amount   float64
		ratios   []int
		expected []int64
	}{
		{100, []int{50, 50}, []int64{5000, 5000}},
		{100, []int{30, 30, 30}, []int64{3334, 3333, 3333}},
		{200, []int{25, 25, 50}, []int64{5000, 5000, 10000}},
		{5, []int{50, 25, 25}, []int64{250, 125, 125}},
	}

	for _, tc := range tcs {
		m := New(tc.amount, "EUR")
		var rs []int64
		split, _ := m.Allocate(tc.ratios...)

		for _, party := range split {
			rs = append(rs, party.amount.val)
		}

		if !reflect.DeepEqual(tc.expected, rs) {
			t.Errorf("Expected allocation of %f for ratios %v to be %v got %v", tc.amount, tc.ratios,
				tc.expected, rs)
		}
	}
}

func TestMoney_Allocate2(t *testing.T) {
	m := New(100, "EUR")
	r, err := m.Allocate()

	if r != nil || err == nil {
		t.Error("Expected err")
	}
}

func TestMoney_Chain(t *testing.T) {
	m := New(10, "EUR")
	om := New(5, "EUR")
	// 10 + 5 = 15 / 5 = 3 * 4 = 12 - 5 = 7
	e := int64(7)

	m, err := m.Add(om)

	if err != nil {
		t.Error(err)
	}

	m = m.Divide(5).Multiply(4)
	m, err = m.Subtract(om)

	if err != nil {
		t.Error(err)
	}

	if m.amount.val != int64(700) {
		t.Errorf("Expected %d got %d", e, m.amount.val)
	}
}

func TestMoney_Format(t *testing.T) {
	tcs := []struct {
		amount   float64
		code     string
		expected string
	}{
		{1, "GBP", "£1.00"},
		{1.01, "GBP", "£1.01"},
		{0.99, "GBP", "£0.99"},
	}

	for _, tc := range tcs {
		m := New(tc.amount, tc.code)
		r := m.Display()

		if r != tc.expected {
			t.Errorf("Expected formatted %f to be %s got %s", tc.amount, tc.expected, r)
		}
	}

}

func TestMoney_Display(t *testing.T) {
	tcs := []struct {
		amount   float64
		code     string
		expected string
	}{
		{1, "AED", "1.00 .\u062f.\u0625"},
		{0.01, "USD", "$0.01"},
	}

	for _, tc := range tcs {
		m := New(tc.amount, tc.code)
		r := m.Display()

		if r != tc.expected {
			t.Errorf("Expected formatted %f to be %s got %s", tc.amount, tc.expected, r)
		}
	}
}

func TestMoney_Allocate3(t *testing.T) {
	pound := New(1, "GBP")
	parties, err := pound.Allocate(33, 33, 33)

	if err != nil {
		t.Error(err)
	}

	if parties[0].Display() != "£0.34" {
		t.Errorf("Expected %s got %s", "£0.34", parties[0].Display())
	}

	if parties[1].Display() != "£0.33" {
		t.Errorf("Expected %s got %s", "£0.33", parties[1].Display())
	}

	if parties[2].Display() != "£0.33" {
		t.Errorf("Expected %s got %s", "£0.33", parties[2].Display())
	}
}

func TestMoney_Comparison(t *testing.T) {
	pound := New(100, "GBP")
	twoPounds := New(200, "GBP")
	twoEuros := New(200, "EUR")

	if r, err := pound.GreaterThan(twoPounds); err != nil || r {
		t.Errorf("Expected %d Greater Than %d == %t got %t", pound.amount.val,
			twoPounds.amount.val, false, r)
	}

	if r, err := pound.LessThan(twoPounds); err != nil || !r {
		t.Errorf("Expected %d Less Than %d == %t got %t", pound.amount.val,
			twoPounds.amount.val, true, r)
	}

	if r, err := pound.LessThan(twoEuros); err == nil || r {
		t.Error("Expected err")
	}

	if r, err := pound.GreaterThan(twoEuros); err == nil || r {
		t.Error("Expected err")
	}

	if r, err := pound.Equals(twoEuros); err == nil || r {
		t.Error("Expected err")
	}

	if r, err := pound.LessThanOrEqual(twoEuros); err == nil || r {
		t.Error("Expected err")
	}

	if r, err := pound.GreaterThanOrEqual(twoEuros); err == nil || r {
		t.Error("Expected err")
	}
}

func TestMoney_Currency(t *testing.T) {
	pound := New(100, "GBP")

	if pound.Currency().Code != "GBP" {
		t.Errorf("Expected %s got %s", "GBP", pound.Currency().Code)
	}
}

func TestMoney_Amount(t *testing.T) {
	pound := New(100, "GBP")

	if pound.Amount() != 100 {
		t.Errorf("Expected %d got %f", 100, pound.Amount())
	}
}
