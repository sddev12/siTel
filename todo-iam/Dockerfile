FROM alpinelinux/golang as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o todo-iam

FROM alpinelinux/golang
WORKDIR /app
COPY --from=build /app/todo-iam .
EXPOSE 3002
CMD ["./todo-iam"]