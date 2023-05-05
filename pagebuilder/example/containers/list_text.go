package containers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anshiii/playground-admin/pagebuilder"
	"github.com/anshiii/playground-admin/presets"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

type TextListItem struct {
	Text  string
	Title string
}

type TextListItems []*TextListItem
type TextList struct {
	ID       uint
	AnchorID string
	Title    string

	Items TextListItems `sql:"type:text;"`
	Text  string
}

func (*TextList) TableName() string {
	return "text_list_content"
}

func (this TextListItems) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *TextListItems) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		return json.Unmarshal([]byte(v), this)
	case []byte:
		return json.Unmarshal(v, this)
	default:
		return errors.New("not supported")
	}
}

func RegisterTextListContainer(pageBuilder *pagebuilder.Builder, db *gorm.DB) {
	container := pageBuilder.RegisterContainer("TextList").
		RenderFunc(func(obj interface{}, input *pagebuilder.RenderInput, ctx *web.EventContext) HTMLComponent {
			props := obj.(*TextList)
			return ListBody(props, input)
		})
	model := container.Model(&TextList{})
	editable := model.Editing("AnchorID", "Title", "Items")

	itemField := pageBuilder.GetPresetsBuilder().NewFieldsBuilder(presets.WRITE).Model(&TextListItem{}).Only("Text", "Title")

	editable.Field("Items").Nested(itemField, &presets.DisplayFieldInSorter{Field: "Text"})
}

func ListBody(props *TextList, input *pagebuilder.RenderInput) HTMLComponent {
	var body = ContainerWrapper(fmt.Sprintf(inflection.Plural(strcase.ToKebab("TextList"))+"_%v", props.ID), props.AnchorID,
		"container-list_content container-lottie", "", "", "", "", true, true, true, "",
		Ul(
			ListItemBody(props.Items, input),
		))
	return body
}

func ListItemBody(items []*TextListItem, input *pagebuilder.RenderInput) HTMLComponent {
	var itemsWrap *HTMLTagBuilder = Li().Class("container-list_content-grid")
	for _, item := range items {
		var itemEle HTMLComponent = Div(
			P().Content(item.Title),
			Span(item.Text),
		)
		itemsWrap.AppendChildren(itemEle)
	}
	return itemsWrap
}
