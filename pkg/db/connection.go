package db

import (
	"fmt"

	. "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetConnection(host, port, user, dbname, password, sslmode string) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)
	connection, err := Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("got an error when tried to make connection with database:%w", err)
	}
	return connection, nil
}
