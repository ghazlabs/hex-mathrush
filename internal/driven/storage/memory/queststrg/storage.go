package queststrg

import (
	"context"
	"database/sql"
	"math/rand"
	"os"
	"time"

	"github.com/ghazlabs/hex-mathrush/internal/core"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/validator.v2"
)

type Storage struct {
	questions []core.Question
}

type Config struct {
	Questions []core.Question `validate:"min=1"`
}

func (c Config) Validate() error {
	return validator.Validate(c)
}

func New(cfg Config) (*Storage, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	if err != nil {
		panic(err.Error())
	}
	var s Storage
	err = s.Init()
	if err != nil {
		panic(err.Error())
	}
	return &s, nil
}

func conn() *sql.DB {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	//db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp(db:"+db_port+")/hex_math")
	//db, err := sql.Open("mysql", "root:test@tcp(host.docker.internal:3306)/hex_math")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (s *Storage) Init() error {
	var quests []core.Question
	db := conn()
	defer db.Close()

	rows, err := db.Query("SELECT id, problem, correct_index FROM questions")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var q core.Question
		var id int

		err := rows.Scan(&id, &q.Problem, &q.CorrectIndex)
		if err != nil {
			panic(err.Error())
		}
		choices, err := db.Query("SELECT choice FROM choices WHERE question_id = ?", id)
		if err != nil {
			panic(err.Error())
		}
		for choices.Next() {
			var choice string
			err := choices.Scan(&choice)
			if err != nil {
				panic(err.Error())
			}

			q.Choices = append(q.Choices, choice)
		}

		quests = append(quests, q)
	}
	s.questions = quests
	return nil
}

func (s *Storage) GetRandomQuestion(ctx context.Context) (*core.Question, error) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	idx := r.Intn(len(s.questions))

	return &s.questions[idx], nil
}
