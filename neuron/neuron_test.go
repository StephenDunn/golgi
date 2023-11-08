package neuron

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

var totalNeurons = 100000
var thresholdMin = 10
var thresholdMax = 100
var strengthMin = 1
var strengthMax = 5
var refactory = 0.25

func TestSetupNeurons_ShouldAssignDownstreamToAllNeurons(t *testing.T) {
	neurons := NewNeurons(totalNeurons, WithThreshold(thresholdMin, thresholdMax), WithRefactory(refactory))
	SetUpNetwork(neurons, strengthMin, strengthMax)

	count := 0
	for _, neu := range neurons {

		if len(neu.axons) == 0 {
			count++
		}
	}

	if count > totalNeurons/10 {
		t.Fatalf("Too many neurons with missing downstream. Total  neurons with missing downstream: %v\n", count)
	} else {
		fmt.Printf("Total neurons with missing downstream: %v\n", count)
	}

	KillNetwork(neurons)
}

func TestMinAndMaxThreshold_ShouldHaveNoNeuronsWithThresholdOutsideMinAndMax(t *testing.T) {
	neurons := NewNeurons(totalNeurons, WithThreshold(thresholdMin, thresholdMax), WithRefactory(refactory))
	thresholds := []int{}

	for _, neu := range neurons {
		thresholds = append(thresholds, neu.threshold)
		if neu.threshold < thresholdMin || neu.threshold > thresholdMax {
			t.Fatalf("Neuron with threshold outside range: %v\n", neu)
		}
	}

	fmt.Printf("Average threshold of neurons: %v\n", averageOfSlice(thresholds))
	KillNetwork(neurons)
}

func TestMinAndMaxStrength_ShouldHaveNoNeuronsWithStrengthOutsideMinAndMax(t *testing.T) {
	neurons := NewNeurons(totalNeurons, WithThreshold(thresholdMin, thresholdMax), WithRefactory(refactory))
	SetUpNetwork(neurons, strengthMin, strengthMax)
	strengths := []int{}

	for _, neu := range neurons {
		for _, axo := range neu.axons {
			strengths = append(strengths, axo.strength)

			if axo.strength < strengthMin || axo.strength > strengthMax {
				t.Fatalf("Axon with strength outside range: %v\n", neu)
			}
		}
	}

	fmt.Printf("Average strength of axons: %v\n", averageOfSlice(strengths))
	KillNetwork(neurons)
}

func TestNeuronFiring_WhenOneNeuronFiresAnotherShouldFire(t *testing.T) {
	neurons := NewNeurons(totalNeurons, WithThreshold(thresholdMin, thresholdMax), WithRefactory(refactory))
	SetUpNetwork(neurons, strengthMin, strengthMax)

	for i := 0; i < 100; i++ {
		neurons[i].receptor <- 1000
	}

	time.Sleep(50 * time.Millisecond)

	triggeredNeurons := []uuid.UUID{}
	for _, neu := range neurons {
		if neu.HasFired {
			triggeredNeurons = append(triggeredNeurons, neu.id)
		}
	}

	if len(triggeredNeurons) == 0 {
		t.Fatalf("No neurons fired\n")
	} else {
		fmt.Printf("Neurons fired: %v\n", len(triggeredNeurons))
	}
	KillNetwork(neurons)
}

func TestNeuronTicker_RecievesTickSynapticThresholdUpdated(t *testing.T) {
	neuron := NewNeuron(WithThreshold(thresholdMin, thresholdMax), WithRefactory(refactory))
	impulse := 40

	neuron.recieveSynapticImpulse(impulse)
	neuron.recieveTicker()

	if neuron.synapticThreshold != 10 {
		t.Fatalf("Synaptic threshold decay incorrect\n")
	} else {
		fmt.Printf("Impulse applied: %v. 1 decay applied at: %v. Synaptic threshold: %v. Correct decay applied. \n", impulse, refactory, neuron.synapticThreshold)
	}
}

func TestNeuronTicker_RecievesTickSynapticHitsZero(t *testing.T) {
	neuron := NewNeuron(WithThreshold(thresholdMin, thresholdMax), WithRefactory(refactory))
	impulse := 10

	neuron.recieveSynapticImpulse(impulse)
	neuron.recieveTicker()
	neuron.recieveTicker()

	if neuron.synapticThreshold != 0 {
		t.Fatalf("Synaptic threshold does not reach zero\n")
	} else {
		fmt.Printf("Impulse applied: %v. 2 decays applied at: %v. Synaptic threshold: %v. Correct decay applied. \n", impulse, refactory, neuron.synapticThreshold)
	}
}

func TestNeuronSynapse_RecievesSynapticTrigger_SynapitcThresholdIncreased(t *testing.T) {
	neuron := NewNeuron(WithThreshold(thresholdMin, thresholdMax), WithRefactory(refactory))
	impulse := 4

	neuron.recieveSynapticImpulse(impulse)

	if neuron.synapticThreshold != impulse {
		t.Fatalf("Synaptic threshold: %v does not equal impulse: %v\n", neuron.synapticThreshold, impulse)
	} else {
		fmt.Printf("Synaptic threshold correctly incrimented by impulse: %v. Symaptic threshold: %v. ", impulse, neuron.synapticThreshold)
	}

	neuron.recieveSynapticImpulse(impulse)

	if neuron.synapticThreshold != impulse*2 {
		t.Fatalf("Synaptic threshold: %v does not equal impulse: %v\n", neuron.synapticThreshold, impulse*2)
	} else {
		fmt.Printf("Synaptic threshold correctly incrimented by impulse: %v. Symaptic threshold: %v\n", impulse, neuron.synapticThreshold)
	}
}

func averageOfSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}

	return sum / len(slice)
}
