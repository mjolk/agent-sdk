package command

import (
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
	api "github.com/trustbloc/agent-sdk/cmd/agent-mobile/pkg/api"
	blindedrouting "github.com/trustbloc/agent-sdk/pkg/controller/command/blindedrouting"
	didclient "github.com/trustbloc/agent-sdk/pkg/controller/command/didclient"
	mediatorclient "github.com/trustbloc/agent-sdk/pkg/controller/command/mediatorclient"
)

func (a *Aries) GetConnectionController() (api.ConnectionController, error) {
	handlers, ok := a.handlers[connection.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", connection.CommandName)
	}
	return &Connection{handlers: handlers}, nil
}
func (a *Aries) GetDIDExchangeController() (api.DIDExchangeController, error) {
	handlers, ok := a.handlers[didexchange.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", didexchange.CommandName)
	}
	return &DIDExchange{handlers: handlers}, nil
}
func (a *Aries) GetIntroduceController() (api.IntroduceController, error) {
	handlers, ok := a.handlers[introduce.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", introduce.CommandName)
	}
	return &Introduce{handlers: handlers}, nil
}
func (a *Aries) GetIssueCredentialController() (api.IssueCredentialController, error) {
	handlers, ok := a.handlers[issuecredential.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", issuecredential.CommandName)
	}
	return &IssueCredential{handlers: handlers}, nil
}
func (a *Aries) GetKMSController() (api.KMSController, error) {
	handlers, ok := a.handlers[kms.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", kms.CommandName)
	}
	return &KMS{handlers: handlers}, nil
}
func (a *Aries) GetLDController() (api.LDController, error) {
	handlers, ok := a.handlers[ld.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", ld.CommandName)
	}
	return &LD{handlers: handlers}, nil
}
func (a *Aries) GetMediatorController() (api.MediatorController, error) {
	handlers, ok := a.handlers[mediator.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", mediator.CommandName)
	}
	return &Mediator{handlers: handlers}, nil
}
func (a *Aries) GetMessagingController() (api.MessagingController, error) {
	handlers, ok := a.handlers[messaging.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", messaging.CommandName)
	}
	return &Messaging{handlers: handlers}, nil
}
func (a *Aries) GetOutOfBandController() (api.OutOfBandController, error) {
	handlers, ok := a.handlers[outofband.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", outofband.CommandName)
	}
	return &OutOfBand{handlers: handlers}, nil
}
func (a *Aries) GetOutOfBandV2Controller() (api.OutOfBandV2Controller, error) {
	handlers, ok := a.handlers[outofbandv2.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", outofbandv2.CommandName)
	}
	return &OutOfBandV2{handlers: handlers}, nil
}
func (a *Aries) GetPresentProofController() (api.PresentProofController, error) {
	handlers, ok := a.handlers[presentproof.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", presentproof.CommandName)
	}
	return &PresentProof{handlers: handlers}, nil
}
func (a *Aries) GetVCWalletController() (api.VCWalletController, error) {
	handlers, ok := a.handlers[vcwallet.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", vcwallet.CommandName)
	}
	return &VCWallet{handlers: handlers}, nil
}
func (a *Aries) GetVDRController() (api.VDRController, error) {
	handlers, ok := a.handlers[vdr.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", vdr.CommandName)
	}
	return &VDR{handlers: handlers}, nil
}
func (a *Aries) GetVerifiableController() (api.VerifiableController, error) {
	handlers, ok := a.handlers[verifiable.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", verifiable.CommandName)
	}
	return &Verifiable{handlers: handlers}, nil
}
func (a *Aries) GetBlindedRoutingController() (api.BlindedRoutingController, error) {
	handlers, ok := a.handlers[blindedrouting.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", blindedrouting.CommandName)
	}
	return &BlindedRouting{handlers: handlers}, nil
}
func (a *Aries) GetDidClientController() (api.DidClientController, error) {
	handlers, ok := a.handlers[didclient.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", didclient.CommandName)
	}
	return &DidClient{handlers: handlers}, nil
}
func (a *Aries) GetMediatorClientController() (api.MediatorClientController, error) {
	handlers, ok := a.handlers[mediatorclient.CommandName]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", mediatorclient.CommandName)
	}
	return &MediatorClient{handlers: handlers}, nil
}
