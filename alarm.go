package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cloudkucooland/go-kasa"
	"github.com/gin-gonic/gin"
	"github.com/szatmary/sonos"
)

func StartAlarmPOST(c *gin.Context) {
	// Start alarm
	err := StartAlarm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Alarm started"})
}

func StartAlarm() error {
	// Turn on light
	device, err := kasa.NewDevice(os.Getenv("TPLINK_KASA_IP"))
	if err != nil {
		return err
	}
	err = device.SetRelayState(true)
	if err != nil {
		return err
	}

	son, err := sonos.NewSonos()
	if err != nil {
		return err
	}
	defer son.Close()

	found, _ := son.Search()
	to := time.After(10 * time.Second)
	for {
		select {
		case <-to:
			return nil
		case zp := <-found:
			fmt.Printf("%s\t%s\t%s", zp.RoomName(), zp.ModelName(), zp.SerialNum())
		}
	}
	return nil

	// Play song
	zp, err := sonos.FindRoom(os.Getenv("SONOS_ROOM"), 20*time.Second)
	if err != nil {
		return err
	}
	if err = zp.SetAVTransportURI(os.Getenv("SONOS_TRANSPORT")); err != nil {
		return err
	}
	if err = zp.Play(); err != nil {
		return err
	}

	return nil
}

func StopAlarmPOST(c *gin.Context) {
	// Stop alarm
	err := StopAlarm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Alarm stopped"})
}

func StopAlarm() error {
	device, err := kasa.NewDevice(os.Getenv("TPLINK_KASA_IP"))
	if err != nil {
		return err
	}

	err = device.SetRelayState(false)
	if err != nil {
		return err
	}

	return nil
}
