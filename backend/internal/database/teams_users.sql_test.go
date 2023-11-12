// Coverage template
package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TeamsUsersSqlSuite struct {
	suite.Suite
}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Christa Moore", TeamID: "Hobart Hayes"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Granville Prosacco", TeamID: "Dell McGlynn"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Afton Wilkinson", TeamID: "Hardy Johnston"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Graciela Crist", TeamID: "Colton Davis"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Webster Price", TeamID: "Brandon Crist"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Shanelle Bednar", TeamID: "Yvette Toy"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Melvin Hilll", TeamID: "Delphine DuBuque"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Ramona West", TeamID: "Hazel Lemke"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Andreane Bins", TeamID: "Rowland Watsica"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Elinor Gusikowski", TeamID: "Cassidy Koelpin"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Noah Corwin", TeamID: "Estevan Kiehn"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Gwendolyn Hoppe", TeamID: "Chandler Raynor"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Abbigail Borer", TeamID: "Aiden Emard"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Emory Sipes", TeamID: "Jaylen Lowe"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Bettie Hahn", TeamID: "Jasen Lang"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Edna Cummerata", TeamID: "Eleonore Orn"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Eveline Boehm", TeamID: "Ali Christiansen"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesAddUsertoTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddUsertoTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddUsertoTeamParams{UserID: "Genevieve Konopelski", TeamID: "Savannah Mayert"}

	q.AddUsertoTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Marcos Cruickshank"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Anibal Schroeder"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Hillary Conroy"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Danyka Moore"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Tyshawn Cummings"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Ahmed Balistreri"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Arnaldo Herman"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Citlalli Mitchell"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Wilber Hills"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Audie Braun"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Jeff Hegmann"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Leda Hansen"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Ignacio Gleason"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Frankie Corkery"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Ansel Jaskolski"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Carter Thiel"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Cordelia McDermott"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllTeamsByUser17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsByUser17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	userID := "Idell Brown"

	q.GetAllTeamsByUser(ctx, userID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Remington Jast"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Maxwell Grady"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Eda Kling"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Francesco Legros"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Turner Friesen"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Terrell Roberts"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Kareem Tremblay"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Bernice Powlowski"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Cathrine Upton"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Kaylee Boehm"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Roman Buckridge"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Stacey Ledner"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Marge Simonis"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Grant Bogan"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Dave Littel"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Donnell Kiehn"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Fabiola Brekke"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetAllUsersFromTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsersFromTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Sigrid Hayes"

	q.GetAllUsersFromTeam(ctx, teamID)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Louvenia Lowe", TeamID: "Dominic Gorczany"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Kayleigh Deckow", TeamID: "Krista Bartell"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Roel Schmitt", TeamID: "Mac Jenkins"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Alva Kuhlman", TeamID: "Ezequiel Thompson"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Malvina White", TeamID: "Shaina Doyle"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Gerson Ankunding", TeamID: "Madelynn Cartwright"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Stefanie Hane", TeamID: "Enrique Kshlerin"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Dylan Bernier", TeamID: "Zoe Turner"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Finn Zieme", TeamID: "May Walker"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Hulda Jacobson", TeamID: "Kathleen Gutmann"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Levi Simonis", TeamID: "Freeda Ferry"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Veda Bailey", TeamID: "Jayden Botsford"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Michel Hills", TeamID: "Merritt Dickinson"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Sammie Pfannerstill", TeamID: "Gabe Kassulke"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Dakota Feeney", TeamID: "Alison Doyle"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Carley Lebsack", TeamID: "Ellie Doyle"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Barrett Casper", TeamID: "Bertram Roob"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesGetOneUserFromTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneUserFromTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneUserFromTeamParams{UserID: "Andre Romaguera", TeamID: "Elena Abshire"}

	q.GetOneUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Elvie Johnson", TeamID: "Vena Wuckert"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Rylan Nicolas", TeamID: "Abby Cole"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Dexter Swift", TeamID: "Winifred Mayer"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Aliyah Kessler", TeamID: "Leilani Kautzer"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Jarrell Johnson", TeamID: "Callie Lynch"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Sterling Bernhard", TeamID: "Abigayle Konopelski"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Gunner Kilback", TeamID: "Domenica Konopelski"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Mazie Keeling", TeamID: "Don Mosciski"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Triston Franecki", TeamID: "Dean Reilly"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Gayle Abbott", TeamID: "Maudie Schaden"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Nigel Konopelski", TeamID: "Sydnie Barrows"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Rylee Johnson", TeamID: "Keyon Wunsch"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Dimitri Eichmann", TeamID: "Joannie Rippin"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Maxwell Farrell", TeamID: "Breana Morissette"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Mariah Nikolaus", TeamID: "Jaeden Koss"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Darwin Hermiston", TeamID: "Percival Harvey"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Omer Jacobson", TeamID: "Alia Sawayn"}

	q.RemoveUserFromTeam(ctx, arg)

}

func (s *TeamsUsersSqlSuite) TestQueriesRemoveUserFromTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveUserFromTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveUserFromTeamParams{UserID: "Estrella Rosenbaum", TeamID: "Wellington Harvey"}

	q.RemoveUserFromTeam(ctx, arg)

}

func TestTeamsUsersSqlSuite(t *testing.T) {
	suite.Run(t, new(TeamsUsersSqlSuite))
}
