package slice_test 
import (
	"testing"
	"fmt"
)
func TestSliceInit(t *testing.T)  {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s1 := []int{3,4}
	t.Log(len(s1), cap(s1), s1[0], s1[1])

	s2 := make([]int, 5)
	s2 = append(s2, 6, 7,8,9)
	t.Log(len(s2), cap(s2))
	// for ide,e := range s2{
	// 	t.Log(ide, e)
	// }
}

func TestSliceGrowing(t *testing.T)  {
	s0 := []int{}
	for i:=0; i<16; i++{
		t.Log(len(s0),cap(s0))
		s0=append(s0, i)
	}
	t.Log(s0[0])
}

func TestSliceDynamiGrow(t *testing.T){
	a := []int{}
	b := []int{1,2,3}
	c := a

	a = append(b, 1)

	fmt.Printf("slice a: %v\n", a)
	fmt.Printf("slice c: %v\n", c)
}