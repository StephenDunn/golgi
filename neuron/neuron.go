package neuron

import (
	"math/rand"

	"github.com/google/uuid"
)

type Neuron struct {
	id        uuid.UUID
	threshold int
	HasFired  bool
	axons     []axon
}

type axon struct {
	neuron   *Neuron
	strength int
}

func NewNeuron(thresholdMin int, thresholdMax int) *Neuron {
	threshold := rand.Intn(thresholdMax-thresholdMin) + thresholdMin
	neu := Neuron{threshold: threshold, id: uuid.New()}
	return &neu
}

func NewNeurons(amount int, thresholdMin int, thresholdMax int) []*Neuron {
	var neurons []*Neuron
	for i := 0; i < amount; i++ {
		neurons = append(neurons, NewNeuron(thresholdMin, thresholdMax))
	}

	return neurons
}

func SetUpNetwork(neurons []*Neuron, strengthMin int, strengthMax int) {
	for _, neu := range neurons {
		set := make(map[int]bool)
		var result []int
		for len(set) < rand.Intn(len(neurons)) {
			value := rand.Intn(len(neurons))
			if !set[value] {
				set[value] = true
				result = append(result, value)
			}
		}

		for _, ix := range result {
			strength := rand.Intn(strengthMax-strengthMin) + strengthMin
			neu.axons = append(neu.axons, axon{strength: strength, neuron: neurons[ix]})
		}
	}
}
