package main

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"strconv"
	// "fmt"
)

//delete an account by name 
func DateleteAcccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Access-Control-Allow-Origin", "*") 
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Headers","content-type")
	writer.Header().Set("content-type","application/json")
	var postbody map[string]string
	body, _ := ioutil.ReadAll(request.Body)
	if(len(body) == 0){
		return
	}
	json.Unmarshal(body, &postbody)
	name := postbody["name"]
	Delete(name)
	res,_ := json.Marshal(1)
	log.Println(name," have been delete...")
	writer.Write(res)
}

func Test(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Access-Control-Allow-Origin", "*") 
	log.Println("test success...")
	writer.Write([]byte("HaHaHaHaHaHa...."))
}

//only update money
func UpdateAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Access-Control-Allow-Origin", "*") 
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Headers","content-type")
	writer.Header().Set("content-type","application/json")
	body, _ := ioutil.ReadAll(request.Body)
	if(len(body) == 0){
		return
	}
	var postbody map[string]string
	json.Unmarshal(body, &postbody)
	name := postbody["name"]
	newmoney,_:= strconv.ParseFloat(postbody["money"],64)
	log.Println(newmoney)
	err :=  UpdateMoney(name, newmoney)
	if err != nil {
		log.Println("some worng happen in updateaccount:",err)
	}else{
		log.Println(name," money update to ",newmoney)
		res,_ := json.Marshal(1)
		writer.Write(res)
	}
}


func GetMoney(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Access-Control-Allow-Origin", "*") 
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Headers","content-type")
	writer.Header().Set("content-type","application/json")
	body, _ := ioutil.ReadAll(request.Body)
	if len(body) == 0 {
		return
	}
	var postbody map[string]string
	json.Unmarshal(body, &postbody)
	name := postbody["name"]
	log.Println(name," apply to look the money...")
	money,err := Getmoney(name)
	if err!=nil {
		log.Println("worng happen in getmoney:",err)
	}else{
		res,_ := json.Marshal(money)
		writer.Write(res)
	}
	return
}


// Create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*") 
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Headers","content-type")
	writer.Header().Set("content-type","application/json")
	var postbody map[string]string
	body, _ := ioutil.ReadAll(request.Body)
	if(len(body) == 0){
		return
	}
	json.Unmarshal(body, &postbody)
	name:=postbody["name"]
	account := AccountDate{
		Name:     name,
		Password: postbody["password"],
		Money:		10.0,
	}
	_,err := Signin(&account)
	if err!=nil {
		res,_ := json.Marshal(-1)
		writer.Write(res)	//already have account with same name
		return
		}else{
			res,_ := json.Marshal(1)
			writer.Write(res)	//register successdly
			log.Println(name," register succedly...")
	}
}


func LoginAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Access-Control-Allow-Origin", "*") 
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Headers","content-type")
	writer.Header().Set("content-type","application/json")
	log.Println("test...")
	var postbody map[string]string
	body, _ := ioutil.ReadAll(request.Body)
	if(len(body) == 0){
		return
	}
	json.Unmarshal(body, &postbody)
	name := postbody["name"]
	SetCookie2(writer,request,name)
	log.Println(name," Login...")
	password:= postbody["password"]
	if name == "blackcardriver" && password=="123456" {
		res,_ := json.Marshal(10)
		writer.Write(res)	//login with Adminstrator 
		log.Println("Administrator login...")
		return
	}
	truepassword,err := GetPassword(name)
	if err!=nil {
		res,_ := json.Marshal(0)
		writer.Write(res)	//can't find account with 'name'
	} else if truepassword == password {
		res,_ := json.Marshal(1)
		writer.Write(res)	//password is right
		err = UpdateLasttime(name)
		if err!=nil {
			log.Println("updatelastime worng:",err)
		}
		log.Println(name," login...")
	} else {
		res,_ := json.Marshal(3)
		writer.Write(res) 	//the password is worng
	}
}


//one page have 20 rows at most
func Getdate(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Access-Control-Allow-Origin", "*") 
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Headers","content-type")
	writer.Header().Set("content-type","application/json")
	var postbody map[string]int
	body, _ := ioutil.ReadAll(request.Body)
	if(len(body) == 0){
	 	return
	}
	json.Unmarshal(body, &postbody)
	page := postbody["page"]
	date,_ := GetDate((page-1)*20)
	result,_:= json.Marshal(date)
	writer.Write(result)
}


//return how the sum of rows and the index of precisl one
func Getstate(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Access-Control-Allow-Origin", "*") 
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Headers","content-type")
	writer.Header().Set("content-type","application/json")
	var postbody map[string]string
	body, _ := ioutil.ReadAll(request.Body)
	if(len(body) == 0){
	 	return
	}
	var state [2]int		//sum and index
	json.Unmarshal(body, &postbody)
	name := postbody["name"]
	state[0],_ = GetSumOfRow()
	state[1],_ = GetRowOfName(name)
	if state[1] > 0 {
		state[1] = (state[1]/20 +1 )
	}
	result,_:= json.Marshal(state)
	writer.Write(result)
}

