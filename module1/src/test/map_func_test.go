package map_func_test
import "testing"
func TestMapWithFunVal(t *testing.T)  {
	m := map[int]func(op int)int {}
	m[1] = func (op int)int  {
		return op
	}
	m[2] = func (op int)int  {
		return op * op
	}
	m[3] = func (op int)int  {
		return op * op * op
	}

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T)  {
	mySet  :=map[int]bool{}
	n := 3
	mySet[3] = true
	if mySet[n]{
		t.Logf("%d is existing in mySet", n)
	}else{
		t.Logf("%d is not existing in mySet", n)
	}
	mySet[1]=true
	mySet[2]=true
	t.Log(mySet)
	delete(mySet, 2)
	t.Log(mySet)
}