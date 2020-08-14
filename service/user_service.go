package user

import (
	"github.com/gin-gonic/gin"

	"github.com/s1250040/go-yumemori/db"
	"github.com/s1250040/go-yumemori/entity"
)

// Service procides user's behavior
type Service struct{}

// User is alias of entity.User struct
type User entity.User

// entity
type Result entity.Result

//
type Sampling struct {
	SamplingTime string
}

type Samplings []Sampling

//

// GetAll is get all User
func (s Service) GetAll() ([]string, error) {
	db := db.GetDB()
	var result []Result
	// if err := db.Raw("select to_char(sampling_start_time, 'yyyy-mm-dd') as ResultDate from public.signal where \"FK_bsb_no\"=211 GROUP BY ResultDate").Scan(&result).Error; err != nil {
	if err := db.Table("public.signal").Select("sampling_start_time").Where("\"FK_bsb_no\" = ?", 211).Scan(&result).Error; err != nil {
		return nil, nil
	}
	// if err := db.Select("to_char(sampling_start_time, 'yyyy/mm/dd') as ResultDate").Where("FK_bsb_no = ? AND sampling_start_time >= ? AND sampling_start_time < ?", "218", "2018/11/1", "2018/11/30").Find(&u).Error; err != nil {
	// 	return u, err
	// }

	var temp []string

	for _, emp := range result {
		var sample = emp.SamplingStartTime.String()
		// var sample = string(emp.SamplingStartTime)
		temp = append(temp, sample)
	}

	return temp, nil
}

// CreateModel is create User model
func (s Service) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a User
func (s Service) GetByID(id string) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateByID is update a User
func (s Service) UpdateByID(id string, c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a User
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
