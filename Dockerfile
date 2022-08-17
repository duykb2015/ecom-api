# Lỗi quá dockẻ ơi

# import modules
FROM golang:1.18-alpine3.16 as modules
# copy two file go.mod and go.sum to the modules folder
COPY go.mod go.sum /modules/
# set working directory to the modules folder
WORKDIR /modules
# download the dependencies
RUN go mod download
# copy all file from workdir (ecom-cms) to the app folder
COPY . /app

WORKDIR /app

EXPOSE 8080

CMD ["/app"]