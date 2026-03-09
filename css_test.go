package csshtml

import (
	"strings"
	"testing"
)

func TestNestedAtrule(t *testing.T) {

	str := `
	@page {
		size: a5;
		@bottom-right-corner {
			border: 4pt solid green;
			border-bottom-color: rebeccapurple;
		}

		/* @top-left-corner {
			border: 1pt solid green;
			border-bottom-color: rebeccapurple;
		} */

	@top-right-corner {
			border: 3pt solid green;
			border-bottom-color: rebeccapurple;
		}

		@bottom-left-corner {
			border: 2pt solid green;
			border-bottom-color: rebeccapurple;
		}

	}`
	toks := tokenizeCSSString(str)
	bl := consumeBlock(toks, false)
	if len(bl.childAtRules[0].childAtRules) != 3 {
		t.Errorf("want 3 child @ rules, got %d", len(bl.childAtRules[0].childAtRules))
	}
}

func TestFontFace(t *testing.T) {
	str := `@font-face {
		font-family: "Trickster";
		src:
		  local("Trickster"),
		  url("trickster-COLRv1.otf") format("opentype") tech(color-COLRv1),
		  url("trickster-outline.otf") format("opentype"),
		  url("trickster-outline.woff") format("woff");
	  }`
	cp := NewCSSParser()
	err := cp.AddCSSText(str)
	if err != nil {
		t.Error(err)
	}
	fontfaces := cp.FontFaces
	if got, want := len(fontfaces), 1; got != want {
		t.Errorf("len(c.FontFaces) = %d, want %d", got, want)
	}
	firstFontFace := fontfaces[0]
	if want, got := 4, len(firstFontFace.Source); got != want {
		t.Errorf("len(c.FontFaces[0].Source) = %d, want %d", got, want)
	}
	if want, got := "color-COLRv1", firstFontFace.Source[1].Tech; got != want {
		t.Errorf(`firstFontFace.Source[1].Tech = %s, want %s`, got, want)
	}
}

func TestConsumeBlock_SimpleRules(t *testing.T) {
	css := `p { color: red; font-size: 12pt; }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	if len(bl.blocks[0].rules) != 2 {
		t.Errorf("got %d rules, want 2", len(bl.blocks[0].rules))
	}
}

func TestConsumeBlock_MultipleSelectors(t *testing.T) {
	css := `h1 { color: blue; } p { color: red; }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 2 {
		t.Fatalf("got %d blocks, want 2", len(bl.blocks))
	}
}

func TestConsumeBlock_ClassSelector(t *testing.T) {
	css := `.highlight { background: yellow; }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	sel := selectorString(bl.blocks[0].componentValues)
	if !strings.Contains(sel, ".highlight") {
		t.Errorf("selector = %q, want to contain '.highlight'", sel)
	}
}

func TestConsumeBlock_IDSelector(t *testing.T) {
	css := `#main { width: 100%; }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	sel := selectorString(bl.blocks[0].componentValues)
	if !strings.Contains(sel, "#main") {
		t.Errorf("selector = %q, want to contain '#main'", sel)
	}
}

func TestConsumeBlock_DescendantSelector(t *testing.T) {
	css := `div p { margin: 0; }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	sel := strings.Join(strings.Fields(selectorString(bl.blocks[0].componentValues)), " ")
	if sel != "div p" {
		t.Errorf("selector = %q, want 'div p'", sel)
	}
}

func TestConsumeBlock_ChildCombinator(t *testing.T) {
	css := `ul > li { list-style: none; }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	sel := strings.Join(strings.Fields(selectorString(bl.blocks[0].componentValues)), " ")
	if sel != "ul > li" {
		t.Errorf("selector = %q, want 'ul > li'", sel)
	}
}

func TestConsumeBlock_RuleWithoutTrailingSemicolon(t *testing.T) {
	css := `p { color: red }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	if len(bl.blocks[0].rules) != 1 {
		t.Errorf("got %d rules, want 1", len(bl.blocks[0].rules))
	}
}

func TestConsumeBlock_EmptyBlock(t *testing.T) {
	css := `p { }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	if len(bl.blocks[0].rules) != 0 {
		t.Errorf("got %d rules, want 0", len(bl.blocks[0].rules))
	}
}

func TestSelectorString_IDNotDoubleHash(t *testing.T) {
	css := `#important { color: red; }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	sel := selectorString(bl.blocks[0].componentValues)
	if sel != "#important" {
		t.Errorf("selector = %q, want %q", sel, "#important")
	}
}

func TestApplyCSS_IDSelector(t *testing.T) {
	htmlStr := `<html><head></head><body><p id="important">text</p></body></html>`
	css := `#important { color: green; }`
	c := NewCSSParser()
	if err := c.AddCSSText(css); err != nil {
		t.Fatal(err)
	}
	doc, err := c.ProcessHTMLChunk(htmlStr)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.ApplyCSS(doc)
	if err != nil {
		t.Fatalf("ApplyCSS failed: %v", err)
	}
	p := doc.Find("#important")
	if val, exists := p.Attr("!color"); !exists || val != "green" {
		t.Errorf("#important !color = %q (exists=%v), want 'green'", val, exists)
	}
}

func TestConsumeBlock_PseudoClass(t *testing.T) {
	css := `a:hover { color: green; }`
	toks := tokenizeCSSString(css)
	bl := consumeBlock(toks, false)
	if len(bl.blocks) != 1 {
		t.Fatalf("got %d blocks, want 1", len(bl.blocks))
	}
	sel := selectorString(bl.blocks[0].componentValues)
	if !strings.Contains(sel, ":hover") {
		t.Errorf("selector = %q, want to contain ':hover'", sel)
	}
}
