package neuron

import (
	"fmt"
	"testing"
)

var totalNeurons = 1000
var threshold = 10

func TestSetupNeurons_ShouldAssignDownstreamToAllNeurons(t *testing.T) {
	neurons := NewNeurons(totalNeurons, threshold)
	SetUpNetwork(neurons)

	count := 0
	for _, neu := range neurons {

		if len(neu.downstreamNeurons) == 0 {
			count++
		}

	}

	if count > totalNeurons/10 {
		t.Fatalf("Too many neurons with missing downstream. Total  neurons with missing downstream: %v\n", count)
	} else {
		fmt.Printf("Total neurons with missing downstream: %v\n", count)
	}
}
