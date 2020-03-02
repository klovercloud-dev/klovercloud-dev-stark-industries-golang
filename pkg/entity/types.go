package entity

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-bongo/bongo"
	"github.com/stark-industries/config"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"

)

type Staff struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id" json:"id"`
	Code               string `bson:"code" json:"code"`
	Name               string `bson:"name" json:"name"`
	Position           string `bson:"position" json:"position"`
	ImageName          string `bson:"imageName" json:"imageName"`
	Image []byte `bson:"image" json:"image"`
}
func (staff *Staff) Save() error {
	err := config.Staff.Save(staff)
	if err != nil {
		return err
	}
	return nil
}
func (staff *Staff) FindById() Staff {
	query := bson.M{"$and": []bson.M{
		{"id": staff.Id},
	},
	}
	temp := Staff{}
	config.Staff.Find(query).Query.One(&temp)
	return temp
}

func (staff *Staff) FindAll() []Staff {
	query := bson.M{}
	tempStaffs := []Staff{}
	config.Staff.Find(query).Query.All(&tempStaffs)
	client,_:=config.GetObjectStorageClientConnection()
	for index,staff:=range tempStaffs{
		output,err:=client.GetObject(&s3.GetObjectInput{
			Bucket: &config.S3Bucket,
			Key:    &staff.ImageName,
		})
		if err==nil{
			body, readErr := ioutil.ReadAll(output.Body)
			if(readErr==nil){
				tempStaffs[index].Image=body
			}
		}


	}

	return tempStaffs
}