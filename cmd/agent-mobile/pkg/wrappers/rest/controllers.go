package rest

import (
	"encoding/json"
	"fmt"
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
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *Connection) RotateDID(request []byte) []byte {
	ep, ok := c.endpoints[connection.RotateDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", connection.RotateDIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", connection.RotateDIDCommandMethod, "connection", err)})
		return m
	}
	return response
}
func (c *Connection) CreateConnectionV2(request []byte) []byte {
	ep, ok := c.endpoints[connection.CreateV2CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", connection.CreateV2CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", connection.CreateV2CommandMethod, "connection", err)})
		return m
	}
	return response
}
func (c *Connection) SetConnectionToDIDCommV2(request []byte) []byte {
	ep, ok := c.endpoints[connection.SetToV2CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", connection.SetToV2CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", connection.SetToV2CommandMethod, "connection", err)})
		return m
	}
	return response
}

type DIDExchange struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *DIDExchange) AcceptExchangeRequest(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.AcceptExchangeRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.AcceptExchangeRequestCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.AcceptExchangeRequestCommandMethod, "didexchange", err)})
		return m
	}
	return response
}
func (c *DIDExchange) AcceptInvitation(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.AcceptInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.AcceptInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.AcceptInvitationCommandMethod, "didexchange", err)})
		return m
	}
	return response
}
func (c *DIDExchange) CreateImplicitInvitation(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.CreateImplicitInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.CreateImplicitInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.CreateImplicitInvitationCommandMethod, "didexchange", err)})
		return m
	}
	return response
}
func (c *DIDExchange) CreateInvitation(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.CreateInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.CreateInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.CreateInvitationCommandMethod, "didexchange", err)})
		return m
	}
	return response
}
func (c *DIDExchange) QueryConnectionByID(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.QueryConnectionByIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.QueryConnectionByIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.QueryConnectionByIDCommandMethod, "didexchange", err)})
		return m
	}
	return response
}
func (c *DIDExchange) QueryConnections(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.QueryConnectionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.QueryConnectionsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.QueryConnectionsCommandMethod, "didexchange", err)})
		return m
	}
	return response
}
func (c *DIDExchange) ReceiveInvitation(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.ReceiveInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.ReceiveInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.ReceiveInvitationCommandMethod, "didexchange", err)})
		return m
	}
	return response
}
func (c *DIDExchange) CreateConnection(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.CreateConnectionCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.CreateConnectionCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.CreateConnectionCommandMethod, "didexchange", err)})
		return m
	}
	return response
}
func (c *DIDExchange) RemoveConnection(request []byte) []byte {
	ep, ok := c.endpoints[didexchange.RemoveConnectionCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didexchange.RemoveConnectionCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didexchange.RemoveConnectionCommandMethod, "didexchange", err)})
		return m
	}
	return response
}

