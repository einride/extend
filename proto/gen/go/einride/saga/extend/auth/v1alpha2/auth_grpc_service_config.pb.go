// Code generated by protoc-gen-go-grpc-service-config. DO NOT EDIT.
package authv1alpha2

// DefaultServiceConfig is the default service config for all services in the package.
// Source: einride/saga/extend/auth/v1alpha2/default_service_config.proto.
const DefaultServiceConfig = `{
  "methodConfig": [
    {
      "name": [
        {}
      ],
      "timeout": "10s",
      "retryPolicy": {
        "maxAttempts": 5,
        "initialBackoff": "0.200s",
        "maxBackoff": "60s",
        "backoffMultiplier": 2,
        "retryableStatusCodes": [
          "UNAVAILABLE",
          "UNKNOWN"
        ]
      }
    }
  ]
}`
