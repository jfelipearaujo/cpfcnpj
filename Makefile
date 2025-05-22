TAG := $(shell git describe --tags --abbrev=0 2>/dev/null)
VERSION := $(shell echo $(TAG) | sed 's/v//')

tag:
	@if [ -z "$(TAG)" ]; then \
        echo "No previous version found. Creating v1.0.0 tag..."; \
        git tag v1.0.0; \
    else \
        echo "Previous version found: $(VERSION)"; \
        read -p "Bump major version (M/m), minor version (R/r), or patch version (P/p)? " choice; \
        if [ "$$choice" = "M" ] || [ "$$choice" = "m" ]; then \
            echo "Bumping major version..."; \
			major=$$(echo $(VERSION) | cut -d'.' -f1); \
            major=$$(expr $$major + 1); \
            new_version=$$major.0.0; \
		elif [ "$$choice" = "R" ] || [ "$$choice" = "r" ]; then \
            echo "Bumping minor version..."; \
			minor=$$(echo $(VERSION) | cut -d'.' -f2); \
            minor=$$(expr $$minor + 1); \
            new_version=$$(echo $(VERSION) | cut -d'.' -f1).$$minor.0; \
		elif [ "$$choice" = "P" ] || [ "$$choice" = "p" ]; then \
            echo "Bumping patch version..."; \
			patch=$$(echo $(VERSION) | cut -d'.' -f3); \
            patch=$$(expr $$patch + 1); \
            new_version=$$(echo $(VERSION) | cut -d'.' -f1).$$(echo $(VERSION) | cut -d'.' -f2).$$patch; \
        else \
            echo "Invalid choice. Aborting."; \
            exit 1; \
        fi; \
        echo "Creating tag for version v$$new_version..."; \
        git tag v$$new_version; \
    fi

test:
	go test -count=1 -race -cover -coverprofile=coverage.txt ./...

lint:
	golangci-lint run
