package main

import (
	"fmt"
	"github.com/nevadex/autogcf/templates"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func Generate() error {
	for i := range RootFiles {
		ext := strings.Split(RootFiles[i], ".")
		ma := MaxAgeOther
		if strings.ToLower(ext[len(ext)-1]) == "html" {
			ma = MaxAgeHTML
		}
		gf := templates.GenSingleFile(RootFiles[i], mime.TypeByExtension("."+ext[len(ext)-1]), "max-age="+ma)
		gf.NoFormat = !ReadableFunctions
		f, err := os.Create(filepath.Join(RootFiles[i], "function.go"))
		if err != nil {
			return err
		}
		err = gf.Render(f)
		if err != nil {
			return err
		}
	}
	fmt.Println("Generated all single-file functions")

	for i := range RootDirs {
		var p, lp, ct, cc string
		for x := range DirFiles[RootDirs[i]] {
			s := DirFiles[RootDirs[i]][x][len(RootDirs[i])+1:]
			p += fmt.Sprintf("\"%v\",", strings.ReplaceAll(s, string(os.PathSeparator), "/"))
			lp += fmt.Sprintf("\"%v\",", strings.ReplaceAll(s, string(os.PathSeparator), "_"))
			ext := strings.Split(s, ".")
			ct += fmt.Sprintf("\"%v\",", mime.TypeByExtension("."+ext[len(ext)-1]))
			ma := MaxAgeOther
			if strings.ToLower(ext[len(ext)-1]) == "html" {
				ma = MaxAgeHTML
			}
			cc += fmt.Sprintf("\"%v\",", "max-age="+ma)
		}

		gf := templates.GenMultiFile(RootDirs[i], p, lp, ct, cc)
		gf.NoFormat = !ReadableFunctions
		f, err := os.Create(filepath.Join(RootDirs[i], "function.go"))
		if err != nil {
			return err
		}
		err = gf.Render(f)
		if err != nil {
			return err
		}
	}
	fmt.Println("Generated all multi-file functions")

	return WaitForConfirm("Start gcloud deployment?")
}
