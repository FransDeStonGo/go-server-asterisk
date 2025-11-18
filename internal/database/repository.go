package database

import (
	"database/sql"
	"time"
)

// Repository простая оболочка для работы с БД
type Repository struct {
    DB *sql.DB
}

// NewRepository создает репозиторий
func NewRepository(db *sql.DB) *Repository {
    return &Repository{DB: db}
}

// SaveCall сохраняет запись о звонке (минимальная реализация)
func (r *Repository) SaveCall(c *Call) (int64, error) {
    if r.DB == nil {
        return 0, nil
    }
    now := time.Now()
    res, err := r.DB.Exec(
        "INSERT INTO calls(from, to, channel, status, created_at, duration) VALUES(?,?,?,?,?,?)",
        c.From, c.To, c.Channel, c.Status, now, c.Duration,
    )
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}
