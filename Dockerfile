# Tahap build
FROM golang:1.22.1-alpine AS build

# Direktori kerja
WORKDIR /goapp

# Copy semua file ke dalam image build
COPY . .

# Unduh dependencies dan lakukan build
RUN go mod download
RUN go build -v -o /goapp/gostore ./cmd/main.go


# Tahap final untuk runtime
FROM alpine:3.14

# Direktori kerja pada image runtime
WORKDIR /app

# Salin file binary hasil build ke image runtime
COPY --from=build /goapp /app/

# Menambahkan binary ke PATH
ENV PATH="/app:${PATH}"

# Membuka port 8081
EXPOSE 8081

# Command yang dijalankan saat container start
ENTRYPOINT [ "gostore" ]

# Instruksi untuk build image
# docker build -t zikrigusli/onlinestore:1 