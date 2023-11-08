package neuron

import (
	"math"
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
	refactory         float64
}

type axon struct {
	neuron   *Neuron
	strength int
}

func (n *Neuron) listen() {
	for {
		select {
		case val := <-n.receptor:
			n.recieveSynapticImpulse(val)
		case <-n.ticker.C:
			n.recieveTicker()
		case <-n.die:
			return
		}
	}
}

func (n *Neuron) recieveSynapticImpulse(val int) {
	if !n.HasFired {
		n.synapticThreshold += val
		if n.synapticThreshold >= n.threshold {
			n.HasFired = true
			for _, axon := range n.axons {
				axon.neuron.receptor <- axon.strength
			}
		}
	}
}

func (n *Neuron) recieveTicker() {
	if n.synapticThreshold > 0 {
		n.synapticThreshold = int(math.Floor(float64(n.synapticThreshold) * n.refactory))
	} else if n.synapticThreshold <= 0 {
		n.HasFired = false
	}
}

func (n *Neuron) kill() {
	close(n.die)
}

type optFunc func(*Neuron)

func WithThreshold(thresholdMin int, thresholdMax int) optFunc {
	return func(n *Neuron) {
		n.threshold = rand.Intn(thresholdMax-thresholdMin) + thresholdMin
	}
}

func WithRefactory(refactory float64) optFunc {
	return func(n *Neuron) {
		n.refactory = refactory
	}
}

func NewNeuron(opts ...optFunc) *Neuron {
	neuron := Neuron{threshold: 0, id: uuid.New(), receptor: make(chan int), die: make(chan struct{}), ticker: time.NewTicker(1 * time.Second), refactory: 0}
	for _, fn := range opts {
		fn(&neuron)
	}

	go neuron.listen()

	return &neuron
}

func NewNeurons(amount int, opts ...optFunc) []*Neuron {
	var neurons []*Neuron
	for i := 0; i < amount; i++ {
		neurons = append(neurons, NewNeuron(opts...))
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
