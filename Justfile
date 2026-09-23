# Automatically load environment variables from .env / direnv
set dotenv-load

# List available commands
default:
    @just --list

# Generate Go code from templ files once
templ-gen:
    templ generate

# Live reload for development with FLAGs
dev +args="":
    wgo -file=.go -file=.templ -xfile=_templ.go templ generate ./cmd/web/ui/... :: go run ./cmd/web {{args}}

# Start and Watch Api server
run-api:
    wgo run ./cmd/api

# Run tests with file watching
watch-test:
    wgo go test -v ./...

# Builds all binaries
build: clean build_front build_back
    @printf "All binaries built!\n"

# Builds the front end
build_front:
    @echo "Building front end..."
    @templ generate
    @go build -o bin/cinemad ./cmd/web
    @echo "Front end built!"

# Builds the back end
build_back:
    @echo "Building back end..."
    @go build -o bin/cinemad_api ./cmd/api
    @echo "Back end built!"

# Starts front and back end
start: start_front start_back

# Starts the front end
start_front: build_front
    @echo "Starting the front end..."
    @STRIPE_KEY="$STRIPE_SECRET_KEY" ./bin/cinemad -port=$FRONTEND_PORT &
    @echo "Front end running!"

# Starts the back end
start_back: build_back
    @echo "Starting the back end..."
    @STRIPE_KEY="$STRIPE_SECRET_KEY" ./bin/cinemad_api -port=$BACKEND_PORT &
    @echo "Back end running!"

# Stops the front end
stop_front:
    @echo "Stopping the front end..."
    -pkill -SIGTERM -f "cinemad -port=$FRONTEND_PORT"
    @echo "Stopped front end"

# Stops the back end
stop_back:
    @echo "Stopping the back end..."
    -pkill -SIGTERM -f "cinemad_api -port=$BACKEND_PORT"
    @echo "Stopped back end"

# Stops the front and back end
stop: stop_front stop_back
    @echo "All applications stopped"

# Clean up build artifacts
clean:
    @echo "Cleaning..."
    rm -rf bin/
    @go clean
    fd -I "_templ.go$" -x rm
    @echo "Cleaned!"

# Tidy module dependencies
tidy:
    @echo '> Tidying module dependencies...'
    go mod tidy
    go mod verify
    @echo '> Formatting .go files...'
    go fmt ./...

# Run quality control checks
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
