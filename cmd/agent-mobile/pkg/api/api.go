package api

type AriesController interface {
	GetConnectionController() (ConnectionController, error)
	GetDIDExchangeController() (DIDExchangeController, error)
	GetIntroduceController() (IntroduceController, error)
	GetIssueCredentialController() (IssueCredentialController, error)
	GetKMSController() (KMSController, error)
	GetLDController() (LDController, error)
	GetMediatorController() (MediatorController, error)
	GetMessagingController() (MessagingController, error)
	GetOutOfBandController() (OutOfBandController, error)
	GetOutOfBandV2Controller() (OutOfBandV2Controller, error)
	GetPresentProofController() (PresentProofController, error)
	GetVCWalletController() (VCWalletController, error)
	GetVDRController() (VDRController, error)
	GetVerifiableController() (VerifiableController, error)
	GetBlindedRoutingController() (BlindedRoutingController, error)
	GetDidClientController() (DidClientController, error)
	GetMediatorClientController() (MediatorClientController, error)
	RegisterHandler(handler Handler, topics string) string
	UnregisterHandler(id string)
}
