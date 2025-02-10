# Html to Pdf Converter Golang

Сервис конвертации Html в Pdf. Для взаимодействия c Chrome DevTools используется библиотека chromedp.

**Request**

### `POST /convert`

**Parameters**

На вход принимается json следующего вида:

| Parameter                          | Type      | Description                                                                                                                                                                                                                                                                                                                                                                 |
|:-----------------------------------|:----------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `html`                             | `string`  | **Required**. Html to convert to pdf                                                                                                                                                                                                                                                                                                                                        |
| `pdfOptions`                       | `object`  | **Optional**.         
| `pdfOptions / scale`                            | `string`  | **Optional**. Scale of the webpage rendering. Defaults to 1. Scale amount must be between 0.1 and 2                                                                                                                                                                                                                                                                         | |
| `pdfOptions / displayHeaderFooter` | `boolean` | **Optional**. Display header and footer. Defaults to `false `                                                                                                                                                                                                                                                                                                               |
| `pdfOptions / headerTemplate`        | `string`  | **Optional**. HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: `date`, `title`, `url`, `pageNumber`, `totalPages`                                                                                                                                                                           |
| `pdfOptions / footerTemplate`        | `string`  | **Optional**. HTML template for the print footer. Should use the same format as the `headerTemplate`                                                                                                                                                                                                                                                                        |
| `pdfOptions / printBackground`       | `boolean` | **Optional**. Print background graphics. Defaults to `false`                                                                                                                                                                                                                                                                                                                |
| `pdfOptions / landscape`             | `boolean` | **Optional**. Paper orientation. Defaults to `false`                                                                                                                                                                                                                                                                                                                        |
| `pdfOptions / pageRanges`            | `string`  | **Optional**. Paper ranges to print, one based, e.g., '1-5, 8, 11-13'. Pages are printed in the document order, not in the order specified, and no more than once. Defaults to empty string, which implies the entire document is printed. The page numbers are quietly capped to actual page count of the document, and ranges beyond the end of the document are ignored. |
| `pdfOptions / paperWidth`            | `integer` | **Optional**. Paper width in inches. Defaults to 8.5 inches.                                                                                                                                                                                                                                                                                                                |
| `pdfOptions / paperHeight`           | `integer` | **Optional**. Paper height in inches. Defaults to 11 inches.                                                                                                                                                                                                                                                                                                                |
| `pdfOptions / marginTop`             | `integer` | **Optional**. Top margin in inches. Defaults to 1cm (~0.4 inches).                                                                                                                                                                                                                                                                                                          |
| `pdfOptions / marginRight`           | `integer` | **Optional**. Right margin in inches. Defaults to 1cm (~0.4 inches).                                                                                                                                                                                                                                                                                                        |
| `pdfOptions / marginBottom`          | `integer` | **Optional**. Bottom margin in inches. Defaults to 1cm (~0.4 inches).                                                                                                                                                                                                                                                                                                       |
| `pdfOptions / marginLeft`            | `integer` | **Optional**. Left margin in inches. Defaults to 1cm (~0.4 inches).                                                                                                                                                                                                                                                                                                         |
| `pdfOptions / preferCSSPageSize`     | `boolean` | **Optional**. Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.                                                                                                                                                                                                                      |

**Response**

При успешном запросе возвращается файл (Content-Type: application/pdf)

При неудачном запросе возвращается json вида:

```
{
    "status": "Error",
    "error": "failed to decode request"
}