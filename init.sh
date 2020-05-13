#!/bin/bash
rm -r ~/.altverifyCLI
rm -r ~/.altverifyD

altverifyD init mynode --chain-id altverify

altverifyCLI config keyring-backend test

altverifyCLI keys add me
altverifyCLI keys add you

altverifyD add-genesis-account $(altverifyCLI keys show me -a) 1000foo,100000000stake
altverifyD add-genesis-account $(altverifyCLI keys show you -a) 1foo

altverifyCLI config chain-id altverify
altverifyCLI config output json
altverifyCLI config indent true
altverifyCLI config trust-node true

altverifyD gentx --name me --keyring-backend test
altverifyD collect-gentxs