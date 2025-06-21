# gator

---

## Overview

Is a guided project from [Boot.dev](https://www.boot.dev/) < https://www.boot.dev/ >.

The project is building an RSS feed aggregator (hence the name 'gator') CLI tool.

So it goes out to an RSS feed and then collects posts and saves them in a Postgres database.
This can be used to follow news feeds.

Since it needs a database you'll need to install Postgres to use this tool.

This project is written in Go so you'll also need 'go install' to install and run the program.
The ['go install' docs are at](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) < https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies >.

---

## Commands

- "register"
  - parameters: <name>
  - description: Registers a user by adding them to the users table in the database.
- "login"
  - parameters: <name>
  - description: Logs a user in.
- "reset"
  - parameters:
  - description: Removes all data from all of the tables.
- "users"
  - parameters:
  - description: Provides a list of all users (think your basic RESTful 'index' method).
- "agg"
  - parameters: <time_between_reqs>
  - description: A long-lived command that fetches posts from an RSS feed at a provided frequency rate and saves them to the database.
- "feeds"
  - parameters:
  - description: Provides a list of all feeds (think your basic RESTful 'index' method)
- "addfeed"
  - parameters: <feed_name> <url>
  - description: Adds a feed to the database that can then have the "agg" command fetch posts from.
- "follow"
  - parameters: <url>
  - description: Allows different users to follow feeds that other users created (keeps us from having reduent feeds throughout the database, while still providing users the ability to check out other feeds.)
- "following"
  - parameters:
  - description: Provides a list of the RSS feeds that the current user is 'following' (I endorse grud brain naming conventions. Keep your 'smarts' to you yourself pls.). Also setting up this functionality was great practice at setting up a joining table to handle many-to-many relationships in the database.
- "unfollow"
  - parameters: <url>
  - description: Allows a user to stop following a specific RSS feed
- "browse"
  - parameters: <limit> (optional)
  - description: Lists out posts that a user has saved to the database. So one can look at them and then decided if they want to go read the post or not.

---

## Usage

Install the program, run it with the commands above and do the things that it can do.

---
