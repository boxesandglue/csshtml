package csshtml

import (
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
