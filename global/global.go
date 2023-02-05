package global

import (
	"log"
	"sort"
)

var (
	RoomPrefixList []string
	roomChan       chan string
)

func init() {
	roomChan = make(chan string, 2)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("global.init() error: " + err.(string))
			}
		}()

		for _room := range roomChan {
			if inArray(_room, RoomPrefixList) {
				continue
			}

			RoomPrefixList = append(RoomPrefixList, _room)
			log.Println("global.init() append room: " + _room)
		}
	}()
}

func InsertRoomPrefix(room string) {
	roomChan <- room
}

func inArray(str string, strArray []string) bool {
	sort.Strings(strArray)
	i := sort.SearchStrings(strArray, str)
	if i < len(strArray) && strArray[i] == str {
		return true
	} else {
		return false
	}
}
