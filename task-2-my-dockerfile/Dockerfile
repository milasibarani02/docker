# Gunakan official golang image sebagai base image
FROM golang:1.20-alpine

# Setel working directory di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum jika ada
# COPY go.mod go.sum ./

# Salin semua file dari project ke container
COPY . .

# Build aplikasi
RUN go build -o wise-word .

# Salin AUTHORS.md ke dalam image Docker
COPY AUTHORS.md /AUTHORS.md

# Ekspos port 80 untuk aplikasi
EXPOSE 80

# Jalankan aplikasi saat container mulai
CMD ["/app/wise-word"]