package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"memrizr/handler"
	"memrizr/repository"
	"memrizr/service"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// your imports here

// will initialize a handler starting from data sources
// which inject into repository layer
// which inject into service layer
// which inject into handler layer
func inject(d *dataSources) (*handler.Server, error) {
	log.Println("Injecting data sources")

	/*
	 * repository layer
	 */
	userRepository := repository.NewUserRepository(d.DB)

	/*
	 * repository layer
	 */
	userService := service.NewUserService(userRepository)

	// load rsa keys
	privKeyFile := os.Getenv("PRIV_KEY_FILE")
	priv, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %w", err)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}
	pubKeyFile := os.Getenv("PUB_KEY_FILE")
	pub, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read public key pem file: %w", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	// load refresh token secret from env variable
	refreshSecret := os.Getenv("REFRESH_SECRET")

	tokenService := service.NewTokenService(privKey, pubKey, refreshSecret)

	server := handler.NewServer(userService, tokenService)

	return server, nil
}
