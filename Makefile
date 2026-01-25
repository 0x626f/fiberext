.PHONY: install
MAKEFLAGS += --no-print-directory


install: setup
	$(eval DEPS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS)))
	@if [ -z "$(DEPS)" ]; then \
		go mod download; \
  	else \
  		go get $(DEPS); \
    fi

%:
	@:
