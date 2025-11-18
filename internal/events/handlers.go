package events

import "log"

// HandleStasisStart обработчик события StasisStart
func HandleStasisStart(evt map[string]interface{}) {
    log.Printf("[events] StasisStart: %v", evt)
}

// HandleChannelStateChange обработчик смены состояния канала
func HandleChannelStateChange(evt map[string]interface{}) {
    log.Printf("[events] ChannelStateChange: %v", evt)
}

// HandleChannelDestroyed обработчик завершения канала
func HandleChannelDestroyed(evt map[string]interface{}) {
    log.Printf("[events] ChannelDestroyed: %v", evt)
}

// HandleChannelDtmf обработчик DTMF событий
func HandleChannelDtmf(evt map[string]interface{}) {
    log.Printf("[events] ChannelDtmfReceived: %v", evt)
}

// HandleChannelHangupRequest обработчик запроса на повесить трубку
func HandleChannelHangupRequest(evt map[string]interface{}) {
    log.Printf("[events] ChannelHangupRequest: %v", evt)
}

// HandleStasisEnd обработчик завершения Stasis
func HandleStasisEnd(evt map[string]interface{}) {
    log.Printf("[events] StasisEnd: %v", evt)
}
