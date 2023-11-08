// Coverage template
package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type OrgsSqlSuite struct {
	suite.Suite
}

func (s *OrgsSqlSuite) TestQueriesCreateOrg0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Alice Marquardt", Name: "Nels Harber", Description: "Weston Jakubowski", OwnerID: "Carley Runolfsdottir", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Heidi Torp", Name: "Pamela Douglas", Description: "Stephania Rogahn", OwnerID: "Sadie Nikolaus", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Margarette Runte", Name: "Pauline Hilll", Description: "Mckenzie Bartoletti", OwnerID: "Jeromy Hettinger", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Quinton Konopelski", Name: "Rachael Bayer", Description: "Marge Goyette", OwnerID: "Nicola Fay", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Gisselle Nicolas", Name: "Moises Mante", Description: "Eliane Lemke", OwnerID: "Columbus Orn", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Aniyah Shields", Name: "Neva Spinka", Description: "Enola Frami", OwnerID: "Greg Davis", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Dennis Mann", Name: "Dawn Herzog", Description: "Berneice Effertz", OwnerID: "Stewart Rice", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Yazmin Sipes", Name: "Juwan Witting", Description: "Henderson Kessler", OwnerID: "Callie Flatley", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Mya King", Name: "Dasia Turcotte", Description: "Lonzo Harber", OwnerID: "Geovanny Morar", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Tanner Hoppe", Name: "Herta Kreiger", Description: "Jayden Grant", OwnerID: "Stephen Rice", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Judson King", Name: "Kiel McLaughlin", Description: "Amelie McDermott", OwnerID: "Kayley Wiza", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Jettie Pfeffer", Name: "Marta Larson", Description: "Cecilia Emmerich", OwnerID: "Edgar Nikolaus", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Annabel Harber", Name: "Pauline Nolan", Description: "Amari Friesen", OwnerID: "Georgette Champlin", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Althea Bernier", Name: "Haylee Larkin", Description: "Shanny Hand", OwnerID: "Vergie Marquardt", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Kristina Labadie", Name: "Percy Daniel", Description: "Trace Paucek", OwnerID: "Erwin Bednar", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Ethan Corkery", Name: "Kenya Balistreri", Description: "Drake Zboncak", OwnerID: "Asha Marks", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Elsa Reynolds", Name: "Ellsworth Simonis", Description: "Amari Bartoletti", OwnerID: "Elliott Rempel", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesCreateOrg17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateOrg17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateOrgParams{ID: "Camille Shields", Name: "Timmy Ankunding", Description: "Asa Koss", OwnerID: "Evelyn Lesch", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Russell Crooks"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Kristy Harris"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Adelia Senger"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Prince Schamberger"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Myrtis Hackett"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Casper Moen"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Malika Schroeder"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Nora Bayer"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Jane Bauch"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Kip Koelpin"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Guido Yundt"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Hoyt Greenfelder"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Amina Gaylord"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Reese Stanton"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Virgil Wiza"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Anderson Lockman"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Darian White"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesDeleteOrg17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteOrg17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lelia Labadie"

	q.DeleteOrg(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetAllOrgs17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllOrgs17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllOrgs(ctx)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Payton Prohaska"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Freida Lemke"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Alvena Hammes"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lina Senger"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Helmer Cummerata"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Michelle Zulauf"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Wilfredo Franecki"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Adeline Ernser"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Dorris Mosciski"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Blanca Jaskolski"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Alfonzo Windler"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lucius Turcotte"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Koby Lebsack"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Alfonso Bechtelar"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Hailie Bergnaum"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Alvera Turner"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Kraig Reilly"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByID17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByID17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Elsa Kerluke"

	q.GetOrgByID(ctx, id)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Granville Hamill"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Bonnie Kling"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Elisa Kessler"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "River Champlin"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Carmel Lehner"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Julianne Spencer"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Makenzie Haley"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Arnoldo Fritsch"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Melyna Boyle"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Rubye Kutch"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Vena Hammes"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jennyfer Bogan"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jessy Lesch"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Trace Hagenes"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Brandt Greenholt"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Agnes Cole"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Randall Johnston"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesGetOrgByName17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetOrgByName17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Albina Runte"

	q.GetOrgByName(ctx, name)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Zora Schulist", Description: "Martina Watsica", UpdatedAt: time.Time{}, ID: "Jolie Turcotte"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Sydney Konopelski", Description: "Kirstin Kohler", UpdatedAt: time.Time{}, ID: "Buford Ledner"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Ora Deckow", Description: "Ernest Cassin", UpdatedAt: time.Time{}, ID: "Virgil Batz"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Preston Klein", Description: "Enos Bahringer", UpdatedAt: time.Time{}, ID: "Lowell Dickens"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Abigale Hartmann", Description: "Doug Brown", UpdatedAt: time.Time{}, ID: "Howard Schamberger"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Alysa McKenzie", Description: "Rene Denesik", UpdatedAt: time.Time{}, ID: "Nigel Mayer"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Kayden Cole", Description: "Maximus Hagenes", UpdatedAt: time.Time{}, ID: "Pablo Parisian"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Keenan Heathcote", Description: "Clifton Witting", UpdatedAt: time.Time{}, ID: "Casandra Kuphal"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Emmanuelle Rosenbaum", Description: "Shyanne Jast", UpdatedAt: time.Time{}, ID: "Merlin Towne"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Gregory Bosco", Description: "Laura Bahringer", UpdatedAt: time.Time{}, ID: "Rylee Parisian"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Winston Frami", Description: "Octavia Cronin", UpdatedAt: time.Time{}, ID: "Gladyce Bernier"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Zachariah Terry", Description: "Kennith Satterfield", UpdatedAt: time.Time{}, ID: "Sydnie Robel"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Curt Runolfsdottir", Description: "Norene Olson", UpdatedAt: time.Time{}, ID: "Effie Schuster"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Heber Huels", Description: "Verna Olson", UpdatedAt: time.Time{}, ID: "Beulah Nicolas"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Marielle Johnston", Description: "Dahlia Terry", UpdatedAt: time.Time{}, ID: "Aryanna Baumbach"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Dejon Miller", Description: "Deon Mosciski", UpdatedAt: time.Time{}, ID: "Garland Mosciski"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Tad Turcotte", Description: "Paolo Abbott", UpdatedAt: time.Time{}, ID: "Manuel Borer"}

	q.UpdateOrg(ctx, arg)

}

func (s *OrgsSqlSuite) TestQueriesUpdateOrg17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateOrg17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateOrgParams{Name: "Mona Hoppe", Description: "Rowena Swift", UpdatedAt: time.Time{}, ID: "Waylon Shields"}

	q.UpdateOrg(ctx, arg)

}

func TestOrgsSqlSuite(t *testing.T) {
	suite.Run(t, new(OrgsSqlSuite))
}
