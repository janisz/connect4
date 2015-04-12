package collect4

import (
	"testing"
)

func TestPerceptron(t *testing.T) {
	p := NewPerceptron(4, 1, []int{3,2})
	activationFunction = Tanh
	actual := p.Compute(NewSimpleVector([]float64{-1, 0.5, 0, 1}))
	expected := NewZeroVector(1)
	if !actual.Equals(expected) {
		t.Errorf("Expected %s but got %s for %s", expected, actual, p)
	}
}

func TestPerceptronComuteWithoutHiddenLayers(t *testing.T) {
	p := NewPerceptron(2, 1, []int{})
	activationFunction = Tanh
	p.layers[0].neurons[0].weights = NewSimpleVector([]float64{1, -0.04930600})
	actual := p.Compute(NewSimpleVector([]float64{0.5, -1}))
	expected := NewSimpleVector([]float64{0.5})
	if !actual.NearlyEquals(expected, 1E-6) {
		t.Errorf("Expected %s but got %s for %s", expected, actual, p)
	}
}

func TestPerceptronComuteLayers(t *testing.T) {
	p := NewPerceptron(2, 1, []int{1})
	activationFunction = Tanh
	p.layers[0].neurons[0].weights = NewSimpleVector([]float64{1, -0.04930600})
	p.layers[1].neurons[0].weights = NewSimpleVector([]float64{-0.200671})
	actual := p.Compute(NewSimpleVector([]float64{0.5, -1}))
	expected := NewSimpleVector([]float64{-0.1})
	if !actual.NearlyEquals(expected, 1E-6) {
		t.Errorf("Expected %s but got %s for %s", expected, actual, p)
	}
}

func TestPerceptronInitialization(t *testing.T) {
	p := NewPerceptron(4, 1, []int{3,2})
	activationFunction = Tanh
	p.Initialize(64)
	actual := p.Compute(NewSimpleVector([]float64{-1, 0.5, 0, 1}))
	expected := NewZeroVector(1)
	if actual.Equals(expected) {
		t.Errorf("NOT Expected %s but got %s for %s", expected, actual, p)
	}
}

func TestPerceptronBackpropagation(t *testing.T) {
	//given:
	activationFunction = Sigmoid
	activationFunctionʼ = Sigmoidʼ
	p := NewPerceptron(2, 1, []int{2})
	p.layers[0].neurons[0].weights = NewSimpleVector([]float64{0.1, 0.8})
	p.layers[0].neurons[1].weights = NewSimpleVector([]float64{0.4, 0.6})
	p.layers[1].neurons[0].weights = NewSimpleVector([]float64{0.3, 0.9})

	input := NewSimpleVector([]float64{0.35, 0.9})

	t.Log(p.Compute(input))

	expected := NewSimpleVector([]float64{0.5})

	//when:
	p.Learn([]Vector{input}, []Vector{expected})

	//then:
	actual := p.Compute(input)
	if !actual.NearlyEquals(expected, -0.18205) {
		t.Errorf("NOT Expected %s but got %s for %s", expected, actual, p)
	}
}
