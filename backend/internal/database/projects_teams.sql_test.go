// Coverage template
package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProjectsTeamsSqlSuite struct {
	suite.Suite
}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Trace Mayert", TeamID: "Rosemarie Mertz"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Spencer Gleason", TeamID: "Ardith Erdman"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Gust Corkery", TeamID: "Kira Huels"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Donald Corwin", TeamID: "Maritza Yundt"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Adell Zulauf", TeamID: "Presley Jerde"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Mariana Ferry", TeamID: "Zoey Bogisich"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Solon Koelpin", TeamID: "Savanna Davis"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Jarvis Jaskolski", TeamID: "Waino Abbott"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Gideon Gleason", TeamID: "Glenna McKenzie"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Hazel Wyman", TeamID: "Torrey Doyle"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Athena Crooks", TeamID: "Lily Purdy"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Waldo Nolan", TeamID: "Thalia Macejkovic"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Vicente Thiel", TeamID: "Janice Shanahan"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Dennis Kohler", TeamID: "Mariam Quitzon"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Damaris DuBuque", TeamID: "Manuela Dicki"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Kyler VonRueden", TeamID: "Deja Rogahn"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Ole Ritchie", TeamID: "Ava Lind"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesAddTeamToProject17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesAddTeamToProject17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := AddTeamToProjectParams{ProjectID: "Abdullah Becker", TeamID: "Hailie Treutel"}

	q.AddTeamToProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Dianna Schumm"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Federico Yundt"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Joanie Trantow"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Haylee Smith"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Terrill Homenick"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Miracle Boyle"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Grace Pfannerstill"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Bertram Lindgren"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Jermain Nikolaus"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Verlie Reilly"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Duncan Lakin"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Mariah Swift"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Noble McLaughlin"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Carlos Fritsch"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Nasir Marks"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Cody Wilderman"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Libbie Stark"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsByTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsByTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	teamID := "Johnnie Ledner"

	q.GetAllProjectsByTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Evans Hauck"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Tod Parker"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Thaddeus Walker"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Charity Keeling"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Natasha Rodriguez"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Blanche Sipes"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Destiny Rodriguez"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Ellis Bernhard"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Emmanuel Hettinger"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Kiley Leannon"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Kris Torphy"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Joanne Rice"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Ward Parisian"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Franco Damore"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Estella Grant"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Enrico Herman"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Theodore Harber"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllTeamsFromProject17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllTeamsFromProject17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	projectID := "Jada Gutmann"

	q.GetAllTeamsFromProject(ctx, projectID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Jordi Hamill", TeamID: "Claudine Schaefer"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Elenora Treutel", TeamID: "Asia Stiedemann"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Nash Steuber", TeamID: "Alysson Ortiz"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Antwon Schiller", TeamID: "Vivien Rippin"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Chet Bruen", TeamID: "Filomena Wolf"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Fausto McKenzie", TeamID: "Davon Zulauf"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "May Ryan", TeamID: "Haley Schinner"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Kenyatta Lesch", TeamID: "Verner McClure"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Horace Padberg", TeamID: "Adela Robel"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Geraldine Lockman", TeamID: "Alexie Gleason"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Cristobal Schneider", TeamID: "Gage Buckridge"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Zachariah Bins", TeamID: "Wava Price"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Penelope Koepp", TeamID: "Domenico Weber"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Lavina Wintheiser", TeamID: "Manuela Keeling"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Alexzander Littel", TeamID: "Bradly Rutherford"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Tyler Conn", TeamID: "Domenick Rau"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Reyna White", TeamID: "Josephine Wisoky"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetOneProjectFromTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOneProjectFromTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := GetOneProjectFromTeamParams{ProjectID: "Vallie Hoppe", TeamID: "Pierre Kuvalis"}

	q.GetOneProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Ethan Hand", TeamID: "Jude Hegmann"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Jace Kuhic", TeamID: "Chadd Shanahan"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Theo Lesch", TeamID: "Gretchen Prohaska"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Brielle Bechtelar", TeamID: "Scottie Goyette"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Cortney Russel", TeamID: "Deja Dickens"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Clemens Cassin", TeamID: "Lowell Nolan"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Izabella Kuhn", TeamID: "Candelario Kessler"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Hilario Goldner", TeamID: "Sonny Conn"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Devin Zboncak", TeamID: "Erick Beier"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Vance Buckridge", TeamID: "Angeline Sawayn"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Lelah Upton", TeamID: "Maddison Carter"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Mckenzie Spencer", TeamID: "Miller Lubowitz"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Saige Cartwright", TeamID: "Seamus Turcotte"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Monserrat Ledner", TeamID: "Nelson Denesik"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Piper Rodriguez", TeamID: "Kelley Quigley"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Linda Macejkovic", TeamID: "Annette Hermiston"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Kali Goyette", TeamID: "Rudolph Cremin"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveProjectFromTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveProjectFromTeam17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveProjectFromTeamParams{ProjectID: "Stephen Stamm", TeamID: "Mateo Yost"}

	q.RemoveProjectFromTeam(ctx, arg)

}

func TestProjectsTeamsSqlSuite(t *testing.T) {
	suite.Run(t, new(ProjectsTeamsSqlSuite))
}
