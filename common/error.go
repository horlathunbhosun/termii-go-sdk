package common

type ErrorMessage struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// IsErrorMessageEmpty checks if the ErrorMessage is empty
func IsErrorMessageEmpty(err ErrorMessage) bool {
	return err.Status == "" && err.StatusCode == 0 && err.Message == ""
}
