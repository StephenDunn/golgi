package main

import (
	"fmt"
	"golgi/cortex"
	_ "golgi/cortex"
)

func main() {
	fmt.Println("App started")
	cortex.Startup()
	// a := neuron.NewNeurons(1000, 10, 100, 0.25)
	// neuron.SetUpNetwork(a, 1, 100)
	// for i, neu := range a {
	// 	fmt.Printf("Index [%v] has fired [%v]\n", i, neu.HasFired)
	// }
}
