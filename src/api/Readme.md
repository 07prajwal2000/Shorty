# Shorty - The URL Shortener
## WORK IN PROGRESS ⚠️ - OPEN TO CONTRIBUTIONS
This is a api server project written in Golang

# Requirements
- Go compiler
- Nodemon (Optional - used for live reloading)
    - You can install by using `npm i -g nodemon`

# How to Run
- Step 1: Install Dependencies `go mod download`
- Step 2: `go run .` or using Nodemon `./hotreload.bat`
- Step 3: Open the `controllers` folder to find all the defined routes

# Technologies used
- Echo (Go HTTP Router)
- Database (maybe Postgres 🤔)
# Routes
- DOMAIN/api/v1/urls/generate = POST
    - ```json
        {
            "originalUrl": string
        }
        ```
- DOMAIN/<URL_ID> = GET