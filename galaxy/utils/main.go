package galaxy

import (

)

var data []byte

var upgrader = websocket.upgrader {

	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}