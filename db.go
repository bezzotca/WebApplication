package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

// Параметры подключения. У lib/pq нет структуры Config — строка подключения
// собирается в формате "ключ=значение".
const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPass = "1234"
	dbName = "public"
)

// openDB создаёт пул соединений и проверяет, что база отвечает.
// Закрывать пул должен вызывающий: defer db.Close().
func openDB() *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	// Или сразу строкой: pq.NewConnector("host=localhost dbname=pqgo")
	c, err := pq.NewConnector(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Создаём пул соединений.
	db := sql.OpenDB(c)

	// sql.OpenDB к базе не подключается — соединение открывается лениво,
	// при первом запросе. Ping устанавливает его сейчас, чтобы ошибка
	// конфигурации всплыла на старте, а не посреди работы.
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}

	return db
}
