package collect4

type Neuron struct {
	weights Vector
	output, gradient, lastGradient float64
}

func NewNeuron(input int) Neuron {
	return Neuron{
		weights: NewSimpleVector(make([]float64, input)),
	}
}

func (n *Neuron) Compute(signal Vector) float64 {
	output := signal.Copy()
	output.MulElements(n.weights)
	n.output = output.Sum()
	return n.output
}