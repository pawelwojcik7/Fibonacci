# syntax=docker/dockerfile:1.2

# Budowanie aplikacji
FROM golang AS builder

# Skopiowanie repozytorium Github
RUN git clone https://github.com/pawelwojcik7/Fibonacci.git
# Przjeście do katalogu budowania aplikacji
WORKDIR 'Fibonacci'
# Zbudowanie aplikacji
RUN CGO_ENABLED=0 go build -a -ldflags '-s' 

# Dystrybucja aplikacji
FROM scratch AS deploy
# Skopiowanie skompilowanej aplikacji 
COPY --from=builder /go/Fibonacci/Fibonacci .
# Wystawienie portu
EXPOSE 8080
# Uruchomienie aplikacji
CMD ["/Fibonacci"]
