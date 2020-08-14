package user

import (
	"strings"

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
	SamplingTime string `json:"start"`
}

// type Samplings []Sampling

//

// GetAll is get all User
func (s Service) GetAll() ([]Sampling, error) {
	db := db.GetDB()
	var result []Result
	// if err := db.Raw("select to_char(sampling_start_time, 'yyyy-mm-dd') as ResultDate from public.signal where \"FK_bsb_no\"=211 GROUP BY ResultDate").Scan(&result).Error; err != nil {
	if err := db.Table("public.signal").Select("sampling_start_time").Where("\"FK_bsb_no\" = ? AND sampling_start_time >= ? AND sampling_start_time < ?", 211, "2018/11/1", "2018/11/30").Scan(&result).Error; err != nil {
		return nil, nil
	}
	// if err := db.Select("to_char(sampling_start_time, 'yyyy/mm/dd') as ResultDate").Where("FK_bsb_no = ? AND sampling_start_time >= ? AND sampling_start_time < ?", "218", "2018/11/1", "2018/11/30").Find(&u).Error; err != nil {
	// 	return u, err
	// }

	var temp []string

	for _, emp := range result {
		var sample = emp.SamplingStartTime.String()
		// var sample = string(emp.SamplingStartTime)
		slice := strings.Split(sample, " ")
		temp = append(temp, slice[0])

	}

	m := make(map[string]struct{})

	newList := make([]string, 0)

	for _, element := range temp {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; !ok {
			m[element] = struct{}{}
			newList = append(newList, element)
		}
	}

	var samplingdata []Sampling
	for _, sdata := range newList {

		var sampling Sampling
		sampling.SamplingTime = sdata
		samplingdata = append(samplingdata, sampling)
	}

	return samplingdata, nil
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
