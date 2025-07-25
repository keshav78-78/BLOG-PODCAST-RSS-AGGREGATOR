📰 Blog & Podcast RSS Aggregator

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/PostgreSQL-14-336791?style=for-the-badge&logo=postgresql" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/sqlc-v2-blue?style=for-the-badge" alt="SQLC">
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge" alt="License: MIT">
</p>

A backend service built with Go for an RSS aggregator. It allows users to subscribe to their favorite blogs and podcasts, fetching new content automatically.

---

✨ Features

👤 User Management: Simple user creation with a unique API key for secure access.

✍️ Feed Management: Authenticated users can add and manage their list of tracked RSS feeds.

🤝 Feed Following: Users can follow and unfollow feeds to customize their content stream.

📖 Post Fetching: Retrieve the latest posts from all followed feeds through a single API call.

⚙️ Background Scraper: A concurrent, time-based scraper runs in the background to keep content up-to-date.

🌐 RESTful API: A clean, versioned API for all client-side interactions.

🗃️ Database Migrations: SQL-based schema migrations for robust and easy database setup.

🛠️ Tech Stack

Language: [Go](https://golang.org/)
Database: [PostgreSQL](https://www.postgresql.org/)
Router: [Chi](https://github.com/go-chi/chi)
Database Driver: [pq](https://github.com/lib/pq)
Query Builder: [sqlc](https://sqlc.dev/) for generating type-safe Go code from raw SQL.
Environment: [godotenv](https://github.com/joho/godotenv) for managing configuration.
UUIDs: [google/uuid](https://github.com/google/uuid) for unique identifiers.

🚀 Getting Started

Follow these instructions to get a local copy up and running for development and testing.

Prerequisites

You'll need the following tools installed on your machine:
Go (version 1.20 or later)
PostgreSQL
[sqlc](https://sqlc.dev/en/latest/overview/install.html)

Installation & Setup

1.  Clone the Repository
    ```sh
    git clone <your-repository-url>
    cd blog-podcast-rss-aggregator
    ```

2.  Install Dependencies
    ```sh
    go mod tidy
    ```

3.  Configure Environment
    Create a `.env` file in the project root. You can copy the example below:
    ```env
    PORT=8080
    DATABASE_URL="postgres://user:password@localhost:5432/rss-aggregator?sslmode=disable"
    ```

4.  Set Up the Database
    Ensure your PostgreSQL server is running.
    Create a new database (e.g., `rss-aggregator`).
    Run the database migrations using a tool like [goose](https://github.com/pressly/goose):
        ```sh
        # Note: You may need to install goose first
        goose -dir sql/schema postgres "$DATABASE_URL" up
        ```

5.  Generate Go Code from SQL
    If you modify any SQL queries in the `sql/queries/` directory, regenerate the Go code:
    ```sh
    sqlc generate
    ```

▶️ Running the Application

Start the server with the following command:

```sh
go run .
The server will be live at http://localhost:8080 (or the port you specified).🔌 API EndpointsAll endpoints are prefixed with /v1.MethodEndpointDescriptionAuthenticationGET/readinessChecks if the service is up and running.NoneGET/errReturns an error for testing purposes.NonePOST/usersCreates a new user and returns an API key.NoneGET/usersGets the current user's details.API KeyPOST/feedsCreates a new feed.API KeyGET/feedsGets all feeds in the system.NonePOST/feed_followsFollows a feed.API KeyGET/feed_followsGets all feed follows for the user.API KeyDELETE/feed_follows/{feedFollowID}Unfollows a specific feed.API KeyGET/postsGets the latest posts for the user.API KeyAuthentication: For protected endpoints, provide your API key in the Authorization header.Authorization: ApiKey YOUR_API_KEY_HERE
📂 Project Structure.
├── sql/
│   ├── queries/      # SQL queries for sqlc to generate Go code
│   └── schema/       # Database migration files (in order)
├── internal/
│   ├── auth/         # Authentication helpers
│   └── database/     # sqlc-generated code and DB models
├── vendor/           # Go module dependencies
├── main.go           # Main application entry point
├── scraper.go        # Logic for the background RSS scraper
├── go.mod            # Go module definitions
└── sqlc.yaml         # Configuration for sqlc
