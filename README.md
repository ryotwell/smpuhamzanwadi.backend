# smpuhamzanwadi.backend

This is the backend for the smpuhamzanwadi project.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/your_username_/smpuhamzanwadi.backend.git
   ```
2. Install Go packages
   ```sh
   go mod tidy
   ```
3. Create a `.env` file and add the following environment variables
   ```
   CORS_ALLOWED_ORIGINS=http://localhost:3000
   DATABASE_URL=postgres://postgres:farid123@localhost:5432/tes_sdu
   ```

### Running the application

```sh
go run main.go
```

### Environment Variables

- `CORS_ALLOWED_ORIGINS` - The allowed origins for CORS
- `DATABASE_URL` - The database connection string
- `PORT` - The port to run the server on

## Built With

* [Go](https://golang.org/) - The language used
* [Gin](https://gin-gonic.com/) - The web framework used
* [GORM](https://gorm.io/) - The ORM used
* [PostgreSQL](https://www.postgresql.org/) - The database used

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
