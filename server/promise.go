package server

import (
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/promise"
)

func GetPromise(client database.Client, id string) (promise.Promise, error) {
	return promise.GetPromise(client, id)
}

func GetPromises(client database.Client) ([]promise.Promise, error) {
	return promise.GetPromises(client)
}

func GetPromiseFavourites(client database.Client) {}

func GetPromiseTag(client database.Client, id string) (promise.Tag, error) {
	return promise.GetTag(client, id)
}

func GetPromiseTags(client database.Client) ([]promise.Tag, error) {
	return promise.GetTags(client)
}

func GetPromiseVotes(client database.Client) ([]promise.Vote, error) {
	return promise.GetVotes(client)
}
