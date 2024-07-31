# Finance
Basic finance web app

## Issues

- [ ] Create the home page
- [ ] Add pico CSS  
- [ ] Configure a local libSQL database  

## Environment

- Instructions to setup the development environment
- Information about the production environment

### Development

Follow the instructions to setup the development environment

#### Go Version

See the file `go.mod` at the project root.

#### Database

By default, the program will open a local connection with a local.db file in the project root.

#### Environment Variables

Create a copy of `.env-example` and rename it to `.env`. Fill the empty values if needed.

#### Run

To start the server use the following command:

```
go run ./src
```

## Production

Not deployed yet