type Introduce struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *Introduce) Actions(request []byte) []byte {
	ep, ok := c.endpoints[introduce.ActionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.ActionsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.ActionsCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) SendProposal(request []byte) []byte {
	ep, ok := c.endpoints[introduce.SendProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.SendProposalCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.SendProposalCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) SendProposalWithOOBInvitation(request []byte) []byte {
	ep, ok := c.endpoints[introduce.SendProposalWithOOBInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.SendProposalWithOOBInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.SendProposalWithOOBInvitationCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) SendRequest(request []byte) []byte {
	ep, ok := c.endpoints[introduce.SendRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.SendRequestCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.SendRequestCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) AcceptProposalWithOOBInvitation(request []byte) []byte {
	ep, ok := c.endpoints[introduce.AcceptProposalWithOOBInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.AcceptProposalWithOOBInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.AcceptProposalWithOOBInvitationCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) AcceptProposal(request []byte) []byte {
	ep, ok := c.endpoints[introduce.AcceptProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.AcceptProposalCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.AcceptProposalCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) AcceptRequestWithPublicOOBInvitation(request []byte) []byte {
	ep, ok := c.endpoints[introduce.AcceptRequestWithPublicOOBInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.AcceptRequestWithPublicOOBInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.AcceptRequestWithPublicOOBInvitationCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) AcceptRequestWithRecipients(request []byte) []byte {
	ep, ok := c.endpoints[introduce.AcceptRequestWithRecipientsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.AcceptRequestWithRecipientsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.AcceptRequestWithRecipientsCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) DeclineProposal(request []byte) []byte {
	ep, ok := c.endpoints[introduce.DeclineProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.DeclineProposalCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.DeclineProposalCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) DeclineRequest(request []byte) []byte {
	ep, ok := c.endpoints[introduce.DeclineRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.DeclineRequestCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.DeclineRequestCommandMethod, "introduce", err)})
		return m
	}
	return response
}
func (c *Introduce) AcceptProblemReport(request []byte) []byte {
	ep, ok := c.endpoints[introduce.AcceptProblemReportCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", introduce.AcceptProblemReportCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", introduce.AcceptProblemReportCommandMethod, "introduce", err)})
		return m
	}
	return response
}

type IssueCredential struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *IssueCredential) Actions(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.ActionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.ActionsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.ActionsCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) SendOffer(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.SendOfferCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.SendOfferCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.SendOfferCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) SendProposal(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.SendProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.SendProposalCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.SendProposalCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) SendProposalV3(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.SendProposalV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.SendProposalV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.SendProposalV3CommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) SendRequest(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.SendRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.SendRequestCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.SendRequestCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) SendRequestV3(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.SendRequestV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.SendRequestV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.SendRequestV3CommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) AcceptProposal(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.AcceptProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.AcceptProposalCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.AcceptProposalCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) AcceptProposalV3(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.AcceptProposalV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.AcceptProposalV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.AcceptProposalV3CommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) DeclineProposal(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.DeclineProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.DeclineProposalCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.DeclineProposalCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) AcceptOffer(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.AcceptOfferCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.AcceptOfferCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.AcceptOfferCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) DeclineOffer(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.DeclineOfferCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.DeclineOfferCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.DeclineOfferCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) NegotiateProposal(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.NegotiateProposalCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.NegotiateProposalCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.NegotiateProposalCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) NegotiateProposalV3(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.NegotiateProposalV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.NegotiateProposalV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.NegotiateProposalV3CommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) AcceptRequest(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.AcceptRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.AcceptRequestCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.AcceptRequestCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) AcceptRequestV3(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.AcceptRequestV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.AcceptRequestV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.AcceptRequestV3CommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) DeclineRequest(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.DeclineRequestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.DeclineRequestCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.DeclineRequestCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) AcceptCredential(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.AcceptCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.AcceptCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.AcceptCredentialCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) DeclineCredential(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.DeclineCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.DeclineCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.DeclineCredentialCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}
func (c *IssueCredential) AcceptProblemReport(request []byte) []byte {
	ep, ok := c.endpoints[issuecredential.AcceptProblemReportCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", issuecredential.AcceptProblemReportCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", issuecredential.AcceptProblemReportCommandMethod, "issuecredential", err)})
		return m
	}
	return response
}

type KMS struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *KMS) CreateKeySet(request []byte) []byte {
	ep, ok := c.endpoints[kms.CreateKeySetCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", kms.CreateKeySetCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", kms.CreateKeySetCommandMethod, "kms", err)})
		return m
	}
	return response
}
func (c *KMS) ImportKey(request []byte) []byte {
	ep, ok := c.endpoints[kms.ImportKeyCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", kms.ImportKeyCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", kms.ImportKeyCommandMethod, "kms", err)})
		return m
	}
	return response
}

type LD struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *LD) AddContexts(request []byte) []byte {
	ep, ok := c.endpoints[ld.AddContextsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", ld.AddContextsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", ld.AddContextsCommandMethod, "ld", err)})
		return m
	}
	return response
}
func (c *LD) AddRemoteProvider(request []byte) []byte {
	ep, ok := c.endpoints[ld.AddRemoteProviderCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", ld.AddRemoteProviderCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", ld.AddRemoteProviderCommandMethod, "ld", err)})
		return m
	}
	return response
}
func (c *LD) RefreshRemoteProvider(request []byte) []byte {
	ep, ok := c.endpoints[ld.RefreshRemoteProviderCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", ld.RefreshRemoteProviderCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", ld.RefreshRemoteProviderCommandMethod, "ld", err)})
		return m
	}
	return response
}
func (c *LD) DeleteRemoteProvider(request []byte) []byte {
	ep, ok := c.endpoints[ld.DeleteRemoteProviderCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", ld.DeleteRemoteProviderCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", ld.DeleteRemoteProviderCommandMethod, "ld", err)})
		return m
	}
	return response
}
func (c *LD) GetAllRemoteProviders(request []byte) []byte {
	ep, ok := c.endpoints[ld.GetAllRemoteProvidersCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", ld.GetAllRemoteProvidersCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", ld.GetAllRemoteProvidersCommandMethod, "ld", err)})
		return m
	}
	return response
}
func (c *LD) RefreshAllRemoteProviders(request []byte) []byte {
	ep, ok := c.endpoints[ld.RefreshAllRemoteProvidersCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", ld.RefreshAllRemoteProvidersCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", ld.RefreshAllRemoteProvidersCommandMethod, "ld", err)})
		return m
	}
	return response
}

