# Gemini chat app with Go

## Install Go
To run this server, Go must be installed on your system.
Check if Go is already installed.
```
go version
```
If Go is not installed, follow the instructions for your operating system from the [official Go installation guide](https://go.dev/doc/install).

## Run the application
You need a Gemini API key to run the server,

If you don't have a Gemini API key ready, you can create a key with one click in [Google AI Studio](https://aistudio.google.com/app/apikey).

1. Navigate to the app directory, `server-go` (i.e. where main.go is located).
2. Run the application with the following command.
```
GOOGLE_API_KEY=<your_api_key> PORT=<your_port> go run .
```

## Environment Variables
* GOOGLE_API_KEY: API key for Gemini service.
* PORT: The port this server is listening on (default 9000).
