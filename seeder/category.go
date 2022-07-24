package seeder

import (
	"hobee/config"
	"hobee/model"
	"fmt"
)

func Category() {
	var db = config.ConnectionDatabase()
	
	var numbers2 = [5][2]string{{"futsal", "/assets/icon/futsal.svg"}, 
								{"badminton", "/assets/icon/badminton.svg"}, 
								{"volley", "/assets/icon/volley.svg"},
								{"basketball", "/assets/icon/basketball.svg"},
								{"zumba", "/assets/icon/zumba.svg"},
							}
	var i int
	fmt.Println(numbers2)
	for i = 0 ; i < 5 ; i++ {

		category := model.Category{}
		category.Name = numbers2[i][0]
		category.Icon =numbers2[i][1]	
	
		err := db.Create(&category).Error
		if err != nil {
			panic(err)
		}
	}


	

	//  return err

}