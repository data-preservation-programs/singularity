# Deal Making Prerequisite

## Find storage providers

As of now, Singularity does not help you find a storage provider that accepts your deal. There are lots of resources you can use to find good quality storage providers, i.e.

* TODO

## Create a Filecoin Wallet

You must create a Filecoin Wallet before making deals. You can't use a ledger wallet or an exchange wallet. To create a Filecoin wallet, you can run below command

```sh
singularity wallet create
```

This will generate a wallet address as well as the private key associated with the wallet. This wallet cannot be used for deal making yet because it is not recognized by the blockchain. Now it's a good time to transfer 0 FIL to this wallet so everybody knows about it.&#x20;

Once this wallet is recorded on chain, the above command will complete and you're ready for dealmaking.

Alternatively, if you already have an existing wallet, you can import them using

```sh
singularity wallet import xxx
```

## \[Optional] Get [datacap](https://docs.filecoin.io/basics/how-storage-works/filecoin-plus/#datacap)

With the current market condition, most storage providers would prefer [verified deals](https://docs.filecoin.io/storage-provider/filecoin-deals/verified-deals/) in contrast to regular deals. If your dataset is more than a few TiB, it's best to apply for datacap with [Filplus govenance team and notaries](https://github.com/filecoin-project/notary-governance)

## Next step

[create-a-deal-schedule.md](create-a-deal-schedule.md "mention")
