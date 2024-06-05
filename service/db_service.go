package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/CorrectRoadH/CasaOS-Debugger/codegen/message_bus"
	_ "github.com/mattn/go-sqlite3"
)

type DBService struct {
	db *sql.DB
}

const (
	DataPath = "/DATA/AppData/casaos-debugger"
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
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='EventHistory'").Scan(&tableName)

	if err == sql.ErrNoRows {
		dbService.Init()
	}

	return dbService
}

func (s *DBService) Init() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS EventHistory (
		uuid TEXT PRIMARY KEY,
		properties TEXT NOT NULL,
		sourceID TEXT NOT NULL,
		name TEXT NOT NULL,
		timestamp DATETIME NOT NULL
	);`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DBService initialized")
}

func convertTimestamp(ts string) (time.Time, error) {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	// 转换为Go的time.Time，然后格式化为SQLite可接受的格式
	return time.Unix(i, 0), nil
}

func (s *DBService) InsertEvent(uuid string, properties map[string]string, sourceID string, name string, timestamp *time.Time) error {
	// 转换为SQLite可接受的格式
	var ts string
	if timestamp == nil {
		ts = time.Now().Format("2006-01-02 15:04:05")
	} else {
		ts = timestamp.Format("2006-01-02 15:04:05")
	}
	// 转换为json字符串
	rawJSON, err := json.Marshal(properties)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("INSERT INTO EventHistory(uuid, properties, sourceID, name, timestamp) VALUES(?, ?, ?, ?, ?)", uuid, string(rawJSON), sourceID, name, ts)
	if err != nil {
		return err
	}
	return nil
}

func (s *DBService) SourceList() ([]string, error) {
	var result []string
	rows, err := s.db.Query("SELECT DISTINCT sourceID FROM EventHistory")
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var sourceID string
		err = rows.Scan(&sourceID)
		if err != nil {
			continue
		}
		result = append(result, sourceID)
	}
	return result, nil
}

func (s *DBService) QueryEvent(name *string, sourceID *string, offset int, length int) ([]message_bus.Event, error) {
	var result []message_bus.Event

	find := []string{}
	if name != nil {
		find = append(find, "name = '"+*name+"'")
	}
	if sourceID != nil {
		find = append(find, "sourceID = '"+*sourceID+"'")
	}
	var where string
	if len(find) > 0 {
		where = " WHERE " + find[0]
		for i := 1; i < len(find); i++ {
			where += " AND " + find[i]
		}
	}
	rows, err := s.db.Query("SELECT uuid, properties, sourceID, name, timestamp FROM EventHistory" + where + " ORDER BY timestamp DESC LIMIT " + strconv.Itoa(length) + " OFFSET " + strconv.Itoa(offset))
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var uuid string
		var properties string
		var sourceID string
		var name string
		var timestamp string
		err = rows.Scan(&uuid, &properties, &sourceID, &name, &timestamp)
		if err != nil {
			continue
		}
		var event message_bus.Event
		event.Name = name
		event.SourceID = sourceID
		event.Uuid = &uuid
		event.Properties = make(map[string]string)
		err = json.Unmarshal([]byte(properties), &event.Properties)
		if err != nil {
			continue
		}
		t, err := convertTimestamp(timestamp)
		if err != nil {
			continue
		}
		event.Timestamp = &t
		result = append(result, event)
	}
	return result, nil
}
