package controllers

import "go-auth/helpers"

var SecretKey = helpers.GoDotEnvVariable("SECRET_KEY")

