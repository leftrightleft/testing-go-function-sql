package main
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
)

func main() {
    cmdName := req.URL.Query()["cmd"][0]
    // Open a connection to the database
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database_name")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Execute a SQL query
    rows, err := db.Query(fmt.Sprintf("SELECT * FROM books WHERE author = '%s'", cmdName))
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    // Iterate over the rows and print the results
    for rows.Next() {
        var id int
        var name string
        var author string
        err = rows.Scan(&id, &name, &author)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println(id, name, author)
    }
}
