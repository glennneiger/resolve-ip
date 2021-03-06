## Modules

<dl>
<dt><a href="#module_resolve-ip">resolve-ip</a></dt>
<dd><p>resolve-ip client library.</p>
</dd>
</dl>

## Functions

<dl>
<dt><a href="#responseLog">responseLog()</a></dt>
<dd><p>Request status log is used to
to output the status of a request returned
by the client.</p>
</dd>
</dl>

<a name="module_resolve-ip"></a>

## resolve-ip
resolve-ip client library.


* [resolve-ip](#module_resolve-ip)
    * [ResolveIP](#exp_module_resolve-ip--ResolveIP) ⏏
        * [new ResolveIP(options)](#new_module_resolve-ip--ResolveIP_new)
        * _instance_
            * [.healthCheck([options], [cb])](#module_resolve-ip--ResolveIP+healthCheck) ⇒ <code>Promise</code>
            * [.locationForIP(ip, [options], [cb])](#module_resolve-ip--ResolveIP+locationForIP) ⇒ <code>Promise</code>
        * _static_
            * [.RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)
                * [.Exponential](#module_resolve-ip--ResolveIP.RetryPolicies.Exponential)
                * [.Single](#module_resolve-ip--ResolveIP.RetryPolicies.Single)
                * [.None](#module_resolve-ip--ResolveIP.RetryPolicies.None)
            * [.Errors](#module_resolve-ip--ResolveIP.Errors)
                * [.BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest) ⇐ <code>Error</code>
                * [.InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError) ⇐ <code>Error</code>
                * [.NotFound](#module_resolve-ip--ResolveIP.Errors.NotFound) ⇐ <code>Error</code>
            * [.DefaultCircuitOptions](#module_resolve-ip--ResolveIP.DefaultCircuitOptions)

<a name="exp_module_resolve-ip--ResolveIP"></a>

### ResolveIP ⏏
resolve-ip client

**Kind**: Exported class  
<a name="new_module_resolve-ip--ResolveIP_new"></a>

#### new ResolveIP(options)
Create a new client object.


| Param | Type | Default | Description |
| --- | --- | --- | --- |
| options | <code>Object</code> |  | Options for constructing a client object. |
| [options.address] | <code>string</code> |  | URL where the server is located. Must provide this or the discovery argument |
| [options.discovery] | <code>bool</code> |  | Use clever-discovery to locate the server. Must provide this or the address argument |
| [options.timeout] | <code>number</code> |  | The timeout to use for all client requests, in milliseconds. This can be overridden on a per-request basis. |
| [options.retryPolicy] | <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code> | <code>RetryPolicies.Single</code> | The logic to determine which requests to retry, as well as how many times to retry. |
| [options.logger] | <code>module:kayvee.Logger</code> | <code>logger.New(&quot;resolve-ip-wagclient&quot;)</code> | The Kayvee  logger to use in the client. |
| [options.circuit] | <code>Object</code> |  | Options for constructing the client's circuit breaker. |
| [options.circuit.forceClosed] | <code>bool</code> |  | When set to true the circuit will always be closed. Default: true. |
| [options.circuit.maxConcurrentRequests] | <code>number</code> |  | the maximum number of concurrent requests the client can make at the same time. Default: 100. |
| [options.circuit.requestVolumeThreshold] | <code>number</code> |  | The minimum number of requests needed before a circuit can be tripped due to health. Default: 20. |
| [options.circuit.sleepWindow] | <code>number</code> |  | how long, in milliseconds, to wait after a circuit opens before testing for recovery. Default: 5000. |
| [options.circuit.errorPercentThreshold] | <code>number</code> |  | the threshold to place on the rolling error rate. Once the error rate exceeds this percentage, the circuit opens. Default: 90. |

<a name="module_resolve-ip--ResolveIP+healthCheck"></a>

#### resolveIP.healthCheck([options], [cb]) ⇒ <code>Promise</code>
Checks if the service is healthy

**Kind**: instance method of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  
**Fulfill**: <code>undefined</code>  
**Reject**: <code>[BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest)</code>  
**Reject**: <code>[InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError)</code>  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.span] | <code>[Span](https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html)</code> | An OpenTracing span - For example from the parent request |
| [options.retryPolicy] | <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code> | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_resolve-ip--ResolveIP+locationForIP"></a>

#### resolveIP.locationForIP(ip, [options], [cb]) ⇒ <code>Promise</code>
Gets the lat/lon for a given IP.

**Kind**: instance method of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  
**Fulfill**: <code>Object</code>  
**Reject**: <code>[BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest)</code>  
**Reject**: <code>[NotFound](#module_resolve-ip--ResolveIP.Errors.NotFound)</code>  
**Reject**: <code>[InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError)</code>  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| ip | <code>string</code> | The IP to try to locate |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.span] | <code>[Span](https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html)</code> | An OpenTracing span - For example from the parent request |
| [options.retryPolicy] | <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code> | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_resolve-ip--ResolveIP.RetryPolicies"></a>

#### ResolveIP.RetryPolicies
Retry policies available to use.

**Kind**: static property of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  

* [.RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)
    * [.Exponential](#module_resolve-ip--ResolveIP.RetryPolicies.Exponential)
    * [.Single](#module_resolve-ip--ResolveIP.RetryPolicies.Single)
    * [.None](#module_resolve-ip--ResolveIP.RetryPolicies.None)

<a name="module_resolve-ip--ResolveIP.RetryPolicies.Exponential"></a>

##### RetryPolicies.Exponential
The exponential retry policy will retry five times with an exponential backoff.

**Kind**: static constant of <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code>  
<a name="module_resolve-ip--ResolveIP.RetryPolicies.Single"></a>

##### RetryPolicies.Single
Use this retry policy to retry a request once.

**Kind**: static constant of <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code>  
<a name="module_resolve-ip--ResolveIP.RetryPolicies.None"></a>

##### RetryPolicies.None
Use this retry policy to turn off retries.

**Kind**: static constant of <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code>  
<a name="module_resolve-ip--ResolveIP.Errors"></a>

#### ResolveIP.Errors
Errors returned by methods.

**Kind**: static property of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  

* [.Errors](#module_resolve-ip--ResolveIP.Errors)
    * [.BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest) ⇐ <code>Error</code>
    * [.InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError) ⇐ <code>Error</code>
    * [.NotFound](#module_resolve-ip--ResolveIP.Errors.NotFound) ⇐ <code>Error</code>

<a name="module_resolve-ip--ResolveIP.Errors.BadRequest"></a>

##### Errors.BadRequest ⇐ <code>Error</code>
BadRequest

**Kind**: static class of <code>[Errors](#module_resolve-ip--ResolveIP.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_resolve-ip--ResolveIP.Errors.InternalError"></a>

##### Errors.InternalError ⇐ <code>Error</code>
InternalError

**Kind**: static class of <code>[Errors](#module_resolve-ip--ResolveIP.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_resolve-ip--ResolveIP.Errors.NotFound"></a>

##### Errors.NotFound ⇐ <code>Error</code>
NotFound

**Kind**: static class of <code>[Errors](#module_resolve-ip--ResolveIP.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_resolve-ip--ResolveIP.DefaultCircuitOptions"></a>

#### ResolveIP.DefaultCircuitOptions
Default circuit breaker options.

**Kind**: static constant of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  
<a name="responseLog"></a>

## responseLog()
Request status log is used to
to output the status of a request returned
by the client.

**Kind**: global function  
