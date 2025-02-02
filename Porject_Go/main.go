package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// Названия наших клиентов
const (
	Client1 = "client1"
	Client2 = "client2"
	Client3 = "client3"
)

// ClientStats хранит счётчики по статус-кодам для каждого клиента
type ClientStats struct {
	StatusCount map[int]int `json:"status_count"`
}

// GlobalStats хранит статистику по всем клиентам
type GlobalStats struct {
	mu      sync.Mutex
	clients map[string]*ClientStats
}

var globalStats = GlobalStats{
	clients: map[string]*ClientStats{
		Client1: {StatusCount: make(map[int]int)},
		Client2: {StatusCount: make(map[int]int)},
		Client3: {StatusCount: make(map[int]int)},
	},
}

// ---- Рейт-лимит на сервере (token bucket, 5 запросов/сек) ----
var tokenBucket = make(chan struct{}, 5)

// initTokenBucket запускает пополнение «ведра» токенов каждые 1 сек
func initTokenBucket() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-ticker.C
			// Добавляем до 5 токенов, если есть место
			for i := 0; i < 5; i++ {
				select {
				case tokenBucket <- struct{}{}:
				default:
					// Канал полон — не добавляем
					break
				}
			}
		}
	}()
}

// handlerHealth — проверка сервера (GET /health)
func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is up\n"))
}

// handlerStats — отдаёт статистику (GET /stats)
func handlerStats(w http.ResponseWriter, r *http.Request) {
	globalStats.mu.Lock()
	defer globalStats.mu.Unlock()

	// Для удобного вывода сделаем структуру
	type Result struct {
		Client        string         `json:"client"`
		StatusCount   map[int]int    `json:"status_count"`
		PositiveTotal int            `json:"positive_total"`
		NegativeTotal int            `json:"negative_total"`
	}

	results := []Result{}

	// Для общей суммы по всем клиентам
	totalStatusCount := make(map[int]int)
	var totalPositive, totalNegative int

	for clientID, stats := range globalStats.clients {
		var pos, neg int
		for code, count := range stats.StatusCount {
			// Добавляем в общий map
			totalStatusCount[code] += count

			// Положительные/отрицательные
			switch code {
			case http.StatusOK, http.StatusAccepted:
				pos += count
			case http.StatusBadRequest, http.StatusInternalServerError, http.StatusTooManyRequests:
				neg += count
			}
		}

		results = append(results, Result{
			Client:        clientID,
			StatusCount:   stats.StatusCount,
			PositiveTotal: pos,
			NegativeTotal: neg,
		})

		totalPositive += pos
		totalNegative += neg
	}

	// Добавим сводную строку "ALL_CLIENTS"
	allClients := Result{
		Client:        "ALL_CLIENTS",
		StatusCount:   totalStatusCount,
		PositiveTotal: totalPositive,
		NegativeTotal: totalNegative,
	}
	results = append(results, allClients)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// handlerRoot — отвечает на GET / (просто привет) и POST / (имитация ответа)
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Go server!\n"))
	case http.MethodPost:
		handlePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed\n"))
	}
}

// handlePost — обрабатывает POST-запросы: рандомный статус (70% положит., 30% отрицат.)
// + ограничение 5 запросов/сек на сервере (tokenBucket)
func handlePost(w http.ResponseWriter, r *http.Request) {
	// Проверяем токен
	select {
	case <-tokenBucket:
		// ок, есть токен, продолжаем
	default:
		// нет токенов => 429
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Server limit exceeded\n"))
		saveStats(r, http.StatusTooManyRequests)
		return
	}

	// Генерируем случайный ответ
	rnd := rand.Intn(100) // [0..99]
	var status int
	if rnd < 70 {
		// 70% — положительные (50% на 200, 50% на 202)
		if rand.Intn(2) == 0 {
			status = http.StatusOK
		} else {
			status = http.StatusAccepted
		}
	} else {
		// 30% — отрицательные (50% на 400, 50% на 500)
		if rand.Intn(2) == 0 {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
	}

	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("Response with status %d\n", status)))

	saveStats(r, status)
}

