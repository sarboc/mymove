package internal

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/auth"
	userop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/users"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/handlers/utils"
	"github.com/transcom/mymove/pkg/models"
)

// ShowLoggedInUserHandler returns the logged in user
type ShowLoggedInUserHandler utils.HandlerContext

// Handle returns the logged in user
func (h ShowLoggedInUserHandler) Handle(params userop.ShowLoggedInUserParams) middleware.Responder {
	session := auth.SessionFromRequestContext(params.HTTPRequest)

	if !session.IsServiceMember() {
		userPayload := internalmessages.LoggedInUserPayload{
			ID: utils.FmtUUID(session.UserID),
		}
		return userop.NewShowLoggedInUserOK().WithPayload(&userPayload)
	}
	// Load Servicemember and first level associations
	serviceMember, err := models.FetchServiceMemberForUser(h.Db, session, session.ServiceMemberID)
	if err != nil {
		h.Logger.Error("Error retrieving service_member", zap.Error(err))
		return userop.NewShowLoggedInUserUnauthorized()
	}

	// Load duty station and transportation office association
	if serviceMember.DutyStationID != nil {
		// Fetch associations on duty station
		dutyStation, err := models.FetchDutyStation(h.Db, *serviceMember.DutyStationID)
		if err != nil {
			return utils.ResponseForError(h.Logger, err)
		}
		serviceMember.DutyStation = dutyStation

		// Fetch duty station transportation office
		transportationOffice, err := models.FetchDutyStationTransportationOffice(h.Db, *serviceMember.DutyStationID)
		if err != nil {
			if errors.Cause(err) != models.ErrFetchNotFound {
				// The absence of an office shouldn't render the entire request a 404
				return utils.ResponseForError(h.Logger, err)
			}
			// We might not have Transportation Office data for a Duty Station, and that's ok
			if errors.Cause(err) != models.ErrFetchNotFound {
				return utils.ResponseForError(h.Logger, err)
			}
		}
		serviceMember.DutyStation.TransportationOffice = transportationOffice
	}

	// Load the latest orders associations and new duty station transport office
	if len(serviceMember.Orders) > 0 {
		orders, err := models.FetchOrderForUser(h.Db, session, serviceMember.Orders[0].ID)
		if err != nil {
			return utils.ResponseForError(h.Logger, err)
		}
		serviceMember.Orders[0] = orders

		newDutyStationTransportationOffice, err := models.FetchDutyStationTransportationOffice(h.Db, orders.NewDutyStationID)
		if err != nil {
			if errors.Cause(err) != models.ErrFetchNotFound {
				// The absence of an office shouldn't render the entire request a 404
				return utils.ResponseForError(h.Logger, err)
			}
			// We might not have Transportation Office data for a Duty Station, and that's ok
			if errors.Cause(err) != models.ErrFetchNotFound {
				return utils.ResponseForError(h.Logger, err)
			}
		}
		serviceMember.Orders[0].NewDutyStation.TransportationOffice = newDutyStationTransportationOffice

		// Load associations on PPM if they exist
		if len(serviceMember.Orders[0].Moves) > 0 {
			if len(serviceMember.Orders[0].Moves[0].PersonallyProcuredMoves) > 0 {
				// TODO: load advances on all ppms for the latest order's move
				ppm, err := models.FetchPersonallyProcuredMove(h.Db, session, serviceMember.Orders[0].Moves[0].PersonallyProcuredMoves[0].ID)
				if err != nil {
					return utils.ResponseForError(h.Logger, err)
				}
				serviceMember.Orders[0].Moves[0].PersonallyProcuredMoves[0].Advance = ppm.Advance
			}
		}
	}

	userPayload := internalmessages.LoggedInUserPayload{
		ID:            utils.FmtUUID(session.UserID),
		ServiceMember: payloadForServiceMemberModel(h.Storage, serviceMember),
	}
	return userop.NewShowLoggedInUserOK().WithPayload(&userPayload)
}
