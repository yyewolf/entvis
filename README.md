# entvis

An [ent](https://entgo.io/) extension to easily implement field visibility for serving REST APIs and controlling data access based on user roles.

[![Go Reference](https://pkg.go.dev/badge/github.com/yyewolf/entvis.svg)](https://pkg.go.dev/github.com/yyewolf/entvis)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyewolf/entvis)](https://goreportcard.com/report/github.com/yyewolf/entvis)
[![License: MIT](https://img.shields.io/badge/License-BSD-3-yellow.svg)](https://opensource.org/license/bsd-3-clause)

## Overview

`entvis` is an extension for the [ent ORM](https://entgo.io/) that provides a simple way to control field visibility in your API responses based on user roles. This makes it easy to implement role-based access control for your REST API data without having to write custom serialization logic for each entity.

The extension generates helper methods that filter entity fields according to the visibility annotations you define, ensuring sensitive data is only exposed to authorized roles.
