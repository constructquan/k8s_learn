package strings_fn

import "testing"
import "strings"
import "strconv"

func TestStringsFunc(t *testing.T)  {
	s := "A,B,C,D"	
	parts := strings.Split(s, ",")
	for _, v := range parts{
		t.Log(v)
	}

	t.Log(strings.Join(parts, "~"))
}

func TestStringConv(t *testing.T){
	s := strconv.Itoa(10)
	t.Log("str" + s)

	if n, err := strconv.Atoi(s); err == nil{
		t.Log(10 + n)
	}
	
}