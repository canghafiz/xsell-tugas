XSELL - Jual beli barang bekas terpercaya
===========

Ini isi backend aplikasinya

Prerequisites
-------------

-   Go 1.20 atau lebih baru
-   PostgreSQL
-   Git

Installation
------------

1.  Clone repository ini

```
https://github.com/canghafiz/xsell-backend.git
cd xsell-backend

```

1.  Install dependencies

```
go mod download

```

1.  Setup Environment Variables

Buat file `.env` di root directory project Anda dan isi dengan konfigurasi berikut:

```
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=nama_database
DB_USERNAME=root
DB_PASSWORD=

# Application Configuration
APP_NAME=YourAppName
APP_PORT=8080
APP_STATUS=Production/Debug

# JWT Configuration
JWT_KEY=your_jwt_secret_key_here

```

### Penjelasan Environment Variables

| Variable | Deskripsi | Default Value | Required |
| --- | --- | --- | --- |
| `DB_HOST` | Host database server | `localhost` | Yes |
| `DB_PORT` | Port database server | `3306` | Yes |
| `DB_DATABASE` | Nama database | - | Yes |
| `DB_USERNAME` | Username database | `root` | Yes |
| `DB_PASSWORD` | Password database | (kosong) | No |
| `APP_NAME` | Nama aplikasi | `YourAppName` | Yes |
| `APP_STATUS` | Environment aplikasi (development/production) | `development` | Yes |
| `APP_PORT` | Port aplikasi | `8080` | Yes |
| `JWT_KEY` | Secret key untuk JWT authentication | - | Yes |

### Contoh File .env

Copy file `.env.example` ke `.env`:

```
cp .env.example .env

```

Lalu sesuaikan nilai-nilai environment variable sesuai kebutuhan Anda.

Database Setup
--------------

1.  Buat database baru:

```
CREATE DATABASE nama_database;

```

1.  Jalankan migration:

```
migrate -database "postgres://username:pw@tcp(localhost:port)/db_name" -path db/migrations up

```

Usage
-----

### Development

Jalankan aplikasi dalam mode development:

```
go run main.go

```

### Build

Build aplikasi:

```
go build -o bin/app main.go

```

### Run Production

Jalankan binary hasil build:

```
./bin/app

```

Project Structure
-----------------

```
.
├── app/
├── controller/
├── db/
├── helper/
├── models/
├── .env
├── .env.example
├── go.mod
├── go.sum
└── main.go

```

API Documentation
-----------------

API documentation dapat diakses di:

-   https://documenter.getpostman.com/view/18098423/2sB3dWr6zU

Contributing
------------

Pull requests are welcome. Untuk perubahan besar, silakan buka issue terlebih dahulu untuk mendiskusikan apa yang ingin Anda ubah.

License
-------

[MIT](https://choosealicense.com/licenses/mit/)

---

© 2025 Hafizarrahman