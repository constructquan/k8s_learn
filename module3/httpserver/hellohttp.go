package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
  	"syscall"
)

func Log(handler http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf( "--> %s %s from %s, code: %d \n",r.Method, r.URL.Path, r.RemoteAddr, http.StatusOK)
        handler(w, r)
    })
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// healthz test with code 200
	w.WriteHeader(200)
}

// get version
func version(w http.ResponseWriter, req *http.Request){
//	os.Setenv("VERSION", "v1.0.2")
        v := os.Getenv("VERSION")
	w.Header().Set("VERSION", v)
  	w.WriteHeader(200)	
}

// response request header 
func header(w http.ResponseWriter, req *http.Request){
	for name, headers := range req.Header {
        for _, h := range headers {
            //io.WriteString(w, fmt.Sprintf( "%v: %v\n", name, h))
		w.Header().Add(name, h)		
        }
    }
}

func Default(w http.ResponseWriter, req *http.Request){
      fmt.Fprintln(w, "Hello world!")

}

func main() {

   mux := http.NewServeMux()
   
   // default request 
   mux.HandleFunc("/", Log(Default))

   // healthz test
   mux.HandleFunc("/healthz",  healthz)    
 
   // get version
   mux.HandleFunc("/version", Log(version))

   // response request header 
   mux.HandleFunc("/header", Log(header))

   server := &http.Server{
      Addr:         ":8080",
      Handler:      mux,
   }
   // bind and start httpserver 
   go server.ListenAndServe()

   // elegant exit httpserver 
   listenSignal(context.Background(), server)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
   sigs := make(chan os.Signal, 1)
   signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

   select {
   case <-sigs:
      log.Println("notify sigs")
      httpSrv.Shutdown(ctx)
      log.Println("http shutdown")
   }
}

