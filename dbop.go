package main
import(
	"database/sql"
	_"github.com/lib/pq"
	"log"
)

var Db *sql.DB	

func init(){
	var err error
	Db,err = sql.Open(
	"postgres",
	"user=postgres dbname=backnetbar password=include123 ")
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
		if err != nil {
    	log.Fatal(err)
	}else{
		log.Println("database connect success...")
	}
}

//register a account and save it in database
func Signin(date *AccountDate)(id int,err error){
	statement := "insert into account(aname,apassword,amoney,lasttime)values($1,$2,$3,LOCALTIMESTAMP)returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		id = -1
		return 
	}
	defer stmt.Close()
	err = stmt.QueryRow(date.Name, date.Password, date.Money).Scan(&id)
	return
} 

//find and delete a account by name
func Delete(name string)(err error){
	Db.Exec("delete from account where aname = $1",name)
	return
}

//find and get the date of a account by name
func Getmoney(name string)(money float64, err error){
	money = 0.0
	err = Db.QueryRow("select amoney from account where aname = $1",name).Scan(&money)
	return
}

//get the password of 'name'
func GetPassword(name string)(password string, err error){
	err = Db.QueryRow("select apassword from account where aname = $1",name).Scan(&password)
	return
}

//check if there 'name' in database
func Ifhavename(name string)(have int){
	err := Db.QueryRow("select id from account where aname = $1",name).Scan(&have)
	if err!=nil {
		return 1
	}
	return 0
}
func UpdateLasttime(name string)(err error){
	_,err = Db.Exec("update account set lasttime=LOCALTIMESTAMP where aname = $1",name)
	return
}

//take and return number of date 
func GetDate(offset int)(account []AccountDate, err error){
	count := offset+1
	rows, err:= Db.Query("select aname, amoney, lasttime from account order by id asc offset $1 limit 20 ",offset)
	if err != nil {
		return
	}
	for rows.Next(){
		temp := AccountDate{Id:count}	//note here
		count ++
		err = rows.Scan(&temp.Name, &temp.Money,&temp.Lasttime)
		if err != nil{
			log.Println(err)
			return
		}
		account = append(account, temp)
	}
	rows.Close()
	return
}

//update the money of 'name'
func UpdateMoney(name string, money float64)(err error){
	_,err = Db.Exec("update account set amoney = $1 where aname = $2",money,name)
	return
}

//return the total number of clien database have save
func GetSumOfRow()(rows int, err error){
	err = Db.QueryRow("select count(*) from account").Scan(&rows)
	return
}

//return the index of 'name'
func GetRowOfName(name string)(row int, err error){
	err = Db.QueryRow("select rownum from (select row_number() over (order by id) as rowNum,* from account)m where m.aname=$1",name).Scan(&row)
	if err != nil {
		row = -1  	//account with 'name' cant be find
	}
	return
}