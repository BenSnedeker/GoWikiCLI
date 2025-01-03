package utils

type Flags struct {
	html       bool
	text       bool
	all        bool
	limited    bool
	summary    bool
	references bool
	wikilinks  bool
	clean      bool
	fancy      bool
}

func NewFlags() Flags {
	return Flags{
		// Raw HTML or tagless string
		html: false,
		text: true,
		// Which sections to return
		all:        false,
		limited:    true,
		summary:    false,
		references: false,
		wikilinks:  false,
		// Raw string or colored string
		clean: false,
		fancy: true,
	}
}

func UpdateTextFlags(arg string, f Flags) Flags {
	switch arg {
	case "--html", "-h":
		f.html = true
		f.text = false
	case "--text", "-t":
		f.html = false
		f.text = true
	}
	return f
}

func UpdateSectionFlags(arg string, f Flags) Flags {
	switch arg {
	case "--all", "-a":
		f.all = true
		f.limited = false
		f.summary = false
		f.references = false
		f.wikilinks = false
	case "--limited", "-l":
		f.all = false
		f.limited = true
		f.summary = false
		f.references = false
		f.wikilinks = false
	case "--summary", "-s":
		f.all = false
		f.limited = false
		f.summary = true
		f.references = false
		f.wikilinks = false
	case "--references", "-r":
		f.all = false
		f.limited = false
		f.summary = false
		f.references = true
		f.wikilinks = false
	case "--wikilinks", "-w":
		f.all = false
		f.limited = false
		f.summary = false
		f.references = false
		f.wikilinks = true
	}
	return f
}

func UpdateVisualFlags(arg string, f Flags) Flags {
	switch arg {
	case "--clean", "-c":
		f.clean = true
		f.fancy = false
	case "--fancy", "-f":
		f.clean = false
		f.fancy = true
	}
	return f
}
