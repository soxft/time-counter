package controller

import (
	"github.com/soxft/time-counter/config"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var RawFile []byte

func Room(c *gin.Context) {

	rawHtml := strings.ReplaceAll(string(RawFile), "{{.room}}", c.Param("room"))
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(rawHtml))

}

func ReadIndex() {
	var err error
	if RawFile, err = os.ReadFile(config.DistPath + "/room.html"); err != nil {
		log.Fatalf("cant find %s/room.html: %s", config.DistPath, err)
	}
}
