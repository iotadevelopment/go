# Modular IOTA node written in go

Package: a "normal" go package that contains logic for a certain topic and that exposes instanciable types (i.e. "network/tcp")

IXI Module: a package that exposes a singleton API (i.e. "parameter" module) but that does not directly modify the behavior of the node

IXI Plugin: a package that directly modified the "behaviour" of a node (i.e. "gossip" plugin)
