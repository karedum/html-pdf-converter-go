package convert

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"html-convert/internal/converter"
	resp "html-convert/internal/lib/api/response"
	"html-convert/internal/lib/browser"
	"io"
	"log/slog"
	"net/http"
)

type Request struct {
	Html    string              `json:"html" validate:"required"`
	Options *browser.PdfOptions `json:"pdfOptions"`
}

func New(ctx context.Context, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")

			render.JSON(w, r, resp.Error("request body can't be empty"))

			return
		}

		if err != nil {
			fmt.Println()
			log.Error("failed to decode request body", "error", err.Error())

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", "error", err.Error())

			render.JSON(w, r, resp.ValidationErrors(validateErr))

			return
		}

		html := req.Html

		pdfOptions := req.Options

		pdf, err := converter.Convert(ctx, html, pdfOptions)

		if err != nil {
			log.Error("failed to convert html")

			render.JSON(w, r, resp.Error("failed to convert html"))

			return
		}

		responseFile(w, pdf)
	}
}

func responseFile(w http.ResponseWriter, pdf []byte) {
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdf)
}
