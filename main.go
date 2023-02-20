package main

import (
	"debug/gosym"
	"embed"
	"fmt"
	"github.com/folays/go_test_symbol-lib"
	"log"
	"os"
)

// https://www.intezer.com/blog/malware-analysis/all-your-go-binaries-are-belong-to-us/
// > [...] the debug package in the standard library, provides sub-packages to parse :
// > - ELF (Linux and Unix)
// > - PE (Windows)
// > - Mach-O (macOS)
//
// > The table PCLNTAB was added in Go 1.2 and holds data needed for Go’s panic messages.
// > The table is used to map between the program counter (location of an assembly instruction) and the source code file and line number,
//
// > Before we can process the table, we need to first find it :
// > - ELF and Mach-O files : very easy because the table is located in its own section called .gopclntab.
// > - PE files, the process is a bit less straightforward. The table is usually located in the .rdata or .text section of the PE file.
// >   The table starts with a magic value that can be used to locate the start of the table.
// >   For Go binaries compiled with 1.2 up to excluding 1.16 of the compiler, the magic value is 0xfffffffb.
// >   For files compiled with 1.16 and later the magic value is 0xfffffffa.
// >   To ensure the match is correct we can use the same checks that the parser function uses to check the table.

// https://lekstu.ga/posts/pclntab-function-recovery/

// https://stackoverflow.com/questions/42554900/how-to-extract-own-symbol-table

//go:embed go.mod
var selfFS embed.FS

func main() {
	pathExe, _ := os.Executable()

	gopclntab, gosymtab, text := mainPlatform(pathExe)

	pcln := gosym.NewLineTable(gopclntab, text)

	tab, err := gosym.NewTable(gosymtab, pcln)
	if err != nil {
		log.Fatalf("parsing %s gosymtab: %v", pathExe, err)
	}

	// [tab] has really one field which is exported, Funcs which is a []Func.

	fmt.Printf("TAB syms %d funcs %d files %d objs %d\n",
		len(tab.Syms),
		len(tab.Funcs),
		len(tab.Files),
		len(tab.Objs),
	)

	for _, fn := range tab.Funcs {
		fmt.Printf("SYM %s\n", fn.Name)
	}

	selfFS.Open("go.mod")

	go_test_symbol_lib.DebugShow()
}
