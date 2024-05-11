package main

import (
       "fmt"
       "log"
       "os"
       "database/sql"
       "github.com/go-sql-driver/mysql"
       )

var db *sql.DB

var errorUnexpectedInput = fmt.Errorf("Expected 1 arg, got: ")

type Website struct {
     Weblink string
}

var cfg = mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "visits",
		AllowNativePasswords: true,
}


func main(){
     w, err := input_chk(os.Args)
     if err != nil {
     	log.Fatal(err)
     }
     db, err := handle(cfg)
     if err != nil {
     	log.Fatal(err)
     }
     site := Website{Weblink: w[0]}
     err = delWebsite(site, db)
     if err != nil {
     	log.Fatal(err)
     }
     //view
     sites, err := getWebsites(db)
     if err != nil {
     	log.Fatal(err)
     }
     fmt.Printf("entries in db: %v", sites)
}

// connect to db
func handle(cfg mysql.Config) (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	fmt.Println("Connected!")
	return db, nil
}

func input_chk(s []string) ([]string, error) {
	if len(s) != 2 {
		return nil, fmt.Errorf("%w %d", errorUnexpectedInput, len(s)-1)
	}
	var slice_s []string
	slice_s = append(slice_s, s[1])
	return slice_s, nil
	//TODO optimize for copies, pass around an encapsulated struct
}

// gets all entries from the websites table
func getWebsites(db *sql.DB) ([]Website, error) {
	var group_sites []Website

	rows, err := db.Query("SELECT * from websites")
	if err != nil {
		return nil, fmt.Errorf("Query error %w", err)
	}
	defer rows.Close() //??

	for rows.Next() {
		var site Website
		if err := rows.Scan(&site.Weblink); err != nil {
			return nil, fmt.Errorf("getWebsites rows: %w", err)
		}
		group_sites = append(group_sites, site)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getWebsites errors: %w", err)
	}
	return group_sites, nil
}

// delWebsite deletes a row from the websites table
func delWebsite(w Website, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM websites where weblink = ?", w.Weblink)
	if err != nil {
		return fmt.Errorf("addWebsite: %w", err)
	}
	return nil
}
