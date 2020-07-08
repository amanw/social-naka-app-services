### Build binary from official Go image
FROM golang:stretch as build
COPY . /app
WORKDIR /app
RUN go build -o /social-naka-app-services .

### Put the binary onto Heroku image
FROM heroku/heroku:16
COPY --from=build /social-naka-app-services /social-naka-app-services

# # Expose port to the outside world
# EXPOSE 8080

CMD ["/social-naka-app-services"]