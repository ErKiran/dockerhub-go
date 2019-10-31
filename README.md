# dockerhub-go [![Build Status][travis-ci-badge]][travis-ci] [![GoDoc][godoc-badge]][godoc]

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

MIT &copy; 2019 [Charles Kenney][profile]

[travis-ci-badge]: https://travis-ci.org/Charliekenney23/dockerhub-go.svg?branch=master
[travis-ci]: https://travis-ci.org/Charliekenney23/dockerhub-go
[godoc-badge]: https://godoc.org/github.com/Charliekenney23/dockerhub-go?status.svg
[godoc]: https://godoc.org/github.com/Charliekenney23/dockerhub-go
[profile]: https://github.com/charliekenney23
