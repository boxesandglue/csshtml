# csshtml

A Go package that parses CSS stylesheets and applies them to HTML documents, producing a DOM tree with computed style attributes on each node.

## Features

- Parse CSS from files or strings
- Apply CSS rules to HTML documents using selector matching
- Support for `@import`, `@font-face`, and `@page` at-rules
- Handles linked stylesheets (`<link href="...">`) in HTML documents
- Returns a [goquery](https://github.com/PuerkitoBio/goquery) Document with style attributes applied

## Usage

```go
package main

import "github.com/speedata/csshtml"

func main() {
    // Create a new CSS parser
    css := csshtml.NewCSSParser()

    // Add CSS rules
    css.AddCSSText(`
        body { font-family: serif; }
        h1 { color: blue; font-size: 24pt; }
    `)

    // Process HTML and apply CSS
    doc, err := css.ProcessHTMLChunk(`
        <html>
        <body>
            <h1>Hello World</h1>
        </body>
        </html>
    `)
    if err != nil {
        panic(err)
    }

    // The returned document has style attributes on matching elements
    // doc.Find("h1") will have style="color: blue; font-size: 24pt;"
}
```

## API

- `NewCSSParser()` - Create a new CSS parser
- `NewCSSParserWithDefaults()` - Create a parser with default browser styles
- `AddCSSText(css string)` - Parse and add CSS rules
- `ProcessHTMLFile(filename string)` - Load HTML file, read linked stylesheets, apply CSS
- `ProcessHTMLChunk(html string)` - Parse HTML string and apply CSS
- `ApplyCSS(doc *goquery.Document)` - Apply collected CSS rules to a document
