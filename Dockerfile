FROM golang:alpine
 
RUN mkdir /app

WORKDIR /app

# ADD go.mod .
# ADD go.sum . 

COPY ./ .

# ADD . . 

# Installs Go dependencies
# RUN go mod download
 
# Builds your app with optional configuration
# RUN go build -o /main.go
 
# RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 4030

# ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

CMD [ "go", "run", "main.go" ]