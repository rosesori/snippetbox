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
  - Any packages which live under this directory can only be imported by code inside the parent of the `internal` directory.
  In this case, any packages that live in `internal` can only be imported by code in our `snippetbox` project directory. This prevents other code bases from importing and relying on the (unversioned and unsupported) packages of the `internal` directory
- `ui` directory contains the user-interface assets used by the web applciation
  - `ui/html` contains HTML templates
  - `ui/static` contains static files (like CSS and images)
