# authentication-api

Handles users auth

Owner: group:default/developer-experience

Docs in Backstage: [docs/index.md](docs/index.md)

Copy `.env.example` to `.env` for local values.

Install: `go mod download`

Install the `.env` pre-commit hook: `make install-hooks`

Run in development: `go run ./cmd/server`

Local Postgres: `docker compose up postgres`

Run the app image and Postgres: `docker compose up --build`

Same checks as CI: `make verify`

Unit tests: `make test`

Integration tests: `make test-integration`

Integration tests start their own dependencies with Testcontainers, so Docker
must be running locally. The same command runs in CI. Do not change
`.github/workflows/ci.yml` to add a dependency.

The scaffold ships one Postgres sample in
`internal/app/postgres_test.go`. That file is a pattern, not the full set
of dependencies this service will need. Replace it with tests of this
service's own components, or keep it and add more files next to it.

### Adding S3, SQS, or other AWS APIs

Use LocalStack. `go test` with `-tags=integration` already picks up every
`//go:build integration` file.

```bash
go get github.com/testcontainers/testcontainers-go/modules/localstack
```

Point the service at `AWS_ENDPOINT_URL` from the container; production leaves
that unset so the AWS SDK talks to real AWS.

- Only Postgres: keep the sample. Change the table and queries to match the schema.
- S3 and SQS, no database: replace `postgres_test.go` with a LocalStack test.
- Postgres, S3, and SQS: keep both files. Testcontainers starts both containers.

Smoke, performance, and regression tests against a deployed environment belong
in CD, not here.
