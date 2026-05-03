APP ?=
APPS := getting-started api

.PHONY: build run all clean $(APPS)

.DEFAULT_GOAL := all

all: $(APPS)

$(APPS):
	@mkdir -p bin
	go build -o ./bin/$@ ./$@

build:
ifndef APP
	$(error APP is required. Usage: make build APP=<app-name>)
endif
ifeq ($(wildcard $(APP)/),)
	$(error Directory '$(APP)' not found. Available: $(APPS))
endif
	@mkdir -p bin
	go build -o ./bin/$(APP) ./$(APP)

run:
ifndef APP
	$(error APP is required. Usage: make run APP=<app-name>)
endif
ifeq ($(wildcard $(APP)/),)
	$(error Directory '$(APP)' not found. Available: $(APPS))
endif
	go run ./$(APP)

clean:
	rm -rf bin
