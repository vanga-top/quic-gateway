## Design

### features

* http proxy gateway
* service function gateway (support for gRpc/spring-cloud etc.)
* use QUIC protocol for communicating to client
* support graceful restart and update
* stateless and distribute server-cluster



### arch



### details

#### server arch

* server interface to control server activities

  * server start/stop/stop right now/restart/update
  *   

* server plugins to build for more protocols and others
* 
