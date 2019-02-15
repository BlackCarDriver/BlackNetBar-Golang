package main

//the detail of an account
type AccountDate struct{
	Id int            `json:"id"`
      Name string       `json:"name"`
	Password string   `json:"password"`
	Money float64     `json:"money"`
	Lasttime string	`json:"lasttime"`
}


