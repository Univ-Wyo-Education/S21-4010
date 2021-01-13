#!/bin/bash

# UTC--2018-03-11T01-56-42.511489695Z--9a6446d642d76a3ac1baf6c6d8c1e5179c58d87f

geth attach "http://192.154.97.75:8545/" <<XXxx
personal.unlockAccount( "0x9a6446d642d76a3ac1baf6c6d8c1e5179c58d87f", "xtwHdVIsKNFdMOcICZz5O+1xs5Q", 150000);
exit
XXxx

