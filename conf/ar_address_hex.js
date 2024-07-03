function arweaveToEVMBytes(arweaveAddress) {
    // Convert base64url to base64
    var base64Address = base64UrlToBase64(arweaveAddress);
    // Decode base64 to binary string
    var binaryString = atob(base64Address);
    // Convert binary string to byte array
    var bytes = new Uint8Array(binaryString.length);
    for (var i = 0; i < binaryString.length; i++) {
        bytes[i] = binaryString.charCodeAt(i);
    }
    // Convert byte array to hex string
    var hexString = byteArrayToHexString(bytes);
    console.log("0x".concat(hexString));
}

function base64UrlToBase64(base64Url) {
    return base64Url.replace(/-/g, '+').replace(/_/g, '/');
}

function byteArrayToHexString(byteArray) {
    return byteArray.reduce(function (acc, _byte) {
        return acc + _byte.toString(16).padStart(2, '0');
    }, '');
}

arweaveToEVMBytes("")