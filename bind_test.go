package binding

import (
	"net/http"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type V struct {
	Id     int    `json:"id"`
	Field  string `json:"field"`
	Arr    []int  `json:"arr"`
	Nested struct {
		InnerField string `json:"inner-field"`
	} `json:"nested"`
}

func (v *V) FieldMap() FieldMap {
	return FieldMap{
		&v.Id:  "id",
		&v.Arr: "arr",
	}
}
func TestBind(t *testing.T) {

	Convey("A request", t, func() {

		Convey("Without a Content-Type", func() {

			Convey("But with a query string", func() {

				Convey("Should invoke the Form deserializer", nil)

			})

			Convey("And without a query string", func() {

				Convey("Should yield an error", nil)

			})

		})

		Convey("With a form-urlencoded Content-Type", func() {

			Convey("Should invoke the Form deserializer", nil)

		})

		Convey("With a multipart/form-data Content-Type", func() {

			Convey("Should invoke the MultipartForm deserializer", nil)

		})

		Convey("With a json Content-Type", func() {

			v := V{}

			r, err := http.NewRequest("POST", "http://www.example.com", strings.NewReader(`{ "field": "value", "arr": [1,2,3], "nested": { "inner-field": "inner-value" }, "id": 1 }`))
			So(err, ShouldEqual, nil)
			r.Header.Add("Content-Type", "application/json; charset=utf-8")
			Bind(r, &v)

			So(v.Id, ShouldEqual, 1)
			t.Logf("v.Arr = %#v; wanted %#v", v.Arr, []int{1, 2, 3})
			So(len(v.Arr), ShouldEqual, 3)

			Convey("Should invoke the Json deserializer", nil)
		})

		Convey("With an unsupported Content-Type", func() {

			Convey("Should yield an error", nil)

		})

		Convey("Missing one required field", func() {

			Convey("A Required error should be produced", nil)

		})

		Convey("Missing multiple required fields", func() {

			Convey("As many Required errors should be produced", nil)

		})
	})

}
