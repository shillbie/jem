package linego

import api "github.com/sakura-rip/linego/talkservice"

/*
PROFILE FUNCTION
*/

// GetProfile return the object of Profile
func (cl *LineClient) GetProfile() (*api.Profile, error) {
	prof, err := cl.Talk.GetProfile(cl.ctx, api.SyncReason_UNKNOWN)
	cl.Profile = prof
	return cl.Profile, err
}

// UpdateProfileName change profile name
func (cl *LineClient) UpdateProfileName(name string) error {
	req := api.NewUpdateProfileAttributesRequest()
	content := api.NewProfileContent()
	content.Value = name
	content.Meta = nil
	req.ProfileAttributes = map[api.ProfileAttribute]*api.ProfileContent{api.ProfileAttribute_DISPLAY_NAME: content}
	cl.reqSeq++
	err := cl.Talk.UpdateProfileAttributes(cl.ctx, cl.reqSeq, req)
	return err
}

// UpdateProfileBio change status message
func (cl *LineClient) UpdateProfileBio(bio string) error {
	req := api.NewUpdateProfileAttributesRequest()
	content := api.NewProfileContent()
	content.Value = bio
	content.Meta = nil
	req.ProfileAttributes = map[api.ProfileAttribute]*api.ProfileContent{api.ProfileAttribute_STATUS_MESSAGE: content}
	cl.reqSeq++
	err := cl.Talk.UpdateProfileAttributes(cl.ctx, cl.reqSeq, req)
	return err
}

/*
Message FUNCTION
*/

// SendMessage send text message
func (cl *LineClient) SendMessage(toMid, text string) (*api.Message, error) {
	msg := api.NewMessage()
	msg.Text = text
	msg.To = toMid
	msg.From_ = cl.Profile.Mid
	msg.RelatedMessageId = "0"
	msg.ContentType = api.ContentType_NONE
	if IsStrInMap(toMid, cl.reqSeqMessage) {
		cl.reqSeqMessage[toMid]++
	} else {
		cl.reqSeqMessage[toMid] = -1
	}
	return cl.Talk.SendMessage(cl.ctx, cl.reqSeqMessage[toMid], msg)
}

// SendContact send contact to toMid
func (cl *LineClient) SendContact(toMid, contactMid string) (*api.Message, error) {
	msg := api.NewMessage()
	msg.To = toMid
	msg.ContentType = api.ContentType_CONTACT
	msg.ContentMetadata = map[string]string{"mid": contactMid}
	if IsStrInMap(toMid, cl.reqSeqMessage) {
		cl.reqSeqMessage[toMid]++
	} else {
		cl.reqSeqMessage[toMid] = 1
	}
	return cl.Talk.SendMessage(cl.ctx, cl.reqSeqMessage[toMid], msg)
}

// SendChatChecked read message
func (cl *LineClient) SendChatChecked(groupID, messageID string) error {
	cl.reqSeq++
	err := cl.Talk.SendChatChecked(cl.ctx, cl.reqSeq, groupID, messageID, 0)
	return err
}

/*
CHAT FUNCTION
*/

// GetChats get chats
func (cl *LineClient) GetChats(chatsMids []string) ([]*api.Chat, error) {
	req := api.NewGetChatsRequest()
	req.ChatMids = chatsMids
	req.WithInvitees = true
	req.WithMembers = true
	res, err := cl.Talk.GetChats(cl.ctx, req)
	if res != nil {
		return res.Chats, err
	}
	return nil, err
}

// AcceptChatInvitation join chat
func (cl *LineClient) AcceptChatInvitation(groupMid string) error {
	req := api.NewAcceptChatInvitationRequest()
	req.ChatMid = groupMid
	req.ReqSeq = cl.reqSeq
	cl.reqSeq++
	_, err := cl.Talk.AcceptChatInvitation(cl.ctx, req)
	return err
}

// AcceptChatInvitationByTicket join chat by ticket
func (cl *LineClient) AcceptChatInvitationByTicket(groupMid, ticketID string) error {
	req := api.NewAcceptChatInvitationByTicketRequest()
	req.ChatMid = groupMid
	req.ReqSeq = cl.reqSeq
	req.TicketId = ticketID
	cl.reqSeq++
	_, err := cl.Talk.AcceptChatInvitationByTicket(cl.ctx, req)
	return err
}

// InviteIntoChat invite friend to chat
func (cl *LineClient) InviteIntoChat(chatMid string, targetMids []string) error {
	req := api.NewInviteIntoChatRequest()
	req.ReqSeq = cl.reqSeq
	cl.reqSeq++
	req.ChatMid = chatMid
	req.TargetUserMids = targetMids
	_, err := cl.Talk.InviteIntoChat(cl.ctx, req)
	return err
}

// ReissueChatTicket get chat invitation ticket
func (cl *LineClient) ReissueChatTicket(chatMid string) (string, error) {
	req := api.NewReissueChatTicketRequest()
	req.GroupMid = chatMid
	req.ReqSeq = cl.reqSeq
	cl.reqSeq++
	res, err := cl.Talk.ReissueChatTicket(cl.ctx, req)
	if res != nil {
		return res.TicketId, err
	}
	return "", err
}

// RejectChatInvitation reject chat
func (cl *LineClient) RejectChatInvitation(chatMid string) error {
	req := api.NewRejectChatInvitationRequest()
	req.ReqSeq = cl.reqSeq
	cl.reqSeq++
	req.ChatMid = chatMid
	_, err := cl.Talk.RejectChatInvitation(cl.ctx, req)
	return err
}

