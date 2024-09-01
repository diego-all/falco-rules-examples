# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY main.go .

# Compilar la aplicación
RUN go build -o api main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copiar el binario desde la etapa de compilación
COPY --from=builder /app/api .

# Exponer el puerto que utiliza la API
EXPOSE 8080

# Comando para ejecutar la API
CMD ["./api"]
