# TAILWINDCSS_OS_ARCH := macos-arm64
TAILWINDCSS_OS_ARCH := linux-x64

.PHONY: benchmark
benchmark:
	go test -bench=.

.PHONY: air
air:
	air

.PHONY: dev
dev:
	make -j2 watch-css air

.PHONY: build-css
build-css: tailwindcss
	# ./tailwindcss -i tailwind.css -o assets/src/styles/app.css --minify
	./tailwindcss -i tailwind.css -o assets/src/styles/app.css

.PHONY: build
build: build-css
	go generate ./...
	go build -o gom cmd/app/main.go

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: lint
lint:
	golangci-lint run

.PHONY: start
start: build-css
	go run ./cmd/app

.PHONY: run
run:
	go run ./cmd/run

tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(TAILWINDCSS_OS_ARCH)
	mv tailwindcss-$(TAILWINDCSS_OS_ARCH) tailwindcss
	chmod a+x tailwindcss
	mkdir -p node_modules/tailwindcss/lib && ln -sf tailwindcss node_modules/tailwindcss/lib/cli.js
	echo '{"devDependencies": {"tailwindcss": "latest"}}' >package.json

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...

.PHONY: watch-css
watch-css: tailwindcss
	# ./tailwindcss -i tailwind.css -o assets/src/styles/app.css --watch --minify
	./tailwindcss -i tailwind.css -o assets/src/styles/app.css --watch
