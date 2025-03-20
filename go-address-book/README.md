# Address Book CLI

A high-performance command-line interface (CLI) application for managing contacts, written in Go. This application provides a simple and efficient way to store and manage contact information using CSV storage with caching.



## Features

- Fast CSV storage with in-memory caching
- Buffered I/O for improved performance
- Batch processing for large datasets
- Thread-safe operations
- Contact management (Add, Update, Delete, Search)
- Test data generation
- Simple and intuitive CLI interface
- Docker support

## Technical Features

- Concurrent-safe operations with mutex locks
- Efficient memory usage with batch processing
- Buffered I/O for better performance
- In-memory caching for faster reads
- Error handling with detailed messages
- Configurable storage paths

## Requirements

For local build:
- Go 1.16 or later

For Docker:
- Docker 20.10 or later

## Installation

### Local Build

1. Clone the repository:
```bash
git clone https://github.com/rushi/address-book-cli.git
cd address-book-cli/go-address-book
```

2. Build the application:
```bash
go build -o address-book cmd/main.go
```

### Docker Build

1. Build using Docker directly:
```bash
docker build -t address-book .
```

Or using Docker Compose:
```bash
docker-compose build
```

## Usage

### Local Run
```bash
./address-book
```

### Docker Run
Using Docker directly:
```bash
docker run -it -v $(pwd)/data:/app/data address-book
```

Using Docker Compose (recommended):
```bash
docker-compose up
```

To run in detached mode:
```bash
docker-compose up -d
```

To attach to a running container:
```bash
docker attach address-book
```

To stop the container:
```bash
docker-compose down
```

The Docker setup provides:
- Data persistence through volume mounting
- Interactive terminal support
- Automatic container restart
- Network isolation
- Timezone configuration
- Container health monitoring

### Health Check
The application includes a health check mechanism that monitors:
- Application process status
- Data directory accessibility
- Container health status

To check container health status:
```bash
docker inspect --format='{{json .State.Health}}' address-book
```

### Available Commands

1. **Add Contact**
   - Add a new contact with first name, last name, email, phone, and address
   - Automatically generates unique ID and timestamps

2. **List Contacts**
   - View all contacts in your address book
   - Displays full contact details including creation and update times

3. **Search Contacts**
   - Search by first name, last name, or email
   - Case-insensitive search
   - Partial match support

4. **Update Contact**
   - Update existing contacts by ID
   - Selective field updates
   - Automatically updates modification timestamp

5. **Delete Contact**
   - Remove contacts by ID
   - Safe deletion with validation

6. **Generate Test Data**
   - Create 10 sample contacts
   - Useful for testing and demonstration

## Configuration

The application uses `config.json` for settings:

```json
{
  "csvPath": "data/contacts.csv"
}
```

The configuration file is automatically created with default values if it doesn't exist.

## Data Storage

Contacts are stored in CSV format with the following structure:
- ID
- First Name
- Last Name
- Email
- Phone
- Address
- Created At
- Updated At

### Storage Features
- Automatic directory creation
- Buffered writing for performance
- Batch processing for large datasets
- In-memory caching
- Thread-safe operations

## Project Structure

```
go-address-book/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── models/           # Data models
│   │   ├── contact.go    # Contact model
│   │   └── addressbook.go# AddressBook model
│   ├── storage/          # Storage implementation
│   │   └── storage.go    # CSV storage
│   ├── config/           # Configuration management
│   │   └── config.go     # Config handling
│   └── generator/        # Test data generation
│       └── generator.go  # Random contact generator
├── data/                 # Data storage directory
│   └── contacts.csv      # Contact database
├── go.mod               # Go module file
└── README.md            # This file
```

## Performance Optimizations

1. **Caching**
   - In-memory cache for frequently accessed data
   - Cache invalidation on updates

2. **Batch Processing**
   - Processes records in batches of 100
   - Reduces memory usage for large datasets

3. **Buffered I/O**
   - Uses buffered readers and writers
   - Reduces system calls
   - Improves I/O performance

4. **Concurrency**
   - Thread-safe operations
   - Read-write mutex for concurrent access

## Error Handling

- Detailed error messages
- Proper error propagation
- Safe error recovery
- Validation checks

## Development

To run tests:
```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 