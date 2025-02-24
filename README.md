# Go Web Application (In progress)

## References

This project is from the Udemy course: https://www.udemy.com/course/building-modern-web-applications-with-go

## Motivation

I bought this course to improve my Golang skills with web applications. The frontend is poor, but my goal was to improve
the backend, design patterns, architecture and tests.

## Overview

### Developed so far

I believe that the commit messages are quite explanatory (I hope so),
but here are some features of the project:

- Main.go is in cmd/web/main.go
- All middlewares are in cmd/web/middleware.go
- All routes are defined in cmd/web/routes.go
- An app wide configuration is defined in the config package
- All handlers are in the handler package which uses repository pattern to obtain access to app configuration
- All models are defined in the models package
- Render package for rendering templates
- Using GOHTML as template engine. All templates are in templates folder.
- State management with sessions
- CSRF protection.

### External Packages

- Chi package for routing.
- NoSurf package for CSRF protection.
- SCS package for session managing.