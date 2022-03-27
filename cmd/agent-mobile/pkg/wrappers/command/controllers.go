package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	command "github.com/hyperledger/aries-framework-go/pkg/controller/command"
	connection "github.com/hyperledger/aries-framework-go/pkg/controller/command/connection"
	didexchange "github.com/hyperledger/aries-framework-go/pkg/controller/command/didexchange"
	introduce "github.com/hyperledger/aries-framework-go/pkg/controller/command/introduce"
	issuecredential "github.com/hyperledger/aries-framework-go/pkg/controller/command/issuecredential"
	kms "github.com/hyperledger/aries-framework-go/pkg/controller/command/kms"
	ld "github.com/hyperledger/aries-framework-go/pkg/controller/command/ld"
	mediator "github.com/hyperledger/aries-framework-go/pkg/controller/command/mediator"
	messaging "github.com/hyperledger/aries-framework-go/pkg/controller/command/messaging"
	outofband "github.com/hyperledger/aries-framework-go/pkg/controller/command/outofband"
	outofbandv2 "github.com/hyperledger/aries-framework-go/pkg/controller/command/outofbandv2"
	presentproof "github.com/hyperledger/aries-framework-go/pkg/controller/command/presentproof"
	vcwallet "github.com/hyperledger/aries-framework-go/pkg/controller/command/vcwallet"
	vdr "github.com/hyperledger/aries-framework-go/pkg/controller/command/vdr"
	verifiable "github.com/hyperledger/aries-framework-go/pkg/controller/command/verifiable"
	didclient "github.com/trustbloc/agent-sdk/pkg/controller/command/didclient"
)

type Connection struct {
	handlers map[string]command.Exec
}

func (c *Connection) RotateDID(request []byte) []byte {
	handler, ok := c.handlers[connection.RotateDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", connection.RotateDIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", connection.RotateDIDCommandMethod, "connection", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Connection) CreateConnectionV2(request []byte) []byte {
	handler, ok := c.handlers[connection.CreateV2CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", connection.CreateV2CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", connection.CreateV2CommandMethod, "connection", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Connection) SetConnectionToDIDCommV2(request []byte) []byte {
	handler, ok := c.handlers[connection.SetToV2CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", connection.SetToV2CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", connection.SetToV2CommandMethod, "connection", err)})
		return m
	}
	return writeResponse(w)
}

type DIDExchange struct {
	handlers map[string]command.Exec
}

