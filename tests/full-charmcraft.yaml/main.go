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
	userConfigIntNoDefault := 3
	postgreSQLUsername := "test-username"
	postgreSQLPassword := "test-password"
	postgreSQLHostname := "test-postgresql"
	postgreSQLPort := 5432
	postgreSQLName := "test-database"
	redisHostname := "redisuri"
	s3Region := "region"
	s3StorageClass := "GLACIER"
	s3Path := "/path/subpath/"
	s3APIVersion := "s3v4"
	samlMetadataURL := "https://login.staging.ubuntu.com/saml/metadata"

	expected := config.CharmConfig{
		Options: config.ConfigOptions{
			BaseURL:     "http://go-app.example.com",
			Port:        9000,
			MetricsPort: &metricsPort,
			MetricsPath: &metricsPath,
			SecretKey:   &secretKey,
			UserConfigOptions: config.UserConfigOptions{
				UserConfigBoolean:         true,
				UserConfigFloat:           1.5,
				UserConfigInt:             2,
				UserConfigIntNoDefault:    &userConfigIntNoDefault,
				UserConfigString:          "newstring",
				UserConfigStringNoDefault: nil,
			},
		},
		Proxy: config.ProxyConfig{
			HTTPProxy:  &httpProxy,
			HTTPSProxy: nil,
			NoProxy:    []string{"127.0.0.1", "localhost", "::1"},
		},
		Integrations: config.Integrations{
			MongoDB: config.MongoDBIntegration{},
			MySQL:   config.MySQLIntegration{},
			PostgreSQL: config.PostgreSQLIntegration{
				DatabaseIntegration: config.DatabaseIntegration{
					ConnectString: "postgresql://test-username:test-password@test-postgresql:5432/test-database?connect_timeout=10",
					Scheme:        "postgresql",
					NetLoc:        "test-username:test-password@test-postgresql:5432",
					Path:          "/test-database",
					Params:        "",
					Query:         "connect_timeout=10",
					Fragment:      "",
					Username:      &postgreSQLUsername,
					Password:      &postgreSQLPassword,
					Hostname:      &postgreSQLHostname,
					Port:          &postgreSQLPort,
					Name:          &postgreSQLName,
				},
			},
			Redis: config.RedisIntegration{
				DatabaseIntegration: config.DatabaseIntegration{
					ConnectString: "http://redisuri",
					Scheme:        "http",
					NetLoc:        "redisuri",
					Path:          "",
					Params:        "",
					Query:         "",
					Fragment:      "",
					Username:      nil,
					Password:      nil,
					Hostname:      &redisHostname,
					Port:          nil,
					Name:          nil,
				},
			},
			S3: config.S3Integration{
				AccessKey:       "access_key",
				SecretKey:       "secret_key",
				Region:          &s3Region,
				StorageClass:    &s3StorageClass,
				Bucket:          "bucket",
				Endpoint:        nil,
				Path:            &s3Path,
				ApiVersion:      &s3APIVersion,
				UriStyle:        nil,
				AddressingStyle: nil,
				Attributes:      nil,
				TLSCAChain:      nil,
			},
			SAML: config.SAMLIntegration{
				EntityID:                "https://login.staging.ubuntu.com",
				MetadataURL:             &samlMetadataURL,
				SingleSignOnRedirectURL: "https://login.staging.ubuntu.com/saml/",
				SigningCertificate:      "MIIFOjCCAyKgAwIBAgIUCWcJxrX+RK7+1rxXeiOk2Pef3NMwDQYJKoZIhvcNAQELBQAwHjEcMBoGA1UEAwwTc2FtbC5jYW5vbmljYWwudGVzdDAeFw0yMzA4MjYxNDQ5NDZaFw0zMzA4MjMxNDQ5NDZaMB4xHDAaBgNVBAMME3NhbWwuY2Fub25pY2FsLnRlc3QwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCvRliUlWDkDSCrRz/H5oTpRfIhMKJ/EwE6aWVue+GxdXivXRSBQFnPquE2XDcPPwEqsmjJ58pZ1D3hhPHQtyY9lhTcBP7TbRo2lx4aT0smCUellwpw4AqqARcPQvBPJ5s+sz4mRk7t8zV0eKjLTvvq8/qWKi1hHW9wJb9hD8Ie3Y3BbI5cc3+K9Nhqq2qNg7Pkmo7dvCDP3F8v+0gDRM6XuNhVq97EdjzUcXbHjj+CbjCeQKxBnhgfdHZEwbvqRv8YuWvJ2PmNUqFgjQRi2x0kXJyJW1OPaW83DiSvkqOaFb1rSt8/tK3ZdQ9YcoXf8bdJyRgrfS6YIJf+1o+ckty7rUEk+LhXH2OGg0rA6MIbTiFrMw4PVTRHJac0q8w4puQzTTKcCTBkj7Wi2sEKcOeauDYueLwGOF7pY4NL1em7xz7zYTsFYcDG5qnjdQAJOLCEJRKWoQezaLW4RTTbCJ0sa7U/lNtym45qwhkM5zR4kyLHV0U0XKTRrD94HslVLPZrxX9MXEos+ys6wWcxovPf5IMg4e3iMqxuwJmIOEMXTM/0GNjlNeovvUvbKz+nA1lIYOnf8BmnNoBUvJoIe02igU7+NBRvlqByYByJBz90JXa0Knh1kK+xCjo4deRtbe1JzWEEJZ+fB58SEn2C/ulJbngnb+uh4x1UgvbkWHH9wQIDAQABo3AwbjAdBgNVHQ4EFgQUVMab5jNo59oiOT2G91sxDCiZPuIwHwYDVR0jBBgwFoAUVMab5jNo59oiOT2G91sxDCiZPuIwDAYDVR0TAQH/BAIwADAeBgNVHREEFzAVghNzYW1sLmNhbm9uaWNhbC50ZXN0MA0GCSqGSIb3DQEBCwUAA4ICAQB5tLn9y3E/MciDkFSh1/PYAyNjs+zGxpm/KH95pKyVWeqFSt09+FnOVI/RUVsx7+FV3++gKfR5flsDbWhmtNAWvujyls7DFuWpf0pDfkrzaul3K1Fu2wF+xVklB+YYVDlP2LYrAdONiVq0Dd5a69tniwTDLjaKhhNJqOTrg/OywmDBa/Ym7kyih5oWP7fdoK9MBeFRrCJsD/831ax+LtNMhp+wgM7cmKJOLqNBbqRDhyPEk9XSK5gJdP1ty/2IEfktsGiPrjR6Wl2WvUj6BlCBFHh7aM4IYjuRo/XU+3yyEb2hvlEYT3rDaNa4ZBtGub6K9/M78CBPULaKlGWAVmdtXH8+04YRxWpyfP0XPPUzkzqAvLs6WEBeD1zEN/+1AtyGEAc0/e7LKwblOSXQC/lkjhiwpIF1UV24wHHDsOoMMvcAF+pJLZkPDKqYAktohhQjtUOOyQ1QxodY6F6ENfimlS0cDP7ngmEswe50VJ52otZvmq1n5aiyeF1hCcIQ8IqeLVrpIMcEA5GTMZgrbnUsjePNoD6kNuXzYrs4+/UIQfl8R51cuONh5jvLVT8SUR/6AYyVeUsK2VjwHFhIR5BRPTGxPsURbOWsjbzsVKFEiCfiG2ciUaexjZClXhRR0wpLqHu5luiR6FMXfE7BrjeJ/VTU37n9580tzlVKk1AA3Q==",
			},
		},
	}

	pretty.Logf("Actual Config %# v\n", cfg)
	if !(reflect.DeepEqual(cfg, expected)) {
		pretty.Logf("Expected Config %# v\n", expected)
		pretty.Pdiff(log.Default(), cfg, expected)
		log.Fatalf("Wrong configuration.")
	}

	if cfg.Integrations.PostgreSQL.IsOptional() {
		log.Fatalf("PostgreSQL integration should not be optional.")
	}

	// PostgreSQL Integration is active, as there are config env vars related to it.
	if !cfg.Integrations.PostgreSQL.IsActive() {
		log.Fatalf("PostgreSQL integration should be active.")
	}

	if !cfg.Integrations.MySQL.IsOptional() {
		log.Fatalf("MySQL integration should be optional.")
	}
	// MySQL Integration is not active, as there are no env variables related to MySQL
	if cfg.Integrations.MySQL.IsActive() {
		log.Fatalf("MySQL integration should not be active.")
	}

	if cfg.Integrations.S3.IsOptional() {
		log.Fatalf("S3 integration should not be optional.")
	}

	if !cfg.Integrations.SAML.IsActive() {
		log.Fatalf("SAML integration should be optional.")
	}

}
