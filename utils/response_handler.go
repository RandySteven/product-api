package utils

import (
	"encoding/json"
	"net/http"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/response"
)

func ResponseHandler(res http.ResponseWriter, httpStatus int, resp response.Response) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpStatus)
	json.NewEncoder(res).Encode(resp)
}
