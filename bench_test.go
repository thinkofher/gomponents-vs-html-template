package bench

import (
	"io"
	"testing"
)

func testProps() props {
	return props{
		Title:    "Amazing Dogs",
		LoggedIn: true,
		Photos: []photo{
			{
				Filename: "lala.png",
				Alt:      "Lovely Gordon Setter female with beautiful eyes.",
			},
			{
				Filename: "lenny.png",
				Alt:      "Young and cuddly Irish Setter male.",
			},
			{
				Filename: "ozzy.png",
				Alt:      "Shiny and dignified Irish Setter male. He loves kisses!",
			},
			{
				Filename: "hauru.png",
				Alt:      "Unkempt, sturdy and big Irish Setter male.",
			},
		},
	}
}

func BenchmarkRenderHTML(b *testing.B) {
	p := testProps()

	b.Run("html/template", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			executeTemplate(io.Discard, p)
		}
	})

	b.Run("github.com/maragudk/gomponents", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			componentDocument(p).Render(io.Discard)
		}
	})
}
