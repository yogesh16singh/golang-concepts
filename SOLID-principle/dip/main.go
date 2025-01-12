package main

import "fmt"

// Abstraction: Database interface
type Database interface {
	SaveData(data string)
}

// Low-level module: MySQLDatabase
type MySQLDatabase struct{}

func (db *MySQLDatabase) SaveData(data string) {
	fmt.Println("Saving data to MySQL Database:", data)
}

// Low-level module: PostgreSQLDatabase
type PostgreSQLDatabase struct{}

func (db *PostgreSQLDatabase) SaveData(data string) {
	fmt.Println("Saving data to PostgreSQL Database:", data)
}

// High-level module: ReportGenerator
type ReportGenerator struct {
	database Database // Depends on the abstraction, not a concrete implementation
}

func (rg *ReportGenerator) Generate(data string) {
	fmt.Println("Generating report with data:", data)
	rg.database.SaveData(data)
}

func main() {
	// Use MySQL
	mysqlDB := &MySQLDatabase{}
	rg := ReportGenerator{database: mysqlDB}
	rg.Generate("MySQL Report Data")

	// Switch to PostgreSQL without modifying ReportGenerator
	postgresDB := &PostgreSQLDatabase{}
	rg.database = postgresDB
	rg.Generate("PostgreSQL Report Data")
}
