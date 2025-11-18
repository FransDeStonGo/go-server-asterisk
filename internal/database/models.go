package database

import "time"

// Call модель звонка
type Call struct {
    ID        int64     `db:"id"`
    From      string    `db:"from"`
    To        string    `db:"to"`
    Channel   string    `db:"channel"`
    Status    string    `db:"status"`
    CreatedAt time.Time `db:"created_at"`
    Duration  int64     `db:"duration"`
}

