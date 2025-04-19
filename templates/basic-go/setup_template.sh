#!/bin/bash

BASE_DIR="templates/basic-go"

echo "📁 Membuat struktur folder untuk template di: $BASE_DIR"

# Daftar folder
mkdir -p $BASE_DIR/cmd
mkdir -p $BASE_DIR/config
mkdir -p $BASE_DIR/internal/handlers
mkdir -p $BASE_DIR/internal/models
mkdir -p $BASE_DIR/internal/routes
mkdir -p $BASE_DIR/internal/service
mkdir -p $BASE_DIR/internal/utils
mkdir -p $BASE_DIR/pkg/middleware

# File: cmd/main.go
cat <<EOF > $BASE_DIR/cmd/main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello from gcraft project!")
}
EOF

# File: config/*.go
for file in app.go config.go fiber.go gorm.go loglrus.go validator.go; do
cat <<EOF > $BASE_DIR/config/$file
package config

// TODO: Implement $file
EOF
done

# File: .env
cat <<EOF > $BASE_DIR/.env
APP_ENV=development
PORT=3000
EOF

# File: .env.example
cat <<EOF > $BASE_DIR/.env.example
APP_ENV=
PORT=
EOF

# File: go.mod
cat <<EOF > $BASE_DIR/go.mod
module your_project_name

go 1.20
EOF

echo "✅ Struktur dan file template berhasil dibuat!"
