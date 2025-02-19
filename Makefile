# Variables
BINARY_NAME=api-tests
GO=go
GOTEST=$(GO) test
GOMOD=$(GO) mod
GOGET=$(GO) get
TEST_REPORT_DIR=reports

# Colores para la salida en consola
CYAN=\033[0;36m
RESET=\033[0m

.PHONY: all build test clean deps help

# Targets
all: clean deps test ## Ejecuta clean, deps y test

build: ## Compila el proyecto
	@echo "$(CYAN)Compilando...$(RESET)"
	$(GO) build -o $(BINARY_NAME) ./cmd/run_tests

test: ## Ejecuta todos los tests
	@echo "$(CYAN)Ejecutando tests...$(RESET)"
	$(GOTEST) -v ./tests/api/...

test-report: ##Ejecuta todos los tests y genera un reporte HTML
	@echo "$(CYAN)Limpiando el directorio de reportes$(RESET)"
	go clean
	rm -f bin/$(BINARY_NAME)
	rm -f $(TEST_REPORT_DIR)/*

	@mkdir -p $(TEST_REPORT_DIR)
	@echo "$(CYAN)Limpiando cache de tests$(RESET)"
	go clean -testcache
	
	@echo "$(CYAN)Ejecutando tests...$(RESET)"
	$(GOTEST) -v -count=1 ./tests/api/...
	@echo "$(CYAN)Generando reporte$(RESET)"
	go run cmd/generate_report/main.go


test-coverage: ## Ejecuta los tests con cobertura
	@echo "$(CYAN)Ejecutando tests con cobertura...$(RESET)"
	$(GOTEST) -coverprofile=coverage.out ./tests/api/...
	$(GO) tool cover -html=coverage.out -o coverage.html

clean: ## Limpia los archivos generados
	@echo "$(CYAN)Limpiando archivos de cobertura...$(RESET)"
	rm -f $(BINARY_NAME)
	rm -f coverage.out
	rm -f coverage.html
	rm -f $(TEST_REPORT_DIR)/*


deps: ## Descarga las dependencias
	@echo "$(CYAN)Downloading dependencies...$(RESET)"
	$(GOMOD) download
	$(GOMOD) tidy

update-deps: ## Actualiza las dependencias
	@echo "$(CYAN)Updating dependencies...$(RESET)"
	$(GOGET) -u ./...
	$(GOMOD) tidy

help: ## Muestra la ayuda
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(CYAN)%-15s$(RESET) %s\n", $$1, $$2}'