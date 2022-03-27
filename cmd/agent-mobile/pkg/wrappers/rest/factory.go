package rest

import (
	"fmt"
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
	rfc0593 "github.com/hyperledger/aries-framework-go/pkg/controller/rest/rfc0593"
	vcwallet "github.com/hyperledger/aries-framework-go/pkg/controller/rest/vcwallet"
	vdr "github.com/hyperledger/aries-framework-go/pkg/controller/rest/vdr"
	verifiable "github.com/hyperledger/aries-framework-go/pkg/controller/rest/verifiable"
	api "github.com/trustbloc/agent-sdk/cmd/agent-mobile/pkg/api"
	blindedrouting "github.com/trustbloc/agent-sdk/pkg/controller/rest/blindedrouting"
	didclient "github.com/trustbloc/agent-sdk/pkg/controller/rest/didclient"
	mediatorclient "github.com/trustbloc/agent-sdk/pkg/controller/rest/mediatorclient"
	"net/http"
)

func (a *Aries) GetConnectionController() (api.ConnectionController, error) {
	endpoints, ok := a.endpoints[connection.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", connection.OperationID)
	}
	return &Connection{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetDIDExchangeController() (api.DIDExchangeController, error) {
	endpoints, ok := a.endpoints[didexchange.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", didexchange.OperationID)
	}
	return &DIDExchange{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetIntroduceController() (api.IntroduceController, error) {
	endpoints, ok := a.endpoints[introduce.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", introduce.OperationID)
	}
	return &Introduce{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetIssueCredentialController() (api.IssueCredentialController, error) {
	endpoints, ok := a.endpoints[issuecredential.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", issuecredential.OperationID)
	}
	return &IssueCredential{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetKMSController() (api.KMSController, error) {
	endpoints, ok := a.endpoints[kms.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", kms.OperationID)
	}
	return &KMS{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetLDController() (api.LDController, error) {
	endpoints, ok := a.endpoints[ld.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", ld.OperationID)
	}
	return &LD{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetMediatorController() (api.MediatorController, error) {
	endpoints, ok := a.endpoints[mediator.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", mediator.OperationID)
	}
	return &Mediator{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetMessagingController() (api.MessagingController, error) {
	endpoints, ok := a.endpoints[messaging.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", messaging.OperationID)
	}
	return &Messaging{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetOutofBandController() (api.OutofBandController, error) {
	endpoints, ok := a.endpoints[outofband.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", outofband.OperationID)
	}
	return &OutofBand{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetOutofBandV2Controller() (api.OutofBandV2Controller, error) {
	endpoints, ok := a.endpoints[outofbandv2.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", outofbandv2.OperationID)
	}
	return &OutofBandV2{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetPresentProofController() (api.PresentProofController, error) {
	endpoints, ok := a.endpoints[presentproof.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", presentproof.OperationID)
	}
	return &PresentProof{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetRfc0593Controller() (api.Rfc0593Controller, error) {
	endpoints, ok := a.endpoints[rfc0593.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", rfc0593.OperationID)
	}
	return &Rfc0593{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetVCWalletController() (api.VCWalletController, error) {
	endpoints, ok := a.endpoints[vcwallet.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", vcwallet.OperationID)
	}
	return &VCWallet{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetVDRController() (api.VDRController, error) {
	endpoints, ok := a.endpoints[vdr.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", vdr.OperationID)
	}
	return &VDR{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetVerifiableController() (api.VerifiableController, error) {
	endpoints, ok := a.endpoints[verifiable.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", verifiable.OperationID)
	}
	return &Verifiable{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetBlindedRoutingController() (api.BlindedRoutingController, error) {
	endpoints, ok := a.endpoints[blindedrouting.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", blindedrouting.OperationID)
	}
	return &BlindedRouting{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetDidClientController() (api.DidClientController, error) {
	endpoints, ok := a.endpoints[didclient.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", didclient.OperationID)
	}
	return &DidClient{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
func (a *Aries) GetMediatorClientController() (api.MediatorClientController, error) {
	endpoints, ok := a.endpoints[mediatorclient.OperationID]
	if !ok {
		return nil, fmt.Errorf("no handlers found for controller [%s]", mediatorclient.OperationID)
	}
	return &MediatorClient{
		Token:      a.Token,
		URL:        a.URL,
		endpoints:  endpoints,
		httpClient: &http.Client{},
	}, nil
}
