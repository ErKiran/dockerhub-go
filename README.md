# dockerhub-go

> A Dockerhub client for Go

## Usage

```go
import "github.com/charliekenney23/dockerhub-go"
```

Construct a new DockerHub client, then use various services on the client to access different parts of the DockerHub API.

```go
client := dockerhub.NewClient(nil)

// login to Dockerhub
err := client.Auth.Login(context.Background(), "username", "password")

// or set an auth token directly
client.SetAuthToken(os.Getenv("DOCKERHUB_API_TOKEN"))
```

## License

MIT &copy; 2019 [Charles Kenney](https://github.com/charliekenney23)
