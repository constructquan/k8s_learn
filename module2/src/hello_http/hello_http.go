package main

import (

	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"net"

)


func Log(handler http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf( "%s %s %d %s\n", r.RemoteAddr, r.Method, http.StatusOK, r.URL)
        handler(w, r)
    })
}


func index(w http.ResponseWriter, r *http.Request){
	//w.Write([]byte("Hello world...GO"))
	// 设置 VERSION
	os.Setenv("VERSION", "v0.9.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os VERSION: %s\n", version)
	// 将request的header设置到response中
	for name, headers := range r.Header{
		for _, h := range headers {
            fmt.Printf( "Header kye: %s, Header value:  %s\n", name, headers)
			w.Header().Set(name, h)		
        }	
	}
	// 打印访问日志ip
	clientip  := getCurrentIP(r)
	log.Printf( "Success! Response code: %d", 200)
	log.Printf("Success: clientip: %s", clientip)

}
//健康检查
func healthz(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w, "working")
}

func getCurrentIP(r *http.Request) string {
	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip
	// 但是要注意的是这两个请求头代表的ip都有可能是伪造的
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
	   // 当请求头不存在即不存在代理时直接获取ip
	   ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
 }

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
	   return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
	   return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
	   return ip
	}
	return ""
 }

func main() {
    mux := http.NewServeMux()
	
	mux.HandleFunc("/", index)
	mux.HandleFunc("healthz", healthz)
	// http.HandleFunc("/header", Log(reqHeader))
	// http.HandleFunc("/version", Log(getVersion))

	// http.HandleFunc("/",  Log(Default))
	// http.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":8080", mux);err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}