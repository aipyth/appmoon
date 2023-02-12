package database

import (
	"fmt"

	"github.com/jackc/pgx"
)

// DBConfig contains the configuration for connecting to the database.
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

// DB is a struct that represents a connection to the database.
type DB struct {
	conn *pgx.Conn
}

// NewDB creates a new connection to the database and returns a DB instance.
func NewDB(config DBConfig) (*DB, error) {
  conn, err := pgx.Connect(pgx.ConnConfig{
  	Host:              config.Host,
  	Port:              uint16(config.Port),
  	Database:          config.DBName,
  	User:              config.User,
  	Password:          config.Password,
  	TLSConfig:         nil,
  	UseFallbackTLS:    false,
  	FallbackTLSConfig: nil,
  })
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &DB{conn: conn}, nil
}

// WriteActivity writes an activity record to the database.
func (db *DB) WriteActivity(activity Activity) error {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	_, err := db.conn.Exec("INSERT INTO activities (class_name, title, timestamp) VALUES ($1, $2, $3)", activity.ClassName, activity.Title, activity.Timestamp)
	if err != nil {
		return fmt.Errorf("unable to write activity to database: %v", err)
	}

	return nil
}

// Close closes the connection to the database.
func (db *DB) Close() error {
	return db.conn.Close()
}
