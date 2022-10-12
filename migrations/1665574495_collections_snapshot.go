package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

// Auto generated migration with the most recent collections configuration.
func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "05fenidlixxgc3h",
				"created": "2022-10-10 07:37:32.487",
				"updated": "2022-10-11 10:17:42.122",
				"name": "challenges",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "rv3x207n",
						"name": "slug",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^[a-z0-9_]+$"
						}
					},
					{
						"system": false,
						"id": "nyns3oak",
						"name": "title",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "8npksrjl",
						"name": "topic",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "y5fxnca9jza0r1u",
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "xqyyp8zk",
						"name": "tags",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 5,
							"collectionId": "vlq1ylcijwbvjoc",
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "jcgfqz3v",
						"name": "lead",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "4qshp7qy",
						"name": "impact",
						"type": "select",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"bigpoint",
								"peanut"
							]
						}
					},
					{
						"system": false,
						"id": "fglyw68y",
						"name": "type",
						"type": "select",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"one-time",
								"recurring",
								"repeatable"
							]
						}
					},
					{
						"system": false,
						"id": "efdzyb4x",
						"name": "frontMatter",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "yhiwpa35",
						"name": "summary",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "a1gyhfav",
						"name": "tips",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "m18mi4nw",
						"name": "content",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "g8dyn9dv",
						"name": "difficulties",
						"type": "json",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "tsfoyn8x",
						"name": "completedText",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "tykxfxhs",
						"name": "image",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": [
								"512x512"
							]
						}
					},
					{
						"system": false,
						"id": "wrekclaf",
						"name": "reminderText",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "1bqondjr",
						"name": "notificationDays",
						"type": "json",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "yygggldn",
						"name": "score",
						"type": "json",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "srcdgpuk",
						"name": "published",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "vlq1ylcijwbvjoc",
				"created": "2022-10-10 13:10:38.479",
				"updated": "2022-10-11 10:17:42.124",
				"name": "tags",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "x1exfwll",
						"name": "tag",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^[a-z_]+$"
						}
					},
					{
						"system": false,
						"id": "3hntbogu",
						"name": "label",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "y5fxnca9jza0r1u",
				"created": "2022-10-10 13:11:14.553",
				"updated": "2022-10-11 10:17:42.125",
				"name": "topics",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "zaqnydxx",
						"name": "topic",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^[a-z_]+$"
						}
					},
					{
						"system": false,
						"id": "xnciefav",
						"name": "label",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "systemprofiles0",
				"created": "2022-10-11 08:42:21.984",
				"updated": "2022-10-11 10:17:42.125",
				"name": "profiles",
				"system": true,
				"schema": [
					{
						"system": true,
						"id": "pbfielduser",
						"name": "userId",
						"type": "user",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "pbfieldname",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pbfieldavatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					}
				],
				"listRule": "userId = @request.user.id",
				"viewRule": "userId = @request.user.id",
				"createRule": "userId = @request.user.id",
				"updateRule": "userId = @request.user.id",
				"deleteRule": null
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		// no revert since the configuration on the environment, on which
		// the migration was executed, could have changed via the UI/API
		return nil
	})
}
