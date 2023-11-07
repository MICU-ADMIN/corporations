```markdown
# Prisma CLI with Go

The Go client for Prisma works slightly differently from the standard Prisma tooling. When using the Go client, you should replace Prisma CLI commands like `prisma ...` with `go run github.com/steebchen/prisma-client-go ...` instead.

For example:

- To re-generate the Go client:
```

go run github.com/steebchen/prisma-client-go generate

```

- To sync the database with your schema for development:
```

go run github.com/steebchen/prisma-client-go db push

```

- To create a Prisma schema from your existing database:
```

go run github.com/steebchen/prisma-client-go db pull

```

- For production use, create a migration locally:
```

go run github.com/steebchen/prisma-client-go migrate dev

```

- To sync your production database with your migrations:
```

go run github.com/steebchen/prisma-client-go migrate deploy

````

Shortcut:
If you primarily work with the Go client and don't have or don't want the NodeJS Prisma CLI installed, you can set up an alias. To do this, edit your `~/.bashrc` or `~/.zshrc` and add the following alias:

```bash
alias prisma="go run github.com/steebchen/prisma-client-go"
````

With this alias, you can use `prisma` commands as usual, and they will invoke the real locally bundled Prisma CLI under the hood.

_Last updated on September 3, 2023_

```

This Markdown format summarizes the provided information about using Prisma CLI with Go and setting up a shortcut using an alias.
```
