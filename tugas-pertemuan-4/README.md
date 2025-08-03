# Tugas Pertemuan 3

Sistem Informasi Mahasiswa Modular

## Contents

- [Description](#description)
- [Installation](#installation)

## Description

Program ini merupakan program untuk melakukan implementasi pembelajaran pada pertemuan 3, meliputi:

- Penerapan Worker, channel, goroutine
-
- Struct dan interface (beserta polymorphism)
- Modular programming menggunakan package
- Scope dan pengaturan akses (private/public)

### Folder Structure

tugas-pertemuan-4/
│
├── cmd/
│ └── main.go
│
├── models/
│ ├── mahasiswa.go
│ ├── tugas.go
│ └── hasil.go
│
├── worker/
│ ├── assignment_worker.go
│ └── grading_worker.go
│
├── config/
│ └── config.go
│
├── docker-compose.yaml
├── init.sql
├── .env
└── go.mod

### Entity & Relasi:

- **Mahasiswa** `(ID, Nama)`
- **Tugas** `(ID, Judul, Deskripsi, MahasiswaID)`
- **Hasil** `(ID, TugasID, Nilai)`

Relasi:

- Mahasiswa **has many** Tugas
- Tugas **has one** Hasil

### Proses:

1. **Worker 1** `assignment_worker` menerima list mahasiswa dan memberikan masing-masing 1 tugas.
2. **Worker 2** `grading_worker` membaca semua tugas, lalu memberi nilai acak (0–100) dan menyimpan ke table `hasil`.

## Installation

1. Open your terminal or command prompt
2. Clone the project

```bash
$ git clone https://github.com/sipamungkas/belajar-golang-hsi
```

3. Move inside the directory and install and link dependency

```bash
$ cd belajar-golang-hsi/tugas-pertemuan-4
```

4. install dependency

```bash
$ go get
```

4. Run program

```bash
$ go run main.go
```
