package internal

import "encoding/json"

type Request struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Response struct {
	Body  string `json:"body"`
	Error string `json:"error"`
}

func ParseRequest(payload []byte) (*Request, error) {
	request := &Request{}
	if err := json.Unmarshal(payload, request); err != nil {
		return nil, err
	}

	return request, nil
}

func SerializeResponse(body string, err error) ([]byte, error) {
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	return json.Marshal(&Response{
		Body:  body,
		Error: errorMessage,
	})
}
