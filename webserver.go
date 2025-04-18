package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RunWebServer(c context.Context) {
	// Prepare list of common middleware.
	middleware := []gin.HandlerFunc{gin.Recovery()} //middlewarefilters.FilterMaliciousIntent(), middlewarelogging.AddContext(), gin.Recovery(), middlewaresecurityheaders.AddSecurityHeaders(ctx)}

	// Register middleware.
	router := gin.New()
	router.Use(middleware...)
	router.POST("/alarm/on", StartAlarmPOST)
	router.POST("/alarm/off", StopAlarmPOST)
	router.Use(FilesGET)

	httpListener, err := net.Listen("tcp", ":19916")
	if err != nil {
		log.Fatalln("Error creating listener: " + err.Error())
	}
	httpServer := &http.Server{
		Handler:           router,
		IdleTimeout:       1 * time.Minute,
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      20 * time.Second,
	}
	if err = httpServer.Serve(httpListener); err != nil {
		log.Fatalln("Error while serving HTTP: " + err.Error())
	}
}
