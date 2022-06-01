package example

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/qor/oss/s3"
	"github.com/qor/qor5/media/oss"
	media_view "github.com/qor/qor5/media/views"
	"github.com/qor/qor5/pagebuilder"
	"github.com/qor/qor5/pagebuilder/containers"
	"github.com/qor/qor5/richeditor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Logger = db.Logger.LogMode(logger.Info)
	return
}

func ConfigPageBuilder(db *gorm.DB) *pagebuilder.Builder {
	sess := session.Must(session.NewSession())

	oss.Storage = s3.New(&s3.Config{
		Bucket:  os.Getenv("S3_Bucket"),
		Region:  os.Getenv("S3_Region"),
		Session: sess,
	})

	err := db.AutoMigrate(
		&containers.WebHeader{},
		&containers.WebFooter{},
		&containers.VideoBanner{},
		&containers.Heading{},
	)
	if err != nil {
		panic(err)
	}
	pb := pagebuilder.New(db)

	media_view.Configure(pb.GetPresetsBuilder(), db)

	richeditor.Plugins = []string{"alignment", "table", "video", "imageinsert"}
	pb.GetPresetsBuilder().ExtraAsset("/redactor.js", "text/javascript", richeditor.JSComponentsPack())
	pb.GetPresetsBuilder().ExtraAsset("/redactor.css", "text/css", richeditor.CSSComponentsPack())

	containers.RegisterHeader(pb)
	containers.RegisterFooter(pb)
	containers.RegisterVideoBannerContainer(pb)
	containers.RegisterHeadingContainer(pb, db)
	return pb
}
