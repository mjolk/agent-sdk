package api

type AriesController interface {
	GetConnectionController() ConnectionController
	GetDIDExchangeController() DIDExchangeController
	GetIntroduceController() IntroduceController
	GetIssueCredentialController() IssueCredentialController
	GetKMSController() KMSController
	GetLDController() LDController
	GetMediatorController() MediatorController
	GetMessagingController() MessagingController
	GetOutofBandController() OutofBandController
	GetOutofBandV2Controller() OutofBandV2Controller
	GetPresentProofController() PresentProofController
	GetRfc0593Controller() Rfc0593Controller
	GetVCWalletController() VCWalletController
	GetVDRController() VDRController
	GetVerifiableController() VerifiableController
	GetBlindedRoutingController() BlindedRoutingController
	GetDidClientController() DidClientController
	GetMediatorClientController() MediatorClientController
}
