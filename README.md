[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

# Go (Golang) and Fiber REST API Boilerplate

**Used libraries:**

-   [gofiber](https://gofiber.io/)
-   [gorm](https://gorm.io/)
-   [jwt-go](https://pkg.go.dev/gopkg.in/dgrijalva/jwt-go.v3?tab=doc)
-   [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)

---

### Features

-   [x] Gofiber Docker Dev Setup with Hot Reload
-   [x] User Auth functionality (Signup, Login, Forgot Password, Reset Password)
-   [x] JWT Authentication
-   [x] REST API
-   [x] Gorm (Golang SQL DB ORM) with Postgres implementation and auto migration
-   [x] MongoDB using the official mongo driver
-   [x] Configs via environmental variables
-   [x] Email notification (Welcome email, Reset password email)
-   [x] Swagger REST API documentation
-   [x] gRPC
-   [x] WebSocket

---

### Run locally

Create `.env` at root, i.e.

```sh
cp .env.example .env
```

Run

```sh
# Terminal 1
docker-compose up        # docker-compose up (Run App With AutoReload)
docker-compose down      # docker-compose down (Shutdown App)

# Terminal 2
swag init -g api/app.go --output ./api/docs # Generates Swagger
```

-   See Swagger Doc `http://localhost:8000/docs`

---

### Todo

-   [ ] Better Input Validations
-   [ ] Custom Error messages
-   [ ] Data Migrations
-   [ ] Logger
-   [ ] Unit tests

maybe?

-   [ ] SMS notification (2FA ,Reset password code)
-   [ ] Redis
-   [ ] GraphQL
-   [ ] Sentry

---

### Contribution

Open to Suggestions and Pull Requests

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
