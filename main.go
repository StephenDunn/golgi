package main

import (
	"fmt"
	"golgi/neuron"
)

func main() {
	a := neuron.NewNeurons(1000, 10, 100)
	neuron.SetUpNetwork(a, 1, 100)
	for i, neu := range a {
		fmt.Printf("Index [%v] has fired [%v]\n", i, neu.HasFired)
	}
}
