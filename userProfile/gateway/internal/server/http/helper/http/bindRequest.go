package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BindRequest(response http.ResponseWriter, request *http.Request, dto interface{}) bool {
	if request.Method == http.MethodGet || request.Method == http.MethodDelete {
		return bindGET(response, request, dto)
	}

	return bindPOST(response, request, dto)
}

func bindGET(response http.ResponseWriter, request *http.Request, dto interface{}) bool {
	if err := request.ParseForm(); err != nil {
		ErrorResponse(err, response, http.StatusUnprocessableEntity)
		return false
	}

	m := map[string]string{}
	for k, v := range request.Form {
		m[k] = v[0]
	}

	data, err := json.Marshal(m)
	if err != nil {
		ErrorResponse(err, response, http.StatusUnprocessableEntity)
		return false
	}

	if err = json.Unmarshal(data, dto); err != nil {
		ErrorResponse(err, response, http.StatusUnprocessableEntity)
		return false
	}

	return true
}

func bindPOST(response http.ResponseWriter, request *http.Request, dto interface{}) bool {
	requestBody := request.Body

	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		ErrorResponse(err, response, http.StatusUnprocessableEntity)
		return false
	}

	err = json.Unmarshal(body, &dto)
	if err != nil {
		ErrorResponse(err, response, http.StatusUnprocessableEntity)
		return false
	}

	return true
}
