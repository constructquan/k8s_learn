package condition_test
import "testing"
func TestCondition(t *testing.T)  {
	if val,err := otherFunction(); err{
		t.Log("err has")
	}else{
		t.Log("err is nil")
		t.Log(val)
	}
}

func otherFunction()(val int, err bool)  {

	var a int = 1
	return a, err
	
}