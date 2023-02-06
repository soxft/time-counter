package controller

import (
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
	if RawFile, err = os.ReadFile("dist/index2.html"); err != nil {
		log.Fatal("fuck no file", err)
	}
}
