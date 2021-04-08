# todo

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type NoteBook struct {
	Name string
	Cpu    string
	Os string
	Ram string
	HardDisk string
	GraphicCard string
}

func insertTodo(n NoteBook {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	result, err := db.Exec(`
		insert into com.notebook  
		values (
			"` + notebook.Name + `", 
			"` + notebook.Cpu + `",
			"` + notebook.Os + `",
			"` + notebook.Ram + `",
			"` + notebook.HardDisk + `",
			"` + notebook.GraphicCard + `"
		)`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	fmt.Println("GOGOGO")
	
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/notebook")

	if err != nil (
		panic(err.Erro())
	)

	fmt.Println("db connected")

	defer db.close()
}

func main() {
	n := NoteBook{
		Name:    "dodo",
		Cpu: "i5",
		Os:   "window",
		HardDisk:     "1T",
		GraphicCard:    "geforce",
	}

	insertNoteBook(com)

	fmt.Println("NoteBook: ", notebook)
}

func selectTodo(Name string) ([]NoteBook, error){
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/notebook")
    if err != nil {
        return nil, err
    }
    defer db.Close()
func insertTodo(n NoteBook) error {
	fmt.Println("insertNoteBook: ", notebook)

	rows, err := db.Query("select * from com.notebook")
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/notebook")
	if err != nil {
		return nil, err
		return err
	}
	defer db.Close()
