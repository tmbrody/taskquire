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

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam0", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam1", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam2", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam3", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam4", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam5", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam6", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam7", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam8", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam9", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam10", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam11", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam12", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam13", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam14", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam15", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam16", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesGetAllProjectsFromTeam17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllProjectsFromTeam17", r)
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

	q.GetAllProjectsFromTeam(ctx, teamID)

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

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Ethan Hand", TeamID: "Jude Hegmann"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Jace Kuhic", TeamID: "Chadd Shanahan"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Theo Lesch", TeamID: "Gretchen Prohaska"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Brielle Bechtelar", TeamID: "Scottie Goyette"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Cortney Russel", TeamID: "Deja Dickens"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Clemens Cassin", TeamID: "Lowell Nolan"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Izabella Kuhn", TeamID: "Candelario Kessler"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Hilario Goldner", TeamID: "Sonny Conn"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Devin Zboncak", TeamID: "Erick Beier"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Vance Buckridge", TeamID: "Angeline Sawayn"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Lelah Upton", TeamID: "Maddison Carter"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Mckenzie Spencer", TeamID: "Miller Lubowitz"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Saige Cartwright", TeamID: "Seamus Turcotte"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Monserrat Ledner", TeamID: "Nelson Denesik"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Piper Rodriguez", TeamID: "Kelley Quigley"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Linda Macejkovic", TeamID: "Annette Hermiston"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Kali Goyette", TeamID: "Rudolph Cremin"}

	q.RemoveTeamFromProject(ctx, arg)

}

func (s *ProjectsTeamsSqlSuite) TestQueriesRemoveTeamFromProject17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesRemoveTeamFromProject17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := RemoveTeamFromProjectParams{ProjectID: "Stephen Stamm", TeamID: "Mateo Yost"}

	q.RemoveTeamFromProject(ctx, arg)

}

func TestProjectsTeamsSqlSuite(t *testing.T) {
	suite.Run(t, new(ProjectsTeamsSqlSuite))
}
