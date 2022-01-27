package api

import (
	"log"
	"net/http"

	"github.com/go-chi/render"

	"github.com/rjNemo/go-pass-gen/passgen"
)

func (s Server) HandleNewPassword(w http.ResponseWriter, r *http.Request) {
	params := &PasswordParams{}
	if err := render.Bind(r, params); err != nil {
		log.Fatal(err)
	}
	opts := passgen.Options{Length: params.Length, WithNumbers: params.WithNumbers}
	password := passgen.New(opts.SetDefaults()).NewPassword()

	render.Status(r, http.StatusAccepted)
	err := render.Render(w, r, &PasswordResponse{Password: password})
	if err != nil {
		log.Fatalf("error: %q", err)
	}
}

type PasswordParams struct {
	Length      int  `json:"length"`
	WithNumbers bool `json:"with_numbers"`
}

func (p PasswordParams) Bind(*http.Request) error {
	return nil
}

type PasswordResponse struct {
	Success  bool    `json:"success"`
	Error    *string `json:"error"`
	Password string  `json:"password"`
}

func (resp *PasswordResponse) Render(http.ResponseWriter, *http.Request) error {
	resp.Success = true
	return nil
}
