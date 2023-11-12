// Coverage template
package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type UsersSqlSuite struct {
	suite.Suite
}

func (s *UsersSqlSuite) TestQueriesCreateUser0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Yasmine Macejkovic", Name: "Jacques Hoeger", Email: "Melyna Bergnaum", Password: "Reece Nienow", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Gage Leannon", Name: "Madeline Schumm", Email: "Milford Marquardt", Password: "Tyrique Paucek", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Donnie Thompson", Name: "Bennie Dibbert", Email: "Misael Nikolaus", Password: "Piper Ledner", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Agustina Dickens", Name: "Vance Glover", Email: "Gabriella Purdy", Password: "Edyth Witting", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Pietro Schumm", Name: "Robyn Zboncak", Email: "Kenyon McGlynn", Password: "Amie Nader", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Kaia Gislason", Name: "Dejah Kerluke", Email: "Selmer Trantow", Password: "Alexandrine Padberg", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Leonie Jenkins", Name: "Dorcas Bergnaum", Email: "Mariam Towne", Password: "Brady Hauck", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Mary Cormier", Name: "Marjolaine Waelchi", Email: "Geovanny Frami", Password: "Eleazar Schneider", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Jarrod Stracke", Name: "Naomi Hilll", Email: "Lenora Predovic", Password: "Reuben Stracke", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Una Smith", Name: "Kameron Rau", Email: "Buck Senger", Password: "Reece Breitenberg", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Jerry Goldner", Name: "Neha Weimann", Email: "Florian Williamson", Password: "Marianne Borer", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Sanford Rogahn", Name: "Marian Nikolaus", Email: "Marcella Murray", Password: "Talon Brakus", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Mae Keebler", Name: "Kameron Legros", Email: "Vesta Legros", Password: "Frida Schimmel", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Vernie Marvin", Name: "Eleanore Smitham", Email: "Ernesto Lakin", Password: "Ora Will", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Ceasar Wolff", Name: "Nelson Leffler", Email: "Mayra Steuber", Password: "Dannie Deckow", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Rachelle Denesik", Name: "Richie Powlowski", Email: "Sylvester Morissette", Password: "Blake Wisoky", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Damien Boyle", Name: "Jefferey White", Email: "Mandy Connelly", Password: "Wilson Lockman", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesCreateUser17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesCreateUser17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := CreateUserParams{ID: "Leopold Schimmel", Name: "Emiliano Stroman", Email: "Xzavier Corwin", Password: "Romaine Grimes", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	q.CreateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Yolanda Strosin"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Jillian Flatley"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Pierce Erdman"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Elody Ziemann"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Hilbert Hilpert"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Demetris Kuphal"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Palma Graham"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Jamir Powlowski"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Ariel Gutkowski"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Geovanni Schuster"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Jerel Collier"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lowell Kunde"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Blake Bosco"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Allie Johnston"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Madonna Willms"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Kaylee Cassin"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Cade Stark"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesDeleteUser17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesDeleteUser17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Murphy Murray"

	q.DeleteUser(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetAllUsers17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetAllUsers17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()

	q.GetAllUsers(ctx)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Rosalyn Daugherty"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Tatum Hettinger"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Noble Gusikowski"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Edgar Heller"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Candice Orn"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Mariela Nitzsche"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Chaim Gorczany"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Deondre Sawayn"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Shanny Batz"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Virgil Ward"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Zora Schmidt"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Larissa Metz"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Valentine Boehm"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Kali Pouros"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Lou Rowe"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Hipolito Ziemann"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Jake Beahan"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByID17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByID17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	id := "Felicia Dickens"

	q.GetUserByID(ctx, id)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jessyca Botsford"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Aileen Schiller"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Litzy Jewess"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Rex Okuneva"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Daron Johnson"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jaycee Cummings"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Darron Bartell"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Roma Gleichner"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Mike Beier"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Kennith McLaughlin"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Zaria Kulas"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Noel Haley"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Serena Jacobs"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Jan Emmerich"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Lois Wisoky"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Freida Flatley"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Aliyah Schaefer"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesGetUserByName17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesGetUserByName17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	name := "Ari Schmitt"

	q.GetUserByName(ctx, name)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser0() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser0", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Ted Borer", Email: "Britney Stamm", Password: "Tavares Corkery", UpdatedAt: time.Time{}, ID: "Antonetta Franecki"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser1", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Estefania Kiehn", Email: "Myrtle Adams", Password: "Ned Ebert", UpdatedAt: time.Time{}, ID: "Kari Auer"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser2", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Millie Braun", Email: "Arthur Deckow", Password: "Arnold Schowalter", UpdatedAt: time.Time{}, ID: "Vilma Trantow"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser3", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Fiona Vandervort", Email: "Mariana Considine", Password: "Manuela Yost", UpdatedAt: time.Time{}, ID: "Mortimer Cassin"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser4() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser4", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Orie Miller", Email: "Eldridge Goodwin", Password: "Ezra Harris", UpdatedAt: time.Time{}, ID: "Heloise Witting"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser5", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Triston Bernier", Email: "Yasmeen Balistreri", Password: "Ezekiel Boyer", UpdatedAt: time.Time{}, ID: "Mabel Botsford"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser6", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Abigail Bernier", Email: "Torrance Thompson", Password: "Marcia Reynolds", UpdatedAt: time.Time{}, ID: "Jedediah Ziemann"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser7() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser7", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Weldon Huel", Email: "Felicity Jewess", Password: "Chadrick Walker", UpdatedAt: time.Time{}, ID: "Cristopher King"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser8() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser8", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Colt Kerluke", Email: "Helene Towne", Password: "Brooklyn Gleichner", UpdatedAt: time.Time{}, ID: "Chadd Kreiger"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser9() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser9", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Christa Stark", Email: "Colin Bartoletti", Password: "Pasquale Bartoletti", UpdatedAt: time.Time{}, ID: "Abagail Ziemann"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser10() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser10", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Greg Watsica", Email: "Destinee Sporer", Password: "Ezekiel Conn", UpdatedAt: time.Time{}, ID: "Luther West"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser11", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Bridget Schumm", Email: "Meda Zulauf", Password: "Bo Kuhn", UpdatedAt: time.Time{}, ID: "Adeline Orn"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser12() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser12", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Vince Wolff", Email: "Dovie Schulist", Password: "Reyes Kirlin", UpdatedAt: time.Time{}, ID: "Dejon Konopelski"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser13() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser13", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Tomasa Conn", Email: "Derek Metz", Password: "Shaniya Nikolaus", UpdatedAt: time.Time{}, ID: "Ezequiel Hintz"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser14", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Alejandrin Ankunding", Email: "Josefa White", Password: "Coy Monahan", UpdatedAt: time.Time{}, ID: "Kirsten Denesik"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser15() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser15", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Layla Greenfelder", Email: "Preston Anderson", Password: "Alessandra Rutherford", UpdatedAt: time.Time{}, ID: "Dino Eichmann"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser16() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser16", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Grace Leannon", Email: "Bruce McClure", Password: "Audra Adams", UpdatedAt: time.Time{}, ID: "Rafael Murazik"}

	q.UpdateUser(ctx, arg)

}

func (s *UsersSqlSuite) TestQueriesUpdateUser17() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in TestQueriesUpdateUser17", r)
		}
	}()

	pointerQ := Queries{db: func() DBTX {
		return nil
	}()}
	q := &pointerQ
	ctx := func() context.Context {
		return nil
	}()
	arg := UpdateUserParams{Name: "Clemens Simonis", Email: "Nicolas Abernathy", Password: "Miracle Gaylord", UpdatedAt: time.Time{}, ID: "Gerhard Crist"}

	q.UpdateUser(ctx, arg)

}

func TestUsersSqlSuite(t *testing.T) {
	suite.Run(t, new(UsersSqlSuite))
}
