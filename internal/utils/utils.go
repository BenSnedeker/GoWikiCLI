package utils

import "github.com/BenSnedeker/GoWiki/gowiki"

func UpdateTextFlags(arg string, f gowiki.Flags) gowiki.Flags {
	switch arg {
	case "--html", "-h":
		f.Html = true
		f.Text = false
	case "--text", "-t":
		f.Html = false
		f.Text = true
	}
	return f
}

func UpdateSectionFlags(arg string, f gowiki.Flags) gowiki.Flags {
	switch arg {
	case "--all", "-a":
		f.All = true
		f.Limited = false
		f.Summary = false
		f.References = false
		f.Wikilinks = false
	case "--limited", "-l":
		f.All = false
		f.Limited = true
		f.Summary = false
		f.References = false
		f.Wikilinks = false
	case "--summary", "-s":
		f.All = false
		f.Limited = false
		f.Summary = true
		f.References = false
		f.Wikilinks = false
	case "--references", "-r":
		f.All = false
		f.Limited = false
		f.Summary = false
		f.References = true
		f.Wikilinks = false
	case "--wikilinks", "-w":
		f.All = false
		f.Limited = false
		f.Summary = false
		f.References = false
		f.Wikilinks = true
	}
	return f
}

func UpdateVisualFlags(arg string, f gowiki.Flags) gowiki.Flags {
	switch arg {
	case "--clean", "-c":
		f.Clean = true
		f.Fancy = false
	case "--fancy", "-f":
		f.Clean = false
		f.Fancy = true
	}
	return f
}
