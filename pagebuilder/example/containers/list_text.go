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
	Id       uint
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
	// case []byte:
	// 	return json.Unmarshal(v, this)
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
	editable := model.Editing("anchorID", "title", "items")

	itemField := pageBuilder.GetPresetsBuilder().NewFieldsBuilder(presets.WRITE).Model(&TextListItem{}).Only("text", "title")

	editable.Field("items").Nested(itemField)
}

func ListBody(props *TextList, input *pagebuilder.RenderInput) HTMLComponent {
	var body = ContainerWrapper(fmt.Sprintf(inflection.Plural(strcase.ToKebab("ListBody"))+"_%v", props.Id), props.AnchorID,
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
			Div().Content(item.Title),
			Div().Content(item.Text),
		)
		itemsWrap.AppendChildren(itemEle)
	}
	return itemsWrap
}
