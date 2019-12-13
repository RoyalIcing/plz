package cmd

import (
	"strings"
)

var markdown = []string{
	"https://tools.ietf.org/html/rfc7763",
	"https://tools.ietf.org/html/rfc7764",
	"https://www.iana.org/assignments/markdown-variants/markdown-variants.xhtml",
}

var timestamps = []string{
	"https://tools.ietf.org/html/rfc3339",
}

var table = map[string][]string{
	"timestamps": timestamps, "timestamp": timestamps, "date": timestamps, "time": timestamps,
	"uri":      []string{"https://tools.ietf.org/html/rfc3986"},
	"http1.1":  []string{"http://tools.ietf.org/html/2616"},
	"http2":    []string{"https://tools.ietf.org/html/rfc7540"},
	"json":     []string{"http://tools.ietf.org/html/rfc8259"},
	"csv":      []string{"https://tools.ietf.org/html/rfc4180"},
	"zlib":     []string{"https://tools.ietf.org/html/rfc1950"},
	"deflate":  []string{"https://tools.ietf.org/html/rfc1951"},
	"gzip":     []string{"https://tools.ietf.org/html/rfc1952"},
	"pdf":      []string{"https://tools.ietf.org/html/rfc8118", "https://www.adobe.com/content/dam/acom/en/devnet/acrobat/pdfs/pdf_reference_1-7.pdf"},
	"tiff":     []string{"https://tools.ietf.org/html/rfc3302"},
	"dicom":    []string{"https://tools.ietf.org/html/rfc3240"},
	"ogg":      []string{"https://tools.ietf.org/html/rfc5334"},
	"jsonrpc":  []string{"https://www.jsonrpc.org/specification"},
	"markdown": markdown, "md": markdown,
	"aria": []string{
		"https://www.w3.org/TR/wai-aria/",
		"https://www.w3.org/TR/wai-aria-practices/",
		"https://www.w3.org/TR/html-aria/",
		"https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA",
	},
}

// Find term in our local table
func Find(term string) []string {
	term = strings.ToLower(term)
	results := table[term]
	return results
}
