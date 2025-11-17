GO ?= CGO_ENABLED=0 go
CPU_NAME := $(shell $(GO) run ./cmd/cpuname)
BENCH_FILE := benches/$(shell $(GO) env GOOS)-$(shell $(GO) env GOARCH)-$(CPU_NAME).txt

.PHONY: all
all: tidy test

.PHONY: clean
clean:
	$(GO) clean
	-rm \
		coverage.txt \
		coverage.xml \
		gl-code-quality-report.json \
		govulncheck.sarif \
		junit.xml \
		README.html \
		golangci-lint.json \
		test.log

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: tidy
tidy:
	test -z "$$(gofmt -l .)"
	$(GO) vet
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest --version
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

cpu.profile:
	$(GO) test -run=^$ -bench=Day10Part1$ -benchmem -memprofile mem.profile -cpuprofile $@

.PHONY: prof
prof: cpu.profile
	$(GO) tool pprof $^

.PHONY: test
test:
	$(GO) test -run=Day -short -vet=all

.PHONY: sast
sast: coverage.xml gl-code-quality-report.json govulncheck.sarif junit.xml

coverage.txt test.log &:
	-$(GO) test -coverprofile=coverage.txt -covermode count -short -v | tee test.log

# Gitlab test report
junit.xml: test.log
	$(GO) run github.com/jstemmer/go-junit-report/v2@latest < $< > $@

# Gitlab coverage report
coverage.xml: coverage.txt
	$(GO) run github.com/boumenot/gocover-cobertura@latest < $< > $@

# Gitlab code quality report
gl-code-quality-report.json: golangci-lint.json
	$(GO) run github.com/banyansecurity/golint-convert@latest < $< > $@

golangci-lint.json:
	-$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --out-format json > $@

# Gitlab dependency report
govulncheck.sarif:
	$(GO) run golang.org/x/vuln/cmd/govulncheck@latest -version
	$(GO) run golang.org/x/vuln/cmd/govulncheck@latest -format=sarif ./... > $@

$(BENCH_FILE): $(wildcard *.go)
	@echo "Running benchmarks and saving to $@..."
	@mkdir -p benches
	$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@

README.html: README.adoc
	asciidoc $^

.PHONY: total
total: $(BENCH_FILE)
	@awk -f total.awk < $(BENCH_FILE)
