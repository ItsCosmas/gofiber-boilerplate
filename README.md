[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

# Go (Golang) and Fiber RESTful API Boilerplate

## 🚧 WORK IN PROGRESS

- Could be used to provide quick bootstrap functionality for your next go and gofiber app.
- I will update(try) this regularly to add functionality and new features.
- Well this is just one way of doing it, not the official or the best 😅.

**Used libraries:**

- [Gofiber](https://gofiber.io/)
- [Gorm](https://gorm.io/)
- [jwt-go](https://github.com/form3tech-oss/jwt-go)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)
- [uuid](https://github.com/google/uuid)
- [Go Mongo Driver](https://go.mongodb.org/mongo-driver)
- [Go Playground Validator](https://github.com/go-playground/validator)

---

### Features

- [x] Gofiber Docker Dev Setup with Hot Reload
- [x] User Auth functionality (Signup, Login, Forgot and Reset Password)
- [x] JWT Authentication
- [x] RESTful API
- [x] Swagger REST API documentation
- [x] Gorm (Golang SQL DB ORM) with Postgres implementation and auto migration
- [x] MongoDB using the official mongo driver
- [x] Configs via environmental variables
- [x] Improved Input Validations(could be better)
- [x] Custom Error messages
- [ ] Email notification (Welcome email, Reset password email)
- [ ] Redis
- [ ] Casbin
- [ ] WebSocket
- [ ] gRPC
- [ ] Improve MongoDB data integrity

---

## Running and Developing locally

1. Create `.env` , i.e.

```sh
cp .env.example .env
```

2. Download and install Swag for generating docs

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

3. Run

- NOTE: You have to generate swagger docs before running the app.

```sh
# Terminal 1
swag init # Generates Swagger Docs

# Terminal 2
docker compose --env-file .env up        # docker compose up (Run App With AutoReload)
docker compose --env-file .env down      # docker compose down (Shutdown App)
```

- API `http://localhost:8000/api/v1`
- Swagger Doc `http://localhost:8000/api/v1/swagger`

---

## Packaging For Production

1. Create `.env` , i.e.

```sh
cp .env.example .env
```

2. Update your `.env` variables for production

- Point to your prod database
- Update JWT issuer, secret key , blah blah
- Basically just follow good production practice

3. Download Swag for generating docs

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

- Generate Swagger Docs. You have to generate swagger docs before packaging the app.

```sh
swag init # Generates Swagger
```

4. Build Your Image

- Permission the build script to run.

```
chmod +x docker-build.sh
```

- You could set the image port on `Dockerfile.prod`
- Run the build script. You must provide a version tag as shown below.

```
./docker-build.sh -v gofiber:1.0.0
```

---

### Todo

- [ ] Data Migrations ?
- [ ] Logger
- [ ] Unit tests

maybe?

- [ ] SMS notification (2FA ,Reset password code)
- [ ] GraphQL
- [ ] Deploy on Kubernetes
- [ ] Write an article

---

### Gotcha's

- Building Swago from source code - `go build -o swag.exe cmd/swag/main.go`

### Contribution

Open to Suggestions and Pull Requests

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
