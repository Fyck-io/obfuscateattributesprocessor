# Obfuscate Attributes Processor

It's not working rb :(

## Building your own Collector

```
git clone git@github.com:open-telemetry/opentelemetry-collector.git
cd opentelemetry-collector/cmd/builder
go build
./builder --config builder-config.yaml --name fyck-collector --output-path ./
```

Notes:

- `go build` should generate the builder bin
- `builder-config.yaml` is the one in this project