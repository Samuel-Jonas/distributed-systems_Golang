# Start from golang base image
FROM golang:1.19.4-alpine3.17

# Setup folders
WORKDIR /app

# copy app
COPY . /app

# download dependencies
RUN go mod download

# build app
RUN go build -o main main.go

# Expose port to the outside world
EXPOSE 5000

RUN chmod +x ./main

# Run the executable
CMD [ "./main" ]