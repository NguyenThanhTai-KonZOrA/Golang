package seed

import (
	"log"

	"github.com/NguyenThanhTai-KonZOrA/Golang/src/models"
	"github.com/jinzhu/gorm"
)

var member = []models.member{
	models.member{
		FullName: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.member{
		FullName: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var book = []models.book{
	models.book{
		Name:   "Title 1",
		Author: "Hello world 1",
		Publication: "ASASAS"
	},8
	models.book{
		Name:   "Title 2",
		Author: "Hello world 2",
		Publication: "ASASAS"
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
