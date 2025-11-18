package events

import "log"

// ProcessEvent — базовая точка входа для обработки событий
func ProcessEvent(evt map[string]interface{}) {
    eventType, ok := evt["type"].(string)
    if !ok {
        log.Printf("[events] WARNING: событие без типа: %v", evt)
        return
    }

    log.Printf("[events] обработка события типа: %s", eventType)

    switch eventType {
    case "StasisStart":
        HandleStasisStart(evt)
    case "ChannelStateChange":
        HandleChannelStateChange(evt)
    case "ChannelDestroyed":
        HandleChannelDestroyed(evt)
    case "ChannelDtmfReceived":
        HandleChannelDtmf(evt)
    case "ChannelHangupRequest":
        HandleChannelHangupRequest(evt)
    case "StasisEnd":
        HandleStasisEnd(evt)
    default:
        log.Printf("[events] неизвестный тип события: %s", eventType)
    }
}