type Mediator struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *Mediator) Register(request []byte) []byte {
	ep, ok := c.endpoints[mediator.RegisterCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", mediator.RegisterCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", mediator.RegisterCommandMethod, "mediator", err)})
		return m
	}
	return response
}
func (c *Mediator) Unregister(request []byte) []byte {
	ep, ok := c.endpoints[mediator.UnregisterCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", mediator.UnregisterCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", mediator.UnregisterCommandMethod, "mediator", err)})
		return m
	}
	return response
}
func (c *Mediator) Connections(request []byte) []byte {
	ep, ok := c.endpoints[mediator.GetConnectionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", mediator.GetConnectionsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", mediator.GetConnectionsCommandMethod, "mediator", err)})
		return m
	}
	return response
}
func (c *Mediator) Reconnect(request []byte) []byte {
	ep, ok := c.endpoints[mediator.ReconnectCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", mediator.ReconnectCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", mediator.ReconnectCommandMethod, "mediator", err)})
		return m
	}
	return response
}
func (c *Mediator) Status(request []byte) []byte {
	ep, ok := c.endpoints[mediator.StatusCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", mediator.StatusCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", mediator.StatusCommandMethod, "mediator", err)})
		return m
	}
	return response
}
func (c *Mediator) BatchPickup(request []byte) []byte {
	ep, ok := c.endpoints[mediator.BatchPickupCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", mediator.BatchPickupCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", mediator.BatchPickupCommandMethod, "mediator", err)})
		return m
	}
	return response
}
func (c *Mediator) ReconnectAll(request []byte) []byte {
	ep, ok := c.endpoints[mediator.ReconnectAllCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", mediator.ReconnectAllCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", mediator.ReconnectAllCommandMethod, "mediator", err)})
		return m
	}
	return response
}

type Messaging struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *Messaging) Services(request []byte) []byte {
	ep, ok := c.endpoints[messaging.RegisteredServicesCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", messaging.RegisteredServicesCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", messaging.RegisteredServicesCommandMethod, "messaging", err)})
		return m
	}
	return response
}
func (c *Messaging) RegisterService(request []byte) []byte {
	ep, ok := c.endpoints[messaging.RegisterMessageServiceCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", messaging.RegisterMessageServiceCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", messaging.RegisterMessageServiceCommandMethod, "messaging", err)})
		return m
	}
	return response
}
func (c *Messaging) UnregisterService(request []byte) []byte {
	ep, ok := c.endpoints[messaging.UnregisterMessageServiceCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", messaging.UnregisterMessageServiceCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", messaging.UnregisterMessageServiceCommandMethod, "messaging", err)})
		return m
	}
	return response
}
func (c *Messaging) RegisterHTTPService(request []byte) []byte {
	ep, ok := c.endpoints[messaging.RegisterHTTPMessageServiceCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", messaging.RegisterHTTPMessageServiceCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", messaging.RegisterHTTPMessageServiceCommandMethod, "messaging", err)})
		return m
	}
	return response
}
func (c *Messaging) Send(request []byte) []byte {
	ep, ok := c.endpoints[messaging.SendNewMessageCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", messaging.SendNewMessageCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", messaging.SendNewMessageCommandMethod, "messaging", err)})
		return m
	}
	return response
}
func (c *Messaging) Reply(request []byte) []byte {
	ep, ok := c.endpoints[messaging.SendReplyMessageCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", messaging.SendReplyMessageCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", messaging.SendReplyMessageCommandMethod, "messaging", err)})
		return m
	}
	return response
}

type OutOfBand struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *OutOfBand) CreateInvitation(request []byte) []byte {
	ep, ok := c.endpoints[outofband.CreateInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", outofband.CreateInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", outofband.CreateInvitationCommandMethod, "outofband", err)})
		return m
	}
	return response
}
func (c *OutOfBand) AcceptInvitation(request []byte) []byte {
	ep, ok := c.endpoints[outofband.AcceptInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", outofband.AcceptInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", outofband.AcceptInvitationCommandMethod, "outofband", err)})
		return m
	}
	return response
}
func (c *OutOfBand) ActionStop(request []byte) []byte {
	ep, ok := c.endpoints[outofband.ActionStopCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", outofband.ActionStopCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", outofband.ActionStopCommandMethod, "outofband", err)})
		return m
	}
	return response
}
func (c *OutOfBand) Actions(request []byte) []byte {
	ep, ok := c.endpoints[outofband.ActionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", outofband.ActionsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", outofband.ActionsCommandMethod, "outofband", err)})
		return m
	}
	return response
}
func (c *OutOfBand) ActionContinue(request []byte) []byte {
	ep, ok := c.endpoints[outofband.ActionContinueCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", outofband.ActionContinueCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", outofband.ActionContinueCommandMethod, "outofband", err)})
		return m
	}
	return response
}

