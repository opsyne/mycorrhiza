// Package interwiki provides interwiki capabilities. Most of them, at least.
package interwiki

import (
	"encoding/json"
	"github.com/bouncepaw/mycomarkup/v5/options"
	"github.com/bouncepaw/mycorrhiza/files"
	"log"
	"os"
)

func Init() {
	var (
		record, err = readInterwiki()
	)
	if err != nil {
		log.Fatalln(err)
	}
	for _, wiki := range record {
		wiki := wiki // This line is required
		wiki.canonize()
		theMap.list = append(theMap.list, &wiki)
		for _, name := range append(wiki.Aliases, wiki.Name) {
			if _, found := theMap.byName[name]; found {
				log.Fatalf("There are multiple uses of the same name ‘%s’\n", name)
			} else {
				theMap.byName[name] = &wiki
			}
		}
	}
	log.Printf("Loaded %d interwiki entries\n", len(theMap.list))
}

func HrefLinkFormatFor(prefix string) (string, options.InterwikiError) {
	if wiki, ok := theMap.byName[prefix]; ok {
		return wiki.LinkHrefFormat, options.Ok
	}
	return "", options.UnknownPrefix
}

func ImgSrcFormatFor(prefix string) (string, options.InterwikiError) {
	if wiki, ok := theMap.byName[prefix]; ok {
		return wiki.ImgSrcFormat, options.Ok
	}
	return "", options.UnknownPrefix
}

func readInterwiki() ([]Wiki, error) {
	var (
		record            []Wiki
		fileContents, err = os.ReadFile(files.InterwikiJSON())
	)
	if os.IsNotExist(err) {
		return record, nil
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileContents, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}