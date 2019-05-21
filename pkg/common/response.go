package common

type UserInfo struct {
	UserID		string 	`json:"user_id"`
	UserName	string	`json:"user_name"`
}

func Response(status int, message string, data interface{}) DataFormat {
	return DataFormat{
		Status: status,
		Message: message,
		Data: data,
	}
}

func ResponseSuccess(data interface{}) DataFormat {
	return Response(1, "success", data)
}

func ResponseFail(data interface{}) DataFormat {
	return Response(0, "fail", data)
}

func ResponseWithError(err string, data interface{}) DataFormat {
	return Response(0, err, data)
}