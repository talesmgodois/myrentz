package acceptance

import (
	"backend/adapters/persistence"
	"backend/config"
	"testing"
)

func TestRepository(t *testing.T) {
	_ = config.GetConfig()
	persistence.ConnectGorm()
	db := persistence.GetDbGorm()
	t.Run("Add item on db", func(t *testing.T) {

		var r = db.Create(&persistence.User{
			RegularPerson: persistence.RegularPerson{
				Name: "Tales",
			},
		})
		if r.Error != nil {
			t.Errorf("Unexpected error %v", r.Error.Error())
		}
		//var userPayload = persistence.User{
		//
		//}
		//var users = persistence.User{}
		//c, err := repo.Create()
		//
		//if err != nil {
		//	t.Errorf("Unexpected error %v", err.Error())
		//}
		//
		//if c == nil {
		//	t.Errorf("c is not supposed to be nil")
		//}
	})
}
