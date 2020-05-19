package polynomial

import "math"

type AlgebraicExpression interface {
	Evaluate(x float64) float64
	Clone() AlgebraicExpression
}

type Polynomial struct {
	Coefficients []float64
}

func (p *Polynomial) Add(p2 *Polynomial) {
	deg := p.Degree()
	for i, c := range p2.Coefficients {
		if i <= deg {
			p.Coefficients[i] += c
		} else {
			p.Coefficients = append(p.Coefficients, p2.Coefficients[i])
		}
	}

}

func (p *Polynomial) Sub(p2 *Polynomial) {
	deg := p.Degree()
	for i, c := range p2.Coefficients {
		if i <= deg {
			p.Coefficients[i] -= c
		} else {
			p.Coefficients = append(p.Coefficients, p2.Coefficients[i])
		}
	}
}

func (p *Polynomial) Multi(p2 *Polynomial) {
	deg1 := p.Degree()
	deg2 := p2.Degree()
	coef := make([]float64, deg1+deg2+1)
	for i, c1 := range p.Coefficients {
		for j, c2 := range p2.Coefficients {
			coef[i+j] += c1 * c2
		}
	}
	p.Coefficients = coef
}

func (p *Polynomial) Derivative() {
	coef := make([]float64, len(p.Coefficients)-1)
	for i, v := range p.Coefficients {
		if i != 0 {
			coef[i-1] = v * float64(i)
		}
	}
	p.Coefficients = coef
}

func (p *Polynomial) DerivativeNew() *Polynomial {
	c := p.Clone()
	c.Derivative()
	return c
}

func (p *Polynomial) Degree() int {
	return len(p.Coefficients) - 1
}

func (p *Polynomial) RootNewton(guess float64, threshhold float64, maxIterations int) (float64, bool) {
	der := p.DerivativeNew()
	xnew := guess - p.Evaluate(guess)/der.Evaluate(guess)
	for i := 0; i < maxIterations; i++ {
		xold := xnew
		xnew = xold - p.Evaluate(xold)/der.Evaluate(xold)
		if math.Abs(xnew-xold) < threshhold {
			return xnew, true
		}
	}
	return 0.0, false
}

func (p Polynomial) Clone() *Polynomial {
	coef := make([]float64, 0)
	for _, f := range p.Coefficients {
		coef = append(coef, f)
	}

	return &Polynomial{coef}
}

func (p Polynomial) Evaluate(x float64) float64 {
	v := 1.0
	sum := 0.0
	for _, c := range p.Coefficients {
		sum += v * c
		v = v * x
	}
	return sum
}
