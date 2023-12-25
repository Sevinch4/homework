package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"main/market/categories"
	"main/market/category_repo"
	"time"
)

func main() {
	db, err := ConnectionDB()
	if err != nil {
		fmt.Println("error is while connecting to database", err)
	}
	defer db.Close()

	categoryRepo := category_repo.NewCategory(db)

	var option int
	fmt.Print(`
				1 - categories
				2 - products
choose option:
`)
	fmt.Scan(&option)
	switch option {
	case 1: //categories
		var catOption int
		fmt.Print(`
			1 - insert categories
			2 - Update categories
			3 - get List
			4 - get by id
			5 - delete
choose option:
`)
		fmt.Scan(&catOption)
		switch catOption {
		case 1: //insert
			var category_name string
			fmt.Print("Input category name: ")
			fmt.Scan(&category_name)
			cat := categories.Category{
				ID:         uuid.New(),
				Name:       category_name,
				Created_at: time.Now(),
				Updated_at: time.Now(),
			}
			if err = categoryRepo.InsertCategory(cat); err != nil {
				fmt.Println("error is while inserting", err)
				return
			}
			fmt.Println("category added successfully")
		case 2: //update
			var (
				category_id   string
				category_name string
			)
			fmt.Print("input id: ")
			fmt.Scan(&category_id)
			fmt.Print("Input category name: ")
			fmt.Scan(&category_name)
			//update method
			id_cat, err := uuid.Parse(category_id)
			if err != nil {
				fmt.Println("error is while parsing id", err)
				return
			}
			c := categories.Category{
				ID:         id_cat,
				Name:       category_name,
				Updated_at: time.Now(),
			}
			if err = categoryRepo.UpdateCategoryByID(c); err != nil {
				fmt.Println("error is while updating data", err)
				return
			}
			fmt.Println("category updated successfully")
		case 3: //get list
			if err = categoryRepo.GetList(); err != nil {
				fmt.Println("error is while calling get list method", err)
				return
			}
		case 4: //get by id
			var category_id string
			fmt.Print("input id: ")
			fmt.Scan(&category_id)

			if err = categoryRepo.GetCategoryByID(category_id); err != nil {
				fmt.Println("error is while calling get category by id ", err)
				return
			}
		case 5: //delete
			var category_id string
			fmt.Print("input id: ")
			fmt.Scan(&category_id)

			if err = categoryRepo.DeleteCategory(category_id); err != nil {
				fmt.Println("error is while calling get category by id ", err)
				return
			}
			fmt.Println("category deleted successfully")
		}

	case 2: //products
		var prodOption int
		fmt.Print(`
			1 - insert products
			2 - Update product
			3 - get List
			4 - get by id
			5 - delete
choose option:
`)
		fmt.Scan(&prodOption)
		switch prodOption {
		case 1: //insert
		case 2: //update
		case 3: //get list
		case 4: //get by id
		case 5: //delete

		}

	}

}

func ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host = localhost port = 5432 password = 1218 database = market sslmode = disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
