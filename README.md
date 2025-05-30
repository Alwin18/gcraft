# 🚀 gcraft

**gcraft** adalah CLI sederhana untuk memulai proyek Go dengan cepat dan efisien. Hanya dengan satu perintah, kamu dapat membuat struktur proyek, menginstal dependensi, dan menjalankan aplikasi.

---

## 📦 Instalasi

**1. Instal langsung dengan:**

```bash
go install github.com/Alwin18/gcraft@latest
```

**2. Clone repository ini dan tambahkan ke PATH:**

```bash
git clone https://github.com/username/gcraft.git
cd gcraft
go build -o gcraft
sudo mv gcraft /usr/local/bin/
```

## Usage

```bash
gcraft create my-app
cd my-app
go mod tidy
go run cmd/main.go
```

## Commands
1. --help
2. create