// saveStats сохраняет статус-код в глобальную структуру для конкретного клиента
func saveStats(r *http.Request, code int) {
	clientID := r.Header.Get("X-Client-Id")
	if clientID == "" {
		clientID = "unknown"
	}

	globalStats.mu.Lock()
	defer globalStats.mu.Unlock()

	// Если вдруг клиент ещё не инициализирован
	if _, exists := globalStats.clients[clientID]; !exists {
		globalStats.clients[clientID] = &ClientStats{StatusCount: make(map[int]int)}
	}

	globalStats.clients[clientID].StatusCount[code]++
}

// ====== Клиентская часть (для демонстрации) ======

// создаём rate limiter (token bucket) на 5 req/sec
func newClientRateLimiter() chan struct{} {
	ch := make(chan struct{}, 5)
	// Заполним 5 токенов в начале
	for i := 0; i < 5; i++ {
		ch <- struct{}{}
	}
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-ticker.C
			for i := 0; i < 5; i++ {
				select {
				case ch <- struct{}{}:
				default:
					// Полно, не добавляем
					break
				}
			}
		}
	}()
	return ch
}

// runClient — запускает N воркеров, каждый шлёт часть запросов
func runClient(clientID string, url string, totalRequests int, workers int) {
	log.Printf("[%s] starting client with %d requests total...\n", clientID, totalRequests)
	clientLimiter := newClientRateLimiter() // чтобы не отправлять >5 req/s
	var wg sync.WaitGroup
	wg.Add(workers)

	requestsPerWorker := totalRequests / workers

	for w := 0; w < workers; w++ {
		go func(workerID int) {
			defer wg.Done()
			for i := 0; i < requestsPerWorker; i++ {
				// Ждём токен, чтобы не превышать 5 req/s
				<-clientLimiter

				req, err := http.NewRequest(http.MethodPost, url, nil)
				if err != nil {
					log.Printf("[%s] worker=%d error creating request: %v", clientID, workerID, err)
					continue
				}
				req.Header.Set("X-Client-Id", clientID)

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Printf("[%s] worker=%d error sending request: %v", clientID, workerID, err)
					continue
				}
				_ = resp.Body.Close()
			}
		}(w)
	}

	wg.Wait()
	log.Printf("[%s] finished sending %d requests.", clientID, totalRequests)
}

// runHealthChecker — клиент, который каждые 5 секунд шлёт GET /health
func runHealthChecker(clientID, url string, stopChan <-chan struct{}) {
	clientLimiter := newClientRateLimiter()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stopChan:
			log.Printf("[%s] Stopping health checker...", clientID)
			return
		case <-ticker.C:
			<-clientLimiter
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("[%s] Health check failed: %v", clientID, err)
				continue
			}
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				log.Printf("[%s] Server is OK (status=%d)", clientID, resp.StatusCode)
			} else {
				log.Printf("[%s] Server NOT OK (status=%d)", clientID, resp.StatusCode)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Загружаем .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Нет .env файла или произошла ошибка при загрузке:", err)
	}

	// 2. Читаем PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
		log.Printf("PORT не задан в .env, используется по умолчанию %s", port)
	}

	// 3. Инициализируем rate limit на сервере
	initTokenBucket()

	// 4. Настраиваем серверные роуты
	http.HandleFunc("/health", handlerHealth)
	http.HandleFunc("/stats", handlerStats)
	http.HandleFunc("/", handlerRoot)

	// 5. Запускаем сервер в горутине
	server := &http.Server{Addr: port}
	go func() {
		log.Printf("Starting server on %s ...", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// 6. Запускаем клиентов (Client1 и Client2 шлют POSTы, Client3 health-check)
	go runClient(Client1, "http://localhost"+port, 100, 2)
	go runClient(Client2, "http://localhost"+port, 100, 2)

	stopHealth := make(chan struct{})
	go runHealthChecker(Client3, "http://localhost"+port+"/health", stopHealth)

	// 7. Ждём пока первые два клиента закончат отправку. Для упрощения — Sleep 30 сек.
	time.Sleep(30 * time.Second)

	// Останавливаем клиента Client3
	close(stopHealth)

	log.Println("All clients finished sending requests. You can now GET /stats to see results.")

	// 8. Подождём ещё чуть-чуть, чтобы можно было запросить /stats
	time.Sleep(5 * time.Second)

	// Завершаем сервер
	if err := server.Close(); err != nil {
		log.Printf("Error closing server: %v", err)
	}
	log.Println("Program finished.")
}
