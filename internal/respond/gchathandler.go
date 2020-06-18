package respond

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MessageRequest struct {
	Type      string    `json:"type"`
	EventTime time.Time `json:"eventTime"`
	Space     struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Type        string `json:"type"`
	} `json:"space"`
	Message struct {
		Name   string `json:"name"`
		Sender struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			AvatarURL   string `json:"avatarUrl"`
			Email       string `json:"email"`
		} `json:"sender"`
		CreateTime   time.Time `json:"createTime"`
		Text         string    `json:"text"`
		ArgumentText string    `json:"argumentText"`
		Thread       struct {
			Name string `json:"name"`
		} `json:"thread"`
		Annotations []struct {
			Length      int `json:"length"`
			StartIndex  int `json:"startIndex"`
			UserMention struct {
				Type string `json:"type"`
				User struct {
					AvatarURL   string `json:"avatarUrl"`
					DisplayName string `json:"displayName"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"user"`
			} `json:"userMention"`
			Type string `json:"type"`
		} `json:"annotations"`
	} `json:"message"`
	User struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		AvatarURL   string `json:"avatarUrl"`
		Email       string `json:"email"`
	} `json:"user"`
}

type MessageResponse struct {
	Text string `json:"text"`
}
type HttpGchatRespondHandler struct {
	Svc *Service
}

func (h HttpGchatRespondHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var m MessageRequest
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		//todo: log
		fmt.Println("error here:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := h.Svc.Respond(m.Message.Text)
	jres := MessageResponse{Text: res}
	err = json.NewEncoder(w).Encode(jres)
	if err != nil {
		//todo: log
		fmt.Println("error here 2:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
