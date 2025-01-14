package pages

import (
	"log"

	"github.com/anshiii/playground-admin/media"
	"github.com/anshiii/playground-admin/media/media_library"
	media_view "github.com/anshiii/playground-admin/media/views"
	"github.com/anshiii/playground-admin/richeditor"
	"github.com/qor5/ui/cropper"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

func Settings(db *gorm.DB) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.PageTitle = "Settings"

		r.Body = h.Div(
			VContainer(

				VRow(
					VCol(
						h.H1("Example of use QMediaBox in any page").Class("text-h5 pt-4 pl-2"),
						media_view.QMediaBox(db).
							FieldName("test").
							Value(&media_library.MediaBox{}).
							Config(&media_library.MediaBoxConfig{
								AllowType: "image",
								Sizes: map[string]*media.Size{
									"thumb": {
										Width:  400,
										Height: 300,
									},
									"main": {
										Width:  800,
										Height: 500,
									},
								},
							}),
					).Cols(6),
				),

				VRow(
					VCol(
						richeditor.RichEditor(db, "Body").Plugins([]string{"alignment", "video", "table", "imageinsert"}).Value(`<p>Could you do an actual logo instead of a font I cant pay you? Can we try some other colors maybe? I cant pay you. You might wanna give it another shot, so make it pop and this is just a 5 minutes job the target audience makes and families aged zero and up will royalties in the company do instead of cash.</p>
						<p>Jazz it up a little I was wondering if my cat could be placed over the logo in the flyer I have printed it out, but the animated gif is not moving I have printed it out, but the animated gif is not moving make it original. Can you make it stand out more? Make it original.</p>`).Label("Body").Placeholder("Place Holder"),
					),
				),

				VRow(
					cropper.Cropper().
						Src("https://agontuk.github.io/assets/images/berserk.jpg").
						Value(cropper.Value{X: 1141, Y: 540, Width: 713, Height: 466}).
						AspectRatio(713, 466).
						Attr("@input", web.Plaid().
							FieldValue("CropperEvent", web.Var("JSON.stringify($event)")).EventFunc(LogInfoEvent).Go()),
				),
			).Fluid(true),
		)
		return
	}
}

const LogInfoEvent = "logInfo"

func LogInfo(ctx *web.EventContext) (r web.EventResponse, err error) {
	log.Println("CropperEvent", ctx.R.FormValue("CropperEvent"))
	return
}
