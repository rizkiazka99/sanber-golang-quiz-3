# Quiz 3 - Golang Bootcamp Batch 68 - SanberCode

## Cara Pengunaan
REST API ini memiliki 2 table utama, yakni Books dan Categories. Sebelum dapat mengakses kedua tabel tersebut, perlu dilakukannya Register dan Login terlebih dahulu karena semua endpoint dari kedua tabel tersebut telah diproteksi oleh JSON Web Token yang diperoleh dengan melakukan Login. Access Token memiliki masa aktif selama 1 jam.

- Base URL: https://sanber-golang-quiz-3-production.up.railway.app/

### Auth
- Register</br>
/api/register</br>
```json
{
    "username":"",
    "password":""
}
```
Password akan dienkripsi menjadi Hash oleh bcrypt setelah endpoint berhasil terpanggil.

- Login</br>
/api/login</br>
```json
{
    "username":"",
    "password":""
}
```
Masukan Username dan Password yang telah diregistrasi, tidak bisa menggunakan random value.
- Output Login:</br>
```json
{
    "data": {
        "id": 20250720184814636,
        "username": "rizkiazka99",
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTMwNDA5MDEsImlkIjoiMjAyNTA3MjAxODQ4MTQ2MzYifQ.MpGLLQw2jqJ_4QlsCs-RooFoRhgk0Zd7cYu-a_T1fxA",
        "token_expiration_time": "2025-07-20T19:48:21.068585943Z"
    }
}
```
Masukan value "access_token" ke dalam key "Authorization" pada Header HTTP dan gunakan Bearer.

### Categories
