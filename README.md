# ECHO Microservice Project

Ini adalah proyek yang menggunakan [Echo Framework](https://echo.labstack.com/) di Go, dengan prinsip-prinsip SOLID, autentikasi JWT, validasi permintaan, dan komunikasi antar-mikroservis dengan gRPC.

## Table of Contents
- [Overview](#overview)
- [Fitur](#fitur)
- [Arsitektur SOLID](#arsitektur-solid)
- [JWT Authentication](#jwt-authentication)
- [Request Validator](#request-validator)
- [gRPC Communication](#grpc-communication)
- [Getting Started](#getting-started)
- [Requirements](#requirements)

## Overview

Proyek ini adalah contoh penerapan `Echo` di Go, yang dibangun dengan arsitektur SOLID untuk memudahkan skalabilitas, perawatan, dan testing. Proyek ini mendukung autentikasi berbasis JWT untuk melindungi endpoint, validasi request yang kuat menggunakan `validator`, dan komunikasi antar layanan menggunakan `gRPC`.

## Fitur

1. **REST API** - Penggunaan Echo untuk pengelolaan routing dan middleware.
2. **Authentication JWT** - Untuk mengamankan akses ke endpoint.
3. **Request Validation** - Menggunakan `github.com/go-playground/validator/v10` untuk validasi data.
4. **gRPC** - Komunikasi antar-mikroservis dengan protokol gRPC.
5. **SOLID Principle** - Untuk menjaga struktur kode yang modular dan mudah dikembangkan.

## Arsitektur SOLID

Prinsip SOLID diterapkan dalam struktur proyek ini untuk memastikan pemisahan tanggung jawab, modularitas, dan fleksibilitas. Implementasi ini memungkinkan pengembang untuk dengan mudah memperbarui atau memperluas fitur tanpa mempengaruhi komponen lainnya.

1. **Single Responsibility Principle** - Setiap layer (handler, service, repository) memiliki satu tanggung jawab spesifik.
2. **Open/Closed Principle** - Kelas dan fungsi mudah diperluas tanpa mengubah fungsionalitas dasar.
3. **Liskov Substitution Principle** - Menggunakan interface agar modul dapat digantikan dengan implementasi lainnya.
4. **Interface Segregation Principle** - Interface spesifik untuk layanan tertentu, sehingga tidak semua fungsi harus diimplementasi.
5. **Dependency Inversion Principle** - Menggunakan dependency injection untuk ketergantungan antar komponen.

## JWT Authentication

Fitur autentikasi JWT digunakan untuk melindungi endpoint yang sensitif. Hanya pengguna dengan token yang valid yang dapat mengakses resource tertentu.

- **Proses Login**: Pengguna mengirim kredensial dan menerima token JWT sebagai respons.
- **Akses Terproteksi**: Setiap endpoint yang membutuhkan autentikasi akan memvalidasi token JWT terlebih dahulu.

## Request Validator

Proyek ini menggunakan [go-playground/validator](https://github.com/go-playground/validator) untuk validasi input data. Validator ini memastikan bahwa data yang masuk sesuai dengan aturan yang telah ditentukan, sehingga dapat mengurangi kesalahan dari data yang tidak valid.

- **Penggunaan `ctx.Validate`**: Semua request diproses dengan `ctx.Validate` untuk memeriksa data sebelum masuk ke layer service.

## gRPC Communication

gRPC digunakan untuk berkomunikasi dengan mikroservis lain, seperti layanan pembayaran. Implementasi gRPC memungkinkan komunikasi antar-layanan yang cepat, aman, dan efisien.

- **Protokol Buffer (proto)**: Mendefinisikan pesan dan layanan untuk komunikasi.
- **Penggunaan gRPC di Payment Service**: Pembaruan status pembayaran di Order Service dilakukan melalui gRPC, memastikan bahwa setiap perubahan status dikirim dan diterima dengan akurat.

## Getting Started

Ikuti langkah-langkah berikut untuk memulai proyek ini.

### Requirements

- **Go 1.23 atau lebih baru**
- **PostgreSQL** untuk penyimpanan data
- **Docker** (opsional, untuk memudahkan setup lingkungan pengembangan)

