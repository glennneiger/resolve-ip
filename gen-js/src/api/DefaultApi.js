/**
 * resolve-ip
 * service that takes in an IP address and converts it to a latitude and longitude
 *
 * OpenAPI spec version: 0.1.2
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/IP', 'opentracing'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/IP'), require('opentracing'));
  } else {
    // Browser globals (root is window)
    if (!root.resolve-ip) {
      root.resolve-ip = {};
    }
    root.resolve-ip.DefaultApi = factory(root.resolve-ip.ApiClient, root.resolve-ip.IP, root.resolve-ip.opentracing);
  }
}(this, function(ApiClient, IP, opentracing) {
  'use strict';

  /**
   * Default service.
   * @module api/DefaultApi
   * @version 0.1.2
   */

  /**
   * Constructs a new DefaultApi. 
   * @alias module:api/DefaultApi
   * @class
   * @param {module:ApiClient} apiClient Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   * @param {Object} options Options for outgoing requests.
   * @param {Object} options.req (Required) Inbound request that triggered this request.
   */
  var exports = function(apiClient, options) {
    this.apiClient = apiClient || ApiClient.instance;
    if (!options || !options.req) {
      throw new Error("req is a required option when constructing an DefaultApi object");
    }
    var span = options.req.span;


    /**
     * Checks if the service is healthy
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}
     */
    this.healthCheck = function() {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var headerParams = {
      };
      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headerParams);
        span.logEvent("GET /healthcheck");
      }
      var formParams = {
      };

      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = null;
      var promise = this.apiClient.callApi(
        '/healthcheck', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
      if (span) {
        // Log a trace event when the request finished
        promise.then(function() {
          span.logEvent("GET /healthcheck finished successfully");
        }, function() {
          span.logEvent("GET /healthcheck failed"); // TODO: with response code
        });
      }
      return promise;
    }


    /**
     * Gets the lat/lon for a given IP.
     * @param {String} ip The IP to try to locate
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/IP}
     */
    this.locationForIP = function(ip) {
      var postBody = null;

      // verify the required parameter 'ip' is set
      if (ip == undefined || ip == null) {
        throw "Missing the required parameter 'ip' when calling locationForIP";
      }


      var pathParams = {
        'ip': ip
      };
      var queryParams = {
      };
      var headerParams = {
      };
      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headerParams);
        span.logEvent("GET /ip/{ip}");
      }
      var formParams = {
      };

      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = IP;
      var promise = this.apiClient.callApi(
        '/ip/{ip}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
      if (span) {
        // Log a trace event when the request finished
        promise.then(function() {
          span.logEvent("GET /ip/{ip} finished successfully");
        }, function() {
          span.logEvent("GET /ip/{ip} failed"); // TODO: with response code
        });
      }
      return promise;
    }
  };

  return exports;
}));
