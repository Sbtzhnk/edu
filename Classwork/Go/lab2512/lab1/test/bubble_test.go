package test

import (
	"lab2512lab1/pkg/bubble"
	helpers "lab2512lab1/pkg/helper"
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	sourceOfTrue, arr := helpers.GenerateSortedAndShuffledArrays(100000)
	bubble.BubbleSort(arr)

	if !reflect.DeepEqual(arr, sourceOfTrue) {
		t.Errorf("bubble sort doesn't work, %v", arr)
	}
	t.Logf("Test bubble %v", sourceOfTrue)
}
