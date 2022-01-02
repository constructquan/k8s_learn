package map_test
import "testing"
func TestMapInit(t *testing.T)  {
	m := map[int]int{1:10, 2:20, 3:30, 4:40}
	t.Log(len(m), m[3])
//Map 访问 key 不存在时，仍会返回零值，所以，不能通过返回nil来判断
	if val, ok := m[3];ok{
		t.Logf("m[3] is %d", val)
	}else{
		t.Log("m[3] is not existing.")
	}

	m1 := make(map[int]int, 10)
	t.Log(m1[2])
}


func TestMapTravel(t *testing.T)  {
	m2 := map[int]int{1:1, 2:2, 3:3,4:4,5:5}

	for k,
}