func (c *DIDExchange) AcceptExchangeRequest(request []byte) []byte {
	handler, ok := c.handlers[didexchange.AcceptExchangeRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.AcceptExchangeRequestCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.AcceptExchangeRequestCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DIDExchange) AcceptInvitation(request []byte) []byte {
	handler, ok := c.handlers[didexchange.AcceptInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.AcceptInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.AcceptInvitationCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DIDExchange) CreateImplicitInvitation(request []byte) []byte {
	handler, ok := c.handlers[didexchange.CreateImplicitInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.CreateImplicitInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.CreateImplicitInvitationCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DIDExchange) CreateInvitation(request []byte) []byte {
	handler, ok := c.handlers[didexchange.CreateInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.CreateInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.CreateInvitationCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DIDExchange) QueryConnectionByID(request []byte) []byte {
	handler, ok := c.handlers[didexchange.QueryConnectionByIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.QueryConnectionByIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.QueryConnectionByIDCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DIDExchange) QueryConnections(request []byte) []byte {
	handler, ok := c.handlers[didexchange.QueryConnectionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.QueryConnectionsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.QueryConnectionsCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DIDExchange) ReceiveInvitation(request []byte) []byte {
	handler, ok := c.handlers[didexchange.ReceiveInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.ReceiveInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.ReceiveInvitationCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DIDExchange) CreateConnection(request []byte) []byte {
	handler, ok := c.handlers[didexchange.CreateConnectionCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.CreateConnectionCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.CreateConnectionCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DIDExchange) RemoveConnection(request []byte) []byte {
	handler, ok := c.handlers[didexchange.RemoveConnectionCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didexchange.RemoveConnectionCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didexchange.RemoveConnectionCommandMethod, "didexchange", err)})
		return m
	}
	return writeResponse(w)
}

type Introduce struct {
	handlers map[string]command.Exec
}

func (c *Introduce) Actions(request []byte) []byte {
	handler, ok := c.handlers[introduce.ActionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.ActionsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.ActionsCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) SendProposal(request []byte) []byte {
	handler, ok := c.handlers[introduce.SendProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.SendProposalCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.SendProposalCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) SendProposalWithOOBInvitation(request []byte) []byte {
	handler, ok := c.handlers[introduce.SendProposalWithOOBInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.SendProposalWithOOBInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.SendProposalWithOOBInvitationCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) SendRequest(request []byte) []byte {
	handler, ok := c.handlers[introduce.SendRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.SendRequestCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.SendRequestCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) AcceptProposalWithOOBInvitation(request []byte) []byte {
	handler, ok := c.handlers[introduce.AcceptProposalWithOOBInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.AcceptProposalWithOOBInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.AcceptProposalWithOOBInvitationCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) AcceptProposal(request []byte) []byte {
	handler, ok := c.handlers[introduce.AcceptProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.AcceptProposalCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.AcceptProposalCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) AcceptRequestWithPublicOOBInvitation(request []byte) []byte {
	handler, ok := c.handlers[introduce.AcceptRequestWithPublicOOBInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.AcceptRequestWithPublicOOBInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.AcceptRequestWithPublicOOBInvitationCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) AcceptRequestWithRecipients(request []byte) []byte {
	handler, ok := c.handlers[introduce.AcceptRequestWithRecipientsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.AcceptRequestWithRecipientsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.AcceptRequestWithRecipientsCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) DeclineProposal(request []byte) []byte {
	handler, ok := c.handlers[introduce.DeclineProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.DeclineProposalCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.DeclineProposalCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) DeclineRequest(request []byte) []byte {
	handler, ok := c.handlers[introduce.DeclineRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.DeclineRequestCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.DeclineRequestCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Introduce) AcceptProblemReport(request []byte) []byte {
	handler, ok := c.handlers[introduce.AcceptProblemReportCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", introduce.AcceptProblemReportCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", introduce.AcceptProblemReportCommandMethod, "introduce", err)})
		return m
	}
	return writeResponse(w)
}

type IssueCredential struct {
	handlers map[string]command.Exec
}

func (c *IssueCredential) Actions(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.ActionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.ActionsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.ActionsCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) SendOffer(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.SendOfferCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.SendOfferCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.SendOfferCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) SendProposal(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.SendProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.SendProposalCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.SendProposalCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) SendProposalV3(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.SendProposalV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.SendProposalV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.SendProposalV3CommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) SendRequest(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.SendRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.SendRequestCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.SendRequestCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) SendRequestV3(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.SendRequestV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.SendRequestV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.SendRequestV3CommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) AcceptProposal(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.AcceptProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.AcceptProposalCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.AcceptProposalCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) AcceptProposalV3(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.AcceptProposalV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.AcceptProposalV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.AcceptProposalV3CommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) DeclineProposal(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.DeclineProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.DeclineProposalCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.DeclineProposalCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) AcceptOffer(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.AcceptOfferCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.AcceptOfferCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.AcceptOfferCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) DeclineOffer(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.DeclineOfferCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.DeclineOfferCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.DeclineOfferCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) NegotiateProposal(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.NegotiateProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.NegotiateProposalCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.NegotiateProposalCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) NegotiateProposalV3(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.NegotiateProposalV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.NegotiateProposalV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.NegotiateProposalV3CommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) AcceptRequest(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.AcceptRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.AcceptRequestCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.AcceptRequestCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) AcceptRequestV3(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.AcceptRequestV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.AcceptRequestV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.AcceptRequestV3CommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) DeclineRequest(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.DeclineRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.DeclineRequestCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.DeclineRequestCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) AcceptCredential(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.AcceptCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.AcceptCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.AcceptCredentialCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) DeclineCredential(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.DeclineCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.DeclineCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.DeclineCredentialCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}
func (c *IssueCredential) AcceptProblemReport(request []byte) []byte {
	handler, ok := c.handlers[issuecredential.AcceptProblemReportCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", issuecredential.AcceptProblemReportCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", issuecredential.AcceptProblemReportCommandMethod, "issuecredential", err)})
		return m
	}
	return writeResponse(w)
}

type KMS struct {
	handlers map[string]command.Exec
}

func (c *KMS) CreateKeySet(request []byte) []byte {
	handler, ok := c.handlers[kms.CreateKeySetCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", kms.CreateKeySetCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", kms.CreateKeySetCommandMethod, "kms", err)})
		return m
	}
	return writeResponse(w)
}
func (c *KMS) ImportKey(request []byte) []byte {
	handler, ok := c.handlers[kms.ImportKeyCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", kms.ImportKeyCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", kms.ImportKeyCommandMethod, "kms", err)})
		return m
	}
	return writeResponse(w)
}

type LD struct {
	handlers map[string]command.Exec
}

func (c *LD) AddContexts(request []byte) []byte {
	handler, ok := c.handlers[ld.AddContextsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", ld.AddContextsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", ld.AddContextsCommandMethod, "ld", err)})
		return m
	}
	return writeResponse(w)
}
func (c *LD) AddRemoteProvider(request []byte) []byte {
	handler, ok := c.handlers[ld.AddRemoteProviderCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", ld.AddRemoteProviderCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", ld.AddRemoteProviderCommandMethod, "ld", err)})
		return m
	}
	return writeResponse(w)
}
func (c *LD) RefreshRemoteProvider(request []byte) []byte {
	handler, ok := c.handlers[ld.RefreshRemoteProviderCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", ld.RefreshRemoteProviderCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", ld.RefreshRemoteProviderCommandMethod, "ld", err)})
		return m
	}
	return writeResponse(w)
}
func (c *LD) DeleteRemoteProvider(request []byte) []byte {
	handler, ok := c.handlers[ld.DeleteRemoteProviderCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", ld.DeleteRemoteProviderCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", ld.DeleteRemoteProviderCommandMethod, "ld", err)})
		return m
	}
	return writeResponse(w)
}
func (c *LD) GetAllRemoteProviders(request []byte) []byte {
	handler, ok := c.handlers[ld.GetAllRemoteProvidersCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", ld.GetAllRemoteProvidersCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", ld.GetAllRemoteProvidersCommandMethod, "ld", err)})
		return m
	}
	return writeResponse(w)
}
func (c *LD) RefreshAllRemoteProviders(request []byte) []byte {
	handler, ok := c.handlers[ld.RefreshAllRemoteProvidersCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", ld.RefreshAllRemoteProvidersCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", ld.RefreshAllRemoteProvidersCommandMethod, "ld", err)})
		return m
	}
	return writeResponse(w)
}

