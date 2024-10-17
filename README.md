# Music Library Service REST API

## Overview

This service processes API requests for song information, enriches the response with probable release date, lyrics, and link using external APIs, and stores the data in a PostgreSQL database. It exposes various RESTful methods for managing music records.

## Project Summary

The Music Library Service REST API is designed to provide a seamless interface for handling song information. It utilizes external APIs to enhance the data with details such as release date, lyrics, and link. The enriched data is stored in a PostgreSQL database, and the service offers RESTful methods for CRUD operations on people records. The project emphasizes clean code, efficient filtering, and pagination, along with comprehensive logging for monitoring and debugging.

## Configs

Before Usage please fill out configs in .env file with your data.

## Usage

* To run this app:
``` 
go run ./cmd/web

```

## REST Methods

1. **Get Songs**
   - Endpoint: `/music`
   - Method: `GET`
   - Retrieves data with filters and pagination.
   - Filtering queries: "group", "song", "releasedate", "text", "link".
   - Pagination queries: "page", "pageSize".

2. **Get Song couplets by ID**
   - Endpoint: `/music/{id}`
   - Method: `GET`
   - Retrieves information for a specific person.
   - Pagination queries: "page", "pageSize".

3. **Add Song**
   - Endpoint: `/music`
   - Method: `POST`
   - Adds a new person to the database.

4. **Update Song**
   - Endpoint: `/music/{id}`
   - Method: `PUT`
   - Modifies information for a specific person.

5. **Delete Song**
   - Endpoint: `/music/{id}`
   - Method: `DELETE`
   - Removes a person from the database.

Enriched data is stored in a PostgreSQL database, and the database structure is created through migrations.

## Logging

The code is extensively covered with debug and info logs to facilitate troubleshooting and monitoring.

## Configuration

Sensitive configuration data is stored in a `.env` file.

## Swagger

Swagger was used to automate the documentation of the RESTful API for this project.

- Endpoint: `/swagger/index.html`