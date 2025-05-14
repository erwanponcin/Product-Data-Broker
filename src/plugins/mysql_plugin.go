package plugins

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type MoldingData struct {
    SerialNumber   string
    MaterialType   string
    CoolingTime    int
    ProductionDate string
}

type MySQLPlugin struct {
    db *sql.DB
}

// Initializes a new MySQL plugin
func NewMySQLPlugin(dsn string) (*MySQLPlugin, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("could not connect to MySQL: %v", err)
    }

    return &MySQLPlugin{
        db: db,
    }, nil
}

// Fetches all data from the MySQL table
func (mp *MySQLPlugin) FetchAllData() ([]MoldingData, error) {
    rows, err := mp.db.Query("SELECT serial_number, material_type, cooling_time, production_date FROM molding_batches")
    if err != nil {
        return nil, fmt.Errorf("could not fetch data: %v", err)
    }
    defer rows.Close()

    var results []MoldingData
    for rows.Next() {
        var data MoldingData
        if err := rows.Scan(&data.SerialNumber, &data.MaterialType, &data.CoolingTime, &data.ProductionDate); err != nil {
            return nil, fmt.Errorf("could not scan row: %v", err)
        }
        results = append(results, data)
    }

    return results, nil
}
