package helper

type response struct {
	Code    int
	Status  string
	Message string
	Data    interface{}
}

func Response(code int, status string, message string, data interface{}) response {
	response := response{}
	response.Code = code
	response.Status = status
	response.Message = message
	response.Data = data

	return response
}

type responseAuth struct {
	Token string
}

func ResponseAuth(token string) responseAuth {
	response := responseAuth{}
	response.Token = token
	
	return response
}
