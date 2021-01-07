[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

# Go (Golang) and Fiber REST API Boilerplate

**Used libraries:**

- [Gofiber](https://gofiber.io/)
- [Gorm](https://gorm.io/)
- [jwt-go](https://pkg.go.dev/gopkg.in/dgrijalva/jwt-go.v3?tab=doc)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)
- [uuid](github.com/google/uuid)
- [Go Mongo Driver](go.mongodb.org/mongo-driver)
- [Go Playground Validator](github.com/go-playground/validator/v10)

---

### Features

- [x] Gofiber Docker Dev Setup with Hot Reload
- [x] User Auth functionality (Signup, Login, Forgot Password, Reset Password)
- [x] JWT Authentication
- [x] REST API
- [x] Swagger REST API documentation
- [x] Gorm (Golang SQL DB ORM) with Postgres implementation and auto migration
- [x] MongoDB using the official mongo driver
- [ ] Redis
- [x] Configs via environmental variables
- [ ] Email notification (Welcome email, Reset password email)
- [ ] gRPC
- [ ] Casbin
- [ ] WebSocket

---

### Run locally

1. Create `.env` at src, i.e.

```sh
cp src/.env.example src/.env
```

2. Download Swag for generating docs

```sh
go get -u github.com/swaggo/swag/cmd/swag
```

3. Run

- NOTE: You have to generate swagger docs before running the app.

```sh
# Terminal 1
swag init -g src/api/app.go --output ./src/api/docs # Generates Swagger

# Terminal 2
docker-compose --env-file ./src/.env up        # docker-compose up (Run App With AutoReload)
docker-compose --env-file ./src/.env down      # docker-compose down (Shutdown App)
```

- API `http://localhost:8000/api/v1`
- Swagger Doc `http://localhost:8000/docs`

---

### Todo

- [ ] Better Input Validations
- [ ] Custom Error messages
- [ ] Data Migrations
- [ ] Logger
- [ ] Unit tests

maybe?

- [ ] SMS notification (2FA ,Reset password code)
- [ ] GraphQL
- [ ] Sentry

---

### Gotcha's

- Building Swago from source code - `go build -o swag.exe cmd/swag/main.go`

### Contribution

Open to Suggestions and Pull Requests

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
