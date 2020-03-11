## Docker Compose Content
* MySQL 5.7
* RabbitMQ
* Elasticsearch 7.5
* Kibana
* News Aplication - REST
* News Aplication - Event

## Running Application
* Rename file `example-env` menjadi `.env`
* Untuk menjalankan seluruh aplikasi cukup run command :
```sh
$ make run
```
* Lalu buat elasticsearch index, bisa menggunakan kibana console dengan mengakses url : http://localhost:5601
lalu jalankan command :
```sh
PUT news
```
atau menggunakan comand CURL :
```sh
$ curl --location --request PUT 'http://localhost:9200/news'
```
* Untuk memberhentikan seluruh service cukup run command :
```sh
$ make stop
```

## Unit testing
untuk menjalankan unit testing run command :
```sh
$ go test $(go list ./...) -cover
```

## API Spesification
Aplikasi news REST API running di port `8080` dengan list endpoint :
* `GET /news` : untuk get data list news, pada endpoint ini terdapat 2 query paramter : `page` dan `limit` (mandatory), example request :
```sh
GET http://localhost:8080/news?limit=10&page=1
```
* `POST /news` : untuk insert data data news, pada endpoint ini terdapat 2 json paramter : `author` dan `body` (mandatory), example request :
```sh
POST http://localhost:8080/news
{
    "author": "testing 5",
    "body": "Ini test article"
}
```
* `GET /infrastructure/healthcheck` : untuk keperluan healthcheck cloud load balancer

## Architecture
aplikasi news terbagi menjadi 2 bagian, `REST server` dan `Event server`, REST server berfungsi untuk menerima request response dan Event server untuk melakukan tugas background processing (worker).
flow yang terjadi pada saat proses create news setelah REST server menerima request, sistem akan melakukan validasi dari json payload yang
dikirim, apakah semua field yang besifat mandatory sudah di fill atau tidak, jika payload sesuai sistem akan melakukan publish message ke `rabbitMQ server`
sebagai data queueing dengan topic `insert-news`, jika proses publish ke `rabbitMQ` sukses REST server akan memberikan response sukses.

lalu `Event server` akan melakukan subscribe ke topic `insert-news`, jika ada message queueing yang di publish saat create news, akan otomatis 
dikonsume oleh `Event server` dan akan di simpan ke dua data source yaitu `MySQL` and `Elasticsearch`, di `Elastcisearch` data yang disimpan hanya `id` dan `created`
sedangkan seluruh field data akan di simpan ke `MySQL`.

jadi pada saat user melakukan request ke endpoint `GET /news`, REST server akan melakukan query ke datasource `Elasticsearch` lalu melakukan populate data
dengan mengambil detail data news ke datasource `MySQL` dan result hasil dari populate data per rows akan di cache ke `redis server` dengan data key news `id` yang memiliki expire time 15 detik, 
jadi pada saat ada request selanjut nya ke `GET /news`, sistem akan melakukan pengecekan ke `redis server`, apakan key news `id` tersebut sudah ada di cache atau belum, kalau ada ambil dari cache, jika belum ambil dari `MySQL`
