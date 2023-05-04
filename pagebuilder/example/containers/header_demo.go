package containers

import (
	"fmt"

	"github.com/anshiii/playground-admin/pagebuilder"
	"github.com/anshiii/playground-admin/presets"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

type HeaderDemoContainer struct {
	title  string
	ID     uint
	random bool
}

func RegisterHeaderDemoContainer(pageBuilder *pagebuilder.Builder, db *gorm.DB) {
	vb := pageBuilder.RegisterContainer("Image").RenderFunc(func(obj interface{}, input *pagebuilder.RenderInput, ctx *web.EventContext) HTMLComponent {
		v := obj.(*ImageContainer)
		return ImageContainerBody(v, input)
	})
	mb := vb.Model(&ImageContainer{}).URIName(inflection.Plural(strcase.ToKebab("Image")))
	eb := mb.Editing("Random", "Random", "AnchorID", "BackgroundColor", "TransitionBackgroundColor", "Image")
	eb.Field("BackgroundColor").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
		return Div(P().Content("hello demo?")).Class("header_demo_car")
	})

}

func HeaderDemoContainerBody(data *HeaderDemoContainer, input *pagebuilder.RenderInput) (body HTMLComponent) {
	body = ContainerWrapper(
		fmt.Sprintf(inflection.Plural(strcase.ToKebab("Image"))+"_%v", data.ID), data.title, "container-title",
		data.title, data.title, "",
		"", data.random, data.random, input.IsEditor, "",
		Div(
			Div().Class("container-image-corner"),
		).Class("container-wrapper"),
	)
	return body
}
