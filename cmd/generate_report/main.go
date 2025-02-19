// cmd/generate_report/main.go

package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "time"
    "GoLang_FRT_API_Tests/pkg/reports"
)

func openLogFile() (*os.File, error) {
    return os.Open("reports/test.log")
}

func newScanner(file *os.File) *bufio.Scanner {
    return bufio.NewScanner(file)
}

func processLogFile(scanner *bufio.Scanner) {
    // Leer los logs y organizarlos por test
    currentTest := ""
    testResults := make(map[string]*reports.TestResult)

    for scanner.Scan() {
        line := scanner.Text()
        // Detectar inicio de nuevo test
        if strings.Contains(line, "🚀 Iniciando") {
            currentTest = strings.TrimSpace(strings.TrimPrefix(line, "🚀 Iniciando"))
            testResults[currentTest] = &reports.TestResult{
                Name:      currentTest,
                Status:    "RUNNING",
                Logs:      []string{},
                Timestamp: time.Now(),
                SubTests:  []*reports.TestResult{},
            }
        } else if strings.Contains(line, "✅ Test") && strings.Contains(line, "completado en") {
            // Detectar fin de test y su resultado
            if result, exists := testResults[currentTest]; exists {
                result.Status = "✅ PASS"

				result.Logs = append(result.Logs, line)
                
                // Extraer duración del test
                parts := strings.Split(line, " ")
                durationStr := parts[len(parts)-1]
                durationStr = strings.TrimSuffix(durationStr, "s")
                duration, err := time.ParseDuration(durationStr + "s")
                if err == nil {
                    result.Duration = duration
                }
            }
            // Resetear currentTest después de que el test termina
            currentTest = ""
        } else if strings.Contains(line, "❌ Error") {
            // Detectar error en el test
            if result, exists := testResults[currentTest]; exists {
                result.Status = "❌ FAIL"
                result.Logs = append(result.Logs, line)
            }
        }else if currentTest != "" && testResults[currentTest] != nil {
            // Agregar línea al log del test actual
            testResults[currentTest].Logs = append(testResults[currentTest].Logs, line)
        }
    }
	// Convertir el mapa a slice para el reporte
    var results []reports.TestResult
    for _, result := range testResults {
        results = append(results, *result)
    }

    // Generar el reporte HTML
    err := reports.GenerateHTMLReport(results, "reports/test-report.html")
    if err != nil {
        log.Fatalf("Error generando reporte HTML: %v", err)
    }

    log.Println("Reporte HTML generado exitosamente en reports/test-report.html")
}

func main() {
    logFile, err := openLogFile()
    if err != nil {
        log.Fatalf("Error abriendo archivo de logs: %v", err)
    }
    defer logFile.Close()
    scanner := newScanner(logFile)
    processLogFile(scanner)
}
