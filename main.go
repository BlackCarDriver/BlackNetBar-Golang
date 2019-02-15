package main

import (
	"net/http"
	"time"
)	
	

func main(){
	mux := http.NewServeMux()
	//mux.HandleFunc("/test",Test)
	//mux.HandleFunc("/delete",DateleteAcccount)
	mux.HandleFunc("/signin",SignupAccount)
	mux.HandleFunc("/login",LoginAccount)
	mux.HandleFunc("/getdate",Getdate)
	mux.HandleFunc("/delete",DateleteAcccount)
	mux.HandleFunc("/update",UpdateAccount)
	mux.HandleFunc("/getmoney",GetMoney)
	mux.HandleFunc("/getstate",Getstate)
	mux.HandleFunc("/test",Test)
	server := &http.Server{
		Addr: "0.0.0.0:5000",
		Handler: mux,
		ReadTimeout:time.Duration(10 * int64(time.Second)),
	}
	server.ListenAndServe()
}