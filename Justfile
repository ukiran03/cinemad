# List available commands
default:
    @just --list

# Generate Go code from templ files once
generate:
    templ generate

# Live reload for development with FLAGs
dev +args="":
    wgo -file=.go -file=.templ -xfile=_templ.go \
    templ generate ./cmd/web/ui/... :: go run ./cmd/web {{args}}

# Run tests with file watching
watch-test:
    wgo go test -v ./...

# Build the production binary
build: generate
    go build -o bin/cinemad ./cmd/web

# Clean up build artifacts
clean:
    rm -rf bin/
    fd -I "_templ.go$" -x rm

# Tidy module dependencies
tidy:
    @echo '> Tidying module dependencies...'
    go mod tidy
    go mod verify
    @echo '> Formatting .go files...'
    go fmt ./...

# run quality control checks
audit:
    @echo '> Checking module dependencies...'
    go mod tidy -diff
    go mod verify
    @echo '> Vetting code...'
    go vet ./...
    go tool staticcheck ./...
    @echo '> Running tests...'
    go test -race -vet=off ./...

# Comprehensive golangci-lint
ci-lint:
    golangci-lint run ./...

# Automatic refactors and formatting
ci-fix:
    golangci-lint run --fix ./...
