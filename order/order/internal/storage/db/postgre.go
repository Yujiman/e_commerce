package db

import (
	"log"
	"sync"
	"time"

	"github.com/Yujiman/e_commerce/goods/order/order/internal/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	once sync.Once
	db   *sqlx.DB
)

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
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go fatalEmptyChannelAfterTime(done, &wg, 60*time.Second)

	wg.Add(1)
	go pingDbInSomeTime(done, &wg, 5*time.Second)

	wg.Wait()
	return db
}

func pingDbInSomeTime(done chan struct{}, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()

	var err error
	db, err = getDb()
	if err == nil {
		done <- struct{}{}
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
			done <- struct{}{}
			return
		}
	}
}

func fatalEmptyChannelAfterTime(done chan struct{}, wg *sync.WaitGroup, duration time.Duration) {
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
	params := config.GetConfig().PostgreConnectionParams

	var err error
	dsn := "postgres://" + params.User + ":" + params.Password + "@" + params.Host + ":" + params.Port + "/" + params.DbName + "?sslmode=disable"
	db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("SELECT 1;")
	if err != nil {
		return nil, err
	}

	return db, nil
}
