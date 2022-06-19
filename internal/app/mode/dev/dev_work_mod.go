package dev

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"link-contractor-api/internal/entities/user"
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
		UserExternalID string `json:"user_external_id"`
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
		if err != nil {
			w.Write(pr.SomethingWentWrong())
			return err
		}

		var up UserPl
		if err := json.Unmarshal(pl, &up); err != nil {
			w.Write(pr.SomethingWentWrong())
			return fmt.Errorf("unmarshal payload: %w", err)
		}

		entityUser := user.User{
			ExternalID: up.UserExternalID,
			SourceType: user.DevSource,
			Name:       "Test",
			Surname:    "Test",
		}

		resp, err := ep.DoAction(r.Context(), entityUser, pl)
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