type Mediator struct {
	handlers map[string]command.Exec
}

func (c *Mediator) Register(request []byte) []byte {
	handler, ok := c.handlers[mediator.RegisterCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", mediator.RegisterCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", mediator.RegisterCommandMethod, "mediator", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Mediator) Unregister(request []byte) []byte {
	handler, ok := c.handlers[mediator.UnregisterCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", mediator.UnregisterCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", mediator.UnregisterCommandMethod, "mediator", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Mediator) Connections(request []byte) []byte {
	handler, ok := c.handlers[mediator.GetConnectionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", mediator.GetConnectionsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", mediator.GetConnectionsCommandMethod, "mediator", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Mediator) Reconnect(request []byte) []byte {
	handler, ok := c.handlers[mediator.ReconnectCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", mediator.ReconnectCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", mediator.ReconnectCommandMethod, "mediator", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Mediator) Status(request []byte) []byte {
	handler, ok := c.handlers[mediator.StatusCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", mediator.StatusCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", mediator.StatusCommandMethod, "mediator", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Mediator) BatchPickup(request []byte) []byte {
	handler, ok := c.handlers[mediator.BatchPickupCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", mediator.BatchPickupCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", mediator.BatchPickupCommandMethod, "mediator", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Mediator) ReconnectAll(request []byte) []byte {
	handler, ok := c.handlers[mediator.ReconnectAllCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", mediator.ReconnectAllCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", mediator.ReconnectAllCommandMethod, "mediator", err)})
		return m
	}
	return writeResponse(w)
}

type Messaging struct {
	handlers map[string]command.Exec
}

func (c *Messaging) Services(request []byte) []byte {
	handler, ok := c.handlers[messaging.RegisteredServicesCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", messaging.RegisteredServicesCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", messaging.RegisteredServicesCommandMethod, "messaging", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Messaging) RegisterService(request []byte) []byte {
	handler, ok := c.handlers[messaging.RegisterMessageServiceCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", messaging.RegisterMessageServiceCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", messaging.RegisterMessageServiceCommandMethod, "messaging", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Messaging) UnregisterService(request []byte) []byte {
	handler, ok := c.handlers[messaging.UnregisterMessageServiceCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", messaging.UnregisterMessageServiceCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", messaging.UnregisterMessageServiceCommandMethod, "messaging", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Messaging) RegisterHTTPService(request []byte) []byte {
	handler, ok := c.handlers[messaging.RegisterHTTPMessageServiceCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", messaging.RegisterHTTPMessageServiceCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", messaging.RegisterHTTPMessageServiceCommandMethod, "messaging", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Messaging) Send(request []byte) []byte {
	handler, ok := c.handlers[messaging.SendNewMessageCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", messaging.SendNewMessageCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", messaging.SendNewMessageCommandMethod, "messaging", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Messaging) Reply(request []byte) []byte {
	handler, ok := c.handlers[messaging.SendReplyMessageCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", messaging.SendReplyMessageCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", messaging.SendReplyMessageCommandMethod, "messaging", err)})
		return m
	}
	return writeResponse(w)
}

type OutofBand struct {
	handlers map[string]command.Exec
}

func (c *OutofBand) CreateInvitation(request []byte) []byte {
	handler, ok := c.handlers[outofband.CreateInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", outofband.CreateInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", outofband.CreateInvitationCommandMethod, "outofband", err)})
		return m
	}
	return writeResponse(w)
}
func (c *OutofBand) AcceptInvitation(request []byte) []byte {
	handler, ok := c.handlers[outofband.AcceptInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", outofband.AcceptInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", outofband.AcceptInvitationCommandMethod, "outofband", err)})
		return m
	}
	return writeResponse(w)
}
func (c *OutofBand) ActionStop(request []byte) []byte {
	handler, ok := c.handlers[outofband.ActionStopCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", outofband.ActionStopCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", outofband.ActionStopCommandMethod, "outofband", err)})
		return m
	}
	return writeResponse(w)
}
func (c *OutofBand) Actions(request []byte) []byte {
	handler, ok := c.handlers[outofband.ActionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", outofband.ActionsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", outofband.ActionsCommandMethod, "outofband", err)})
		return m
	}
	return writeResponse(w)
}
func (c *OutofBand) ActionContinue(request []byte) []byte {
	handler, ok := c.handlers[outofband.ActionContinueCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", outofband.ActionContinueCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", outofband.ActionContinueCommandMethod, "outofband", err)})
		return m
	}
	return writeResponse(w)
}

