package main

import (
	"fmt"
	"golgi/neuron"
)

func main() {
	a := neuron.NewNeurons(1000, 100)
	neuron.SetUpNetwork(a)
	for i, neu := range a {
		fmt.Printf("Index [%v] has fired [%v]\n", i, neu.HasFired)
	}
}
