package degree

import (
	"fmt"

	"github.com/akigugale/alt-verify/x/degree/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the degree type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateDegree:
			return handleMsgCreateDegree(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName,  msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handde<Action> does x
func handleMsgCreateDegree(ctx sdk.Context, k Keeper, msg MsgCreateDegree) (*sdk.Result, error) {
	var degree = types.Degree{
		Creator: msg.Creator,
		Student: msg.Student,
		Subject: msg.Subject,
		Batch  : msg.Batch,
	}

	_, err := k.GetDegree(ctx, degree.Student)
	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Degree for that student already exists")
	}
	k.SetDegree(ctx, degree)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateDegree),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Student),
			sdk.NewAttribute(types.AttributeCreator, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeSubject, msg.Subject),
			sdk.NewAttribute(types.AttributeBatch, string(msg.Batch)),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
