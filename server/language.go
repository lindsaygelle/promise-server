package server

import (
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/language"
)

// GetLanguageTag returns a language.Tag.
func GetLanguageTag(client database.Client, id string) (language.Tag, error) {
	return language.GetTag(client, id)
}

// GetLanguageTags returns a slice of language.Tag.
func GetLanguageTags(client database.Client) ([]language.Tag, error) {
	return language.GetTags(client)
}

// GetLanguage returns a language.Language.
func GetLanguage(client database.Client, id string) (language.Language, error) {
	return language.GetLanguage(client, id)
}

// GetLanguages returns a slice of language.Language.
func GetLanguages(client database.Client) ([]language.Language, error) {
	return language.GetLanguages(client)
}
