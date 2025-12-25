package test

import (
	helpers "lab2512lab1/pkg/helper"
	"lab2512lab1/pkg/insert"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	sourceOfTrue, arr := helpers.GenerateSortedAndShuffledArrays(100000)
	insert.InsertionSort(arr)
	if !reflect.DeepEqual(arr, sourceOfTrue) {
		t.Errorf("insert sort doesn't work, %v", arr)
	}
	t.Logf("Test insert %v", sourceOfTrue)
}
