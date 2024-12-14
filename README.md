# Snippetbox

This repository follows the exercises from the book "Let's Go" by Alex Edwards
to create a SnippetBox application.

## Development

### Run locally

To run the application locally, run the following command:

```bash
go run .
```

### Project structure

- `cmd` contains application-specific code for the executable applications in the project
- `internal` contains the supporting, non-application-specific code used in the project
- `ui` directory contains the user-interface assets used by the web applciation
  - `ui/html` contains HTML templates
  - `ui/static` contains static files (like CSS and images)
