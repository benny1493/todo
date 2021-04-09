package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type NoteBook struct {
	name        string
	cpu         string
	os          string
	harddisk    string
	graphiccard string
}

func main() {

	notebook := NoteBook{
		name:        "lenovo",
		cpu:         "i5",
		os:          "window",
		harddisk:    "1T",
		graphiccard: "geforce",
	}
	insertNoteBook(notebook)

	fmt.Println("NoteBook: ", notebook)
}

func insertNoteBook(notebook NoteBook) {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/com")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec(`
		insert into com.notebook  
		values (
			"` + notebook.name + `", 
			"` + notebook.cpu + `",
			"` + notebook.os + `",
			"` + notebook.harddisk + `",
			"` + notebook.graphiccard + `"
		)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func updateNoteBook(name string, notebook NoteBook) error {
	fmt.Println("target: ", name)
	fmt.Println("data: ", notebook)

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/com")
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("DB Connected")

	_, err = db.Query(`
		update 
			todo 
		set 
			user_id = "` + notebook.name + `", 
			start_date="` + notebook.cpu + `", 
			end_date="` + notebook.os + `", 
			title="` + notebook.harddisk + `",
			status="` + notebook.graphiccard + `" 
		where 
			user_id="` + name + `"`)

	if err != nil {
		return err
	}

	fmt.Println("data insert success")

	return nil
}

func deleteNoteBook(name string) error {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/com")
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("DB Connected")

	_, err = db.Query(`delete from notebook where user_id="` + name + `"`)

	if err != nil {
		return err
	}

	fmt.Println("data insert success")

	return nil
}
