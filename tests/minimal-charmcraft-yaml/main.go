package main

import (
	"go-app/config"
	"log"
	"reflect"

	"github.com/caarlos0/env/v11"
	"github.com/kr/pretty"
)

func main() {
	var cfg config.CharmConfig
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("Error parsing configuration: %v", err)
	}

	metricsPort := 9001
	metricsPath := "/metrics"
	secretKey := "onerandomkey"
	httpProxy := "http://proxy.example.com:3128"
	expected := config.CharmConfig{
		Options: config.ConfigOptions{
			BaseURL:     "http://go-app.example.com",
			Port:        9000,
			MetricsPort: &metricsPort,
			MetricsPath: &metricsPath,
			SecretKey:   &secretKey,
		},
		Proxy: config.ProxyConfig{
			HTTPProxy:  &httpProxy,
			HTTPSProxy: nil,
			NoProxy:    []string{"127.0.0.1", "localhost", "::1"},
		},
	}

	pretty.Logf("Actual Config %# v\n", cfg)
	if !(reflect.DeepEqual(cfg, expected)) {
		pretty.Logf("Expected Config %# v\n", expected)
		pretty.Pdiff(log.Default(), cfg, expected)
		log.Fatalf("Wrong configuration.")
	}

}
