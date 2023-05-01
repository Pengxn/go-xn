// error/message/warning container
let errorContainer, messageContainer, warningContainer;

document.addEventListener('DOMContentLoaded', () => {
    errorContainer = document.getElementById("error");
    messageContainer = document.getElementById("message");
    warningContainer = document.getElementById("warning");

    document.getElementById("signup").addEventListener("click", e => {
        e.preventDefault()
        const username = document.getElementById("username").value
        const displayName = document.getElementById("displayName").value

        submit(username, displayName)
    })
})

showError = (error) => {
    console.error(error);
    errorContainer.innerText = error;
    errorContainer.style.display = "block";
}
showMessage = (message) => {
    console.info(message);
    messageContainer.innerText = message;
    messageContainer.style.display = "block";
}
showWarning = (warning) => {
    console.warn(warning);
    warningContainer.innerText = warning;
    warningContainer.style.display = "block";
}

// webauthn register
submit = (username, displayName) => {
    resp = fetch("/admin/register/begin", {
        method: "POST",
        headers: {"Content-Type": "application/x-www-form-urlencoded"},
        body: "username="+username+"&displayName="+displayName,
    }).then(resp => {
        resp.json().then(data => {
            creation = data.creation
            creation.publicKey.challenge = coerceToArrayBuffer(creation.publicKey.challenge)
            creation.publicKey.user.id = coerceToArrayBuffer(creation.publicKey.user.id)
            navigator.credentials.create(creation).then(credential => {
                crea = JSON.stringify({
                    id: credential.id,
                    type: credential.type,
                    // rawId: new TextDecoder("utf-8").decode(new Uint8Array(credential.rawId)),
                    rawId: coerceToBase64Url(new Uint8Array(credential.rawId)),
                    response: {
                        AttestationObject: coerceToBase64Url(new Uint8Array(credential.response.attestationObject)),
                        clientDataJson: coerceToBase64Url(new Uint8Array(credential.response.clientDataJSON)),
                        clientDataJsonString: new TextDecoder("utf-8").decode(new Uint8Array(credential.response.clientDataJSON))
                    }
                })
                data = JSON.stringify({
                    "username": username,
                    "displayName": displayName,
                    "credential": JSON.parse(crea)
                })
                console.log(data)
                fetch("/admin/register/finish", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: data
                }).then(resp => {
                    resp.json().then(data => {
                        showWarning(JSON.stringify(data))
                    })
                }).catch(err => showError(err))
            }).catch(err => showError(err))
        })
    }).catch(err => showError(err))
}

// ArrayBuffer <=> String 转换
coerceToBase64Url = value => {
    const uint8Array = (() => {
        if (Array.isArray(value)) return Uint8Array.from(value);
        if (value instanceof ArrayBuffer) return new Uint8Array(value);
        if (value instanceof Uint8Array) return value;
        console.warn('Could not coerce to string:', value);
        throw new Error('Could not coerce to string');
    })();

    let string = '';
    for (let i = 0; i < uint8Array.byteLength; i++) {
        string += String.fromCharCode(uint8Array[i]);
    }

    const base64String = btoa(string);
    return base64ToBase64Url(base64String);
}
base64ToBase64Url = base64 => {
    return base64.replace(/\+/g, '-').replace(/\//g, '_').replace(/=*$/g, '');
}
coerceToArrayBuffer = value => {
    if (typeof value === 'string') {
        const base64 = base64UrlToBase64(value);
        const string = atob(base64);
        const bytes = new Uint8Array(string.length);
        for (let i = 0; i < string.length; i++) {
            bytes[i] = string.charCodeAt(i);
        }

        return bytes;
    }

    console.warn('Could not coerce to string:', value);
    throw new TypeError('Could not coerce to ArrayBuffer');
}
base64UrlToBase64 = base64Url => {
    return base64Url.replace(/-/g, '+').replace(/_/g, '/');
}
