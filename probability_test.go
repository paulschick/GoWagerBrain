// Copyright (c) Paul Schick
// SPDX-License-Identifier: MPL-2.0

package GoWagerBrain

import "testing"

type testParams struct {
	Odds     float64
	Decimals int
	Expected float64
}

type runParams struct {
	t      *testing.T
	params []testParams
	fn     func(f float64, i int) float64
	fnName string
}

func (p *runParams) runImpliedProbTest() {
	for _, param := range p.params {
		result := p.fn(param.Odds, param.Decimals)
		if result != param.Expected {
			p.t.Errorf("%s(%f, %d) = %f; want %f", p.fnName, param.Odds, param.Decimals, result, param.Expected)
		}
	}
}

func TestDecimalImpliedWinProb(t *testing.T) {
	params := []testParams{
		{2.4, 4, 0.4167},
		{1.33, 2, 0.75},
	}
	p := runParams{t, params, DecimalImpliedWinProb, "DecimalImpliedWinProb"}
	p.runImpliedProbTest()
}

func TestAmericanImpliedWinProb(t *testing.T) {
	params := []testParams{
		{140, 4, 0.4167},
		{-300, 2, 0.75},
	}
	p := runParams{t, params, AmericanImpliedWinProb, "AmericanImpliedWinProb"}
	p.runImpliedProbTest()
}
