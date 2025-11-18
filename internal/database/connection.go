package database

import (
	"database/sql"
)

// OpenDB открывает соединение с БД по пути (строка пути, например './callcenter.db')
// Использует имя драйвера "sqlite" — драйвер должен быть подключен в main или через go.mod, если требуется.
func OpenDB(path string) (*sql.DB, error) {
    return sql.Open("sqlite", path)
}
