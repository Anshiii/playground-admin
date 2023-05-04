package containers

import (
	"fmt"

	"github.com/anshiii/playground-admin/pagebuilder"
	"github.com/anshiii/playground-admin/presets"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/qor5/ui/vuetify"
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
	container := pageBuilder.RegisterContainer("Demo-Container Name").
		RenderFunc(func(obj interface{}, input *pagebuilder.RenderInput, ctx *web.EventContext) HTMLComponent {
			headerDemoContainer := obj.(*HeaderDemoContainer)
			return HeaderDemoContainerBody(headerDemoContainer, input)
		})

	container.Model(&HeaderDemoContainer{}).URIName(inflection.Plural(strcase.ToKebab("Demo-Container Name"))).Editing("title", "random").Field("random").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
		return vuetify.VSelect().
			Items([]bool{true, false}).
			Value(field.Value(obj)).
			Label(field.Label).
			FieldName(field.FormKey)
	})

}

func HeaderDemoContainerBody(data *HeaderDemoContainer, input *pagebuilder.RenderInput) (body HTMLComponent) {
	body = ContainerWrapper(
		fmt.Sprintf(inflection.Plural(strcase.ToKebab("Image"))+"_%v", data.ID), data.title, "container-title",
		data.title, data.title, "",
		"", data.random, data.random, input.IsEditor, "",
		Div(
			Div().Content("hello world").Class("container-image-corner"),
		).Class("container-wrapper"),
	)
	return body
}
