# <h1>golang_restapi</h1>

docker hub link => https://cloud.docker.com/repository/docker/x3xmermaid/golang_api

untuk menjalankan docker golang api menggunakan docker mysql
1. jalankan perintah berikut guna untuk mendownload container mysql ke dalam docker:<br>
</n>=> docker run --name mysqldb -v /my/own/datadir:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

2. cek ip container mysql yang telah dialankan :<br>
docker inspect mysqldb | grep IPAddr

3. buka database.go<br>
pada line 9 ganti ip mysql sesuai dengan ip yang telah didapatkan<br>
" db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/todo") "
 
4. jalankan perintah berikut untuk masuk kedalam mysql pada docker<br>
docker exec -it mysqldb bin/bash

5. login ke dalam mysql<br>
mysql -u root -proot
					
6. buat database baru dengan nama todo<br>
create database todo

7. pada cmd masuk ke dalam folder database kemudian jalankan perintah berikut untuk mengimport database ke dalam mysql docker<br>
docker exec -i mysqldb mysql -u root -proot todo < todo.sql

8. build dockerfile<br>
docker build -t golang_api .

9. jalankan image docker yag telah dibuat<br>
docker run -p 1:8080 golang_api
