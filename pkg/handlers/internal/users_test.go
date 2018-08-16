package internal

import (
	"net/http/httptest"

	userop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/users"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *utils.HandlerSuite) TestUnknownLoggedInUserHandler() {
	unknownUser := testdatagen.MakeDefaultUser(suite.db)

	req := httptest.NewRequest("GET", "/users/logged_in", nil)
	req = suite.authenticateUserRequest(req, unknownUser)

	params := userop.ShowLoggedInUserParams{
		HTTPRequest: req,
	}

	handler := ShowLoggedInUserHandler(NewHandlerContext(suite.db, suite.logger))

	response := handler.Handle(params)

	okResponse, ok := response.(*userop.ShowLoggedInUserOK)
	suite.True(ok)
	suite.Equal(okResponse.Payload.ID.String(), unknownUser.ID.String())
}

func (suite *utils.HandlerSuite) TestServiceMemberLoggedInUserHandler() {
	firstName := "Joseph"
	sm := testdatagen.MakeExtendedServiceMember(suite.db, testdatagen.Assertions{
		ServiceMember: models.ServiceMember{
			FirstName: &firstName,
		},
	})

	req := httptest.NewRequest("GET", "/users/logged_in", nil)
	req = suite.authenticateRequest(req, sm)

	params := userop.ShowLoggedInUserParams{
		HTTPRequest: req,
	}

	handler := ShowLoggedInUserHandler(NewHandlerContext(suite.db, suite.logger))

	response := handler.Handle(params)

	okResponse, ok := response.(*userop.ShowLoggedInUserOK)
	suite.True(ok)
	suite.Equal(okResponse.Payload.ID.String(), sm.UserID.String())
	suite.Equal("Joseph", *okResponse.Payload.ServiceMember.FirstName)
}

func (suite *utils.HandlerSuite) TestServiceMemberNoTransportationOfficeLoggedInUserHandler() {
	firstName := "Joseph"
	sm := testdatagen.MakeExtendedServiceMember(suite.db, testdatagen.Assertions{
		ServiceMember: models.ServiceMember{
			FirstName: &firstName,
		},
	})

	// Remove transportation office info from current station
	station := sm.DutyStation
	station.TransportationOfficeID = nil
	suite.mustSave(&station)

	req := httptest.NewRequest("GET", "/users/logged_in", nil)
	req = suite.authenticateRequest(req, sm)

	params := userop.ShowLoggedInUserParams{
		HTTPRequest: req,
	}

	handler := ShowLoggedInUserHandler(NewHandlerContext(suite.db, suite.logger))

	response := handler.Handle(params)

	okResponse, ok := response.(*userop.ShowLoggedInUserOK)
	suite.True(ok)
	suite.Equal(okResponse.Payload.ID.String(), sm.UserID.String())
}
