package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kostiantyn-Liapkalo/comment-service/store"
	_ "github.com/lib/pq" // Драйвер для PostgreSQL
)

func handleCommentRequest(cs *store.CommentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Обробляємо запит на отримання коментарів
			comments := cs.GetComments()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(comments)
		case http.MethodPost:
			// Обробляємо запит на додавання нового коментаря
			var comment store.Comment
			if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}
			if err := cs.AddComment(&comment); err != nil {
				http.Error(w, "Failed to add comment", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleLikeRequest(cs *store.CommentStore, like bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if like {
			cs.LikeComment(id)
		} else {
			cs.DislikeComment(id)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	// Ініціалізація та підключення до бази даних PostgreSQL
	dsn := "user=postgres password=your-password dbname=comments sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Створення об'єкта CommentStore
	cs := store.NewCommentStore(db)

	// Реєстрація обробників для відповідних HTTP-шляхів
	http.HandleFunc("/comments", handleCommentRequest(cs))
	http.HandleFunc("/comments/like", handleLikeRequest(cs, true))
	http.HandleFunc("/comments/dislike", handleLikeRequest(cs, false))

	// Запуск HTTP-сервера
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
