package money

import (
	"fmt"
	"math"
	"strings"
)

// Formatter stores Money formatting information
type Formatter struct {
	Fraction int
	Decimal  string
	Thousand string
	Grapheme string
	Template string
}

// NewFormatter creates new Formatter instance
func NewFormatter(fraction int, decimal, thousand, grapheme, template string) *Formatter {
	return &Formatter{
		Fraction: fraction,
		Decimal:  decimal,
		Thousand: thousand,
		Grapheme: grapheme,
		Template: template,
	}
}

// Format returns string of formatted integer using given currency template
func (f *Formatter) Format(amount int64) string {
	var mant string

	// Work with absolute amount value
	template := fmt.Sprintf("%%.%df", f.Fraction)
	sas := strings.Split(fmt.Sprintf(template, float64(amount)/math.Pow10(f.Fraction)), ".")
	exp := sas[0]

	if len(sas) > 1 {
		mant = sas[1]
	} else {
		mant = ""
	}

	if exp[0:1] == "-" {
		exp = exp[1:]
	}

	if f.Thousand != "" {
		for i := len(exp) - 3; i > 0; i -= 3 {
			exp = exp[:i] + f.Thousand + exp[i:]
		}
	}

	sa := exp
	if len(mant) > 0 {
		sa += "." + mant
	}
	sa = strings.Replace(f.Template, "1", sa, 1)
	sa = strings.Replace(sa, "$", f.Grapheme, 1)

	// Add minus sign for negative amount
	if amount < 0 {
		sa = "-" + sa
	}

	return sa
}

// abs return absolute value of given integer
func (f Formatter) abs(amount int64) int64 {
	if amount < 0 {
		return -amount
	}

	return amount
}
