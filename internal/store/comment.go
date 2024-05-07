package store

import (
	"database/sql"
	"errors"
)

// Comment представляє структуру коментарія з лайками і дизлайками
type Comment struct {
	ID       int64  `json:"id"`
	Content  string `json:"content"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}

// CommentStore містить методи для роботи з коментарями в базі даних
type CommentStore struct {
	db *sql.DB
}

// NewCommentStore створює новий об'єкт CommentStore
func NewCommentStore(db *sql.DB) *CommentStore {
	return &CommentStore{db: db}
}

// AddComment додає новий коментар до бази даних
func (cs *CommentStore) AddComment(comment *Comment) error {
	query := `INSERT INTO comments (content) VALUES ($1) RETURNING id, likes, dislikes`
	err := cs.db.QueryRow(query, comment.Content).Scan(&comment.ID, &comment.Likes, &comment.Dislikes)
	if err != nil {
		return err
	}
	return nil
}

// GetComment отримує коментар з бази даних за його ID
func (cs *CommentStore) GetComment(id int64) (*Comment, error) {
	query := `SELECT id, content, likes, dislikes FROM comments WHERE id = $1`
	comment := &Comment{}
	err := cs.db.QueryRow(query, id).Scan(&comment.ID, &comment.Content, &comment.Likes, &comment.Dislikes)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}
	return comment, nil
}

// UpdateComment оновлює лайки або дизлайки коментаря в базі даних
func (cs *CommentStore) UpdateComment(id int64, likes, dislikes int) error {
	query := `UPDATE comments SET likes = $1, dislikes = $2 WHERE id = $3`
	_, err := cs.db.Exec(query, likes, dislikes, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteComment видаляє коментар з бази даних за його ID
func (cs *CommentStore) DeleteComment(id int64) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := cs.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