type OutofBandV2 struct {
	handlers map[string]command.Exec
}

func (c *OutofBandV2) CreateInvitation(request []byte) []byte {
	handler, ok := c.handlers[outofbandv2.CreateInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", outofbandv2.CreateInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", outofbandv2.CreateInvitationCommandMethod, "outofbandv2", err)})
		return m
	}
	return writeResponse(w)
}
func (c *OutofBandV2) AcceptInvitation(request []byte) []byte {
	handler, ok := c.handlers[outofbandv2.AcceptInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", outofbandv2.AcceptInvitationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", outofbandv2.AcceptInvitationCommandMethod, "outofbandv2", err)})
		return m
	}
	return writeResponse(w)
}

type PresentProof struct {
	handlers map[string]command.Exec
}

func (c *PresentProof) Actions(request []byte) []byte {
	handler, ok := c.handlers[presentproof.ActionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.ActionsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.ActionsCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) SendRequestPresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.SendRequestPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.SendRequestPresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.SendRequestPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) SendRequestPresentationV3(request []byte) []byte {
	handler, ok := c.handlers[presentproof.SendRequestPresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.SendRequestPresentationV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.SendRequestPresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) AcceptRequestPresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.AcceptRequestPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.AcceptRequestPresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.AcceptRequestPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) AcceptRequestPresentationV3(request []byte) []byte {
	handler, ok := c.handlers[presentproof.AcceptRequestPresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.AcceptRequestPresentationV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.AcceptRequestPresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) NegotiateRequestPresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.NegotiateRequestPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.NegotiateRequestPresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.NegotiateRequestPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) NegotiateRequestPresentationV3(request []byte) []byte {
	handler, ok := c.handlers[presentproof.NegotiateRequestPresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.NegotiateRequestPresentationV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.NegotiateRequestPresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) AcceptProblemReport(request []byte) []byte {
	handler, ok := c.handlers[presentproof.AcceptProblemReportCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.AcceptProblemReportCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.AcceptProblemReportCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) DeclineRequestPresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.DeclineRequestPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.DeclineRequestPresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.DeclineRequestPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) SendProposePresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.SendProposePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.SendProposePresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.SendProposePresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) SendProposePresentationV3(request []byte) []byte {
	handler, ok := c.handlers[presentproof.SendProposePresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.SendProposePresentationV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.SendProposePresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) AcceptProposePresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.AcceptProposePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.AcceptProposePresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.AcceptProposePresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) AcceptProposePresentationV3(request []byte) []byte {
	handler, ok := c.handlers[presentproof.AcceptProposePresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.AcceptProposePresentationV3CommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.AcceptProposePresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) DeclineProposePresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.DeclineProposePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.DeclineProposePresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.DeclineProposePresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) AcceptPresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.AcceptPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.AcceptPresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.AcceptPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}
