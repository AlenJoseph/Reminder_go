package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	//  tm "github.com/buger/goterm"
	// "time"
)

func main() {

	// getTime()

	fmt.Println("Welcome to Reminder Cli")
	fmt.Println("What do you want to do?")
	fmt.Println("1.Add Reminder")
	fmt.Println("2.Edit Reminder")
	fmt.Println("3.Delete Reminder")
	fmt.Println("4.View Reminder")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 1:
		addReminder()
	case 2:
		editReminder()
	case 3:
		deleteReminder()
	case 4:
		viewReminder()
	default:
		fmt.Printf("Invalid Entry\n")
	}

}
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "Reminder"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// func getTime(){
// 	// tm.Clear() // Clear current screen

//     for {
//         // By moving cursor to top-left position we ensure that console output
//         // will be overwritten each time, instead of adding new.
//         tm.MoveCursor(1,1)

//         tm.Println("Current Time:", time.Now().Format(time.RFC1123))

//         tm.Flush() // Call it every time at the end of rendering

//         time.Sleep(time.Second)
//     }
// }
func addReminder() {
	fmt.Println("What do you want to remind")
	var rem string
	fmt.Scanln(&rem)
	fmt.Println("When do you want to be reminded")
	var dt string
	fmt.Scanln(&dt)
	fmt.Println("added")
}
func editReminder() {
	fmt.Print("edit")
}
func deleteReminder() {
	db := dbConn()
	var id int
	viewReminder()
	fmt.Println("Enter the id of the reminder you want to delete")
	fmt.Scan(&id)
	delForm, err := db.Prepare("DELETE FROM Reminder_Table WHERE Reminder_Table_id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	fmt.Print("delete")
}
func viewReminder() {
	db := dbConn()
	var (
		id  int
		rem string
		dt  string
	)
	rows, err := db.Query("SELECT * FROM Reminder_Table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &rem, &dt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, rem, dt)

		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
	}

}
