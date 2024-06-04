package service

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type DBService struct {
	db *sql.DB
}

const (
	DataPath = "/DATA/AppData/zimaos-status"
)

func NewDBService() *DBService {
	// create DATA path if not exsit
	if _, err := os.Stat(DataPath); os.IsNotExist(err) {
		os.MkdirAll(DataPath, os.ModePerm)
	}

	db, err := sql.Open("sqlite3", "file:"+DataPath+"/data.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}

	dbService := &DBService{
		db: db,
	}

	// init
	var tableName string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='CPUData'").Scan(&tableName)

	if err == sql.ErrNoRows {
		dbService.Init()
	}

	return dbService
}

func (s *DBService) Init() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS CPUData (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME NOT NULL,
		percent REAL NOT NULL
	);`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}

	sqlStmt = `
	CREATE TABLE IF NOT EXISTS MemData (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME NOT NULL,
		percent REAL NOT NULL
	);`
	_, err = s.db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DBService initialized")
}

func convertTimestamp(ts string) (int64, error) {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return 0, err
	}
	// 转换为Go的time.Time，然后格式化为SQLite可接受的格式

	return i, nil
}
