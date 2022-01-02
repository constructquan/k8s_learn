package string_test
import (
	"testing"
	"unsafe"
	"fmt"
	"unicode/utf8"
	"strings"
	"sync"
	//"time"
)
func TestStringUnicode(t *testing.T)  {
	s := "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	t.Log(c[0])
	t.Log("rune size: ", unsafe.Sizeof(c[0]))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T)  {
	s := "中华人民共和国"
	for _, c := range s{
		t.Logf("%[1]c %[1]x", c)
	}
	t.Log(len(s))
}

func TestString(t *testing.T){
	var s1 = string("abcdef")

	t.Log(s1)
    t.Logf("%T %p\n", s1,&s1)
	s1 = "XYZ"
	t.Log(s1)
	t.Logf("%T %p\n", s1, &s1)
	t.Logf("长度： %d\n", len(s1))
	var s2 = string(`中国最棒,yes.`)
	
	fmt.Printf( s2+"\n")
	fmt.Printf("长度： %d\n", len(s2))

	//遍历的是字节流序列
	for i:=0;i<len(s2);i++{
		fmt.Printf("%x ", s2[i])
	}
	fmt.Print("\n")

	//遍历unicode，码点
	for _, c:=range s2{
		fmt.Printf("%x-", c)

		fmt.Printf("%c ", c)
	}
	

}

func TestEncodeDecode(t *testing.T)  {
	//uncicode 码点和 rune 可以对应。
	//但是由于直接使用码点，在存储效率上太低。每个都需要占用4个字节，而且和ASCII码不兼容。
	//在不同系统、机器上传输需要考虑系统位数兼容问题。
	//所以，使用UTF-8编码方式，对unicode字符集进行编码。


	//rune -> []byte
	encodeRune()

	//[]byte -> rune
	decodeRune()
}

func encodeRune()  {
	var r rune = 0x4E2D
	fmt.Printf("The unicode charactor is %c\n", r)//中
	buf := make([]byte,3)
	_ = utf8.EncodeRune(buf, r)
	fmt.Printf("utf-8 representation is %x\n", buf)
}

func decodeRune()  {
	var buf = []byte{0xe4, 0xb8, 0xad} 
	r, _ := utf8.DecodeRune(buf) //对buf的进行utf-8 解码
	fmt.Printf("the unicode charactor after decoding {0xe4, 0xb8, 0xad} is %x\n", r)

}

func TestStringBuilder(t *testing.T)  {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var b strings.Builder

	for n:=0;n<1000;n++{
		wg.Add(1)	
		go func(){
			mu.Lock()
			defer mu.Unlock()
			b.WriteString("a")
			wg.Done()			
		}()		
	}
	wg.Wait()
	//time.Sleep(time.Second *1)
	fmt.Printf("string.builder 长度：%d\n", len(b.String()))
}

func Hello(a string){
	fmt.Printf("Hello:%p\n", &a)
	fmt.Print(unsafe.Sizeof(a))
}

func TestHello(t *testing.T){
	a := "chensir"
	fmt.Printf("TestHell: %p\n", &a)
	Hello(a)


	fmt.Print("\n\n---\n")

	s1 := "你好时间"
	s2 := s1
	fmt.Printf("s1-%s %p\n", s1, unsafe.Pointer(&s1))
	fmt.Printf("s2-%s %p\n", s2, unsafe.Pointer(&s2))
	var n int = 9
	fmt.Print(unsafe.Sizeof(n))
}