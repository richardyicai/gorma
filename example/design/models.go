package design

import (
	"github.com/bketelsen/gorma"
	. "github.com/bketelsen/gorma/dsl"
)

var sg = StorageGroup("MyStorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		Model("Bottle", BottlePayload, func() {
			Description("This is the bottle model")
			Field("ID", gorma.PKInteger, func() {
				Description("This is the ID PK field")
			})
			Field("Vintage", gorma.Integer, func() {
				SQLTag("index")
			})
			Field("CreatedAt", gorma.Timestamp, func() {})
			Field("UpdatedAt", gorma.Timestamp, func() {})
			Field("DeletedAt", gorma.NullableTimestamp, func() {})
		})
	})
})
