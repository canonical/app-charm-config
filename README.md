# app-charm-config

`app-charm-config` is a tool to generate Go structs for a 12 factor Go app generated using rockcraft and charmcraft
go-framework extensions.


## Usage

Create a Kubernetes charm for a Go project following the [tutorial (TODO)](https://juju.is/docs/sdk/write-your-first-kubernetes-charm-for-a-go-app) using
the rockcraft and charmcraft go-framework.

You can choose to install the tool with `go install github.com/canonical/app-charm-config@latest` or run it directly without install.

You can add a directive similar to the following one to one of your Go files, replacing the placeholders:
`//go:generate app-charm-config -c <charmcraft_yaml_file> -o <output_go_file> -p <package_name>`

Without installing, you can create a directive like the following one:
`//go:generate go run github.com/canonical/app-charm-config@latest -c <charmcraft_yaml_file> -o <output_go_file> -p <package_name>`

You can also use the tool directly without a `go:generate` directive. Once the Go source code file with the structs is generated,
you can use it with the "github.com/caarlos0/env/v11" library like:
```
	import (
	    ...
	    "github.com/caarlos0/env/v11"
	    "your-go-app/configpackage"
	)
	...
	var cfg configpackage.CharmConfig
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("Error parsing configuration: %v", err)
	}
```

You can see the contents of the `CharmConfig` struct looking at the generated Go file.
