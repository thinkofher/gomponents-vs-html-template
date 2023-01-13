# gomponents vs html/template

This is very simple and naive benchmark, which compares [html/template](//pkg.go.dev/html/template) and [github.com/maragudk/gomponents](//github.com/maragudk/gomponents), libraries for HTML rendering.

## Results

Tested on MacBook Pro M1 16GB.

    goos: darwin
    goarch: arm64
    pkg: github.com/thinkofher/gomponents-vs-html-template
    BenchmarkRenderHTML/html/template-8         	  144690	      8290 ns/op
    BenchmarkRenderHTML/github.com/maragudk/gomponents-8         	  157592	      7532 ns/op
    PASS
    ok  	github.com/thinkofher/gomponents-vs-html-template	3.770s
