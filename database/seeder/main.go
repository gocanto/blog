package main

import (
	"github.com/gocanto/blog/bootstrap"
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/database/seeder/seeds"
	"github.com/gocanto/blog/env"
	"github.com/gocanto/blog/webkit/cli"
	"sync"
	"time"
)

var environment *env.Environment
var textColour *cli.TextColour

func init() {
	secrets, _ := bootstrap.Spark("./.env")

	environment = secrets
}

func main() {
	dbConnection := bootstrap.MakeDbConnection(environment)
	logs := bootstrap.MakeLogs(environment)

	defer (*logs).Close()
	defer (*dbConnection).Close()

	// [1] --- Truncate the DB.
	truncateDB(dbConnection, environment)
	cli.MakeTextColour("DB Truncated successfully ...", cli.Green).Println()
	time.Sleep(2 * time.Second)

	// [2] --- Create the Seeder Runner.
	seeder := seeds.MakeSeeder(dbConnection)

	// [3] --- Seed users and posts sequentially because the below seeders depend on them.
	UserA, UserB := seeder.SeedUsers()
	posts := seeder.SeedPosts(UserA, UserB)

	categoriesChan := make(chan []database.Category)
	tagsChan := make(chan []database.Tag)

	go func() {
		defer close(categoriesChan)

		cli.MakeTextColour("Seeding categories ...", cli.Yellow).Println()
		categoriesChan <- seeder.SeedCategories()
	}()

	go func() {
		defer close(tagsChan)

		cli.MakeTextColour("Seeding tags ...", cli.Magenta).Println()
		tagsChan <- seeder.SeedTags()
	}()

	// [4] Use channels to concurrently seed categories and tags since they are main dependencies.
	categories := <-categoriesChan
	tags := <-tagsChan

	// [5] Use a WaitGroup to run independent seeding tasks concurrently.
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()

		cli.MakeTextColour("Seeding comments ...", cli.Blue).Println()
		seeder.SeedComments(posts...)
	}()

	go func() {
		defer wg.Done()

		cli.MakeTextColour("Seeding likes ...", cli.Cyan).Println()
		seeder.SeedLikes(posts...)
	}()

	go func() {
		defer wg.Done()

		cli.MakeTextColour("Seeding posts-categories ...", cli.Gray).Println()
		seeder.SeedPostsCategories(categories, posts)
	}()

	go func() {
		defer wg.Done()

		cli.MakeTextColour("Seeding posts-tags ...", cli.Magenta).Println()
		seeder.SeedPostTags(tags, posts)
	}()

	go func() {
		defer wg.Done()

		cli.MakeTextColour("Seeding views ...", cli.Yellow).Println()
		seeder.SeedPostViews(posts, UserA, UserB)
	}()

	wg.Wait()

	cli.MakeTextColour("DB seeded as expected.", cli.Green).Println()
}
