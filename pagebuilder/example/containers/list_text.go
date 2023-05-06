package containers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anshiii/playground-admin/media/media_library"
	"github.com/anshiii/playground-admin/pagebuilder"
	"github.com/anshiii/playground-admin/presets"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

type TextItem struct {
	Text  string
	Title string
	Image media_library.MediaBox `sql:"type:text;"`
}

type TextListItems []*TextItem
type TextList struct {
	ID       uint
	AnchorID string
	Title    string

	Items TextListItems `sql:"type:text;"`
	Text  string
}

func (*TextList) TableName() string {
	return "text_list_content_but_with_image"
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

	itemField := pageBuilder.GetPresetsBuilder().NewFieldsBuilder(presets.WRITE).Model(&TextItem{}).Only("Text", "Title", "Image")

	editable.Field("Items").Nested(itemField, &presets.DisplayFieldInSorter{Field: "Text"})
}

func ListBody(props *TextList, input *pagebuilder.RenderInput) HTMLComponent {
	var wrapEle = Ul()
	for _, item := range props.Items {
		var contentEle = ListItemBody(item)
		wrapEle.AppendChildren(contentEle)
	}
	var body = ContainerWrapper(fmt.Sprintf(inflection.Plural(strcase.ToKebab("TextList"))+"_%v", props.ID), props.AnchorID,
		"container-list_content container-lottie", "", "", "", "", true, true, true, "",
		wrapEle)
	return body
}

func ListItemBody(item *TextItem) HTMLComponent {
	var itemEle HTMLComponent = Li(
		Img(item.Image.Url),
		P().Text(item.Title).Class("text-list-item-title"),
		Span(item.Text).Class("text-list-item-text"),
	)
	return itemEle
}
