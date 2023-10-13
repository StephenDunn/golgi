package neuron

import (
	"math/rand"
)

type Neuron struct {
	threshold         int
	HasFired          bool
	downstreamNeurons []*Neuron
}

func NewNeuron(threshold int) *Neuron {
	neu := Neuron{threshold: threshold}
	return &neu
}

func NewNeurons(amount int, threshold int) []Neuron {
	var neurons []Neuron
	for i := 0; i < amount; i++ {
		neurons = append(neurons, *NewNeuron(threshold))
	}

	return neurons
}

func SetUpNetwork(neurons *[]Neuron) {
	for _, neu := range *neurons {
		set := make(map[int]bool)
		var result []int
		for len(set) < rand.Intn(len(*neurons)) {
			value := rand.Intn(len(*neurons))
			if !set[value] {
				set[value] = true
				result = append(result, value)
			}
		}

		for _, ix := range result {
			neu.downstreamNeurons = append(neu.downstreamNeurons, &((*neurons)[ix]))
		}
	}
}
