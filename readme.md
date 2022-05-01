# loan

**loan** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://starport.com).

## Get started

```
starport chain serve -r
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Add loan

```
loand tx loan request-loan 100token 2token 200token 500 --from bob -y
```

### Approve loan

```
loand tx loan approve-loan 0 --from alice -y
```

### Query bank module

```
loand query bank balances <alice_address>
```

### Repay loan

```
loand tx loan repay-loan 0 --from bob -y
```

### Liquidate loan

```
loand tx loan liquidate-loan 0 --from alice -y
```

### Cancel loan

```
loand tx loan cancel-loan 0 --from bob -y
```