type OutOfBandV2 struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *OutOfBandV2) CreateInvitation(request []byte) []byte {
	ep, ok := c.endpoints[outofbandv2.CreateInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", outofbandv2.CreateInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", outofbandv2.CreateInvitationCommandMethod, "outofbandv2", err)})
		return m
	}
	return response
}
func (c *OutOfBandV2) AcceptInvitation(request []byte) []byte {
	ep, ok := c.endpoints[outofbandv2.AcceptInvitationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", outofbandv2.AcceptInvitationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", outofbandv2.AcceptInvitationCommandMethod, "outofbandv2", err)})
		return m
	}
	return response
}

type PresentProof struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *PresentProof) Actions(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.ActionsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.ActionsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.ActionsCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) SendRequestPresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.SendRequestPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.SendRequestPresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.SendRequestPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) SendRequestPresentationV3(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.SendRequestPresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.SendRequestPresentationV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.SendRequestPresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) AcceptRequestPresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.AcceptRequestPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.AcceptRequestPresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.AcceptRequestPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) AcceptRequestPresentationV3(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.AcceptRequestPresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.AcceptRequestPresentationV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.AcceptRequestPresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) NegotiateRequestPresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.NegotiateRequestPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.NegotiateRequestPresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.NegotiateRequestPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) NegotiateRequestPresentationV3(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.NegotiateRequestPresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.NegotiateRequestPresentationV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.NegotiateRequestPresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) AcceptProblemReport(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.AcceptProblemReportCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.AcceptProblemReportCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.AcceptProblemReportCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) DeclineRequestPresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.DeclineRequestPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.DeclineRequestPresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.DeclineRequestPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) SendProposePresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.SendProposePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.SendProposePresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.SendProposePresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) SendProposePresentationV3(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.SendProposePresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.SendProposePresentationV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.SendProposePresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) AcceptProposePresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.AcceptProposePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.AcceptProposePresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.AcceptProposePresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) AcceptProposePresentationV3(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.AcceptProposePresentationV3CommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.AcceptProposePresentationV3CommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.AcceptProposePresentationV3CommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) DeclineProposePresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.DeclineProposePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.DeclineProposePresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.DeclineProposePresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) AcceptPresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.AcceptPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.AcceptPresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.AcceptPresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}
func (c *PresentProof) DeclinePresentation(request []byte) []byte {
	ep, ok := c.endpoints[presentproof.DeclinePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", presentproof.DeclinePresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", presentproof.DeclinePresentationCommandMethod, "presentproof", err)})
		return m
	}
	return response
}

