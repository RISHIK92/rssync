# rssync

A simple Go application to fetch RSS feed items, track the last updated date, and send email notifications for new items.

## Features

- Fetches and parses RSS feeds
- Tracks the last updated date to avoid duplicate notifications
- Sends email notifications for new RSS items

## Architecture Overview

The application is organized into a modular structure for clarity and maintainability:

### 1. Entry Point (`cmd/main.go`)

- Orchestrates the workflow: fetches RSS, checks for new items, and triggers email notifications.
- Calls the RSS parser and mailer modules.

### 2. RSS Parsing (`internals/rss/`)

- `parser.go` handles fetching and parsing RSS feeds using Go's `net/http` and `encoding/xml` packages.
- Filters items based on the last updated date (from state management).
- Returns only new items for further processing.

### 3. State Management (`internals/state/`)

- `state.go` manages reading and updating the last processed date.
- Uses a simple file (`last_updated.txt`) to persist the last update timestamp between runs.
- Ensures only new RSS items are processed and notified.

### 4. Email Notification (`internals/mailer/`)

- `smtp.go` contains logic to send emails using SMTP.
- Sends an email for each new RSS item, with details like title, link, publication date, and description.
- Requires SMTP credentials to be configured by the user.

### Data Flow

1. **Start**: `main.go` is executed.
2. **Fetch**: RSS feed is fetched and parsed by the RSS module.
3. **Filter**: Only items newer than the last update are selected.
4. **Notify**: Each new item triggers an email via the mailer module.
5. **Update State**: The latest publication date is saved for the next run.

### Error Handling

- Network, parsing, and file I/O errors are handled gracefully with error messages.
- Email sending errors are logged but do not halt the entire process.

## Project Structure

```
cmd/
  main.go            # Entry point of the application
internals/
  mailer/            # Email sending logic
  rss/               # RSS parsing logic
  state/             # State management (last updated date)
go.mod, go.sum       # Go module files
last_updated.txt     # Stores the last updated date
```

## Usage

1. **Clone the repository:**
   ```sh
   git clone <repo-url>
   cd rssync
   ```
2. **Configure email sending:**
   - Update the email logic in `internals/mailer/smtp.go` with your SMTP credentials.
3. **Run the application:**
   ```sh
   go run ./cmd/main.go
   ```

## How it works

- The app fetches the RSS feed from a specified URL.
- It parses the feed and checks for new items since the last update (tracked in `last_updated.txt`).
- For each new item, it sends an email notification.
- The last updated date is updated after processing.

## Requirements

- Go 1.18 or higher
- Internet connection
- SMTP credentials for sending emails
