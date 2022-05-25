package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	errorHelper "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/error"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/utils"
)

func ErrorResponse(err error, w http.ResponseWriter, statusCode int) {
	msg := err.Error()

	// Trim message
	msg = strings.ReplaceAll(msg, "\n", "")

	// Delete Rpc Error Description
	if strings.HasPrefix(msg, "rpc error:") {
		i := strings.Index(msg, "desc = ")
		msg = msg[i+7:]
	}

	// Group errors if many
	message := make(map[string]interface{})

	oneMsg, msgMap := errorHelper.SplitErrorsToMap(msg)
	message["error"] = oneMsg
	if msgMap != nil {
		message["errors"] = msgMap
		delete(message, "error")
	}

	if statusCode < 100 || statusCode >= 600 {
		utils.LogPrintf("Send wrong http error msg=:%s, code=%d\n", message, statusCode)
		w.WriteHeader(522)
		_, err = w.Write([]byte("{\"error\": \"" + msg + "\"}"))
		if err != nil {
			utils.LogPrintln("Error from ResponseWriter: " + err.Error())
			return
		}
		return
	}

	log.Println(statusCode)
	// Write error
	w.WriteHeader(statusCode)

	errResp, _ := json.Marshal(&message)

	_, err = w.Write(errResp)
	if err != nil {
		utils.LogPrintln("Error from ResponseWriter: " + err.Error())
		return
	}
}
