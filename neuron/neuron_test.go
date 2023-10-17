package neuron

import (
	"fmt"
	"testing"
)

func TestSetupNeurons_ShouldAssignDownstreamToAllNeurons(t *testing.T) {
	neurons := NewNeurons(1000, 10)
	SetUpNetwork(neurons)

	for i, neu := range neurons {
		if len(neu.downstreamNeurons) == 0 {
			fmt.Println(neu)
			t.Fatalf("Neuron with missing downstream. Index [%v]", i)
		}
	}
}
