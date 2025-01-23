# Radar

# TODO:

## GUI
 - Think about fixies main menu buttons positions on Macos
 - Reinit model on run (dogswim away)
   - Stop modeling on exit to main menu??
 - Set app icon
 - Update app name (now Desctop)
 - Support two radiuses: For fighting and hissing

## Profiler run

```bash
go run ./cmd/engine/engine.go
go tool pprof -http=:8080 cpu.pprof

```