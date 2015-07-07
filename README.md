# Example Route Service

An example route service for Cloud Foundry.

## Route Service Overview

The Route Service feature is currently in development, the proposal can be found in this [Google Doc](https://docs.google.com/document/d/1bGOQxiKkmaw6uaRWGd-sXpxL0Y28d3QihcluI15FiIA/edit#heading=h.8djffzes9pnb).

This example route service uses the new headers/features that have been added to the GoRouter. For example:

- `X-CF-Forwarded-URL`: A header that contains the original URL that the GoRouter received.
- `X-CF-Proxy-Signature`: A header that the GoRouter uses to determine if a request has gone through the route service.

## Getting Started

- Download this repository and `cf push` to your chosen CF deployment.
- Push your app which will be associated with the route service.
- Use the [rtr CLI](https://github.com/cloudfoundry-incubator/routing-api-cli) to register this example as the `route_service_url` of your chosen app.
- Tail the logs of this route service in order to verify that requests to your app go through the route service. The example logging route service will log requests and responses to and from your app.
