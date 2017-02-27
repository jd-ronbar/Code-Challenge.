package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "hello/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"bytes"
	"net/url"
	"strconv"
	"database/sql"
	"log"
)

func init() {
	truncateDb()
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func truncateDb() {
	db, err := sql.Open("mysql",
		"root:1234509876@tcp(127.0.0.1:3306)/hello")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("TRUNCATE TABLE  driver;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	var data = url.Values{}
	data.Set("Id", "4231")
	data.Set("Name", "Jhon Doe")
	data.Set("License_number", "3485404")
	r, _ := http.NewRequest("POST", "/drivers/add", bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test /drivers/add Endpoint\n", t, func() {
		Convey("Status Code Should Be 302", func() {
			So(w.Code, ShouldEqual, 302)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeZeroValue)
		})
	})

	r, _ = http.NewRequest("GET", "/drivers", nil)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test /drivers Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("Recently added driver should be listed in the resutl", func() {
			r, _ := regexp.Compile(".*Jhon Doe.*")
			So(r.MatchString(w.Body.String()), ShouldBeTrue)
		})


	})
}

