## debug: run app on local machine with debug mode
.PHONY: run
debug:
	@echo "Running app on local machine with debug mode"
	@APP_ENV="local" DEBUG="true" go run main.go