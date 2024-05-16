package main

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

func GetBooksBySubject(ctx context.Context, req *BookReqDTO) (*GetBooksRespDTO, error) {
	redis := NewServRedis()

	var resp *GetBooksRespDTO

	dataRedis, err := redis.GetData(ctx, req.Subject)

	if err != nil {
		log.Printf("unable to GET data from redis. error: %v", err)
	}

	if dataRedis != "" {
		// get data from redis if is there
		_ = json.Unmarshal([]byte(dataRedis), &resp)

		log.Println("data from redis")
		return resp, nil

	}

	resp, err = GetIntegrationBooksBySubject(req.Subject)
	log.Println("data not from redis")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	redisData, _ := json.Marshal(resp)
	ttl := time.Duration(2) * time.Minute

	// set data to redis
	err = redis.SetData(context.Background(), req.Subject, redisData, ttl)
	if err != nil {
		log.Printf("unable to SET data. error: %v", err)
		return nil, err
	}

	return resp, nil
}
