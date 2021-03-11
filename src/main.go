package main

import (
	"log"
	"net/url"

	"rds/remotedevice"

	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	NewDevice    bool   `long:"new" description:"Device settings json configutaor"`
	SettingsJSON string `short:"j" long:"json" description:"Simulator settings json"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {

	// // parse flags
	// _, err := parser.Parse()
	// if err != nil {
	// 	// switch flagsErr := err.(type) {
	// 	// case flags.ErrorType:
	// 	// 	if flagsErr == flags.ErrHelp {
	// 	// 		os.Exit(0)
	// 	// 	}
	// 	// 	os.Exit(1)
	// 	// default:
	// 	// 	os.Exit(1)
	// 	// }
	// 	os.Exit(0)
	// }

	// //read settings json
	// configParser := viper.New()
	// configParser.SetConfigType("json")
	// configParser.SetConfigFile(options.SettingsJSON)
	// err = configParser.ReadInConfig()
	// if err != nil {
	// 	// panic(fmt.Errorf("Fatal error config file: %s \n", err))
	// }

	// ledType := configParser.Get("device.led.type")

	// fmt.Println("options.File")
	// fmt.Printf("%v\n", len(options.SettingsJSON))

	//parse settings json
	// TODO: Add wrong configuration error

	conncetionURL := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/",
	}

	simulatedDevice := remotedevice.Device{
		ID: "42",
		LedChanel: remotedevice.LedData{
			Type: remotedevice.Single,
		},
	}

	// Establish connection
	err := simulatedDevice.Connect(conncetionURL.String())
	if err != nil {
		log.Fatal("Connection err: ", err)
	}

	// Start device
	err = simulatedDevice.Init()

	//Config window
	raylib.SetConfigFlags(raylib.FlagWindowResizable)

	raylib.InitWindow(800, 400, "Test window")
	raylib.SetWindowMinSize(100, 100)
	raylib.SetTargetFPS(60)

	// Draw window
	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		if err == nil {

			// bc := raylib.Color{
			// 	R: uint8(simulatedDevice.LedChanel.Data.R),
			// 	G: uint8(simulatedDevice.LedChanel.Data.G),
			// 	B: uint8(simulatedDevice.LedChanel.Data.B),
			// 	A: ,
			// }

			raylib.ClearBackground(raylib.Color{
				R: uint8(simulatedDevice.LedChanel.Data.R),
				G: uint8(simulatedDevice.LedChanel.Data.G),
				B: uint8(simulatedDevice.LedChanel.Data.B),
				A: uint8(255),
			})
			// switch ledType {
			// case remotedevice.Single:
			// 	{

			// 	}
			// }
		} else { // Display error message
			raylib.ClearBackground(raylib.DarkGray)
			textSize := 42
			raylib.DrawText(err.Error(),
				int32((raylib.GetScreenWidth()-450)/2),
				int32((raylib.GetScreenHeight()-textSize)/2),
				int32(textSize),
				raylib.RayWhite,
			)
		}
		raylib.EndDrawing()
	}

}
