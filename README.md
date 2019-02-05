# Modular IOTA node written in go

This implementation of a IOTA node tries to achieve the maximum possible performance of GO while at the same time offering maximum modularity.

The code is structured into the following  

**Package:**

This is a "normal" go package that contains logic for a certain topic and that exposes instanciable types (i.e. "network/tcp")

**IXI Module:**

A package that exposes a singleton API (i.e. "parameter" module) but that does not directly modify the behavior of the node

**IXI Plugin:**

A package that directly modifies the "behaviour" of a node (i.e. "gossip" plugin)
