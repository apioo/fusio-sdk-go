
// ConsumerTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    
    
    "github.com/apioo/sdkgen-go"
    
    "net/http"
    
    
)

type ConsumerTag struct {
    internal *sdkgen.TagAbstract
}

func (client *ConsumerTag) Identity() *ConsumerIdentityTag {
    return NewConsumerIdentityTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Account() *ConsumerAccountTag {
    return NewConsumerAccountTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Transaction() *ConsumerTransactionTag {
    return NewConsumerTransactionTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Webhook() *ConsumerWebhookTag {
    return NewConsumerWebhookTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Token() *ConsumerTokenTag {
    return NewConsumerTokenTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Scope() *ConsumerScopeTag {
    return NewConsumerScopeTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Plan() *ConsumerPlanTag {
    return NewConsumerPlanTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Payment() *ConsumerPaymentTag {
    return NewConsumerPaymentTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Page() *ConsumerPageTag {
    return NewConsumerPageTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Log() *ConsumerLogTag {
    return NewConsumerLogTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Grant() *ConsumerGrantTag {
    return NewConsumerGrantTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) Event() *ConsumerEventTag {
    return NewConsumerEventTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *ConsumerTag) App() *ConsumerAppTag {
    return NewConsumerAppTag(client.internal.HttpClient, client.internal.Parser)
}




func NewConsumerTag(httpClient *http.Client, parser *sdkgen.Parser) *ConsumerTag {
	return &ConsumerTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
