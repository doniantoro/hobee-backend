package seeder

import (
	"hobee/config"
	"hobee/model"
)

func Category() {
	var db = config.ConnectionDatabase()
	var arr = [5][2]string{{"futsal", "/assets/icon/futsal.svg"}, 
								{"badminton", "/assets/icon/badminton.svg"}, 
								{"volley", "/assets/icon/volley.svg"},
								{"basketball", "/assets/icon/basketball.svg"},
								{"zumba", "/assets/icon/zumba.svg"},
							}
	var i int
	for i = 0 ; i < 5 ; i++ {

		category := model.Category{}
		category.Name = arr[i][0]
		category.Icon =arr[i][1]	
	
		err := db.Create(&category).Error
		if err != nil {
			panic(err)
		}
	}


	

	//  return err

}