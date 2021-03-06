package model

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

var (
	COL_PLATFORM = "platforms"
)

// 服务器
type Platform struct {
	Id          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Desc        string        `bson:"desc"`
	Mask        string        `bson:"mask"`
	DownloadUrl string        `bson:"download_url"`
	UpdatedAt   time.Time     `bson:"updated_at"`
	CreatedAt   time.Time     `bson:"created_at"`

	Servers []Server // 关联的服务器集合
}

func ColPlatform(s *mgo.Session) *mgo.Collection {
	return s.DB(DB_NAME).C(COL_PLATFORM)
}

// 查询所有平台
func FindAllPlatform() []Platform {
	s := Session()
	defer s.Close()
	colPlatform := ColPlatform(s)
	var platforms []Platform
	colPlatform.Find(nil).All(&platforms)
	return platforms
}

//通过id查询平台
func FindPlatform(id string) *Platform {
	revel.INFO.Printf("id = %s \n", id)
	s := Session()
	defer s.Close()
	colPlatform := ColPlatform(s)
	objectId := bson.ObjectIdHex(id)
	platform := &Platform{}
	colPlatform.Find(bson.M{"_id": objectId}).One(platform)
	colServer := ColServer(s)
	colServer.Find(bson.M{"platform_id": objectId}).All(&platform.Servers)
	return platform
}
