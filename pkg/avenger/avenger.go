package avenger

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-bongo/bongo"
	"github.com/stark-industries/config"
	"github.com/stark-industries/pkg/apis/dto"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
)

type Avenger struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id" json:"id"`
	Code               string `bson:"code" json:"code"`
	Name               string `bson:"name" json:"name"`
	ImageName          string `bson:"imageName" json:"imageName"`
}
func (avenger *Avenger) Save() error {
	err := config.Avenger.Save(avenger)
	if err != nil {
		return err
	}
	return nil
}
func (avenger *Avenger) FindById() Avenger {
	query := bson.M{"$and": []bson.M{
		{"id": avenger.Id},
	},
	}
	temp := Avenger{}
	config.Avenger.Find(query).Query.One(&temp)
	return temp
}

func (avenger *Avenger) FindAll() []dto.AvengerDto {
	query := bson.M{}
	avengers := []Avenger{}
	avengersDto := []dto.AvengerDto{}
	config.Avenger.Find(query).Query.All(&avengers)
	client,_:=config.GetObjectStorageClientConnection()
	for _,avenger:=range avengers {
		avengerDto:=dto.AvengerDto{
			Id:        avenger.Id,
			Code:      avenger.Code,
			Name:      avenger.Name,
			ImageName: avenger.ImageName,
			Image:nil,
		}
		output,err:=client.GetObject(&s3.GetObjectInput{
			Bucket: &config.S3Bucket,
			Key:    &avenger.ImageName,
		})
		if err==nil{
			body, readErr := ioutil.ReadAll(output.Body)
			if(readErr==nil){
				avengerDto.Image=body
				log.Println(body,"_______________")
			}
		}

		avengersDto = append(avengersDto,avengerDto)
	}
	return avengersDto
}