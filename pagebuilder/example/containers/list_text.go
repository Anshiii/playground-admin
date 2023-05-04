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

type TextList struct {
	ID       uint
	anchorID string
	title    string
	items    []*TextListItem
}

type TextListItem struct {
	text  string
	title string
}

func RegisterTextListContainer(pageBuilder *pagebuilder.Builder, db *gorm.DB) {
	container := pageBuilder.RegisterContainer("TextList").
		RenderFunc(func(obj interface{}, input *pagebuilder.RenderInput, ctx *web.EventContext) HTMLComponent {
			props := obj.(*TextList)
			return ListBody(props, input)
		})
	model := container.Model(&TextList{})
	editable := model.Editing("anchorID", "title", "items")

	itemField := pageBuilder.GetPresetsBuilder().NewFieldsBuilder(presets.WRITE).Model(&TextListItem{}).Only("text", "title")

	editable.Field("items").Nested(itemField)
}

func ListBody(props *TextList, input *pagebuilder.RenderInput) HTMLComponent {
	var body = ContainerWrapper(fmt.Sprintf(inflection.Plural(strcase.ToKebab("ListBody"))+"_%v", props.ID), props.anchorID,
		"container-list_content container-lottie", "", "", "", "", true, true, true, "",
		Ul(
			ListItemBody(props.items, input),
		))
	return body
}

func ListItemBody(items []*TextListItem, input *pagebuilder.RenderInput) HTMLComponent {
	var itemsWrap *HTMLTagBuilder = Li().Class("container-list_content-grid")
	for _, item := range items {
		var itemEle HTMLComponent = Div(
			Div().Content(item.title),
			Div().Content(item.text),
		)
		itemsWrap.AppendChildren(itemEle)
	}
	return itemsWrap
}
