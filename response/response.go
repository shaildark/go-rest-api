package response

func APIResponse(code int64, message string, data any, payload any, extra any) map[string]any {
	if data == nil {
		data = map[string]any{}
	}

	if payload == nil {
		payload = map[string]any{}
	}

	if extra == nil {
		extra = map[string]any{}
	}

	return map[string]any{
		"code":    code,
		"message": message,
		"data":    data,
		"payload": payload,
		"extra":   extra,
	}
}