// GetChat return one chat mid
func (cl *LineClient) GetChat(chatID string) (*api.Chat, error) {
	req := api.NewGetChatsRequest()
	req.ChatMids = []string{chatID}
	req.WithInvitees = true
	req.WithMembers = true
	res, err := cl.Talk.GetChats(cl.ctx, req)
	if res != nil {
		if len(res.Chats) > 0 {
			return res.Chats[0], err
		}
	}
	return nil, err
}

// UpdateChatName change chat name
func (cl *LineClient) UpdateChatName(chatID string, name string) error {
	chat := &api.Chat{}
	chat.ChatMid = chatID
	chat.ChatName = name
	req := api.NewUpdateChatRequest()
	req.Chat = chat
	req.UpdatedAttribute = api.UpdatedAttribute_NAME
	_, err := cl.Talk.UpdateChat(cl.ctx, req)
	return err

}

// UpdateChatURL change chat url
func (cl *LineClient) UpdateChatURL(chatID string, typeVar bool) error {
	chat := &api.Chat{}
	chat.Extra.GroupExtra.PreventedJoinByTicket = typeVar
	chat.ChatMid = chatID
	req := api.NewUpdateChatRequest()
	req.Chat = chat
	req.UpdatedAttribute = api.UpdatedAttribute_PREVENTED_JOIN_BY_TICKET
	_, err := cl.Talk.UpdateChat(cl.ctx, req)
	return err
}

// DeleteOtherFromChat kickout some one from chat
func (cl *LineClient) DeleteOtherFromChat(toMid, targetMid string) error {
	req := api.NewDeleteOtherFromChatRequest()
	req.ChatMid = toMid
	req.ReqSeq = cl.reqSeq
	req.TargetUserMids = []string{targetMid}
	cl.reqSeq++
	_, err := cl.Talk.DeleteOtherFromChat(cl.ctx, req)
	return err
}

// DeleteSelfFromChat leave from chat
func (cl *LineClient) DeleteSelfFromChat(toMid string) error {
	req := api.NewDeleteSelfFromChatRequest()
	req.ChatMid = toMid
	req.ReqSeq = cl.reqSeq
	cl.reqSeq++
	req.LastSeenMessageDeliveredTime = 0
	req.LastSeenMessageId = ""
	req.LastMessageDeliveredTime = 0
	req.LastMessageId = ""
	_, err := cl.Talk.DeleteSelfFromChat(cl.ctx, req)
	return err
}

// GetAllChatMids get all chat mids
func (cl *LineClient) GetAllChatMids() (*api.GetAllChatMidsResponse, error) {
	req := api.NewGetAllChatMidsRequest()
	req.WithInvitedChats = true
	req.WithMemberChats = true
	return cl.Talk.GetAllChatMids(cl.ctx, req, api.SyncReason_UNKNOWN)
}

// CancelChatInvitation cancel user
func (cl *LineClient) CancelChatInvitation(groupID, targetMid string) error {
	req := api.NewCancelChatInvitationRequest()
	req.ChatMid = groupID
	req.ReqSeq = cl.reqSeq
	cl.reqSeq++
	req.TargetUserMids = []string{targetMid}
	_, err := cl.Talk.CancelChatInvitation(cl.ctx, req)
	return err
}

// FindChatByTicket find chat by ticket
func (cl *LineClient) FindChatByTicket(ticketID string) (*api.Chat, error) {
	req := api.NewFindChatByTicketRequest()
	req.TicketId = ticketID
	res, err := cl.Talk.FindChatByTicket(cl.ctx, req)
	if res != nil {
		return res.Chat, err
	}
	return nil, err
}

/*
Settings
*/

// GetSettings get settings
func (cl *LineClient) GetSettings() (*api.Settings, error) {
	return cl.Talk.GetSettings(cl.ctx, api.SyncReason_UNKNOWN)
}

// UpdateSettings update settings
func (cl *LineClient) UpdateSettings(attr []api.SettingAttribute, settings *api.Settings) error {
	_, err := cl.Talk.UpdateSettingsAttributes2(cl.ctx, cl.reqSeq, attr, settings)
	return err
}

// DisableE2ee disable e2ee
func (cl *LineClient) DisableE2ee() error {
	set := api.NewSettings()
	set.E2eeEnable = false
	err := cl.UpdateSettings([]api.SettingAttribute{api.SettingAttribute_E2EE_ENABLE}, set)
	return err
}

/*
Contact
*/

// FindAndAddContactsByMid add friend
func (cl *LineClient) FindAndAddContactsByMid(targetMid string) (map[string]*api.Contact, error) {
	return cl.Talk.FindAndAddContactsByMid(cl.ctx, cl.reqSeq, targetMid, api.ContactType_MID, `{"screen":"homeTab","spec":"native"}`)
}

// GetContacts get contact with list
func (cl *LineClient) GetContacts(targetMid []string) ([]*api.Contact, error) {
	return cl.Talk.GetContacts(cl.ctx, targetMid)
}

// GetContact get contact with list
func (cl *LineClient) GetContact(targetMid string) (*api.Contact, error) {
	return cl.Talk.GetContact(cl.ctx, targetMid)
}

// CoteHan update display name over ridden
func (cl *LineClient) CoteHan(mid, cote string) error {
	return cl.Talk.UpdateContactSetting(cl.ctx, cl.reqSeq, mid, api.ContactSettingAttribute_CONTACT_SETTING_DISPLAY_NAME_OVERRIDE, cote)
}

// GetAllContactIds get all list of mid
func (cl *LineClient) GetAllContactIds() ([]string, error) {
	return cl.Talk.GetAllContactIds(cl.ctx, api.SyncReason_UNKNOWN)
}

/*
OTHER
*/

// Noop nothing to do
func (cl *LineClient) Noop() {
	_ = cl.Talk.Noop(cl.ctx)
}
