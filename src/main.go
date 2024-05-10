package main

//dont forget to export DBUSER and DBPASS
import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"os"
	//"errors"
)

type request struct {
	req string
	//TODO error storage
}

var db *sql.DB

var errorUnexpectedInput = fmt.Errorf("Expected 1 argument, got: ")

type Websites struct {
	Weblink string
}

// connect to db
func handle(cfg mysql.Config) (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
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
func getWebsites(db *sql.DB) ([]Websites, error) {
	var group_sites []Websites

	rows, err := db.Query("SELECT * from websites")
	if err != nil {
		return nil, fmt.Errorf("Query error %w", err)
	}
	defer rows.Close() //??

	for rows.Next() {
		var site Websites
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

// queries single websites table row
func getWebsiteByStamp(db *sql.DB) (Websites, error) {
	var w Websites
	weblink := "%test%"

	row := db.QueryRow("SELECT * from websites where weblink like ?", weblink)
	if err := row.Scan(&w.Weblink); err != nil {
		return w, fmt.Errorf("getWebsiteByStamp error with row: %w", err)
	}
	return w, nil
}

// inserts a row into the websites table, no id
func addWebsite(w Websites, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO websites (weblink) VALUES (?)", w.Weblink)
	//result above - removed
	if err != nil {
		return fmt.Errorf("addWebsite: %w", err)
	}
	/*
	   id, err := result.lastInstertedId()//retrieves the id of the database row
	   if err != nil {
	   	return fmt.Errorf("addWebsite id gen: %w", err"
	   }
	*/
	return nil
}

// warning - will panic for zero arguments for some reason
func main() {
	fmt.Println("Hello world!", os.Args[1])
	/*
			   input_slice := os.Args
		           if len(input_slice) != 2 {
		               return
		           }
		           r := request{req: os.Args[1]}
	*/
	var err error
	str, err := input_chk(os.Args)
	if err != nil {
		fmt.Println(fmt.Errorf("Input error: %w", err))
		return
	}
	r := request{req: str[0]}
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "visits",
		AllowNativePasswords: true,
	}
	db, err = handle(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	sites, err := getWebsites(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("entries in db: %v\n", sites)
	site, err := getWebsiteByStamp(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("entry in db: %v\n", site)
	/*adds stuff to db
	w := Websites{Weblink: "https://go.dev"}
	err = addWebsite(w, db)
	if err != nil {
		log.Fatal(err)
	}
	sites, err = getWebsites(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("entries in db: %v\n", sites)
	*/
	// TODO: pretty struct fill
}

//TODO1.1 read user input and react to it (filter) /v1 - cmd, v2 stdin(165)
//Input is a single text string
//TODO1.2 store the string entered in a /v1 - file, /v2 - database

//TODO2.1 display a line from the database
//TODO2.2 open websites entered. Feed to firefox
//TODO2.3 log opened website or error

//TODO3.1 find where mozilla history is and pull data from there