func (c *PresentProof) DeclinePresentation(request []byte) []byte {
	handler, ok := c.handlers[presentproof.DeclinePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", presentproof.DeclinePresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", presentproof.DeclinePresentationCommandMethod, "presentproof", err)})
		return m
	}
	return writeResponse(w)
}

type Rfc0593 struct {
	handlers map[string]command.Exec
}

type VCWallet struct {
	handlers map[string]command.Exec
}

func (c *VCWallet) CreateProfile(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.CreateProfileCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.CreateProfileCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.CreateProfileCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) UpdateProfile(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.UpdateProfileCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.UpdateProfileCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.UpdateProfileCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) ProfileExists(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.ProfileExistsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.ProfileExistsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.ProfileExistsCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Open(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.OpenCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.OpenCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.OpenCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Close(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.CloseCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.CloseCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.CloseCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Add(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.AddCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.AddCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.AddCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Remove(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.RemoveCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.RemoveCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.RemoveCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Get(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.GetCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.GetCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.GetCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) GetAll(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.GetAllCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.GetAllCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.GetAllCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Query(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.QueryCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.QueryCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.QueryCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Issue(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.IssueCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.IssueCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.IssueCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Prove(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.ProveCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.ProveCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.ProveCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Verify(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.VerifyCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.VerifyCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.VerifyCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Derive(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.DeriveCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.DeriveCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.DeriveCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) CreateKeyPair(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.CreateKeyPairCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.CreateKeyPairCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.CreateKeyPairCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) Connect(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.ConnectCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.ConnectCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.ConnectCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) ProposePresentation(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.ProposePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.ProposePresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.ProposePresentationCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) PresentProof(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.PresentProofCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.PresentProofCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.PresentProofCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) ProposeCredential(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.ProposeCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.ProposeCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.ProposeCredentialCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) RequestCredential(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.RequestCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.RequestCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.RequestCredentialCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VCWallet) ResolveCredentialManifest(request []byte) []byte {
	handler, ok := c.handlers[vcwallet.ResolveCredentialManifestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vcwallet.ResolveCredentialManifestCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vcwallet.ResolveCredentialManifestCommandMethod, "vcwallet", err)})
		return m
	}
	return writeResponse(w)
}

type VDR struct {
	handlers map[string]command.Exec
}

func (c *VDR) SaveDID(request []byte) []byte {
	handler, ok := c.handlers[vdr.SaveDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vdr.SaveDIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vdr.SaveDIDCommandMethod, "vdr", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VDR) GetDIDRecords(request []byte) []byte {
	handler, ok := c.handlers[vdr.GetDIDsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vdr.GetDIDsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vdr.GetDIDsCommandMethod, "vdr", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VDR) GetDID(request []byte) []byte {
	handler, ok := c.handlers[vdr.GetDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vdr.GetDIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vdr.GetDIDCommandMethod, "vdr", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VDR) ResolveDID(request []byte) []byte {
	handler, ok := c.handlers[vdr.ResolveDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vdr.ResolveDIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vdr.ResolveDIDCommandMethod, "vdr", err)})
		return m
	}
	return writeResponse(w)
}
func (c *VDR) CreateDID(request []byte) []byte {
	handler, ok := c.handlers[vdr.CreateDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", vdr.CreateDIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", vdr.CreateDIDCommandMethod, "vdr", err)})
		return m
	}
	return writeResponse(w)
}

type Verifiable struct {
	handlers map[string]command.Exec
}

func (c *Verifiable) ValidateCredential(request []byte) []byte {
	handler, ok := c.handlers[verifiable.ValidateCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.ValidateCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.ValidateCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) SaveCredential(request []byte) []byte {
	handler, ok := c.handlers[verifiable.SaveCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.SaveCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.SaveCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) GetCredential(request []byte) []byte {
	handler, ok := c.handlers[verifiable.GetCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.GetCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.GetCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) GetCredentialByName(request []byte) []byte {
	handler, ok := c.handlers[verifiable.GetCredentialByNameCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.GetCredentialByNameCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.GetCredentialByNameCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) GetCredentials(request []byte) []byte {
	handler, ok := c.handlers[verifiable.GetCredentialsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.GetCredentialsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.GetCredentialsCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) SignCredential(request []byte) []byte {
	handler, ok := c.handlers[verifiable.SignCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.SignCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.SignCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) DeriveCredential(request []byte) []byte {
	handler, ok := c.handlers[verifiable.DeriveCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.DeriveCredentialCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.DeriveCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) SavePresentation(request []byte) []byte {
	handler, ok := c.handlers[verifiable.SavePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.SavePresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.SavePresentationCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) GetPresentation(request []byte) []byte {
	handler, ok := c.handlers[verifiable.GetPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.GetPresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.GetPresentationCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) GetPresentations(request []byte) []byte {
	handler, ok := c.handlers[verifiable.GetPresentationsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.GetPresentationsCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.GetPresentationsCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) GeneratePresentation(request []byte) []byte {
	handler, ok := c.handlers[verifiable.GeneratePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.GeneratePresentationCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.GeneratePresentationCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) GeneratePresentationByID(request []byte) []byte {
	handler, ok := c.handlers[verifiable.GeneratePresentationByIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.GeneratePresentationByIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.GeneratePresentationByIDCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) RemoveCredentialByName(request []byte) []byte {
	handler, ok := c.handlers[verifiable.RemoveCredentialByNameCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.RemoveCredentialByNameCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.RemoveCredentialByNameCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}
func (c *Verifiable) RemovePresentationByName(request []byte) []byte {
	handler, ok := c.handlers[verifiable.RemovePresentationByNameCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", verifiable.RemovePresentationByNameCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", verifiable.RemovePresentationByNameCommandMethod, "verifiable", err)})
		return m
	}
	return writeResponse(w)
}

type BlindedRouting struct {
	handlers map[string]command.Exec
}

type DidClient struct {
	handlers map[string]command.Exec
}

func (c *DidClient) CreateOrbDID(request []byte) []byte {
	handler, ok := c.handlers[didclient.CreateOrbDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didclient.CreateOrbDIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didclient.CreateOrbDIDCommandMethod, "didclient", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DidClient) ResolveOrbDID(request []byte) []byte {
	handler, ok := c.handlers[didclient.ResolveOrbDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didclient.ResolveOrbDIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didclient.ResolveOrbDIDCommandMethod, "didclient", err)})
		return m
	}
	return writeResponse(w)
}
func (c *DidClient) CreatePeerDID(request []byte) []byte {
	handler, ok := c.handlers[didclient.CreatePeerDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("no handlers found for controller [%s]", didclient.CreatePeerDIDCommandMethod)})
		return m
	}
	w := &bytes.Buffer{}
	r := bytes.NewReader(request)
	err := handler(w, r)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Errorf("Aries error \n command: [%s]\n controller: [%s]\n error: [%w]", didclient.CreatePeerDIDCommandMethod, "didclient", err)})
		return m
	}
	return writeResponse(w)
}

type MediatorClient struct {
	handlers map[string]command.Exec
}
