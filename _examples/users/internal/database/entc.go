package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/yyewolf/entvis"
)

func main() {
	err := entc.Generate("./internal/database/ent/schema", &gen.Config{}, entc.Extensions(entvis.NewViewExtension()))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
