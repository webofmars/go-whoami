# go-whoami

A cute gopher displaying the hostname where it is running.

Useful for testing load balancing, ingress rules, DNS resolution and cache behaviour in Kubernetes or Docker environments.

![Docker Image](https://img.shields.io/docker/v/webofmars/go-whoami?label=docker%20hub)
![Image Size](https://img.shields.io/docker/image-size/webofmars/go-whoami)

## Features

- Displays the hostname of the container/pod it runs in
- Single static binary, final image based on `scratch` (~6.5 MB)
- Multi-arch images (`linux/amd64`, `linux/arm64`)
- Sends anti-caching headers (`Cache-Control: no-store`, `Pragma`, `Expires`) so browsers always show fresh content

## Quick start

### Docker

```bash
docker run -d -p 3000:3000 webofmars/go-whoami:latest
```

Then open <http://localhost:3000>.

### Kubernetes

```bash
kubectl create deployment whoami --image=webofmars/go-whoami:latest --replicas=3
kubectl expose deployment whoami --port=3000
```

Scale up and refresh your browser (or `curl` the service) to see the hostname change across pods.

## Build from source

### Requirements

- Go 1.22+
- Docker (for container builds)

### Local build

```bash
go build -o app .
./app
```

The server listens on port **3000**.

### Container build

```bash
docker build -t webofmars/go-whoami:latest .
```

Multi-arch build & push:

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --provenance true --sbom true \
  --push \
  -t webofmars/go-whoami:latest .
```

## Project structure

```
.
├── whoami.go                # main application
├── go.mod                   # Go module definition
├── Dockerfile               # multi-stage build (golang → scratch)
├── templates/
│   └── index.html.tmpl      # page template (hostname injected at startup)
└── html/                    # static assets (CSS, images)
```

## How it works

At startup, the app renders `templates/index.html.tmpl` with the OS hostname into `html/index.html`, then serves the `html/` directory via `http.FileServer`. A middleware adds `Cache-Control: no-cache, no-store, must-revalidate`, `Pragma: no-cache` and `Expires: 0` headers to every response to prevent browser caching issues (especially with Chromium-based browsers).

## Configuration

| Setting | Value | Notes |
|---------|-------|-------|
| Port    | `3000` | hardcoded |

## Contributing

Issues and pull requests are welcome. Dependency updates for Go modules and Docker base images are handled by [Dependabot](.github/dependabot.yml).

## License

See [LICENSE](LICENSE) if present, otherwise all rights reserved by the author.
