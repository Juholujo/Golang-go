package main

import "fmt"

// Опции сервера
type Options struct {
	MaxConnections int
	Timeout        int
	Protocol       string
	Port           int
	Logging        bool
}

// Сервер с опциями
type Server struct {
	Opts Options
}

// Конструктор сервера
func newServer(maxConn int, timeout int, protocol string, port int, logging bool) *Server {
	return &Server{
		Opts: Options{
			MaxConnections: maxConn,
			Timeout:        timeout,
			Protocol:       protocol,
			Port:           port,
			Logging:        logging,
		},
	}
}

func main() {
	// Создаем сервер с заданными параметрами
	server := newServer(100, 30, "http", 8080, true)

	// Параметры сервера
	fmt.Printf("Server options: %+v\n", server.Opts)

	// Изменяем параметры вручную
	server.Opts.MaxConnections = 200
	fmt.Printf("Updated server options: %+v\n", server.Opts)
}
