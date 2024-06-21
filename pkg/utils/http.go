package utils

func GetMessage(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func GetMessageError(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}
