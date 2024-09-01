# Descripción

**Go's Built-in Fuzzing:** Desde **Go 1.18**, Go incluye soporte nativo para fuzzing. Puedes utilizar el marco de fuzzing integrado en Go, que es más adecuado para programas escritos en Go. Aquí tienes un ejemplo básico:

    go test -fuzz=FuzzVulnerableHandler -fuzztime=10s


    GOMAXPROCS=1 go test -fuzz=FuzzVulnerableHandler -fuzztime=30s