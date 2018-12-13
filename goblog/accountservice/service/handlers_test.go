package service

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/cloudenterprise/goblog/accountservice/dbclient"
	"github.com/cloudenterprise/goblog/accountservice/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAccountWrongPath(t *testing.T) {

	Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetAccount(t *testing.T) {

	// Create a mock instance of our boltDB
	mockRepo := &dbclient.MockBoltClient{}

	// Mock 1: For "123" as input, return a proper Account struct and nil error
	mockRepo.On("QueryAccount", "123").Return(model.Account{ID: "123", Name: "Person_123"}, nil)

	// Mock 2: For "456" as input, return an empty Account struct and actual error
	mockRepo.On("QueryAccount", "456").Return(model.Account{}, fmt.Errorf("Some error"))

	// Set package global DBClient var to mockRepo
	DBClient = mockRepo

	// Mock 1 Test
	Convey("Given a HTTP request for /accounts/123", t, func() {
		req := httptest.NewRequest("GET", "/accounts/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)

				account := model.Account{}
				json.Unmarshal(resp.Body.Bytes(), &account)
				So(account.ID, ShouldEqual, "123")
				So(account.Name, ShouldEqual, "Person_123")
			})
		})
	})

	// Mock 2 Test
	Convey("Given a HTTP request for /account/456", t, func() {
		req := httptest.NewRequest("GET", "/accounts/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}
