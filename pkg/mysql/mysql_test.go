package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/fortytw2/dockertest"
)

func TestMysqlConnect(t *testing.T) {
	container, err := dockertest.RunContainer("mysql:5.7", "3306", func(addr string) error {
		db, err := sql.Open("mysql", fmt.Sprintf("root:rahasia@tcp(%v)/mysql", addr))
		if err != nil {
			return err
		}

		err = db.Ping()
		if err != nil {
			return err
		}

		return nil
	}, "-e", "MYSQL_ROOT_PASSWORD=rahasia")
	defer container.Shutdown()

	if err != nil {
		t.Fatalf("could not start mysql, %s", err)
	}

	// setup config
	os.Setenv("DATABASE_DRIVER", "mysql")
	os.Setenv("DATABASE_DSN", fmt.Sprintf("root:rahasia@tcp(%v)/mysql", container.Addr))

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Error establishing connection %v", r)
			}
		}()

		mysql := newConnection()
		mysql.Close()
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Error establishing connection %v", r)
			}
		}()

		mysql := GetConnection()
		mysql.Close()
	}()
}
