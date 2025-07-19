# Soal 2 HTTP Validasi

## Contents

- [Description](#description)
- [Installation](#installation)

## Description

Program ini merupakan http server sederhana menggunakan package bawaan net/http dengan endpoint `/validate`. Pada endpoint ini akan menerima query param berupa email dan age. Kemudian akan dilakukan validasi sebagai berikut:

- Validasi bahwa `email` tidak kosong dan `age` ≥ 18
- Return response JSON: `{ "status": "ok" }` atau `{ "error": "message" }`
- Log setiap request yang masuk (gunakan `logrus`)
- Gunakan `go mod` dan struktur idiomatik
- Buat repo GitHub dan gunakan commit meaningful

## Installation

1. Open your terminal or command prompt
2. Clone the project

```bash
$ git clone https://github.com/sipamungkas/belajar-golang-hsi
```

3. Move inside the directory and install and link dependency

```bash
$ cd belajar-golang-hsi/soal-2-http-validasi
```

4. Run program

```bash
$ go run main.go
```

5. Check the http response using
```bash
$ curl -X GET 'http://localhost:8000/validate?email=ragil@sipamungkas.com&age=18'
```
