package converter

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"html-convert/internal/lib/browser"
)

func Convert(ctx context.Context, html string, options *browser.PdfOptions) ([]byte, error) {

	var pdf []byte
	newCtx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	err := chromedp.Run(newCtx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx)
		}),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			printToPdfOptions := page.PrintToPDF()

			if options != nil {
				browser.SetPdfOptions(printToPdfOptions, options)
			}

			buf, _, err := printToPdfOptions.Do(ctx)
			if err != nil {
				return err
			}
			pdf = buf
			return nil
		}),
	)

	if err != nil {
		return pdf, err
	}
	return pdf, nil
}
