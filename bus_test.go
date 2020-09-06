package events

import (
	"reflect"
	"testing"
)

func TestNew_Lifecycle(t *testing.T) {
	receives := make([][]int, 0)

	d := New()

	d.On("event", func(x int, y int) { receives = append(receives, []int{1, x, y}) })
	d.On("event", func(x int) { receives = append(receives, []int{2, x}) })
	d.On("event", func(x int, y string, z string) { receives = append(receives, []int{3, x, len(y)}) })

	d.Fire("event", 123, 123)
	d.Fire("event", 234, 234)
	d.Fire("event", "string", "value")

	assertEquals(t, [][]int{
		{1, 123, 123},
		{2, 123},
		{3, 123, 0},
		{1, 234, 234},
		{2, 234},
		{3, 234, 0},
		{1, 0, 0},
		{2, 0},
		{3, 0, 5},
	}, receives)
}

func assertEquals(t testing.TB, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}