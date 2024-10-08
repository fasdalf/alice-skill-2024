package main

import (
	"flag"
	"os"
)

var (
	flagRunAddr  string
	flagLogLevel string
	// переменная будет содержать параметры соединения с СУБД
	flagDatabaseURI string
)

func parseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&flagLogLevel, "l", "info", "log level")
	// обрабатываем аргумент -d
	flag.StringVar(&flagDatabaseURI, "d", "", "database URI")
	flag.Parse()

	// для случаев, когда в переменной окружения RUN_ADDR присутствует непустое значение,
	// переопределим адрес запуска сервера,
	// даже если он был передан через аргумент командной строки
	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		flagRunAddr = envRunAddr
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		flagLogLevel = envLogLevel
	}
	// обрабатываем переменную окружения DATABASE_URI
	if envDatabaseURI := os.Getenv("DATABASE_URI"); envDatabaseURI != "" {
		flagDatabaseURI = envDatabaseURI
	}
}
