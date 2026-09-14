[![Explore in Constellation](https://img.shields.io/badge/Explore%20in-Constellation-blue)](https://constellation.speedata.de)

# csshtml

> **This package has moved.** csshtml is now part of
> [htmlbag](https://github.com/boxesandglue/htmlbag) as of htmlbag
> v0.0.54, and this repository is archived. Every identifier kept its
> name, so switching means changing the import path from
> `github.com/boxesandglue/csshtml` to `github.com/boxesandglue/htmlbag`
> and the qualifier from `csshtml.` to `htmlbag.`. v0.0.22 is the final
> release; older versions stay available through the Go module proxy.

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

## Ecosystem

csshtml is part of a broader ecosystem of PDF, typesetting and publishing technologies.

**[Explore the constellation →](https://constellation.speedata.de)**
