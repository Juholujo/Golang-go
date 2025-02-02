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

// Определяем тип функции настройки
type Option func(*Options)

// Функция-конструктор сервера с использованием функций настройки
func newServer(opts ...Option) *Server {
	// Задаем значения по умолчанию
	defaultOptions := Options{
		MaxConnections: 100,
		Timeout:        30,
		Protocol:       "http",
		Port:           8080,
		Logging:        true,
	}

	// Применяем функции настройки
	for _, opt := range opts {
		opt(&defaultOptions)
	}

	return &Server{Opts: defaultOptions}
}

// Функции настройки
func withMaxConn(maxConn int) Option {
	return func(o *Options) {
		o.MaxConnections = maxConn
	}
}

func withTimeout(timeout int) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func withProtocol(protocol string) Option {
	return func(o *Options) {
		o.Protocol = protocol
	}
}

func withPort(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func withLogging(logging bool) Option {
	return func(o *Options) {
		o.Logging = logging
	}
}

func main() {
	// Создаем сервер с заданными параметрами
	server := newServer(
		withMaxConn(200),
		withTimeout(60),
		withProtocol("https"),
		withPort(9090),
		withLogging(false),
	)

	// Параметры сервера
	fmt.Printf("Server options: %+v\n", server.Opts)
}
