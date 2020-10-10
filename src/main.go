package main

import (
	"github.com/gufranmirza/provider-microservice/cmd"
)

// @title API Documentation
// @version 2.0
// @description API Documentation

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /provider-api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cmd.Execute()
}
