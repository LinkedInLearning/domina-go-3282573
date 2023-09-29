package main

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
    interrupcion := make(chan os.Signal, 1)
    signal.Notify(interrupcion, os.Interrupt)

    url := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}

    conexion, _, error := websocket.DefaultDialer.Dial(url.String(), nil)
    if error != nil {
		fmt.Println("Error al conectarse al servidor WebSocket")
		return
	}
    defer conexion.Close()

    mensaje := []byte("¡Hola, servidor WebSocket!")

    error = conexion.WriteMessage(websocket.TextMessage, mensaje)
    if error != nil {
        fmt.Println("Error al escribir el mensaje")
        return
    }
   
    _, mensajeServidor, error := conexion.ReadMessage()
    if error != nil {
        fmt.Println("Error al leer el mensaje")
        return
    }
    fmt.Printf("Mensaje recibido del Servidor: %s\n", mensajeServidor)
    
    select {
        case <-interrupcion:
            fmt.Println("Interrupción recibida, cerrando la conexión...")
            error := conexion.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if error != nil {
                fmt.Println("Error Cerrando la conexion.")
            }
            return
        case <-time.After(10 * time.Second):
            fmt.Println("Tiempo de espera agotado.")
            return
        }
}
