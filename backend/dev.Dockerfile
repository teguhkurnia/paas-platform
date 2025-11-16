FROM golang:alpine3.22

WORKDIR /app

# Install Railpacks untuk build app
RUN curl -sSL https://railpack.com/install.sh | sh

# Install 'air', alat hot-reload untuk Go
RUN go install github.com/air-verse/air@latest

# Salin file dependensi & download
COPY go.mod go.sum ./
RUN go mod download

# Salin sisa kode (ini hanya untuk build awal)
COPY . .

# Perintah 'air' akan menonton file dan me-restart server secara otomatis
CMD ["air", "-c", ".air.toml"]
