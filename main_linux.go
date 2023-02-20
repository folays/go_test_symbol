package main

func mainPlatform(pathExe string) (gopclntab []byte, gosymtab []byte, text uint64) {
	var (
		err  error
		file *elf.File
	)

	if file, err = elf.Open(pathExe); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if gopclntab, err = file.Section(".gopclntab").Data(); err != nil {
		log.Fatalf("reading %s gopclntab: %v", file, err)
	}
	gosymtab, _ = file.Section(".gosymtab").Data()

	//text = file.Section("__text").Addr

	return
}
