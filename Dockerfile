# Etapa 1: build da aplicação
FROM golang:1.24.2-alpine AS builder

# Cria diretório da app
WORKDIR /app

# Copia os arquivos go
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila a aplicação para um binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Etapa 2: imagem final enxuta
FROM alpine:latest

# Define diretório da aplicação
WORKDIR /root/

# Copia o binário da etapa anterior
COPY --from=builder /app/server .

# Expõe a porta usada pela aplicação
EXPOSE 8080

# Comando de start
CMD ["./server"]