type VCWallet struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *VCWallet) CreateProfile(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.CreateProfileCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.CreateProfileCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.CreateProfileCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) UpdateProfile(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.UpdateProfileCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.UpdateProfileCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.UpdateProfileCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) ProfileExists(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.ProfileExistsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.ProfileExistsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.ProfileExistsCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Open(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.OpenCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.OpenCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.OpenCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Close(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.CloseCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.CloseCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.CloseCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Add(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.AddCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.AddCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.AddCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Remove(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.RemoveCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.RemoveCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.RemoveCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Get(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.GetCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.GetCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.GetCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) GetAll(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.GetAllCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.GetAllCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.GetAllCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Query(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.QueryCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.QueryCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.QueryCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Issue(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.IssueCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.IssueCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.IssueCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Prove(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.ProveCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.ProveCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.ProveCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Verify(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.VerifyCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.VerifyCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.VerifyCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Derive(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.DeriveCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.DeriveCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.DeriveCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) CreateKeyPair(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.CreateKeyPairCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.CreateKeyPairCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.CreateKeyPairCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) Connect(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.ConnectCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.ConnectCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.ConnectCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) ProposePresentation(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.ProposePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.ProposePresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.ProposePresentationCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) PresentProof(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.PresentProofCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.PresentProofCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.PresentProofCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) ProposeCredential(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.ProposeCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.ProposeCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.ProposeCredentialCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) RequestCredential(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.RequestCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.RequestCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.RequestCredentialCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}
func (c *VCWallet) ResolveCredentialManifest(request []byte) []byte {
	ep, ok := c.endpoints[vcwallet.ResolveCredentialManifestCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vcwallet.ResolveCredentialManifestCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vcwallet.ResolveCredentialManifestCommandMethod, "vcwallet", err)})
		return m
	}
	return response
}

type VDR struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *VDR) SaveDID(request []byte) []byte {
	ep, ok := c.endpoints[vdr.SaveDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vdr.SaveDIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vdr.SaveDIDCommandMethod, "vdr", err)})
		return m
	}
	return response
}
func (c *VDR) GetDIDRecords(request []byte) []byte {
	ep, ok := c.endpoints[vdr.GetDIDsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vdr.GetDIDsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vdr.GetDIDsCommandMethod, "vdr", err)})
		return m
	}
	return response
}
func (c *VDR) GetDID(request []byte) []byte {
	ep, ok := c.endpoints[vdr.GetDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vdr.GetDIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vdr.GetDIDCommandMethod, "vdr", err)})
		return m
	}
	return response
}
func (c *VDR) ResolveDID(request []byte) []byte {
	ep, ok := c.endpoints[vdr.ResolveDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vdr.ResolveDIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vdr.ResolveDIDCommandMethod, "vdr", err)})
		return m
	}
	return response
}
func (c *VDR) CreateDID(request []byte) []byte {
	ep, ok := c.endpoints[vdr.CreateDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", vdr.CreateDIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", vdr.CreateDIDCommandMethod, "vdr", err)})
		return m
	}
	return response
}

type Verifiable struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *Verifiable) ValidateCredential(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.ValidateCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.ValidateCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.ValidateCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) SaveCredential(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.SaveCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.SaveCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.SaveCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) GetCredential(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.GetCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.GetCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.GetCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) GetCredentialByName(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.GetCredentialByNameCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.GetCredentialByNameCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.GetCredentialByNameCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) GetCredentials(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.GetCredentialsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.GetCredentialsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.GetCredentialsCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) SignCredential(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.SignCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.SignCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.SignCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) DeriveCredential(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.DeriveCredentialCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.DeriveCredentialCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.DeriveCredentialCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) SavePresentation(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.SavePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.SavePresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.SavePresentationCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) GetPresentation(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.GetPresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.GetPresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.GetPresentationCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) GetPresentations(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.GetPresentationsCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.GetPresentationsCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.GetPresentationsCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) GeneratePresentation(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.GeneratePresentationCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.GeneratePresentationCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.GeneratePresentationCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) GeneratePresentationByID(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.GeneratePresentationByIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.GeneratePresentationByIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.GeneratePresentationByIDCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) RemoveCredentialByName(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.RemoveCredentialByNameCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.RemoveCredentialByNameCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.RemoveCredentialByNameCommandMethod, "verifiable", err)})
		return m
	}
	return response
}
func (c *Verifiable) RemovePresentationByName(request []byte) []byte {
	ep, ok := c.endpoints[verifiable.RemovePresentationByNameCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", verifiable.RemovePresentationByNameCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", verifiable.RemovePresentationByNameCommandMethod, "verifiable", err)})
		return m
	}
	return response
}

type BlindedRouting struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

type DidClient struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}

func (c *DidClient) CreateOrbDID(request []byte) []byte {
	ep, ok := c.endpoints[didclient.CreateOrbDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didclient.CreateOrbDIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didclient.CreateOrbDIDCommandMethod, "didclient", err)})
		return m
	}
	return response
}
func (c *DidClient) ResolveOrbDID(request []byte) []byte {
	ep, ok := c.endpoints[didclient.ResolveOrbDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didclient.ResolveOrbDIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didclient.ResolveOrbDIDCommandMethod, "didclient", err)})
		return m
	}
	return response
}
func (c *DidClient) CreatePeerDID(request []byte) []byte {
	ep, ok := c.endpoints[didclient.CreatePeerDIDCommandMethod]
	if !ok {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("no handlers found for controller [%s]", didclient.CreatePeerDIDCommandMethod)})
		return m
	}
	response, err := exec(c.URL, c.Token, c.httpClient, ep, request)
	if err != nil {
		m, _ := json.Marshal(Response{Error: fmt.Sprintf("Aries error \n command: [%s]\n controller: [%s]\n error: [%s]", didclient.CreatePeerDIDCommandMethod, "didclient", err)})
		return m
	}
	return response
}

type MediatorClient struct {
	httpClient httpClient
	endpoints  map[string]endpoint
	URL        string
	Token      string
}
