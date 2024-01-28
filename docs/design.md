# siTel

## Introduction
Simple Telemetry is an open source educational platform designed to help users learn about observability principles and practices in microservice architecture. The application aims to provide a simple yet comprehensive environment for users to understand key concepts and techniques related to monitoring, logging, tracing, and troubleshooting in distributed systems.

The application itself is a simple todo list solution that comprises so far of: 
- A stripped back registration, login and session management system 

and aims to comprise of:
- A todo list system for creating, editing and deleting simple todo lists
- A user profile system for creating and editing user profile details
- A simulated stripped back subscription system

More features can be put forward and added as the application grows over time. 

An essential aspect of the design is to ensure that the application is intentionally simplified, allowing users to concentrate on implementing basic observability principles to enhance their understanding and proficiency in the subject matter.

## Architecture Overview
### todo-frontend
A front end built with NextJS

### todo-iam
IAM REST API writen in Go with Echo

Handles registration and login backend fuctionality

Uses MongoDB `users` collection for the user account datastore

### todo-session
Session management REST API written in NodeJS with Express

Handles session management creating and validating user session with session IDs

Uses redis for storing session IDs with TTL configuration

### MongoDB
Datastore for user accounts

Used by the **todo-iam** service

### redis
Cache for session IDs

Used by the **todo-session** service

## Notes
Simple Telemetry is still **work in progress** and contribution to the project is welcome

More detailed documentation can be found in the `docs` folder as it is created