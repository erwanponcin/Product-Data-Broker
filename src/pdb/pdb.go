package main

import(
    "fmt"
	"bufio"
    "log"
	"os"
	"strings"

	"database/sql"
    _ "github.com/go-sql-driver/mysql"

	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
    // Connexion to MySql
    dsn := "root:root@tcp(127.0.0.1:3306)/products"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	// Connexion to Mongo
    clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
    client, err := mongo.Connect(nil, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to databases!")


	fmt.Println("Choose a product code:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	

	// request MySQL
	request:="SELECT * FROM iron_products WHERE code='" + input + "'"

	rows, err := db.Query(request)
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

		fmt.Printf("ID:              %d\n", id)
		fmt.Printf("Code:            %s\n", code)
		fmt.Printf("Name:            %s\n", name)
		fmt.Printf("Category:        %s\n", category)
		fmt.Printf("Price Per Tonne: %.2f\n", price_per_tonne)
		fmt.Printf("Stock Tonne:     %.2f\n", stock_tonne)
	}

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

	// request Mongo
    collection := client.Database("products_datas").Collection("iron_products")

    var result bson.M
    err = collection.FindOne(nil, bson.M{"product_code": input}).Decode(&result)
    if err != nil {
        log.Fatal("Document non trouvé : ", err)
    }

    for key, value := range result {
        fmt.Printf("%s: %v\n", key, value)
    }
}