
// SystemTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    
    
    
    
    
    "net/http"
    
    
)

type SystemTag struct {
    internal *sdkgen.TagAbstract
}

func (client *SystemTag) Payment() *SystemPaymentTag {
    return NewSystemPaymentTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *SystemTag) Meta() *SystemMetaTag {
    return NewSystemMetaTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *SystemTag) Connection() *SystemConnectionTag {
    return NewSystemConnectionTag(client.internal.HttpClient, client.internal.Parser)
}




func NewSystemTag(httpClient *http.Client, parser *sdkgen.Parser) *SystemTag {
	return &SystemTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
