// Coverage template
package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ProjectsSqlSuite struct {
	suite.Suite
}

func (s *ProjectsSqlSuite) TestQueriesCreateProject0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Vinnie Klocko", Name: "Rolando Macejkovic", Description: "Granville Konopelski", OrgID: "Greyson Walker", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Dudley Abernathy", Name: "Jacinto Murray", Description: "Karl Paucek", OrgID: "Elliott Walsh", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Magnus Dickinson", Name: "Jillian Fritsch", Description: "Mollie Hagenes", OrgID: "Sister Dibbert", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Giles Ryan", Name: "Omer Doyle", Description: "Travon Terry", OrgID: "Dusty Crooks", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Kolby Little", Name: "Jamal Monahan", Description: "Dan Altenwerth", OrgID: "Santina Luettgen", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Yasmin Gleason", Name: "Yasmine Waelchi", Description: "Maximilian Considine", OrgID: "Genesis Bosco", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Nathaniel Bogisich", Name: "Jerrold Witting", Description: "Royce Fisher", OrgID: "Laurel Kuhic", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Chasity Nikolaus", Name: "London Stehr", Description: "Gianni Bernhard", OrgID: "Lukas Collins", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Jovanny Funk", Name: "Efren Stehr", Description: "Avis Pfannerstill", OrgID: "Sarina Fahey", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Fleta Wiegand", Name: "Vincenzo Denesik", Description: "Bertrand Gibson", OrgID: "Felicia Jewess", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Winnifred Rosenbaum", Name: "Myrtle Morar", Description: "Irving Hudson", OrgID: "Hallie Hayes", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Kayli Kilback", Name: "Maybelle Schultz", Description: "Candice White", OrgID: "Margarita Batz", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Reilly Jakubowski", Name: "Daphney Hayes", Description: "Loraine White", OrgID: "Delpha Thompson", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Maybelle Doyle", Name: "Dane Jaskolski", Description: "Dino Stamm", OrgID: "Pattie Corwin", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Michaela Heathcote", Name: "Katheryn Bernier", Description: "Gaylord Feest", OrgID: "Cesar Bernhard", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Oda Bashirian", Name: "Brice Murray", Description: "Baron Mraz", OrgID: "Jared Ondricka", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Jeffery Dach", Name: "Art Smitham", Description: "Mose Kerluke", OrgID: "Jevon Hirthe", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesCreateProject17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateProject17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateProjectParams{ID: "Brennon Bailey", Name: "Florine Strosin", Description: "Granville Crist", OrgID: "Elnora Mertz", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Melissa Hermiston"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Toney Braun"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Dortha Littel"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Kara Lind"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Abelardo Romaguera"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Rylee Braun"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Rosa Block"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Coy Hettinger"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Delaney Kovacek"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Carolyn Langosh"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Frieda Sipes"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Dominic Pollich"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Rosie Johnson"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Liza Kuhn"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Rosella Armstrong"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Leonora Schamberger"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lawson Lubowitz"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesDeleteProject17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteProject17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Geovany Ruecker"

	q.DeleteProject(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Linda Mohr"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Ansley Haag"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Maxie Macejkovic"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Marcelo Padberg"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Alfredo Goldner"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Stuart Kertzmann"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Daphne Crona"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Chanelle Becker"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Quinn Auer"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Amber Carroll"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Barrett Goldner"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Christa Schmidt"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Robyn McGlynn"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Kareem Schneider"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Rudolph Collier"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Violette Volkman"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Consuelo Greenfelder"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByID17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByID17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Monty Gusikowski"

	q.GetProjectByID(ctx, id)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Samir Tremblay"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Isobel Nolan"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Kristoffer Bauch"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Remington Koch"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Austyn Jenkins"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Reynold Denesik"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Lera Krajcik"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Allan Spencer"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Brown Hand"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Ella Jerde"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Dane Stiedemann"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Uriah Tromp"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jakayla Hayes"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Ena Hilpert"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Orval Bins"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Erwin Bechtelar"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Ashleigh Lowe"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectByName17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectByName17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Aida Kovacek"

	q.GetProjectByName(ctx, name)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Amya Hane"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Loraine Kuhlman"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Lila Hudson"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Elmira Huel"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Emile Williamson"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Savion Bauch"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Devonte Jacobs"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Cristal Wuckert"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Raphaelle Kihn"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Mavis Heaney"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Allie Hoppe"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Dasia Barton"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Leilani Langosh"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Jaylon Cronin"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Gonzalo Senger"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Shania Runte"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Jairo Bradtke"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesGetProjectsByOrgID17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetProjectsByOrgID17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	orgID := "Annette Moen"

	q.GetProjectsByOrgID(ctx, orgID)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Brando Jenkins", Description: "Lavinia Koelpin", UpdatedAt: time.Time{}, ID: "Macie Yost"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Jennings Bosco", Description: "Dejon Sanford", UpdatedAt: time.Time{}, ID: "Ashley Yundt"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Ubaldo Johns", Description: "Donnell Fritsch", UpdatedAt: time.Time{}, ID: "Koby Ebert"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Blanche Corwin", Description: "Gus Schulist", UpdatedAt: time.Time{}, ID: "Mac Hoppe"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Bobbie Bosco", Description: "Elsie Jerde", UpdatedAt: time.Time{}, ID: "Kobe Schmeler"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Makenna Runolfsson", Description: "Dana Dicki", UpdatedAt: time.Time{}, ID: "Anderson Williamson"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Florine Grimes", Description: "Jayson Durgan", UpdatedAt: time.Time{}, ID: "Solon Hoeger"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Layla Williamson", Description: "Nia Rowe", UpdatedAt: time.Time{}, ID: "Lurline Mertz"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Karelle Champlin", Description: "Isai Johnston", UpdatedAt: time.Time{}, ID: "Jordon West"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Meda Spencer", Description: "Winona Metz", UpdatedAt: time.Time{}, ID: "Monty Beier"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Max Carter", Description: "Aileen Treutel", UpdatedAt: time.Time{}, ID: "Mollie Johns"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Ulises Flatley", Description: "Loren Lang", UpdatedAt: time.Time{}, ID: "Golda Satterfield"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Ernie Stiedemann", Description: "Shany Lesch", UpdatedAt: time.Time{}, ID: "Cydney Dach"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Florencio Kunze", Description: "Prince Reynolds", UpdatedAt: time.Time{}, ID: "Johnnie Lowe"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Rubie Emmerich", Description: "Corene Wiza", UpdatedAt: time.Time{}, ID: "Tressie Bernhard"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Trisha Jerde", Description: "Christiana Langosh", UpdatedAt: time.Time{}, ID: "Lyric Watsica"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Roy Pagac", Description: "Earl Mraz", UpdatedAt: time.Time{}, ID: "Jaylen Hessel"}

	q.UpdateProject(ctx, arg)

}

func (s *ProjectsSqlSuite) TestQueriesUpdateProject17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateProject17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateProjectParams{Name: "Victoria Rice", Description: "Bryon Bahringer", UpdatedAt: time.Time{}, ID: "Archibald Mann"}

	q.UpdateProject(ctx, arg)

}

func TestProjectsSqlSuite(t *testing.T) {
	suite.Run(t, new(ProjectsSqlSuite))
}
