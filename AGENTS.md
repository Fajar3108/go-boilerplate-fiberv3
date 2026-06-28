# Tech Stack
- Language: Go (Golang)
- Framework: Fiber v3
- Database: Neon Postgre
- ORM: GORM
- Config: Viper (reading from .env)
- Authentication: JWT (golang-jwt/v5) with Refresh Token Rotation
- Validation: Go Playground Validator v10
- Email: Gomail v2
- Docker: make sure build this project on docker and run this project on docker

# Projects Structure
.
├── cmd/
│   ├── api/main.go         # API application entry point
├── config/                 # Viper setup & .env constants
├── database/               # Thread-safe GORM singleton initializer
├── internal/
│   ├── action/             # Atomic, reusable business logic (e.g., CreateNewUser)
│   ├── controller/         # HTTP (Fiber) layer. Only part that knows about HTTP.
│   ├── request/            # Structs for request body validation.
│   ├── resource/           # Structs for response transformation (DTOs).
│   ├── router/             # Fiber route definitions (Auth, Category, etc).
│   └── service/            # Business logic orchestrators (e.g., AuthService).
├── pkg/
│   ├── error-handler/      # Custom error definitions & global error handler.
│   ├── file-storage/       # Framework-agnostic file upload logic.
│   ├── helpers/            # Utility functions (UUID, slug, response).
│   ├── mail/               # Email sending helper.
│   ├── middleware/         # Custom middleware (e.g., JWT).
│   ├── model/              # GORM data entities (User, Category, etc).
│   ├── token/              # JWT generation & parsing.
│   └── validation/         # Thread-safe Validator singleton.
├── storage/                # (gitignored) Uploaded files are stored here.
├── .env.example            # Configuration file template.
├── go.mod                  # Go dependencies.
└── makefile                # Helper commands (serve, migrate).

# Go AI Agent Directives

You are an expert Go (Golang) developer. Your code must be simple, readable, and highly idiomatic. You prioritize the Go standard library over third-party packages, and you write explicit, robust code rather than "clever" abstractions.

## 1. Project Architecture & Layout
* **Follow standard Go layout:** Place entry points in `cmd/`, private business logic in `internal/`, and highly reusable components in `pkg/`.
* **Accept Interfaces, Return Structs:** Consumers should define the interfaces they need. Do not proactively create broad, Java-style interfaces unless there are multiple concrete implementations immediately required.
* **Dependency Injection:** Pass dependencies (like database connections or loggers) explicitly into constructors. Do not use global state or `init()` functions for application logic.

## 2. Code Style & Idioms
* **Formatting:** All generated code must be perfectly formatted according to `gofmt` / `goimports`.
* **Pointers:** Pass by value by default. Only pass by pointer when you need to mutate the receiver or when the struct is massive. Do not use pointers for basic types (strings, ints) unless mapping to a nullable database column.
* **Naming:** Use short, concise variable names (e.g., `req` instead of `request`, `b` instead of `buffer`). Acronyms should be all caps (e.g., `UserID`, not `UserId`).

## 3. Error Handling (Strict)
* **Never swallow errors:** Every error must be checked and handled immediately. 
* **Wrap errors with context:** Use `fmt.Errorf("failed to fetch user %d: %w", id, err)`. Do not just return the raw `err`.
* **No Panics:** Never use `panic` or `log.Fatal` in library/internal code. Only use them during the initial boot sequence in `main.go`.
* **Sentinel Errors:** Define sentinel errors at the top of the package (e.g., `var ErrNotFound = errors.New("not found")`) for errors that consumers need to check via `errors.Is`.

## 4. Concurrency
* **Keep it simple:** Do not use goroutines or channels unless asynchronous execution is explicitly required. Synchronous code is easier to read and debug.
* **No Leaks:** Every goroutine must have a clear exit path. Always use `context.Context` to manage cancellation and timeouts across boundaries.
* **Synchronization:** Prefer `sync.Mutex` for simple state protection, and use channels for passing ownership of data.

## 5. Testing
* **Table-Driven Tests:** Write all tests using the table-driven test pattern. 
* **Location:** Place tests alongside the code they test (e.g., `user_test.go` next to `user.go`).
* **Assertions:** Use the standard `testing` package. You may use a lightweight assertion library like `testify/require` only if it is already present in the `go.mod`.

## 6. AI Operational Rules
* Before writing code, review `go.mod` to understand existing dependencies.
* Do not invent imaginary libraries or functions. 
* If a request requires complex domain knowledge, outline your proposed architecture in comments or markdown before writing the implementation.
* If a step in a multi-step task fails (e.g., compiler error), stop and report the error immediately. Do not blindly proceed to the next step.

## ❌ Do Not
- Do not hardcode credentials in any `.go` file.
- Do not output the contents of the `.env` file or the user's connection string in any response.
- Do not use any deprecated or alternative Go Postgres drivers like `lib/pq`.