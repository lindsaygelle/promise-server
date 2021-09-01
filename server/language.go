package server

import (
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/language"
)

// LanguageTag returns a language.Tag.
func LanguageTag(client database.Client, id string) (language.Tag, error) {
	return language.GetTag(client, id)
}

// LanguageTags returns a slice of language.Tag.
func LanguageTags(client database.Client) ([]language.Tag, error) {
	return language.GetTags(client)
}

// Language returns a language.Tag
func Language(client database.Client, id string) (language.Tag, error) {
	return language.GetTag(client, id)
}

// Languages returns a slice of language.Language.
func Languages(client database.Client) ([]language.Language, error) {
	return language.GetLanguages(client)
}
