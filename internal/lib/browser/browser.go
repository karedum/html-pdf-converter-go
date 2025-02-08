package browser

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"reflect"
)

type PdfOptions struct {
	Landscape           *bool    `json:"landscape" validate:"omitempty,boolean"`
	DisplayHeaderFooter *bool    `json:"displayHeaderFooter" validate:"omitempty,boolean"`
	PrintBackground     *bool    `json:"printBackground" validate:"omitempty,boolean"`
	Scale               *float64 `json:"scale" validate:"omitempty,numeric"`
	PaperWidth          *float64 `json:"paperWidth" validate:"omitempty,numeric"`
	PaperHeight         *float64 `json:"paperHeight" validate:"omitempty,numeric"`
	MarginTop           *float64 `json:"marginTop" validate:"omitempty,numeric"`
	MarginBottom        *float64 `json:"marginBottom" validate:"omitempty,numeric"`
	MarginLeft          *float64 `json:"marginLeft" validate:"omitempty,numeric"`
	MarginRight         *float64 `json:"marginRight" validate:"omitempty,numeric"`
	PageRanges          *string  `json:"pageRanges" validate:"omitempty,min=1"`
	HeaderTemplate      *string  `json:"headerTemplate" validate:"omitempty,min=1"`
	FooterTemplate      *string  `json:"footerTemplate" validate:"omitempty,min=1"`
	PreferCSSPageSize   *bool    `json:"preferCSSPageSize" validate:"omitempty,boolean"`
}

func MustRun(chromeAddress string) (context.Context, context.CancelFunc) {
	dockerURL := fmt.Sprintf("wss://%s", chromeAddress)
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), dockerURL)

	ctx, cancel := chromedp.NewContext(allocatorContext)

	err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
	)

	if err != nil {
		log.Fatalf("cannot run chrome: %s", err)
	}

	return ctx, cancel
}

func SetPdfOptions(params *page.PrintToPDFParams, options *PdfOptions) {

	v := reflect.ValueOf(options).Elem()
	typeOfT := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() != true {
			value := v.Field(i).Elem().Interface()
			switch typeOfT.Field(i).Name {
			case "Landscape":
				params.Landscape = value.(bool)
			case "DisplayHeaderFooter":
				params.DisplayHeaderFooter = value.(bool)
			case "PrintBackground":
				params.PrintBackground = value.(bool)
			case "Scale":
				params.Scale = value.(float64)
			case "PaperWidth":
				params.PaperWidth = value.(float64)
			case "PaperHeight":
				params.PaperHeight = value.(float64)
			case "MarginTop":
				params.MarginTop = value.(float64)
			case "MarginBottom":
				params.MarginBottom = value.(float64)
			case "MarginLeft":
				params.MarginLeft = value.(float64)
			case "MarginRight":
				params.MarginRight = value.(float64)
			case "PageRanges":
				params.PageRanges = value.(string)
			case "HeaderTemplate":
				params.HeaderTemplate = value.(string)
			case "FooterTemplate":
				params.FooterTemplate = value.(string)
			case "PreferCSSPageSize":
				params.PreferCSSPageSize = value.(bool)
			}
		}
	}
}
