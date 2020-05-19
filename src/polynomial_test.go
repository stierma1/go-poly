package polynomial

import (
	"reflect"
	"testing"
)

func TestPolynomial_Add(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	type args struct {
		p2 *Polynomial
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			p.Add(tt.args.p2)
		})
	}
}

func TestPolynomial_Sub(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	type args struct {
		p2 *Polynomial
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			p.Sub(tt.args.p2)
		})
	}
}

func TestPolynomial_Multi(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	type args struct {
		p2 *Polynomial
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			p.Multi(tt.args.p2)
		})
	}
}

func TestPolynomial_Derivative(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			p.Derivative()
		})
	}
}

func TestPolynomial_DerivativeNew(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Polynomial
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			if got := p.DerivativeNew(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polynomial.DerivativeNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolynomial_Degree(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			if got := p.Degree(); got != tt.want {
				t.Errorf("Polynomial.Degree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolynomial_RootNewton(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	type args struct {
		guess         float64
		threshhold    float64
		maxIterations int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			got, got1 := p.RootNewton(tt.args.guess, tt.args.threshhold, tt.args.maxIterations)
			if got != tt.want {
				t.Errorf("Polynomial.RootNewton() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Polynomial.RootNewton() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPolynomial_Clone(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Polynomial
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			if got := p.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polynomial.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolynomial_Evaluate(t *testing.T) {
	type fields struct {
		Coefficients []float64
	}
	type args struct {
		x float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			"Test",
			fields{[]float64{1.0, 1.0}},
			args{1.0},
			2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Polynomial{
				Coefficients: tt.fields.Coefficients,
			}
			if got := p.Evaluate(tt.args.x); got != tt.want {
				t.Errorf("Polynomial.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
