package internal

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
	"github.com/google/uuid"
)

func GetResult(conn redis.Conn, id string) (*Response, error) {
	reply, err := redis.String(conn.Do("GET", id))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}

		return nil, err
	}

	response := &Response{}
	if err := json.Unmarshal([]byte(reply), response); err != nil {
		return nil, err
	}

	return response, err
}

func PublishWork(conn redis.Conn, url string) (string, error) {
	id := uuid.New().String()

	request, err := json.Marshal(&Request{
		ID:  id,
		URL: url,
	})

	if err != nil {
		return "", err
	}

	if _, err := conn.Do("PUBLISH", "work", string(request)); err != nil {
		return "", err
	}

	return id, nil
}
