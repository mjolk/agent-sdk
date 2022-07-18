package rest

import (
	connection1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/connection"
	didexchange1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/didexchange"
	introduce1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/introduce"
	issuecredential1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/issuecredential"
	kms1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/kms"
	ld1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/ld"
	mediator1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/mediator"
	messaging1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/messaging"
	outofband1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/outofband"
	outofbandv21 "github.com/hyperledger/aries-framework-go/pkg/controller/command/outofbandv2"
	presentproof1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/presentproof"
	vcwallet1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/vcwallet"
	vdr1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/vdr"
	verifiable1 "github.com/hyperledger/aries-framework-go/pkg/controller/command/verifiable"
	connection "github.com/hyperledger/aries-framework-go/pkg/controller/rest/connection"
	didexchange "github.com/hyperledger/aries-framework-go/pkg/controller/rest/didexchange"
	introduce "github.com/hyperledger/aries-framework-go/pkg/controller/rest/introduce"
	issuecredential "github.com/hyperledger/aries-framework-go/pkg/controller/rest/issuecredential"
	kms "github.com/hyperledger/aries-framework-go/pkg/controller/rest/kms"
	ld "github.com/hyperledger/aries-framework-go/pkg/controller/rest/ld"
	mediator "github.com/hyperledger/aries-framework-go/pkg/controller/rest/mediator"
	messaging "github.com/hyperledger/aries-framework-go/pkg/controller/rest/messaging"
	outofband "github.com/hyperledger/aries-framework-go/pkg/controller/rest/outofband"
	outofbandv2 "github.com/hyperledger/aries-framework-go/pkg/controller/rest/outofbandv2"
	presentproof "github.com/hyperledger/aries-framework-go/pkg/controller/rest/presentproof"
	vcwallet "github.com/hyperledger/aries-framework-go/pkg/controller/rest/vcwallet"
	vdr "github.com/hyperledger/aries-framework-go/pkg/controller/rest/vdr"
	verifiable "github.com/hyperledger/aries-framework-go/pkg/controller/rest/verifiable"
	didclient1 "github.com/trustbloc/agent-sdk/pkg/controller/command/didclient"
	blindedrouting "github.com/trustbloc/agent-sdk/pkg/controller/rest/blindedrouting"
	didclient "github.com/trustbloc/agent-sdk/pkg/controller/rest/didclient"
	"net/http"
)

type endpoint struct {
	Path   string
	Method string
}

