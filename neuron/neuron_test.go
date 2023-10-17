package neuron

import (
	"fmt"
	"testing"
)

var totalNeurons = 10000
var thresholdMin = 10
var thresholdMax = 100
var strengthMin = 1
var strengthMax = 50

func TestSetupNeurons_ShouldAssignDownstreamToAllNeurons(t *testing.T) {
	neurons := NewNeurons(totalNeurons, thresholdMin, thresholdMax)
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
}

func TestMinAndMaxThreshold_ShouldHaveNoNeuronsWithThresholdOutsideMinAndMax(t *testing.T) {
	neurons := NewNeurons(totalNeurons, thresholdMin, thresholdMax)
	thresholds := []int{}

	for _, neu := range neurons {
		thresholds = append(thresholds, neu.threshold)
		if neu.threshold < thresholdMin || neu.threshold > thresholdMax {
			t.Fatalf("Neuron with threshold outside range: %v\n", neu)
		}
	}

	fmt.Printf("Average threshold of neurons: %v\n", averageOfSlice(thresholds))
}

func TestMinAndMaxStrength_ShouldHaveNoNeuronsWithStrengthOutsideMinAndMax(t *testing.T) {
	neurons := NewNeurons(totalNeurons, thresholdMin, thresholdMax)
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
}

func averageOfSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}

	return sum / len(slice)
}
