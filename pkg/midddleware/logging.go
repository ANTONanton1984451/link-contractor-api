package midddleware

import "net/http"

type (
	Logger interface {
		Error(args ...interface{})
		Errorf(tmpl string, args ...interface{})

		Info(args ...interface{})
		Infof(tmpl string, args ...interface{})
	}

	HandlerFunc func(w http.ResponseWriter, r *http.Request) error
)

// NewLogMiddleware логирование входящих запросов
func NewLogMiddleware(logger Logger, handlerFunc HandlerFunc) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logger.Infof("got request with method %s and path %s", request.Method, request.RequestURI)
		err := handlerFunc(writer, request)
		if err != nil {
			logger.Errorf("error while execute request: %s", err)
		}
	})
}
