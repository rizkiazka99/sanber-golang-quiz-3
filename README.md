# Quiz 3 - Golang Bootcamp Batch 68 - SanberCode

## Cara Pengunaan
REST API ini memiliki 2 table utama, yakni Books dan Categories. Sebelum dapat mengakses kedua tabel tersebut, perlu dilakukannya Register dan Login terlebih dahulu karena semua endpoint dari kedua tabel tersebut telah diproteksi oleh JSON Web Token yang diperoleh dengan melakukan Login. Access Token memiliki masa aktif selama 1 jam.

- Base URL: https://sanber-golang-quiz-3-production.up.railway.app/

### Auth
- Register</br>
(POST) /api/register</br>
```json
Request Body:
{
    "username":"",
    "password":""
}
```
Password akan dienkripsi menjadi Hash oleh bcrypt setelah endpoint berhasil terpanggil.

- Login</br>
(POST) /api/login</br>
```json
Request Body:
{
    "username":"",
    "password":""
}
```
Masukan Username dan Password yang telah diregistrasi, tidak bisa menggunakan random value.
```json
Output:
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
- Create Category</br>
(POST) /api/categories</br>
```json
Request Body:
{
    "name": "Technology"
}
```
- Get All Categories</br>
(GET) /api/categories</br>
- Get Category by ID</br>
(GET) /api/categories/:id</br>
- Get Books by Category ID</br>
(GET) /api/categories/:id/books</br>
- Update Category</br>
(PUT) /api/categories/:id</br>
- Delete Category</br>
(DELETE) /api/categories/:id</br>

### Books
- Create Book</br>
(POST) /api/books</br>
```json
Request Body:
{
  "title": "Computer Hardware 101",
  "category_id": 20250721024035261,
  "description": "A beginner's guide to computers and the process of building one",
  "image_url": "https://example.com/images/great-adventure.jpg",
  "release_year": 2021,
  "price": 15,
  "total_page": 320
}
```
- Get All Books</br>
(GET) /api/books</br>
- Get Book by ID</br>
(GET) /api/books/:id</br>
- Update Book</br>
(PUT) /api/books/:id</br>
- Delete Category</br>
(DELETE) /api/books/:id</br>

