# rest-api tutorial 

## Install commands
### Build the app
`mod go build -o bin/go-rest-api internal/main.go`

### Run the app
`mod go run internal/main.go`

### Generate Go code
`mod go generate github.com/albacanete/learning-go/rest-api/internal github.com/albacanete/learning-go/rest-api/pkg/swagger`

### Validate swagger
`swagger validate pkg/swagger/swagger.yml`

### Doc for swagger
`docker run -i yousan/swagger-yaml-to-html < pkg/swagger/swagger.yml > doc/index.html`

## API endpoints

1. /healthz
2. /hello/{user}
3. /gopher/{name}
    Try `localhost:8080/gopher/dr-who` :)