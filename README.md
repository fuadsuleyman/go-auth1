********************************* About DB Start *************************************
<!-- users tablenin ici .sql fayli -->
CREATE TABLE users
(
    id serial not null unique,
    usertype smallint not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
)

<!-- docker -de postgres qoshmaq ucun addimlar -->
- docker pull postgres (terminalda)
- docker run --name=auth-db -e POSTGRES_PASSWORD='fuaddauf' -p 5436:5432 -d --rm postgres (terminalda)

<!-- mende migrate install olmadigi ucun asagidakilari eledim -->
<!-- asagidaki komndalari root olaraq etmek lazimdi ona gore evver terminalda sudo -s -->
- curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
- echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
- apt-get update
- apt-get install -y migrate

<!-- ve yene davam edirem -->
- migrate create -ext sql -dir ./schema -seq init (terminalda)
<!-- yuxaridaki komandadan sonra schema papkasi ve icinde 2 .sql fayl yaranir -->

<!-- .sql fayllarinda deyisiklik etdikden sonra asagidaki komanda ile migrate edirik -->
- migrate -path ./schema -database 'postgres://postgres:fuaddauf@localhost:5436/postgres?sslmode=disable' up
<!-- mende error (Dirty database version 1. Fix and force version.) cixdi, ona gore asagidaki kimi eledim -->
- migrate -path ./schema -database 'postgres://postgres:fuaddauf@localhost:5436/postgres?sslmode=disable' force 1 up
<!-- sonra yeniden adi migrate ... up -i eledim terminalda seriya cixdi -->

<!-- teminalda girib data bazaya baxmaq ucun -->
- docker ps -a (elmekle postgres-in container adini gotururuk)
- docker exec -it a2f7eec2073d bin/bash
- psql -U postgres
- \d  // etmekle table adlari gorunur

<!-- table-leri silmek ucun asagidaki -->
- migrate -path ./schema -database 'postgres://postgres:fuaddauf@localhost:5436/postgres?sslmode=disable' down
********************************* About DB Start *************************************

********************************* DB Connect from Project *************************************
<!-- sql-le ishlemek ucun asagidakini yukledim -->
- go get -u github.com/jmoiron/sqlx

<!-- passwordu config file-da saxlamaq duzgun deyil .env-le islemek ucun asagidakini yukledim -->
- go get github.com/joho/godotenv
********************************* DB Connect from Project Finish *************************************

********************************* For Logging *************************************
- go get -u github.com/sirupsen/logrus
******* END *******

***************************** Registration Notes *********************************

******* END *******

***************************** login Notes *********************************
- go get -u github.com/dgrijalva/jwt-go

******* END *******

Can not find edende export GO111MODULE=on 

### Docker run migrate
/home/fuad/golang/rest-api/register1/schema

### it is not work
docker run -v /home/fuad/golang/rest-api/register1/schema:/migrations --network host migrate/migrate
    -path=/migrations/ -database 'postgres://postgres:fuaddauf@localhost:5436/postgres?sslmode=disable' up

docker run -v ./schema:/migrations --network host migrate/migrate
    -path=/migrations/ -database 'postgres://postgres:fuaddauf@localhost:5432/postgres?sslmode=disable' up