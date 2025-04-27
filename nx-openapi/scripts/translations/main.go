package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/amphipath/msea-scouter-utils/translations"
)

//go:embed dictionary.json
var rawDict []byte

//go:embed library.json
var rawLib []byte

func loadDict() *translations.PossibleValues {
	p := translations.PossibleValues{}
	json.Unmarshal(rawDict, &p)
	return &p
}

func loadLibrary() *translations.TranslationLibrary {
	l := translations.TranslationLibrary{}
	json.Unmarshal(rawDict, &l)
	return &l
}

func main() {
	dict := loadDict()

	lib := loadLibrary()

	for k, v := range dict.CharacterClass {
		if len(k) > 0 {
			key := fmt.Sprintf("job_%s", k)
			lib.AddKeyIfAbsent(translations.CategoryJob, key, translations.LanguageMSEA, k)
			lib.AddKeyIfAbsent(translations.CategoryJob, key, translations.LanguageKMS, v)
		}
	}

	b, _ := json.MarshalIndent(lib, "", "  ")
	os.WriteFile("./output.json", b, 0644)
}
