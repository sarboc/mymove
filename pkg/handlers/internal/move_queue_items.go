package internal

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/auth"
	queueop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/queues"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/handlers/utils"
	"github.com/transcom/mymove/pkg/models"
	"go.uber.org/zap"
)

func payloadForMoveQueueItem(MoveQueueItem models.MoveQueueItem) *internalmessages.MoveQueueItem {
	MoveQueueItemPayload := internalmessages.MoveQueueItem{
		ID:               utils.FmtUUID(MoveQueueItem.ID),
		CreatedAt:        utils.FmtDateTime(MoveQueueItem.CreatedAt),
		Edipi:            swag.String(MoveQueueItem.Edipi),
		Rank:             MoveQueueItem.Rank,
		CustomerName:     swag.String(MoveQueueItem.CustomerName),
		Locator:          swag.String(MoveQueueItem.Locator),
		Status:           swag.String(MoveQueueItem.Status),
		PpmStatus:        MoveQueueItem.PpmStatus,
		OrdersType:       swag.String(MoveQueueItem.OrdersType),
		MoveDate:         utils.FmtDatePtr(MoveQueueItem.MoveDate),
		CustomerDeadline: utils.FmtDate(MoveQueueItem.CustomerDeadline),
		LastModifiedDate: utils.FmtDateTime(MoveQueueItem.LastModifiedDate),
		LastModifiedName: swag.String(MoveQueueItem.LastModifiedName),
	}
	return &MoveQueueItemPayload
}

// ShowQueueHandler returns a list of all MoveQueueItems in the moves queue
type ShowQueueHandler utils.HandlerContext

// Handle retrieves a list of all MoveQueueItems in the system in the moves queue
func (h ShowQueueHandler) Handle(params queueop.ShowQueueParams) middleware.Responder {
	session := auth.SessionFromRequestContext(params.HTTPRequest)

	if !session.IsOfficeUser() {
		return queueop.NewShowQueueForbidden()
	}

	lifecycleState := params.QueueType

	MoveQueueItems, err := models.GetMoveQueueItems(h.db, lifecycleState)
	if err != nil {
		h.logger.Error("Loading Queue", zap.String("State", lifecycleState), zap.Error(err))
		return utils.ResponseForError(h.logger, err)
	}

	MoveQueueItemPayloads := make([]*internalmessages.MoveQueueItem, len(MoveQueueItems))
	for i, MoveQueueItem := range MoveQueueItems {
		MoveQueueItemPayload := payloadForMoveQueueItem(MoveQueueItem)
		MoveQueueItemPayloads[i] = MoveQueueItemPayload
	}
	return queueop.NewShowQueueOK().WithPayload(MoveQueueItemPayloads)
}
