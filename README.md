# gator

An RSS Feed AggreGATOR that follows feeds for multiple users.


## Requirements

1. Go
1. PostgeSQL (installed locally, not as a container)


## Installation

```bash
go install github.com/mgwinsor/gator@latest
```

## Usage

There are a total of 11 Gator commands:

To get started, first create a `gatorconfig.json` file at `$XDG_CONFIG/gator`.
This file should contain the connection string for the database as well as the
name of the current user (optional, as the user can be registered through
gator):

```json
{"db_url":"postgres://user:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"user1"}
```

To register a new user, pass the user's name to the `register` command.

```bash
gator register <user>
```

If multiple users are registered, you can login to the current user with the
`login` command. Note that this is not password protected!

```bash
gator login <user>
```

If you want to reset all data saved in the database for all users, use the `rest` command.

```bash
gator reset
```

To get a list of all users as well as which one is logged in, run the `users` command.

```bash
gator users
```

Continuously fetch new feeds at a specified interval (e.g. `1m`, `1h`, etc.)

```bash
gator agg <duration>
```

Add a feed to fetch with the `agg` command by specifying the name and URL of the feed.

```bash
gator addfeed <feed_name> <feed_url>
```

Get a list of the tracked feeds across all users:

```bash
gator feeds
```

Get a list of the feeds followed for the current user:

```bash
gator following
```

Unfollow a certain feed for the current user:

```bash
gator unfollow <feed_url>
```

View all the posts that have been fetched for the current user:

```bash
gator browse
```
