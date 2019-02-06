package main

import (
"log"
"fmt"
"database/sql"
	_ "github.com/go-sql-driver/mysql"
//  tm "github.com/buger/goterm"
// "time"
)
func main() {

	// getTime()
	dbCon()
	fmt.Println("Welcome to Reminder Cli")
	fmt.Println("What do you want to do?")
	fmt.Println("1.Add Reminder")
	fmt.Println("2.Edit Reminder")
	fmt.Println("3.Delete Reminder")
	fmt.Println("4.View Reminder")
    var input int
	fmt.Scanln(&input)
	switch input{
	case 1 : addReminder()
	case 2 : editReminder()
	case 3 : deleteReminder()
	case 4 :viewReminder()
	default:
	   fmt.Printf("Invalid Entry\n" );
 }
 
	
    
}
func dbCon(){
	db, err := sql.Open("mysql",
	"root:root@tcp(localhost:3306)/Reminder")
    if err != nil {
	log.Fatal(err)
     }
    defer db.Close()
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
func addReminder(){
	fmt.Println("What do you want to remind")
	var rem string
	fmt.Scanln(&rem)
	fmt.Println("When do you want to be reminded")
	var dt string
	fmt.Scanln(&dt)
	fmt.Println("added")
}
func editReminder(){
	fmt.Print("edit")
}
func deleteReminder(){
	fmt.Print("delete")
}
func viewReminder(){
	fmt.Print("view")
}