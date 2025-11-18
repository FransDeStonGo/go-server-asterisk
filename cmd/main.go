package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "sync"
    "syscall"

    asterisk "go-server-asterisk/internal/websocket"
    "go-server-asterisk/internal/config"
    "go-server-asterisk/internal/database"
    "go-server-asterisk/internal/events"

    _ "modernc.org/sqlite"
)

func main() {
    log.Println("🚀 Запуск Call Center Server...")

    cfg := config.LoadFromEnv()

    db, err := database.OpenDB(cfg.DBPath)
    if err != nil {
        log.Fatalf("Ошибка открытия БД: %v", err)
    }
    defer db.Close()

    // Простая инициализация таблицы
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS calls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        "from" TEXT,
        "to" TEXT,
        channel TEXT,
        status TEXT,
        created_at DATETIME,
        duration INTEGER
    )`)
    if err != nil {
        log.Fatalf("Ошибка создания таблицы: %v", err)
    }

    repo := database.NewRepository(db)
    _ = repo

    ariCfg := asterisk.ARIConfig{
        URL:      cfg.ARIURL,
        Username: cfg.ARIUsername,
        Password: cfg.ARIPassword,
        AppName:  cfg.ARIAppName,
    }

    client, err := asterisk.ConnectAndListen(ariCfg)
    if err != nil {
        log.Fatalf("Ошибка подключения к ARI: %v", err)
    }
    defer client.Close()

    eventsCh, errCh := client.ReadEvents()

    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        for {
            select {
            case <-ctx.Done():
                log.Println("Получен сигнал остановки")
                return
            case err := <-errCh:
                if err != nil {
                    log.Printf("WebSocket error: %v", err)
                    stop()
                    return
                }
            case evt, ok := <-eventsCh:
                if !ok {
                    log.Println("Канал событий закрыт")
                    stop()
                    return
                }
                // Обрабатываем событие
                events.ProcessEvent(evt)
            }
        }
    }()

    log.Println("✅ Слушатель событий запущен")
    wg.Wait()
    log.Println("👋 Завершение работы приложения")
}

