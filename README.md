# contract-caller

```
./contract-caller.exe caller \
--chain-id 11155111 \
--chain-rpc "https://endpoints.omniatech.io/v1/eth/sepolia/public" \
--master-db-host localhost \
--master-db-port 5432 \
--master-db-user postgres \
--master-db-password your_password \
--master-db-name contract_caller \
--private-key privatekey \
--dapplink-vrf-address 0x54368b23FF0FC70DE6ea08e3d4BaC51393f42f36 \
--dapplink-vrf-factory-address 0xa8a6476f8FFE08E440F974f4dfF02Bc668d34961 \
--caller_address 0x8E4B0d162CF3fB58f195302cce3ac018b471A8D8 \
--slave-db-enable false
```