func EndPoints() map[string]map[string]endpoint {
	endpoints := make(map[string]map[string]endpoint)

	endpoints[connection.OperationID] = map[string]endpoint{
		connection1.CreateV2CommandMethod: {
			Method: http.MethodPost,
			Path:   connection.CreateV2,
		},
		connection1.RotateDIDCommandMethod: {
			Method: http.MethodPost,
			Path:   connection.RotateDID,
		},
		connection1.SetToV2CommandMethod: {
			Method: http.MethodPost,
			Path:   connection.SetToV2,
		},
	}
	endpoints[didexchange.OperationID] = map[string]endpoint{
		didexchange1.AcceptExchangeRequestCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.AcceptExchangeRequest,
		},
		didexchange1.AcceptInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.AcceptInvitation,
		},
		didexchange1.CreateConnectionCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.CreateConnection,
		},
		didexchange1.CreateImplicitInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.CreateImplicitInvitation,
		},
		didexchange1.CreateInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.CreateInvitation,
		},
		didexchange1.QueryConnectionByIDCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.QueryConnectionByID,
		},
		didexchange1.QueryConnectionsCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.QueryConnections,
		},
		didexchange1.ReceiveInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.ReceiveInvitation,
		},
		didexchange1.RemoveConnectionCommandMethod: {
			Method: http.MethodPost,
			Path:   didexchange.RemoveConnection,
		},
	}
	endpoints[introduce.OperationID] = map[string]endpoint{
		introduce1.AcceptProblemReportCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.AcceptProblemReport,
		},
		introduce1.AcceptProposalCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.AcceptProposal,
		},
		introduce1.AcceptProposalWithOOBInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.AcceptProposalWithOOBInvitation,
		},
		introduce1.AcceptRequestWithPublicOOBInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.AcceptRequestWithPublicOOBInvitation,
		},
		introduce1.AcceptRequestWithRecipientsCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.AcceptRequestWithRecipients,
		},
		introduce1.ActionsCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.Actions,
		},
		introduce1.DeclineProposalCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.DeclineProposal,
		},
		introduce1.DeclineRequestCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.DeclineRequest,
		},
		introduce1.SendProposalCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.SendProposal,
		},
		introduce1.SendProposalWithOOBInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.SendProposalWithOOBInvitation,
		},
		introduce1.SendRequestCommandMethod: {
			Method: http.MethodPost,
			Path:   introduce.SendRequest,
		},
	}
	endpoints[issuecredential.OperationID] = map[string]endpoint{
		issuecredential1.AcceptCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.AcceptCredential,
		},
		issuecredential1.AcceptOfferCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.AcceptOffer,
		},
		issuecredential1.AcceptProblemReportCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.AcceptProblemReport,
		},
		issuecredential1.AcceptProposalCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.AcceptProposal,
		},
		issuecredential1.AcceptProposalV3CommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.AcceptProposalV3,
		},
		issuecredential1.AcceptRequestCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.AcceptRequest,
		},
		issuecredential1.AcceptRequestV3CommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.AcceptRequestV3,
		},
		issuecredential1.ActionsCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.Actions,
		},
		issuecredential1.DeclineCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.DeclineCredential,
		},
		issuecredential1.DeclineOfferCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.DeclineOffer,
		},
		issuecredential1.DeclineProposalCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.DeclineProposal,
		},
		issuecredential1.DeclineRequestCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.DeclineRequest,
		},
		issuecredential1.NegotiateProposalCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.NegotiateProposal,
		},
		issuecredential1.NegotiateProposalV3CommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.NegotiateProposalV3,
		},
		issuecredential1.SendOfferCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.SendOffer,
		},
		issuecredential1.SendProposalCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.SendProposal,
		},
		issuecredential1.SendProposalV3CommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.SendProposalV3,
		},
		issuecredential1.SendRequestCommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.SendRequest,
		},
		issuecredential1.SendRequestV3CommandMethod: {
			Method: http.MethodPost,
			Path:   issuecredential.SendRequestV3,
		},
	}
	endpoints[kms.OperationID] = map[string]endpoint{
		kms1.CreateKeySetCommandMethod: {
			Method: http.MethodPost,
			Path:   kms.CreateKeySet,
		},
		kms1.ImportKeyCommandMethod: {
			Method: http.MethodPost,
			Path:   kms.ImportKey,
		},
	}
	endpoints[ld.OperationID] = map[string]endpoint{
		ld1.AddContextsCommandMethod: {
			Method: http.MethodPost,
			Path:   ld.AddContexts,
		},
		ld1.AddRemoteProviderCommandMethod: {
			Method: http.MethodPost,
			Path:   ld.AddRemoteProvider,
		},
		ld1.DeleteRemoteProviderCommandMethod: {
			Method: http.MethodPost,
			Path:   ld.DeleteRemoteProvider,
		},
		ld1.GetAllRemoteProvidersCommandMethod: {
			Method: http.MethodPost,
			Path:   ld.GetAllRemoteProviders,
		},
		ld1.RefreshAllRemoteProvidersCommandMethod: {
			Method: http.MethodPost,
			Path:   ld.RefreshAllRemoteProviders,
		},
		ld1.RefreshRemoteProviderCommandMethod: {
			Method: http.MethodPost,
			Path:   ld.RefreshRemoteProvider,
		},
	}
	endpoints[mediator.OperationID] = map[string]endpoint{
		mediator1.BatchPickupCommandMethod: {
			Method: http.MethodPost,
			Path:   mediator.BatchPickup,
		},
		mediator1.GetConnectionsCommandMethod: {
			Method: http.MethodPost,
			Path:   mediator.GetConnections,
		},
		mediator1.ReconnectAllCommandMethod: {
			Method: http.MethodPost,
			Path:   mediator.ReconnectAll,
		},
		mediator1.ReconnectCommandMethod: {
			Method: http.MethodPost,
			Path:   mediator.Reconnect,
		},
		mediator1.RegisterCommandMethod: {
			Method: http.MethodPost,
			Path:   mediator.Register,
		},
		mediator1.StatusCommandMethod: {
			Method: http.MethodPost,
			Path:   mediator.Status,
		},
		mediator1.UnregisterCommandMethod: {
			Method: http.MethodPost,
			Path:   mediator.Unregister,
		},
	}
	endpoints[messaging.OperationID] = map[string]endpoint{
		messaging1.RegisterHTTPMessageServiceCommandMethod: {
			Method: http.MethodPost,
			Path:   messaging.RegisterHTTPMessageService,
		},
		messaging1.RegisterMessageServiceCommandMethod: {
			Method: http.MethodPost,
			Path:   messaging.RegisterMessageService,
		},
		messaging1.RegisteredServicesCommandMethod: {
			Method: http.MethodPost,
			Path:   messaging.RegisteredServices,
		},
		messaging1.SendNewMessageCommandMethod: {
			Method: http.MethodPost,
			Path:   messaging.SendNewMessage,
		},
		messaging1.SendReplyMessageCommandMethod: {
			Method: http.MethodPost,
			Path:   messaging.SendReplyMessage,
		},
		messaging1.UnregisterMessageServiceCommandMethod: {
			Method: http.MethodPost,
			Path:   messaging.UnregisterMessageService,
		},
	}
	endpoints[outofband.OperationID] = map[string]endpoint{
		outofband1.AcceptInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   outofband.AcceptInvitation,
		},
		outofband1.ActionContinueCommandMethod: {
			Method: http.MethodPost,
			Path:   outofband.ActionContinue,
		},
		outofband1.ActionStopCommandMethod: {
			Method: http.MethodPost,
			Path:   outofband.ActionStop,
		},
		outofband1.ActionsCommandMethod: {
			Method: http.MethodPost,
			Path:   outofband.Actions,
		},
		outofband1.CreateInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   outofband.CreateInvitation,
		},
	}
	endpoints[outofbandv2.OperationID] = map[string]endpoint{
		outofbandv21.AcceptInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   outofbandv2.AcceptInvitation,
		},
		outofbandv21.CreateInvitationCommandMethod: {
			Method: http.MethodPost,
			Path:   outofbandv2.CreateInvitation,
		},
	}
	endpoints[presentproof.OperationID] = map[string]endpoint{
		presentproof1.AcceptPresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.AcceptPresentation,
		},
		presentproof1.AcceptProblemReportCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.AcceptProblemReport,
		},
		presentproof1.AcceptProposePresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.AcceptProposePresentation,
		},
		presentproof1.AcceptProposePresentationV3CommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.AcceptProposePresentationV3,
		},
		presentproof1.AcceptRequestPresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.AcceptRequestPresentation,
		},
		presentproof1.AcceptRequestPresentationV3CommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.AcceptRequestPresentationV3,
		},
		presentproof1.ActionsCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.Actions,
		},
		presentproof1.DeclinePresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.DeclinePresentation,
		},
		presentproof1.DeclineProposePresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.DeclineProposePresentation,
		},
		presentproof1.DeclineRequestPresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.DeclineRequestPresentation,
		},
		presentproof1.NegotiateRequestPresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.NegotiateRequestPresentation,
		},
		presentproof1.NegotiateRequestPresentationV3CommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.NegotiateRequestPresentationV3,
		},
		presentproof1.SendProposePresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.SendProposePresentation,
		},
		presentproof1.SendProposePresentationV3CommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.SendProposePresentationV3,
		},
		presentproof1.SendRequestPresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.SendRequestPresentation,
		},
		presentproof1.SendRequestPresentationV3CommandMethod: {
			Method: http.MethodPost,
			Path:   presentproof.SendRequestPresentationV3,
		},
	}
	endpoints[vcwallet.OperationID] = map[string]endpoint{
		vcwallet1.AddCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Add,
		},
		vcwallet1.CloseCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Close,
		},
		vcwallet1.ConnectCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Connect,
		},
		vcwallet1.CreateKeyPairCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.CreateKeyPair,
		},
		vcwallet1.CreateProfileCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.CreateProfile,
		},
		vcwallet1.DeriveCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Derive,
		},
		vcwallet1.GetAllCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.GetAll,
		},
		vcwallet1.GetCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Get,
		},
		vcwallet1.IssueCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Issue,
		},
		vcwallet1.OpenCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Open,
		},
		vcwallet1.PresentProofCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.PresentProof,
		},
		vcwallet1.ProfileExistsCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.ProfileExists,
		},
		vcwallet1.ProposeCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.ProposeCredential,
		},
		vcwallet1.ProposePresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.ProposePresentation,
		},
		vcwallet1.ProveCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Prove,
		},
		vcwallet1.QueryCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Query,
		},
		vcwallet1.RemoveCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Remove,
		},
		vcwallet1.RequestCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.RequestCredential,
		},
		vcwallet1.ResolveCredentialManifestCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.ResolveCredentialManifest,
		},
		vcwallet1.UpdateProfileCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.UpdateProfile,
		},
		vcwallet1.VerifyCommandMethod: {
			Method: http.MethodPost,
			Path:   vcwallet.Verify,
		},
	}
	endpoints[vdr.OperationID] = map[string]endpoint{
		vdr1.CreateDIDCommandMethod: {
			Method: http.MethodPost,
			Path:   vdr.CreateDID,
		},
		vdr1.GetDIDCommandMethod: {
			Method: http.MethodPost,
			Path:   vdr.GetDID,
		},
		vdr1.GetDIDsCommandMethod: {
			Method: http.MethodPost,
			Path:   vdr.GetDIDs,
		},
		vdr1.ResolveDIDCommandMethod: {
			Method: http.MethodPost,
			Path:   vdr.ResolveDID,
		},
		vdr1.SaveDIDCommandMethod: {
			Method: http.MethodPost,
			Path:   vdr.SaveDID,
		},
	}
	endpoints[verifiable.OperationID] = map[string]endpoint{
		verifiable1.DeriveCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.DeriveCredential,
		},
		verifiable1.GeneratePresentationByIDCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.GeneratePresentationByID,
		},
		verifiable1.GeneratePresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.GeneratePresentation,
		},
		verifiable1.GetCredentialByNameCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.GetCredentialByName,
		},
		verifiable1.GetCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.GetCredential,
		},
		verifiable1.GetCredentialsCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.GetCredentials,
		},
		verifiable1.GetPresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.GetPresentation,
		},
		verifiable1.GetPresentationsCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.GetPresentations,
		},
		verifiable1.RemoveCredentialByNameCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.RemoveCredentialByName,
		},
		verifiable1.RemovePresentationByNameCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.RemovePresentationByName,
		},
		verifiable1.SaveCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.SaveCredential,
		},
		verifiable1.SavePresentationCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.SavePresentation,
		},
		verifiable1.SignCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.SignCredential,
		},
		verifiable1.ValidateCredentialCommandMethod: {
			Method: http.MethodPost,
			Path:   verifiable.ValidateCredential,
		},
	}
	endpoints[blindedrouting.OperationID] = map[string]endpoint{}
	endpoints[didclient.OperationID] = map[string]endpoint{
		didclient1.CreateOrbDIDCommandMethod: {
			Method: http.MethodPost,
			Path:   didclient.CreateOrbDID,
		},
		didclient1.CreatePeerDIDCommandMethod: {
			Method: http.MethodPost,
			Path:   didclient.CreatePeerDID,
		},
		didclient1.ResolveOrbDIDCommandMethod: {
			Method: http.MethodPost,
			Path:   didclient.ResolveOrbDID,
		},
	}
	return endpoints
}
