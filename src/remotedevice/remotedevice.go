package remotedevice

import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Specs struct {
}

type LedType string

const (
	Single LedType = "single"
	Strip  LedType = "strip"
	Array  LedType = "array"
	Ring   LedType = "ring"
)

type RGBData struct {
	R int8
	G int8
	B int8
}

type LedData struct {
	Type LedType
	// Position      [2]float32
	// Specification Specs
	Data *RGBData
}

type Device struct {
	ID         string
	LedChanel  LedData
	connection *websocket.Conn
}

func (d *Device) Connect(host string) error {

	// establish connnection
	con, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		return err
	}

	err = con.WriteMessage(websocket.TextMessage, []byte("rd:"+d.ID))
	if err != nil {
		return err
	}

	_, msg, err := con.ReadMessage()

	if err != nil {
		return err
	}

	if bytes.Compare(msg, []byte("recv")) != 0 {
		return errors.New("Wrong responce")
	}

	d.connection = con

	if d.connection == nil {
		log.Panic("Connection gone at assigmnet")
	}
	// start
	return nil
}

func (d *Device) Init() error {

	switch d.LedChanel.Type {
	case Single:
		{
			d.LedChanel.Data = new(RGBData)
			*d.LedChanel.Data = RGBData{0, 0, 0}
		}
	}

	// start reading
	go d.dataStream()

	return nil
}

func (d *Device) dataStream() {

	//fill ledData with received data
	switch d.LedChanel.Type {
	case Single:
		{
			if d.connection == nil {
				log.Panic("Connection gone")
				break
			}
			for {
				_, msg, err := d.connection.ReadMessage()
				if err != nil {
					log.Print("Data loop:", err)
					break
				}
				fmt.Println("msg", msg)
				// parse message
				d.LedChanel.Data.R = int8(msg[0])
				d.LedChanel.Data.G = int8(msg[1])
				d.LedChanel.Data.B = int8(msg[2])
			}
		}
	}

}
