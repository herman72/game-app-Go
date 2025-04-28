package httpmsg

import (
	"game-app-go/pkg/errmsg"
	"game-app-go/pkg/richerror"
	"net/http"
)



func Error(err error)(message string, code int){
	switch re := err.(type){
	case richerror.RichError:
		msg := re.Message()
		//  we should not expose unexpected erorrs
		kind := mapkindToHTTPStatusCode(re.Kind())
		if kind >= 500 {
			msg = errmsg.ErrorMsgSomethingWentWrong
		}
		return msg, kind
	default:
		return err.Error(), http.StatusBadRequest
	}

}

func mapkindToHTTPStatusCode(kind richerror.Kind)int {
	switch kind {
		case richerror.KindInvalid:
			return http.StatusUnprocessableEntity
		case richerror.KindNotFound:
			return http.StatusNotFound
		case richerror.KindForbidden:
			return http.StatusForbidden
		case richerror.KindUnexpected:
			return http.StatusInternalServerError
		default:
			return http.StatusBadRequest
	}
}
