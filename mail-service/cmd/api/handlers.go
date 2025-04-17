package main

import "net/http"

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From string `json:"from"`
		To string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var payload mailMessage

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From: payload.From,
		To: payload.To,
		Subject: payload.Subject,
		Data: payload.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp = jsonResponse{
		Error: false,
		Message: "E-mail sent to " + payload.To,
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
