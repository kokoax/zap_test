package main

import (
	zap "go.uber.org/zap"
	log "log"
)

func log_test(logger *log.Logger) {
	logger.Println("A walrus appears")
}

func main() {
	zap_logger, _ := zap.NewProduction()
	child := zap_logger.With(
		zap.String("OK", "2"),
	)
	std_logger := zap.NewStdLog(child)

	log_test(std_logger)
}
