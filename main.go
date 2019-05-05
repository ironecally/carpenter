package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	log.Println("asd")
	initGithub(ctx)
	err := getRepoInfo(ctx, "ironecally", "carpenter")
	log.Println(err)
}
