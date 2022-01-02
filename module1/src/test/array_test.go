package array_test
import (
	"testing"
)
func TestArrayInit(t *testing.T)  {
	var arr [3]int
	arr[0] = 1
    t.Log(arr[0], arr[1], arr[2])

	var arr1 = [...]int {2,3,4,5,6}
	t.Log(arr1[3])

	arr2 := [3]string {"I", "am", "tall"}
	t.Log(arr2[2])

	//把整个数据都打印出来：
	t.Log(arr1, arr2)

}

func TestArrayTravel(t *testing.T)  {
	arr := [3]string {"I", "am", "tall"}
	for k,val := range arr{
		t.Log(k, val)
	}
}

func TestArraySection(t *testing.T)  {
	arr := [...]string{"You", "run", "faster", "than", "foxes"}
	for idex, e := range arr{
		t.Log(idex, e)
	}

	var arr_sec = arr[3:]
	t.Log(arr_sec)
}