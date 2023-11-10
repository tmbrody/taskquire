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

type TeamsSqlSuite struct {
	suite.Suite
}

func (s *TeamsSqlSuite) TestQueriesCreateTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Tiara Rolfson", Name: "Haylie Goyette", Description: sql.NullString{String: "Maureen Sporer", Valid: false}, OwnerID: "Haley Friesen", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Ola VonRueden", Name: "Itzel Konopelski", Description: sql.NullString{String: "Dorian Huel", Valid: false}, OwnerID: "Cesar Hahn", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Sophie Ryan", Name: "Kaden Schowalter", Description: sql.NullString{String: "Janelle Kunze", Valid: true}, OwnerID: "Jamir Sipes", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Eino Price", Name: "Providenci Bednar", Description: sql.NullString{String: "Heather McKenzie", Valid: true}, OwnerID: "Cierra Prosacco", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Esteban Zemlak", Name: "Garnet Rogahn", Description: sql.NullString{String: "Magali Kub", Valid: false}, OwnerID: "Lolita Cruickshank", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Vivianne Wuckert", Name: "Orie Mills", Description: sql.NullString{String: "Alene Maggio", Valid: true}, OwnerID: "Mafalda Wolff", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Linnea Morissette", Name: "Ruben Runolfsson", Description: sql.NullString{String: "Gayle Hoeger", Valid: false}, OwnerID: "Jazlyn Gulgowski", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Timmothy Greenfelder", Name: "Emilio Hauck", Description: sql.NullString{String: "Karl Champlin", Valid: false}, OwnerID: "Aliyah Ankunding", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Idell Bayer", Name: "Stacey Stoltenberg", Description: sql.NullString{String: "Werner Jacobson", Valid: true}, OwnerID: "Stacey Weimann", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Ahmed Prohaska", Name: "Austen Beier", Description: sql.NullString{String: "Jalyn Mertz", Valid: false}, OwnerID: "Ramiro Kertzmann", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Manley Kling", Name: "Vickie Torphy", Description: sql.NullString{String: "Jamal Durgan", Valid: false}, OwnerID: "Kaci Conn", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "George Larkin", Name: "Harmony Toy", Description: sql.NullString{String: "Christelle Kutch", Valid: false}, OwnerID: "Cicero Cruickshank", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Jack Boyle", Name: "Cullen Carroll", Description: sql.NullString{String: "Reuben Mosciski", Valid: false}, OwnerID: "Maxime Romaguera", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Randall Schumm", Name: "Willow Sauer", Description: sql.NullString{String: "Marquise Wunsch", Valid: false}, OwnerID: "Demarco Raynor", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Perry Jacobs", Name: "Judson Hammes", Description: sql.NullString{String: "Tito Raynor", Valid: false}, OwnerID: "Bridget Bayer", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Verna Gerhold", Name: "Savanna Rodriguez", Description: sql.NullString{String: "Caleigh Funk", Valid: false}, OwnerID: "Major Zulauf", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Kim Mayert", Name: "Blaze Kiehn", Description: sql.NullString{String: "Sally Quigley", Valid: true}, OwnerID: "Orval Nitzsche", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesCreateTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateTeamParams{ID: "Camylle Okuneva", Name: "Vernice Bogisich", Description: sql.NullString{String: "Effie Gibson", Valid: false}, OwnerID: "Emanuel Homenick", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Randi Donnelly"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Bella Bins"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Jayde Kerluke"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Watson Hermiston"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Madonna West"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Tyshawn Kris"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Blaze Dooley"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Alessandra Corwin"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lyla Kiehn"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Cristian Flatley"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Paris Reichel"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Retha Prosacco"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Elmer Towne"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Franco Kirlin"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Earl Schneider"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Vernon White"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Gordon Hudson"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesDeleteTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Pablo Morar"

	q.DeleteTeam(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetAllTeams17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeams17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllTeams(ctx)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Keira Cruickshank"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Bernardo Rau"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Jennifer Koss"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Marilie Muller"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Mason Hills"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Meghan Gorczany"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Luther Breitenberg"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Isabella Johns"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Aubree Bins"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Gwen Christiansen"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Mathias Kihn"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Ford Emmerich"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Janice Hickle"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Geovanny Kreiger"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Reyna Erdman"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lia Ondricka"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Della Hodkiewicz"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByID17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByID17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Paolo Dicki"

	q.GetTeamByID(ctx, id)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Demetrius Mueller"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Moriah Beatty"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Raul Bergstrom"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Ramona Mohr"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Janick Pacocha"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jane Kessler"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Pearline DuBuque"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Lydia Witting"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Erin Nienow"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Ashlynn Wintheiser"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Cecelia Harvey"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Fredrick Rempel"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Abbigail Becker"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Dortha Barton"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Sigmund Boehm"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Gregory Lynch"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jordon Moore"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesGetTeamByName17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetTeamByName17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Susie Pouros"

	q.GetTeamByName(ctx, name)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Kelvin Waters", Description: sql.NullString{String: "Emelie Grimes", Valid: true}, UpdatedAt: time.Time{}, ID: "Araceli Steuber"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Meaghan Daniel", Description: sql.NullString{String: "Aracely Ondricka", Valid: false}, UpdatedAt: time.Time{}, ID: "Berniece Schuster"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Saige Wiegand", Description: sql.NullString{String: "Neva Stanton", Valid: true}, UpdatedAt: time.Time{}, ID: "Hester Hermann"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Jaren Hamill", Description: sql.NullString{String: "Wilburn Jakubowski", Valid: false}, UpdatedAt: time.Time{}, ID: "Einar Dooley"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Lori Mills", Description: sql.NullString{String: "Rebecca Hyatt", Valid: false}, UpdatedAt: time.Time{}, ID: "Sim Maggio"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Sydnie Cartwright", Description: sql.NullString{String: "Lamont Terry", Valid: false}, UpdatedAt: time.Time{}, ID: "Sonny Heller"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Bianka Mraz", Description: sql.NullString{String: "Idella Wolf", Valid: false}, UpdatedAt: time.Time{}, ID: "Will Haley"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Mateo Jacobi", Description: sql.NullString{String: "Danny Grant", Valid: false}, UpdatedAt: time.Time{}, ID: "Filiberto Luettgen"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Issac Swift", Description: sql.NullString{String: "Abigale Botsford", Valid: true}, UpdatedAt: time.Time{}, ID: "Josephine Koch"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Hassan Konopelski", Description: sql.NullString{String: "Dandre Simonis", Valid: false}, UpdatedAt: time.Time{}, ID: "Doug McDermott"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Betty Hackett", Description: sql.NullString{String: "Jess Metz", Valid: true}, UpdatedAt: time.Time{}, ID: "Shanie Mann"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Jamel Jerde", Description: sql.NullString{String: "Yadira Grant", Valid: false}, UpdatedAt: time.Time{}, ID: "Katrina Effertz"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Selena Brekke", Description: sql.NullString{String: "Dameon Ryan", Valid: true}, UpdatedAt: time.Time{}, ID: "Alexie Berge"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Jana Morar", Description: sql.NullString{String: "April Roberts", Valid: false}, UpdatedAt: time.Time{}, ID: "Art Carroll"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Eleazar Sporer", Description: sql.NullString{String: "Chasity Mraz", Valid: false}, UpdatedAt: time.Time{}, ID: "Cristopher Cassin"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Clyde Pfannerstill", Description: sql.NullString{String: "Joesph Bernier", Valid: false}, UpdatedAt: time.Time{}, ID: "Eloise Larkin"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "Samson Swaniawski", Description: sql.NullString{String: "Horace Emmerich", Valid: true}, UpdatedAt: time.Time{}, ID: "Mathew Hermann"}

	q.UpdateTeam(ctx, arg)

}

func (s *TeamsSqlSuite) TestQueriesUpdateTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateTeamParams{Name: "George Maggio", Description: sql.NullString{String: "Ava Funk", Valid: true}, UpdatedAt: time.Time{}, ID: "Elton Schiller"}

	q.UpdateTeam(ctx, arg)

}

func TestTeamsSqlSuite(t *testing.T) {
	suite.Run(t, new(TeamsSqlSuite))
}
