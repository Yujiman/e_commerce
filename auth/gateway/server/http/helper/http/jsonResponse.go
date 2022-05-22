package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func EmptyResponse(responseWriter http.ResponseWriter) {
	responseWriter.WriteHeader(http.StatusOK)

	empty, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		ErrorResponse(err, responseWriter, http.StatusInternalServerError)
		return
	}

	_, err = responseWriter.Write(empty)
	if err != nil {
		ErrorResponse(err, responseWriter, http.StatusInternalServerError)
		return
	}
}

func JsonResponse(responseWriter http.ResponseWriter, body interface{}) {

	responseWriter.WriteHeader(http.StatusOK)

	var jsonByte []byte
	var err error

	switch v := body.(type) {
	case proto.Message: // If body from gRPC response
		mOpt := protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		}

		jsonByte, err = mOpt.Marshal(body.(proto.Message))
		if err != nil {
			ErrorResponse(err, responseWriter, http.StatusInternalServerError)
			return
		}

	case map[string]interface{}: // if custom body
		jsonByte, err = json.Marshal(body)
		if err != nil {
			ErrorResponse(err, responseWriter, http.StatusInternalServerError)
			return
		}
	default:
		ErrorResponse(
			fmt.Errorf("wrong body response type: %T!\n", v),
			responseWriter,
			http.StatusInternalServerError,
		)
		return
	}

	_, err = responseWriter.Write(jsonByte)
	if err != nil {
		ErrorResponse(err, responseWriter, http.StatusInternalServerError)
		return
	}
}
