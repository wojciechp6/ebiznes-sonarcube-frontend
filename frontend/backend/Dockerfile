# Używamy oficjalnego obrazu Go
FROM golang:1.24

RUN apt-get update
RUN apt-get install -y build-essential gcc

# Kopiujemy pliki go.mod i go.sum do kontenera
COPY go.mod go.sum ./

# Pobieramy zależności
RUN go mod download

# Kopiujemy wszystkie pliki źródłowe do kontenera
COPY . .

# Tworzymy katalog na dane SQLite
RUN mkdir -p data

ENV CGO_ENABLED=1

# Budujemy aplikację (plik wykonywalny o nazwie "backend")
RUN go build -o backend .

# Ustawiamy domyślny punkt wejścia
CMD ["./backend"]

# Eksponujemy port, na którym nasłuchuje aplikacja
EXPOSE 8080