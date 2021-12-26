package stream

import (
	"math/rand"
)

type RandomNumberServer struct {
	UnimplementedRandomServiceServer
}

func (server *RandomNumberServer) AddRandomNumber(request *RandomNumberRequest, stream RandomService_AddRandomNumberServer) error {
	var randomInt int64 = rand.Int63()
	for i := 0; i < 100; i++ {
		var response *RandomNumberResponse = &RandomNumberResponse{
			OriginalNumber: request.Number,
			TotalNumber:    randomInt + int64(i),
		}

		var streamErr error = stream.Send(response)
		if streamErr != nil {
			panic(streamErr)
		}
	}

	return nil
}

func (server *RandomNumberServer) SubstractRandomNumber(request *RandomNumberRequest, stream RandomService_SubstractRandomNumberServer) error {
	var randomInt int64 = rand.Int63()
	for i := 0; i < 100; i++ {
		var response *RandomNumberResponse = &RandomNumberResponse{
			OriginalNumber: request.Number,
			TotalNumber:    randomInt - int64(i),
		}

		var streamErr error = stream.Send(response)
		if streamErr != nil {
			panic(streamErr)
		}
	}

	return nil
}
