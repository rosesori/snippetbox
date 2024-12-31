# Snippetbox

Snippetbox is a web application which lets people paste and share snippets of text (like [GitHub Gists](https://gist.github.com/) or [Pastebin](https://pastebin.com/).

This project was made by following the exercises from the book ["Let's Go" by Alex Edwards](https://lets-go.alexedwards.net/).

## Development

### Run locally

To run the application locally, run the following command:

```bash
$ go run ./cmd/web
```

Add the `-help` flag to see the available command-line flags:

```bash
$ go run ./cmd/web -help

Usage of /var/folders/my/mq4mv0s15kj_2x85csjvwl_40000gn/T/go-build4245412781/b001/exe/web:
  -addr string
        HTTP network address (default ":4000")
```

### Project structure

- `cmd` contains application-specific code for the executable applications in the project
- `internal` contains the supporting, non-application-specific code used in the project
  - Any packages which live under this directory can only be imported by code inside the parent of the `internal` directory.
  In this case, any packages that live in `internal` can only be imported by code in our `snippetbox` project directory. This prevents other code bases from importing and relying on the (unversioned and unsupported) packages of the `internal` directory
- `ui` directory contains the user-interface assets used by the web applciation
  - `ui/html` contains HTML templates
  - `ui/static` contains static files (like CSS and images)

