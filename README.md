# MpesaOverlay

MpesaOverlay is modern, scalable, secure open source and patent-free API overlay for the Daraja API by Safaricom written in Go.

It accepts connections over various API (i.e. REST, GraphQL, GRPC, Pub/Sub), thus making a seamless bridge between the client and the Daraja API. It is used as a middleware for building complex Payments solutions.

**Features**
- Interface bridging (i.e. REST, GraphQL, GRPC, Pub/Sub)
- Fine-grained access control
- Platform logging and instrumentation support
- Container-based deployment using Docker

## Architecture

### Components

MpesaOverlay platform is comprised of the following services:
- users	Manages platform's users and auth concerns
- things	Manages platform's things, channels and access policies
- rest-adapter	Provides a REST interface for accessing communication channels
- graphql-adapter	Provides a GraphQL interface for accessing communication channels
- grpc-adapter	Provides a gRPC interface for accessing communication channels
- pubsub-adapter	Provides a Pub/Sub interface for accessing communication channels
- mainflux-cli	Command line interface

![Mpesa Overlay Architecture](assets/architecture.png)
