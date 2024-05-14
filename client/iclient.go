package client

import ltime "time"

type IClient interface {
	WithHttpClient(httpClient IHttpClient) IClient
}

type IHttpClient interface {
	WithRetryCount(pretryCount int) IHttpClient
	WithDelay(pdelay ltime.Duration) IHttpClient
	WithSleep(psleep ltime.Duration) IHttpClient
}
