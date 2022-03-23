module example.com/dependency_fetcher

go 1.16

require github.com/onsi/gomega v1.17.0

require github.com/gorilla/mux v1.8.0

require gopkg.in/yaml.v2 v2.4.0

replace golang.org/x/text => golang.org/x/text v0.3.7

require (
	github.com/onsi/ginkgo/v2 v2.1.3
	github.com/tdewolff/minify/v2 v2.9.17
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
)

replace github.com/onsi/ginkgo => ../../../../../ginkgo_v2
