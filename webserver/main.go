package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ── METRICS ──
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"path", "method", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "How long each request took",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	httpErrorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of error responses",
		},
		[]string{"path", "error_type"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(httpErrorsTotal)
}

// ── CHAOS ENGINE ──
// Yeh function randomly decide karta hai ki request mein kya problem aayegi
func chaos(path string) (slowdown bool, err bool) {
	n := rand.Intn(10) // 0 to 9

	// 20% chance of slow response
	slowdown = n >= 8

	// 30% chance of error
	err = n >= 7

	return
}

// ── MIDDLEWARE ──
func metricsMiddleware(next http.HandlerFunc, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		slow, fail := chaos(path)

		// Simulate slow response — 1 to 3 seconds delay
		if slow {
			delay := time.Duration(1000+rand.Intn(2000)) * time.Millisecond
			log.Printf("⚠️  SLOW REQUEST on %s — delaying %v", path, delay)
			time.Sleep(delay)
		}

		// Simulate error — return 500
		if fail {
			log.Printf("❌ ERROR on %s — returning 500", path)
			httpErrorsTotal.WithLabelValues(path, "internal_server_error").Inc()
			httpRequestsTotal.WithLabelValues(path, r.Method, "500").Inc()
			http.Error(w, "Internal Server Error — something went wrong!", http.StatusInternalServerError)
			return
		}

		// Normal request
		next(w, r)

		duration := time.Since(start).Seconds()
		httpRequestsTotal.WithLabelValues(path, r.Method, "200").Inc()
		httpRequestDuration.WithLabelValues(path).Observe(duration)

		log.Printf("✅ OK %s — took %.2fs", path, duration)
	}
}

// ── HANDLERS ──
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello! Everything is fine here.")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintln(w, "POST request successful")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// ── LOAD SIMULATOR ──
// Yeh automatically requests bhejta rehta hai taaki Grafana mein data dikhe
func simulateLoad() {
	go func() {
		for {
			// Har 500ms mein ek request
			time.Sleep(500 * time.Millisecond)
			_, err := http.Get("http://localhost:8080/hello")
			if err != nil {
				log.Println("Load simulator error:", err)
			}
		}
	}()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", metricsMiddleware(formHandler, "/form"))
	http.HandleFunc("/hello", metricsMiddleware(helloHandler, "/hello"))
	http.Handle("/metrics", promhttp.Handler())

	// Auto load simulator start karo
	simulateLoad()

	fmt.Println("🚀 Server started at :8080")
	fmt.Println("📊 Metrics at :8080/metrics")
	fmt.Println("⚡ Load simulator running...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}