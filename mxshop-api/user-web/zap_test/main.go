package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	// logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	url := "https://imooc.com"
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", 100),
	)
	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch URL",
	// 	// Structured context as loosely typed key-value pairs.
	// 	"url", url,
	// 	"attempt", 3,
	// )
	// sugar.Infof("failed to fetch URL: %s", url)
}
