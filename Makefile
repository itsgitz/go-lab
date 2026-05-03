APP ?=

APPS := $(shell sed -n '/^use (/,/^)/p' go.work | sed '/^[[:space:]]*$/d' | sed 's/[[:space:]]*\.\///')

.PHONY: build run all clean $(APPS)

.DEFAULT_GOAL := all

all: $(APPS)

$(APPS):
	@mkdir -p bin
	go build -o ./bin/$@ ./$@

build:
ifdef APP
	@mkdir -p bin
	go build -o ./bin/$(APP) ./$(APP)
else
	@echo "Usage: make build APP=<app-name>"; exit 1
endif

run:
ifdef APP
	go run ./$(APP)
else
	@echo "Usage: make run APP=<app-name>"; exit 1
endif

clean:
	rm -rf bin
