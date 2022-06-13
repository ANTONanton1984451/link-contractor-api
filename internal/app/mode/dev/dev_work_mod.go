package dev

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"link-contractor-api/internal/entrypoint"
	"link-contractor-api/pkg/midddleware"
	"net/http"

	"go.uber.org/zap"
)

type (
	Presenter interface {
		SomethingWentWrong() []byte
	}

	UserPl struct {
		UserExternalID int64 `json:"user_external_id"`
	}
)

func StartWorking(ep entrypoint.Entrypoint, workPort string, logger *zap.Logger, pr Presenter) error {

	mux := http.NewServeMux()

	mux.Handle("/form", midddleware.NewLogMiddleware(logger.Sugar(), func(writer http.ResponseWriter, request *http.Request) error {
		form, err := readForm()
		if err != nil {
			return fmt.Errorf("read form: %w", err)
		}
		writer.Write(form)
		return nil
	}))

	mux.Handle("/dev", midddleware.NewLogMiddleware(logger.Sugar(), func(w http.ResponseWriter, r *http.Request) error {
		body := r.Body
		defer body.Close()

		pl, err := io.ReadAll(body)
		var up UserPl
		if err := json.Unmarshal(pl, &up); err != nil {
			return fmt.Errorf("unmarshal payload: %w", err)
		}

		if err != nil {
			w.Write(pr.SomethingWentWrong())
			return err
		}

		resp, err := ep.DoAction(r.Context(), up.UserExternalID, pl)
		if err != nil {
			w.Write(pr.SomethingWentWrong())
			return err
		}

		w.Write(resp)
		return nil
	}))
	fmt.Println("Start listening on port" + workPort)
	return http.ListenAndServe(workPort, mux)
}

const _formPath = `./web/dev/form.html`

func readForm() ([]byte, error) {
	return ioutil.ReadFile(_formPath)
}
