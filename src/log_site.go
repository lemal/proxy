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

/*
type request struct {
	req string
	//TODO error storage
}
*/

var db *sql.DB

var errorUnexpectedInput = fmt.Errorf("Expected 1 argument, got: ")

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

// warning - will panic for zero arguments for some reason
func main() {  
	var err error
	str, err := input_chk(os.Args)
	if err != nil {
	   log.Fatal(err)
	}
	w := Website{Weblink: str[0]}
	
	db, err = handle(cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	err = addWebsite(w, db)
	if err != nil {
		log.Fatal(err)
	}
	
	sites, err := getWebsites(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("entries in db: %v\n", sites)
	/*
	site, err := getWebsiteByStamp(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("entry in db: %v\n", site)
	*/
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
//Input is a single text string v
//TODO1.2 store the string entered in a /v1 - file, /v2 - database v
//TODO1.3 add the entered string to the database v

//TODO2.0 add a program to clear the database

//TODO2.1 display a line from the db v
//TODO2.2 find the line by like
//TODO2.3 log stuff into a file

//TODO3.1 find where mozilla history is and pull data from there
//TODO3.2 launch the website ?? if it is like a website launcher

//TODO Check that the export happened. Or even do the export here (alt)


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

// queries single websites table row
func getWebsiteByStamp(db *sql.DB) (Website, error) {
	var w Website
	weblink := "%test%"

	row := db.QueryRow("SELECT * from websites where weblink like ?", weblink)
	if err := row.Scan(&w.Weblink); err != nil {
		return w, fmt.Errorf("getWebsiteByStamp error with row: %w", err)
	}
	return w, nil
}

// inserts a row into the websites table, no id
func addWebsite(w Website, db *sql.DB) error {
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

