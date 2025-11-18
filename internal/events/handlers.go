package events

import "log"

func HandleStasisStart(evt map[string]interface{}) {
    channel := evt["channel"].(map[string]interface{})
    
    channelName := channel["name"].(string)
    
    caller := channel["caller"].(map[string]interface{})
    callerNum := caller["number"].(string)
    
    dialplan := channel["dialplan"].(map[string]interface{})
    calleeNum := dialplan["exten"].(string)
    
    log.Printf("📞 НОВЫЙ ЗВОНОК: %s (%s) -> %s", 
        callerNum, channelName, calleeNum)
    
    // Сохранить в БД
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
