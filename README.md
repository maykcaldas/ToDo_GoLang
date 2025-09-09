# GoToDo

Just messing around with Go.
README was AI generated lol

## Features

- **Add Tasks**: Create new tasks with priority levels (low, medium, high)
- **Remove Tasks**: Delete tasks you've completed
- **List Tasks**: View all your current tasks
- **Persistent Storage**: Tasks are saved in JSON format

## Project Structure

```text
.
├── cmd/todo/           # Main application entry point
├── internal/           # Private application packages
│   └── utils/          # JSON handling and task management utilities
├── data/               # Storage directory for task data
├── bin/                # Compiled binaries
├── tests/              # Test files
```

## Getting Started

### Prerequisites

- Go 1.25.1 or higher

### Running the Application

```bash
go run cmd/todo/main.go
```

### Building the Application

```bash
go build -o bin/todo cmd/todo/main.go
```

Or use the Makefile:

```bash
make build
```

### Using the Application

Once running, the application will prompt you with command options:

1. **add** - Add a new task
   - Enter the task description
   - Enter priority (low/medium/high)

2. **remove** - Remove a task
   - Enter the task to remove

3. **list** - View all tasks
   - Shows all tasks with their priorities

## Development

### Project Layout

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):

- `/cmd` - Main applications for this project
- `/internal` - Private application and library code
- `/pkg` - Library code that's safe to use by external applications
- `/tests` - Additional external test apps and test data
- `/data` - Application data storage

### Data Structure

Tasks are stored in a JSON file with the following structure:

```json
{
  "task1": "high",
  "task2": "medium",
  "task3": "low"
}
```

## Makefile Commands

- `make build` - Build the application
- `make run` - Build and run the application
- `make test` - Run tests
- `make clean` - Clean build artifacts
