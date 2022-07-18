package api

type ConnectionController interface {
	RotateDID([]byte) []byte
	CreateConnectionV2([]byte) []byte
	SetConnectionToDIDCommV2([]byte) []byte
}
type DIDExchangeController interface {
	AcceptExchangeRequest([]byte) []byte
	AcceptInvitation([]byte) []byte
	CreateImplicitInvitation([]byte) []byte
	CreateInvitation([]byte) []byte
	QueryConnectionByID([]byte) []byte
	QueryConnections([]byte) []byte
	ReceiveInvitation([]byte) []byte
	CreateConnection([]byte) []byte
	RemoveConnection([]byte) []byte
}
type IntroduceController interface {
	Actions([]byte) []byte
	SendProposal([]byte) []byte
	SendProposalWithOOBInvitation([]byte) []byte
	SendRequest([]byte) []byte
	AcceptProposalWithOOBInvitation([]byte) []byte
	AcceptProposal([]byte) []byte
	AcceptRequestWithPublicOOBInvitation([]byte) []byte
	AcceptRequestWithRecipients([]byte) []byte
	DeclineProposal([]byte) []byte
	DeclineRequest([]byte) []byte
	AcceptProblemReport([]byte) []byte
}
type IssueCredentialController interface {
	Actions([]byte) []byte
	SendOffer([]byte) []byte
	SendProposal([]byte) []byte
	SendProposalV3([]byte) []byte
	SendRequest([]byte) []byte
	SendRequestV3([]byte) []byte
	AcceptProposal([]byte) []byte
	AcceptProposalV3([]byte) []byte
	DeclineProposal([]byte) []byte
	AcceptOffer([]byte) []byte
	DeclineOffer([]byte) []byte
	NegotiateProposal([]byte) []byte
	NegotiateProposalV3([]byte) []byte
	AcceptRequest([]byte) []byte
	AcceptRequestV3([]byte) []byte
	DeclineRequest([]byte) []byte
	AcceptCredential([]byte) []byte
	DeclineCredential([]byte) []byte
	AcceptProblemReport([]byte) []byte
}
type KMSController interface {
	CreateKeySet([]byte) []byte
	ImportKey([]byte) []byte
}
type LDController interface {
	AddContexts([]byte) []byte
	AddRemoteProvider([]byte) []byte
	RefreshRemoteProvider([]byte) []byte
	DeleteRemoteProvider([]byte) []byte
	GetAllRemoteProviders([]byte) []byte
	RefreshAllRemoteProviders([]byte) []byte
}
type MediatorController interface {
	Register([]byte) []byte
	Unregister([]byte) []byte
	Connections([]byte) []byte
	Reconnect([]byte) []byte
	Status([]byte) []byte
	BatchPickup([]byte) []byte
	ReconnectAll([]byte) []byte
}
type MessagingController interface {
	Services([]byte) []byte
	RegisterService([]byte) []byte
	UnregisterService([]byte) []byte
	RegisterHTTPService([]byte) []byte
	Send([]byte) []byte
	Reply([]byte) []byte
}
type OutOfBandController interface {
	CreateInvitation([]byte) []byte
	AcceptInvitation([]byte) []byte
	ActionStop([]byte) []byte
	Actions([]byte) []byte
	ActionContinue([]byte) []byte
}
type OutOfBandV2Controller interface {
	CreateInvitation([]byte) []byte
	AcceptInvitation([]byte) []byte
}
type PresentProofController interface {
	Actions([]byte) []byte
	SendRequestPresentation([]byte) []byte
	SendRequestPresentationV3([]byte) []byte
	AcceptRequestPresentation([]byte) []byte
	AcceptRequestPresentationV3([]byte) []byte
	NegotiateRequestPresentation([]byte) []byte
	NegotiateRequestPresentationV3([]byte) []byte
	AcceptProblemReport([]byte) []byte
	DeclineRequestPresentation([]byte) []byte
	SendProposePresentation([]byte) []byte
	SendProposePresentationV3([]byte) []byte
	AcceptProposePresentation([]byte) []byte
	AcceptProposePresentationV3([]byte) []byte
	DeclineProposePresentation([]byte) []byte
	AcceptPresentation([]byte) []byte
	DeclinePresentation([]byte) []byte
}
type VCWalletController interface {
	CreateProfile([]byte) []byte
	UpdateProfile([]byte) []byte
	ProfileExists([]byte) []byte
	Open([]byte) []byte
	Close([]byte) []byte
	Add([]byte) []byte
	Remove([]byte) []byte
	Get([]byte) []byte
	GetAll([]byte) []byte
	Query([]byte) []byte
	Issue([]byte) []byte
	Prove([]byte) []byte
	Verify([]byte) []byte
	Derive([]byte) []byte
	CreateKeyPair([]byte) []byte
	Connect([]byte) []byte
	ProposePresentation([]byte) []byte
	PresentProof([]byte) []byte
	ProposeCredential([]byte) []byte
	RequestCredential([]byte) []byte
	ResolveCredentialManifest([]byte) []byte
}
type VDRController interface {
	SaveDID([]byte) []byte
	GetDIDRecords([]byte) []byte
	GetDID([]byte) []byte
	ResolveDID([]byte) []byte
	CreateDID([]byte) []byte
}
type VerifiableController interface {
	ValidateCredential([]byte) []byte
	SaveCredential([]byte) []byte
	GetCredential([]byte) []byte
	GetCredentialByName([]byte) []byte
	GetCredentials([]byte) []byte
	SignCredential([]byte) []byte
	DeriveCredential([]byte) []byte
	SavePresentation([]byte) []byte
	GetPresentation([]byte) []byte
	GetPresentations([]byte) []byte
	GeneratePresentation([]byte) []byte
	GeneratePresentationByID([]byte) []byte
	RemoveCredentialByName([]byte) []byte
	RemovePresentationByName([]byte) []byte
}
type BlindedRoutingController interface{}
type DidClientController interface {
	CreateOrbDID([]byte) []byte
	ResolveOrbDID([]byte) []byte
	CreatePeerDID([]byte) []byte
}
type MediatorClientController interface{}
