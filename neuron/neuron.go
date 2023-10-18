package neuron

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Neuron struct {
	id                uuid.UUID
	threshold         int
	HasFired          bool
	axons             []axon
	receptor          chan int
	synapticThreshold int
	die               chan struct{}
	ticker            *time.Ticker
}

type axon struct {
	neuron   *Neuron
	strength int
}

func (n *Neuron) listen() {
	for {
		select {
		case val := <-n.receptor:
			if !n.HasFired {
				n.HasFired = true
				n.synapticThreshold += val
				if n.synapticThreshold >= n.threshold {
					for _, axon := range n.axons {
						axon.neuron.receptor <- axon.strength
					}
				}
			}
		case <-n.ticker.C:
			if n.synapticThreshold > 0 {
				n.synapticThreshold -= 1
			} else if n.synapticThreshold <= 0 {
				n.HasFired = false
			}

		case <-n.die:
			return
		}
	}
}

func (n *Neuron) kill() {
	close(n.die)
}

func NewNeuron(thresholdMin int, thresholdMax int) *Neuron {
	threshold := rand.Intn(thresholdMax-thresholdMin) + thresholdMin
	neu := Neuron{threshold: threshold, id: uuid.New(), receptor: make(chan int), die: make(chan struct{}), ticker: time.NewTicker(1 * time.Second)}

	go neu.listen()

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

func KillNetwork(neurons []*Neuron) {
	for _, neu := range neurons {
		neu.kill()
	}
}
