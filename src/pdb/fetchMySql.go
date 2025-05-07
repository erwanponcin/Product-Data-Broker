package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Connexion à MySQL
    dsn := "root:root@tcp(127.0.0.1:3306)/products"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Test de la connexion
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connecté à MySQL avec succès!")

    // Exemple de requête SQL
    rows, err := db.Query("SELECT * FROM iron_products WHERE id=1") // Remplace par le nom de ta table
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var id int
		var code string
        var name string
		var category string
		var price_per_tonne float64
		var stock_tonne float64
        err := rows.Scan(&id, &code, &name, &category, &price_per_tonne, &stock_tonne)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Code: %s, Name: %s, Category: %s, Price Per Tonne: %.2f, Stock Tonne: %.2f\n", id, code, name, category, price_per_tonne, stock_tonne)
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}
