package storage

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type found struct {
	Found int
}

func Migrate(db *gorm.DB) {
	d := db.Exec("CREATE TABLE IF NOT EXISTS migrations (name VARCHAR NOT NULL);")
	if d.Error != nil {
		log.Fatal(d.Error)
	}

	files, err := ioutil.ReadDir("./migrations")
	if err != nil {
		log.Fatal(err)
	}

	//Run new migrations
	migrated := 0
	for _, f := range files {
		sql, err := ioutil.ReadFile(fmt.Sprintf("./migrations/%s", f.Name()))
		if err != nil {
			log.Fatal(err)
		}

		// initialize -1. 0 is a valid number.
		found := found{-1}
		d := db.Raw("SELECT count(1) as found FROM migrations WHERE name=?", f.Name()).Scan(&found)
		if d.Error != nil {
			log.Fatal(d.Error)
		}

		if found.Found == 0 {
			d = db.Exec(string(sql))
			if d.Error != nil {
				log.Fatal(d.Error)
			}

			d = db.Exec("INSERT INTO migrations(name) VALUES(?)", f.Name())
			if d.Error != nil {
				log.Fatal(d.Error)
			}

			migrated++
			log.Println("Migrated: ", f.Name())
		}
	}

	if migrated == 0 {
		log.Println("No migrations")
	}
}
