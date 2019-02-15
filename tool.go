package main

import(
	"net/http"
	"log"
)

//not use as now
func SetCookie(w http.ResponseWriter, r *http.Request){
	log.Println("setcookie right")
	SetCookie2(w,r,"right!!!")
}

//not use as now
func SetCookie2(w http.ResponseWriter, r *http.Request,name string){
	log.Println(name)
	cookie := http.Cookie{
		Name: "username",
		Value: name,
		HttpOnly: true,
	}
	http.SetCookie(w,&cookie)
	// w.Header().Set("Set-Cookie",cookie.String())
}

//not use as now
func GetCookie(w http.ResponseWriter, r *http.Request){
	// ck,err := r.Cookie("username")
	// if err!=nil {
	// 	log.Println(err)
	// }else{
	// 	log.Println(ck.Value)
	// }
	h:=r.Header["Cookie"]
		log.Println(h)
}

