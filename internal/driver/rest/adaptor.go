package rest

import{
	"github.com/go-sql-driver/mysql"
}

db *sql.DB

type game struct{
	Id string
	PlayerName string
	Scenario string 
	Score int
	CountCorrect int
	QuestionId int
	QuestionTimeout int
}


type question{
	Id int
	problem string
	choices []string
	correctIndex int
}

func conn(){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hex_math")
	if err != nil{
		panic(err.Error())
	}
	return db
}

func GetQuestion(id int){
	db := conn()
	defer db.Close()
	var q question
	err := db.QueryRow("SELECT * FROM questions WHERE id = ", id).Scan(&q.Id, &q.problem, &q.choices, &q.correctIndex)
	if err != nil{
		panic(err.Error())
	}
	return q
}

func GetPlayer(playerName string){
	db := conn()
	defer db.Close()
	var p game
	err := db.QueryRow("SELECT * FROM games WHERE player_name = ", playerName).Scan(&p.Id, &p.PlayerName, &p.Scenario, &p.Score, &p.CountCorrect, &p.QuestionId, &p.QuestionTimeout)
	if err != nil{
		panic(err.Error())
	}
	return p
}



