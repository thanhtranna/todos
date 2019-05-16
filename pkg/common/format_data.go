package common

type DataFormat struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Format(status int, message string, data interface{}) DataFormat {
	return DataFormat{
		Status: status,
		Message: message,
		Data: data,
	}
}
