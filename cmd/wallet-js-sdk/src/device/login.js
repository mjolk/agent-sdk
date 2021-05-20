/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

import axios from 'axios';

const client = axios.create({
    withCredentials: true
})

/**
 * DeviceRegister provides device registration features
 * @param agent instance
 * @class
 */

export class DeviceLogin {
    constructor(agent) {
        this.agent = agent
        this.serverURL = this.agent['edge-agent-server']
    }

    async login() {
        client.get(`${this.serverURL}/device/login/begin`,
            null,
            function (data) {
                return data
            },
            'json')
            .then((credentialRequestOptions) => {
                credentialRequestOptions.data.publicKey.challenge = bufferDecode(credentialRequestOptions.data.publicKey.challenge);
                credentialRequestOptions.data.publicKey.allowCredentials.forEach(function (listItem) {
                    listItem.id = bufferDecode(listItem.id)
                });

                return navigator.credentials.get({
                    publicKey: credentialRequestOptions.data.publicKey
                })
            })
            .then((assertion) => {
                let {rawId} = assertion
                let {authenticatorData, clientDataJSON, signature, userHandle} = assertion.response

                client.post(`${this.serverURL}/device/login/finish`,
                    JSON.stringify({
                        id: assertion.id,
                        rawId: bufferEncode(rawId),
                        type: assertion.type,
                        response: {
                            authenticatorData: bufferEncode(authenticatorData),
                            clientDataJSON: bufferEncode(clientDataJSON),
                            signature: bufferEncode(signature),
                            userHandle: bufferEncode(userHandle),
                        },
                    }),
                    function (data) {
                        return data
                    },
                    'json')
            })
            // eslint-disable-next-line no-unused-vars
            .then((success) => {
                window.location.href = "/dashboard";
            })
            .catch((error) => {
                console.error(error);
            })
    }
}

function bufferDecode(value) {
    return Uint8Array.from(atob(value), c => c.charCodeAt(0));
}


function bufferEncode(value) {
    return btoa(String.fromCharCode.apply(null, new Uint8Array(value)))
        .replace(/\+/g, "-")
        .replace(/\//g, "_")
        .replace(/=/g, "");
}
