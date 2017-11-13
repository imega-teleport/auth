// +build acceptance

package acceptance

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {
	maxRetries := 60
	log.Println("Starting tests")

	db, err := sql.Open("mysql", getDSN())

	if err != nil {
		log.Fatalf("Failed to open DB %s", err)
	}

	for {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("err %s", err)
		if maxRetries == 0 {
			break
		}
		log.Printf("iteration")
		maxRetries--
		<-time.After(time.Duration(1 * time.Second))
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database")
	}

	os.Exit(m.Run())
}
