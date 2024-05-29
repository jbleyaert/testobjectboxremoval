//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen clean .
//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

package main

// ToBeRemoved is an entity that should be removed from the objectbox-model.json.
// type ToBeRemoved struct {
// 	Id   uint64 `objectbox:"id"`
// 	Name string
// }

type ToBeKept struct {
	Id   uint64 `objectbox:"id"`
	Name string
}
