package pkg

import "testing"
import "reflect"

func TestInsertionSort_1_3_2(t *testing.T) {
	doTest(
		[]float64{1., 3., 2.},
		[]float64{1., 2., 3.},
		InsertionSort, t)
}

func TestInsertionSort_2_3_1_2(t *testing.T) {
	doTest(
		[]float64{2., 3., 1., 2.},
		[]float64{1., 2., 2., 3.},
		InsertionSort, t)
}

func TestInsertionSort_1_6_3_4_5_2(t *testing.T) {
	doTest(
		[]float64{1., 6., 3., 4., 5., 2.},
		[]float64{1., 2., 3., 4., 5., 6.},
		InsertionSort, t)

}

func doTest(input, expected []float64,
	sorting_function func([]float64) []float64,
	t *testing.T) {
	t_copy := make([]float64, len(input))
	copy(t_copy, input)
	actual := InsertionSort(t_copy)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(`
	InsertionSort: 
		input 		= %v	
		expected	= %v 
		actual 		= %v`, input, expected, actual)
	}
}
