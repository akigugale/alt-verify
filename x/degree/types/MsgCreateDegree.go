package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Msg Create Degree
var _ sdk.Msg = &MsgCreateDegree{}

// MsgCreateDegree - struct for MsgCreateDegree
type MsgCreateDegree struct {
	Creator sdk.AccAddress `json:"address" yaml:"address"` // address of the degree creator
	Student sdk.AccAddress `json:"student" yaml:"student"` // address of student
	Subject string `json:"subject" yaml:"subject"` // address of the degree creator
	Batch uint16 `json:"batch" yaml:"batch"` // year
}

// creates a new MsgCreateDegree instance
func NewMsgCreateDegree(creator, student sdk.AccAddress, subject string, batch uint16) MsgCreateDegree {
	return MsgCreateDegree{
		Creator: creator,
		Student: student,
		Subject: subject,
		Batch  : batch,
	}
}

// CreateDegreeConst is CreateDegree Constant
const CreateDegreeConst = "CreateDegree"

// nolint
func (msg MsgCreateDegree) Route() string { return RouterKey }
func (msg MsgCreateDegree) Type() string  { return CreateDegreeConst }
func (msg MsgCreateDegree) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateDegree) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateDegree) ValidateBasic() error {
	if msg.Creator.Empty() {
		errmsg := fmt.Sprintf("Missing University address - %s .", msg.Creator.String())
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errmsg)
	}
	if msg.Student.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing Student address")
	}
	if msg.Batch == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "batch cannot be empty")
	}
	if msg.Subject == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subject cannot be empty")
	}
	return nil
}
