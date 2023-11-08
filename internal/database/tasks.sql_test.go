// Coverage template
package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TasksSqlSuite struct {
	suite.Suite
}

func (s *TasksSqlSuite) TestQueriesCreateTask0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Holden Bayer", Name: "Sven Dooley", Description: sql.NullString{String: "Susan Kuhn", Valid: false}, ProjectID: "Bridget Lesch", TeamID: "Stephon Koelpin", OwnerID: "Jayme Leffler", ParentID: "Teresa Kihn", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Augustus Johnson", Name: "Taya Boehm", Description: sql.NullString{String: "Neva Ortiz", Valid: false}, ProjectID: "Adan Casper", TeamID: "Lionel Pacocha", OwnerID: "Shane Schimmel", ParentID: "Adelle Carroll", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Jo Watsica", Name: "Dahlia Farrell", Description: sql.NullString{String: "Vicente Bednar", Valid: true}, ProjectID: "Ana Baumbach", TeamID: "Hester Wiza", OwnerID: "Cleora Little", ParentID: "Joanie Cole", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Luella Bartell", Name: "Jed DuBuque", Description: sql.NullString{String: "Eugenia Rowe", Valid: false}, ProjectID: "Margarett Jerde", TeamID: "Deshawn Heller", OwnerID: "Mossie Goyette", ParentID: "Eden Pacocha", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Verda Davis", Name: "Frederik Schoen", Description: sql.NullString{String: "Marquis Runolfsson", Valid: true}, ProjectID: "Lloyd Cassin", TeamID: "Kevin Littel", OwnerID: "Tremaine Roob", ParentID: "Odie Kuphal", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Ebba Damore", Name: "Don Medhurst", Description: sql.NullString{String: "Royce Hessel", Valid: false}, ProjectID: "Kyleigh Ernser", TeamID: "Magdalena Rippin", OwnerID: "Fredy Runolfsson", ParentID: "Elinor Wisoky", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Alyson Herzog", Name: "Guillermo Schmitt", Description: sql.NullString{String: "Toney Gislason", Valid: false}, ProjectID: "Bartholome Schaefer", TeamID: "Christ Miller", OwnerID: "Saul Kunze", ParentID: "Clementine Hermiston", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Ollie Marvin", Name: "Brown Waelchi", Description: sql.NullString{String: "Burley Corwin", Valid: true}, ProjectID: "Kianna Quitzon", TeamID: "Wanda Bayer", OwnerID: "Colby Hegmann", ParentID: "Mayra Cole", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Sydney Rau", Name: "Reagan Rogahn", Description: sql.NullString{String: "Lenna Konopelski", Valid: true}, ProjectID: "Brant Mraz", TeamID: "Fredy Funk", OwnerID: "Clarissa Shields", ParentID: "Price Bruen", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Delaney Runolfsdottir", Name: "Scot Ryan", Description: sql.NullString{String: "Lawrence Adams", Valid: false}, ProjectID: "Emery White", TeamID: "Milton Wiza", OwnerID: "Newton Marks", ParentID: "Telly Baumbach", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Mossie Gerlach", Name: "Raina Rolfson", Description: sql.NullString{String: "Danial Nolan", Valid: true}, ProjectID: "Lonie Price", TeamID: "Garret McDermott", OwnerID: "Gwen Lang", ParentID: "Leanne Mitchell", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Misty Daniel", Name: "Frankie Hessel", Description: sql.NullString{String: "Brayan Treutel", Valid: true}, ProjectID: "Taryn Daugherty", TeamID: "Alfred Glover", OwnerID: "Mozelle Torp", ParentID: "Shayne Smitham", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Nikko Conn", Name: "Lesley Grady", Description: sql.NullString{String: "Kelsie Gleichner", Valid: true}, ProjectID: "Kayden Dietrich", TeamID: "Aurelio Olson", OwnerID: "Elmira Schultz", ParentID: "Kale Huel", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Joanny Grimes", Name: "King Willms", Description: sql.NullString{String: "Arden Rau", Valid: true}, ProjectID: "Stacey Huels", TeamID: "Marcelina Reichel", OwnerID: "Alfredo Graham", ParentID: "Carson Maggio", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Adelia Harber", Name: "Jarod Yost", Description: sql.NullString{String: "Americo Waters", Valid: true}, ProjectID: "Terrell Stroman", TeamID: "Wilson Spinka", OwnerID: "Idell Ferry", ParentID: "Dorothy Heathcote", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Abbigail Bayer", Name: "Chance Bosco", Description: sql.NullString{String: "Danyka Parker", Valid: true}, ProjectID: "Kellen Will", TeamID: "Ola Lakin", OwnerID: "Gabriella Ankunding", ParentID: "Kendrick Effertz", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Grover Lynch", Name: "Isabell Rutherford", Description: sql.NullString{String: "Kristina Block", Valid: false}, ProjectID: "Daphnee Macejkovic", TeamID: "Curtis Schaden", OwnerID: "Mathew Fisher", ParentID: "Tillman Yundt", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesCreateTask17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTask17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTaskParams{ID: "Loy Lynch", Name: "Anastasia Moen", Description: sql.NullString{String: "Golda Aufderhar", Valid: true}, ProjectID: "Marcelle Schultz", TeamID: "Lulu Konopelski", OwnerID: "Lisa Larson", ParentID: "Bryon Oberbrunner", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lucious Upton"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Hayden Lehner"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Garfield Lang"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Enid Tillman"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Ike Leannon"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Kenyatta Lubowitz"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Adolph Goldner"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Jack Brakus"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Letha Hartmann"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Hulda Veum"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Ronaldo Kub"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Josianne Wolf"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Casandra Maggio"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Tatum Gorczany"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lauretta Christiansen"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lucinda Gusikowski"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Alison Gleichner"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesDeleteTask17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTask17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Berenice Mante"

	q.DeleteTask(ctx, id)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Angel Harris"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Merlin Ferry"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Earnestine West"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Dimitri Huel"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Roger Lemke"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Ashlynn Cormier"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Liana Shanahan"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Tania Hackett"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Else Cruickshank"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Rosemary Moen"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Erling Mraz"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Alessandro Windler"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Efren Ortiz"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Neil Cronin"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Libbie Boehm"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Marcelino Reynolds"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Kellen Corwin"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetSubtasksByParentID17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetSubtasksByParentID17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	parentID := "Keaton Reichert"

	q.GetSubtasksByParentID(ctx, parentID)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Molly Grant"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Leonel Braun"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Cooper Klein"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Nyasia Crona"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Emily Ryan"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Dock Bayer"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Anabelle Moen"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Peggie Ferry"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Dolores Murray"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Doris Streich"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Kristoffer Denesik"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Alexa DuBuque"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Carrie Green"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jacey Ernser"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Holly Gibson"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Orville Stokes"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Ebony Reynolds"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTaskByName17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTaskByName17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Odessa Kirlin"

	q.GetTaskByName(ctx, name)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Thelma Hammes"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Dawson Lind"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Jaron Langworth"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Robert Crona"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Ray Hermiston"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Hipolito Auer"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Felicita Halvorson"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Lottie Davis"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Cecile Stark"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Zion Schmidt"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Kailee Harvey"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Aurelio Kessler"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Oma Sawayn"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Justice Kling"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Aimee Morissette"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Daryl Medhurst"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Aubrey Treutel"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesGetTasksByProjectID17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTasksByProjectID17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Cloyd Lang"

	q.GetTasksByProjectID(ctx, projectID)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Eldred Wiegand", Description: sql.NullString{String: "Kirk Steuber", Valid: true}, UpdatedAt: time.Time{}, ID: "Lennie Ratke"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Garrick Lynch", Description: sql.NullString{String: "Holly Lowe", Valid: true}, UpdatedAt: time.Time{}, ID: "Gust Goodwin"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Roy Balistreri", Description: sql.NullString{String: "Rodrick Kirlin", Valid: false}, UpdatedAt: time.Time{}, ID: "Quentin Smitham"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Barry Franecki", Description: sql.NullString{String: "Vivienne Wolff", Valid: false}, UpdatedAt: time.Time{}, ID: "Deontae Grimes"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Megane Hahn", Description: sql.NullString{String: "Zoey Kirlin", Valid: false}, UpdatedAt: time.Time{}, ID: "Devon Christiansen"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Marisol Schaden", Description: sql.NullString{String: "Fae Davis", Valid: true}, UpdatedAt: time.Time{}, ID: "Ava Volkman"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Helga Boyle", Description: sql.NullString{String: "Myah Stokes", Valid: false}, UpdatedAt: time.Time{}, ID: "Omari Russel"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Theron Rowe", Description: sql.NullString{String: "Sedrick Reynolds", Valid: true}, UpdatedAt: time.Time{}, ID: "Aileen Schumm"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Ebony Rohan", Description: sql.NullString{String: "Nathaniel Stehr", Valid: false}, UpdatedAt: time.Time{}, ID: "Macy Jast"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Garrison Gutkowski", Description: sql.NullString{String: "Rowan Heaney", Valid: true}, UpdatedAt: time.Time{}, ID: "Shany Bartell"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Demario Pagac", Description: sql.NullString{String: "Emmitt Murazik", Valid: false}, UpdatedAt: time.Time{}, ID: "Cecelia Vandervort"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Desmond Davis", Description: sql.NullString{String: "Ivory Bailey", Valid: true}, UpdatedAt: time.Time{}, ID: "Cindy Lemke"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Darian Ledner", Description: sql.NullString{String: "Alexandre Morissette", Valid: false}, UpdatedAt: time.Time{}, ID: "Benny Lind"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Karlie Olson", Description: sql.NullString{String: "Johathan Treutel", Valid: true}, UpdatedAt: time.Time{}, ID: "Sophie Veum"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Beaulah Klein", Description: sql.NullString{String: "Art Roob", Valid: false}, UpdatedAt: time.Time{}, ID: "Maymie Farrell"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Craig Wyman", Description: sql.NullString{String: "Bobbie Conroy", Valid: true}, UpdatedAt: time.Time{}, ID: "Claudia Ruecker"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Jeramy Shields", Description: sql.NullString{String: "Raymond Pacocha", Valid: false}, UpdatedAt: time.Time{}, ID: "Esperanza Miller"}

	q.UpdateTask(ctx, arg)

}

func (s *TasksSqlSuite) TestQueriesUpdateTask17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTask17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTaskParams{Name: "Tyrell Ullrich", Description: sql.NullString{String: "Alena Brown", Valid: true}, UpdatedAt: time.Time{}, ID: "Destiny McDermott"}

	q.UpdateTask(ctx, arg)

}

func TestTasksSqlSuite(t *testing.T) {
	suite.Run(t, new(TasksSqlSuite))
}
