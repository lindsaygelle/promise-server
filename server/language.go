package server

import (
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/language"
)

// LanguageTags returns a slice of language.Tag.
func LanguageTags(c database.Client) ([]language.Tag, error) {
	return language.GetTags(c)
}

// Languages returns a slice of language.Language.
func Languages(c database.Client) ([]language.Language, error) {
	return language.GetLanguages(c)
}
