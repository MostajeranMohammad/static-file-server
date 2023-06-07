## Running App

Run in dev mode:

    air

Build app:

    go build -o ./bin/app ./cmd

Generate Swagger Open Api specs:

    swag  init  -g  internal/application/app.go