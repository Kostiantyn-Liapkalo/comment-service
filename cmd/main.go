package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/your-username/comment-service/internal/store"
    _ "github.com/lib/pq" // Драйвер для PostgreSQL
)

func handleCommentRequest(cs *store.CommentStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Обработка запросов на создание, получение и удаление комментариев
    }
}

func handleLikeRequest(cs *store.CommentStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Обработка запросов на лайк комментария
    }
}

func handleDislikeRequest(cs *store.CommentStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Обработка запросов на дизлайк комментария
    }
}

func main() {
    // Инициализация и подключение к базе данных PostgreSQL
    dsn := "user=postgres password=your-password dbname=comments sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // Создание объекта CommentStore
    cs, err := store.NewCommentStore(db)
    if err != nil {
        log.Fatalf("Failed to create comment store: %v", err)
    }

    // Регистрация обработчиков на соответствующие HTTP-пути
    http.HandleFunc("/comments", handleCommentRequest(cs))
    http.HandleFunc("/comments/like", handleLikeRequest(cs))
    http.HandleFunc("/comments/dislike", handleDislikeRequest(cs))

    // Запуск HTTP-сервера
    fmt.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}