package db

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Yujiman/e_commerce/auth/role/internal/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var once sync.Once
var db *sqlx.DB

func GetDbConnection() *sqlx.DB {
	once.Do(func() {
		db = getDbWithTicker()
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(20)
		db.SetConnMaxIdleTime(5 * time.Minute)
	})
	return db
}

func getDbWithTicker() *sqlx.DB {
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go fatalEmptyChannelAfterTime(done, &wg, 60*time.Second)

	wg.Add(1)
	go pingDbInSomeTime(done, &wg, 5*time.Second)

	wg.Wait()
	return db
}

func pingDbInSomeTime(done chan bool, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()

	var err error
	db, err = getDb()
	if err == nil {
		done <- true
		return
	}
	log.Printf("Get db error=%v\n", err)

	ticker := time.NewTicker(duration)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			var err error
			db, err = getDb()
			if err != nil {
				log.Printf("Get db error=%v\n", err)
				continue
			}
			done <- true
			return
		}
	}
}

func fatalEmptyChannelAfterTime(done chan bool, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	tickerFail := time.NewTicker(duration)
	for {
		select {
		case <-done:
			return
		case <-tickerFail.C:
			log.Panicln("Db init error")
			return
		}
	}
}

func getDb() (*sqlx.DB, error) {
	params := config.GetPostgreConnectionParams()

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		params.User, params.Password, params.Host, params.Port, params.DbName,
	)

	var err error
	db, err = sqlx.Open("pgx", connString)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("SELECT 1;")
	if err != nil {
		return nil, err
	}

	return db, nil
}
