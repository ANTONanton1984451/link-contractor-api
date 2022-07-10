package vk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"link-contractor-api/internal/entities/user"
	"link-contractor-api/internal/entrypoint"
	"net/http"
	url "net/url"
	"strconv"

	"go.uber.org/zap"
)

type (
	Presenter interface {
		SomethingWentWrong() []byte
	}
)

func StartWorking(ep entrypoint.Entrypoint, cfg Config, log *zap.Logger, pr Presenter) error {
	mux := http.NewServeMux()
	logger := log.Sugar()

	mux.HandleFunc("/vk", middleWare(cfg, logger, pr, func(ctx context.Context, event Event) (handleResult, error) {
		switch event.Type {
		case _confirmationType:
			return handleResult{responseMessage: cfg.ConfirmToken}, nil
		case _messageNew:
			userID := strconv.FormatUint(event.Object.Message.FromID, 10)

			us := user.User{
				ExternalID: userID,
				SourceType: user.VkSource,
			}
			// в случае отправления - peer_id == айдишник юзера
			resp, err := ep.DoAction(ctx, us, []byte(event.Object.Message.Text))
			if err != nil {
				return handleResult{}, fmt.Errorf("do entrypoint action: %w", err)
			}

			return handleResult{sendMessage: string(resp)}, nil
		}
		return handleResult{}, nil
	}))

	fmt.Println(fmt.Sprintf("start work on %s port", cfg.WorkPort))
	return http.ListenAndServe(cfg.WorkPort, mux)
}

func middleWare(cfg Config, logger *zap.SugaredLogger, pr Presenter, handleFunc func(ctx context.Context, event Event) (handleResult, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		pl, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("ok"))
			logger.Errorf("read body: %s", err)
		}

		logger.Infof("request with payload %s", string(pl))

		var event Event
		if err = json.Unmarshal(pl, &event); err != nil {
			w.Write([]byte("ok"))
			logger.Errorf("unmarshal event(%s): %w", pl, err)
		}

		result, err := handleFunc(r.Context(), event)

		if err != nil && event.Type == _confirmationType {
			logger.Errorf("handle event: %s", err)
			w.Write([]byte("ok"))
			return
		}
		if err != nil && event.Type == _messageNew {
			sendMessage(logger, cfg.AccessToken, cfg.ApiVersion, cfg.ApiURL, strconv.FormatUint(event.Object.Message.FromID, 10), pr.SomethingWentWrong())
			w.Write([]byte("ok"))
			return
		}

		if result.responseMessage != "" {
			w.Write([]byte(result.responseMessage))
			return
		}

		if result.sendMessage != "" {
			sendMessage(logger, cfg.AccessToken, cfg.ApiVersion, cfg.ApiURL, strconv.FormatUint(event.Object.Message.FromID, 10), []byte(result.sendMessage))
		}
		w.Write([]byte(`ok`))

	}
}

type handleResult struct {
	sendMessage, responseMessage string
}

func sendMessage(logger *zap.SugaredLogger, token, apiVersion, vkURL, userID string, message []byte) {
	query := url.Values{}
	query.Set("message", string(message))
	query.Set("peer_id", userID)
	query.Set("user_id", userID)
	query.Set("access_token", token)
	query.Set("v", apiVersion)
	// todo разобраться с этим
	query.Set("random_id", "0")

	requestUrl := &url.URL{}
	requestUrl.Host = vkURL
	requestUrl.Scheme = "https"
	requestUrl.RawQuery = query.Encode()
	requestUrl.Path = "method/messages.send"

	logger.Infof("make request, url %s, message %s, user_id %d", requestUrl.String(), string(message), userID)
	_, err := http.Get(requestUrl.String())
	if err != nil {
		logger.Errorf("execute request: %s", err)
	}
}
