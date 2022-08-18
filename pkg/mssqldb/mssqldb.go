package mssqldb

import (
	"database/sql"
	"flag"
	"fmt"

	mssql "github.com/denisenkom/go-mssqldb"
)

// Connection ...
type Connection struct {
	Debug    bool
	Password string
	Port     int
	Server   string
	User     string
	DB       string
}

// Connect ...
func (m Connection) Connect() (*sql.DB, error) {
	flag.Parse()

	if m.Debug {
		fmt.Printf(" password:%s\n", m.Password)
		fmt.Printf(" port:%d\n", m.Port)
		fmt.Printf(" server:%s\n", m.Server)
		fmt.Printf(" user:%s\n", m.User)
	}

	connStr := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=30&encrypt=disable",
		m.User, m.Password, m.Server, m.Port, m.DB,
	)
	if m.Debug {
		fmt.Printf(" connString:%s\n", connStr)
	}

	// Create a new connector object by calling NewConnector
	connector, err := mssql.NewConnector(connStr)
	if err != nil {
		return nil, err
	}

	// Use SessionInitSql to set any options that cannot be set with the dsn string
	// With ANSI_NULLS set to ON, compare NULL data with = NULL or <> NULL will return 0 rows
	connector.SessionInitSQL = "SET ANSI_NULLS ON"

	// Pass connector to sql.OpenDB to get a sql.DB object
	db := sql.OpenDB(connector)

	return db, err
}
