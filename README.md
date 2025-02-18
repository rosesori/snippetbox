# Snippetbox

Snippetbox is a web application which lets someone paste and share snippets of text like [GitHub Gists](https://gist.github.com/) or [Pastebin](https://pastebin.com/).

This project was made by following the exercises from the book ["Let's Go" by Alex Edwards](https://lets-go.alexedwards.net/). Some comments taken from the book for educational documentation.

## Development

### Run locally

To run the application locally, run the following command:

```bash
$ go run ./cmd/web

time=2025-01-24T23:32:53.265-08:00 level=INFO source=/Users/rosesoriano/Code/snippetbox/cmd/web/main.go:59 msg="starting server" addr=:4000
...
```

You can add the `-help` flag to see the available command-line flags:

```bash
$ go run ./cmd/web -help

Usage of /var/folders/my/mq4mv0s15kj_2x85csjvwl_40000gn/T/go-build4245412781/b001/exe/web:
  -addr string
        HTTP network address (default ":4000")
```

To launch the MySQL Command-Line Client, run the following, and enter your password:

```bash
$ mysql -D snippetbox -u web -p

Enter password:
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 17
Server version: 9.1.0 Homebrew

Copyright (c) 2000, 2024, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
```

### Project structure

- `cmd` contains application-specific code for the executable applications in the project
- `internal` contains the supporting, non-application-specific code used in the project
  - Any packages which live under this directory can only be imported by code inside the parent of the `internal` directory.
  In this case, any packages that live in `internal` can only be imported by code in our `snippetbox` project directory. This prevents other code bases from importing and relying on the (unversioned and unsupported) packages of the `internal` directory
- `ui` directory contains the user-interface assets used by the web application
  - `ui/html` contains HTML templates
  - `ui/static` contains static files (like CSS and images)
