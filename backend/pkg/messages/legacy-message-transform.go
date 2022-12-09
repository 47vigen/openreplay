package messages

func transformDeprecated(msg Message) Message {
	switch m := msg.(type) {
	case *JSExceptionDeprecated:
		return &JSException{
			Name:     m.Name,
			Message:  m.Message,
			Payload:  m.Payload,
			Metadata: "{}",
		}
	case *SessionEndDeprecated:
		return &SessionEnd{
			Timestamp:     m.Timestamp,
			EncryptionKey: "",
		}
	case *Fetch:
		return &NetworkRequest{
			Type:      "fetch",
			Method:    m.Method,
			URL:       m.URL,
			Request:   m.Request,
			Response:  m.Response,
			Status:    m.Status,
			Timestamp: m.Timestamp,
			Duration:  m.Duration,
		}
	}
	return msg
}
