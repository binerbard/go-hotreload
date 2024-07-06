# Go-Hotreload

**go-hotreload** adalah proyek Go yang mendukung hot-reloading untuk aplikasi server-side. Proyek ini memonitor perubahan dalam direktori proyek dan secara otomatis mengkompilasi ulang dan menjalankan ulang server setiap kali perubahan terdeteksi.

## Fitur

**Hot Reloading:** Secara otomatis mengkompilasi ulang dan menjalankan ulang server Go saat perubahan terdeteksi.

**Integrasi dengan fsnotify:** Menggunakan fsnotify untuk mendeteksi perubahan file.

**Pengaturan dan Penggunaan Sederhana:** Mudah dikonfigurasi dan digunakan.

## Persyaratan

Go 1.20 atau lebih baru

### Clone repository:

```console
git clone https://github.com/username/go-hotreload.git
cd go-hotreload
```

### Inisialisasi dan Instalasi Dependensi:

```console
go mod tidy
```

## Penggunaan

### Jalankan Aplikasi:

```console
go run main.go
```

Aplikasi akan memonitor perubahan dalam direktori proyek dan secara otomatis mengkompilasi ulang dan menjalankan ulang server saat perubahan terdeteksi.
