FROM golang AS builder
RUN git clone https://github.com/pawelwojcik7/Fibonacci.git
RUN cd Fibonacci
WORKDIR 'Fibonacci'
RUN CGO_ENABLED=0 go build -a -ldflags '-s' 

FROM scratch AS deply
COPY --from=builder /go/Fibonacci/Fibonacci .
EXPOSE 8080
CMD ["/Fibonacci"]
