package asterisk

import "log"

// HandleEvent обрабатывает одно событие ARI.
func HandleEvent(evt map[string]interface{}) {
    // Простая маршрутизация по типу события, если есть поле "type" или "event"
    if t, ok := evt["type"]; ok {
        log.Printf("[ARI event] type=%v", t)
        return
    }
    if e, ok := evt["event"]; ok {
        log.Printf("[ARI event] event=%v", e)
        return
    }
    log.Printf("[ARI event] unknown: %v", evt)
}
