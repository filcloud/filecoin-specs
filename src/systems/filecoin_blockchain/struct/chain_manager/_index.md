---
title: Chain Manager
---

The _Chain Manager_ is a central component in the blockchain system. It tracks and updates competing subchains received by a given node in order to select the appropriate blockchain head: the latest block of the heaviest subchain it is aware of in the system.

In so doing, the _chain manager_ is the central subsystem that handles bookkeeping for numerous other systems in a Filecoin node and exposes convenience methods for use by those systems, enabling systems to sample randomness from the chain for instance, or to see which block has been finalized most recently.

The chain manager interfaces and functions are included here, but we expand on important details below for clarity.


# Chain Expansion

## Incoming blocks and semantic validation

Once a block has been received and syntactically validated by the {{<sref chainsync>}}, it must be semantically validated by the chain manager for inclusion on a given chain.

A semantically valid block:

- must be from a valid miner with power on the chain
- must only have valid parents in the tipset, meaning
  - that each parent itself must be a valid block
  - that they all have the same parents themselves
  - that they are all at the same height (i.e. include the same number of tickets)
- must have valid tickets:
  - the first ticket in its tickets array was generated from the smallest ticket in the parent tipset
  - any following ticket in the array was generated from the prior one (meaning all tickets in the ticket array have been generated by the same miner)
- must have a valid timestamp, meaning
  - must be later than the latest parent block timestamp, and must have been mined within the appropriate mining period


  must be later than the earliest parent block time plus appreopriate delay, which is BLOCK_DELAY (30s by default) * len(blk.Tickets).
  see [BLock Validation](mining.md#block-validation)
- must only have valid state transitions:
  - all messages in the block must be valid
  - the execution of each message, in the order they are in the block, must produce a receipt matching the corresponding one in the receipt set of the block, see [the state machine spec](state-machine.md).
- the resulting state root after all messages are applied, must match the one in the block


{{% notice info %}}
Once the block passes validation, it must be added to the local datastore, regardless whether it is understood as the best tip at this point. Future blocks from other miners may be mined on top of it and in that case we will want to have it around to avoid refetching.
{{% /notice %}}

{{% notice info %}}
To make certain validation checks simpler, blocks should be indexed by height and by parent set. That way sets of blocks with a given height and common parents may be quickly queried. It may also be useful to compute and cache the resultant aggregate state of blocks in these sets, this saves extra state computation when checking which state root to start a block at when it has multiple parents.
{{% /notice %}}

The following requires having and processing (executing) the messages

- Messages can be checked by verifying the messages hash correctly to the value.
- MessageAggregateSig can be checked by verifying the messages sign correctly
- MessageReceipts can only be checked by executing the messages
- StateRoot is the result of the execution of the messages, and can only be verified by executing them

## Block reception algorithm

Chain selection is a crucial component of how the Filecoin blockchain works. Every chain has an associated weight accounting for the number of blocks mined on it and so the power (storage) they track. It is always preferable to mine atop a heavier Tipset rather than a lighter one. While a miner may be foregoing block rewards earned in the past, this lighter chain is likely to be abandoned by other miners forfeiting any block reward earned as miners converge on a final chain. For more on this, see [chain selection](expected-consensus.md#chain-selection) in the Expected Consensus spec.

However, ahead of finality, a given subchain may be abandoned in order of another, heavier one mined in a given round. In order to rapidly adapt to this, the chain manager must maintain and update all subchains being considered up to finality.

That is, for every incoming block, even if the incoming block is not added to the current heaviest tipset, the chain manager should add it to the appropriate subchain it is tracking, or keep track of it independently until either:
- it is able to do so, through the reception of another block in that subchain
- it is able to discard it, as that block was mined before finality

We give an example of how this could work in the block reception algorithm.

### ChainTipsManager

The Chain Tips Manager is a subcomponent of Filecoin consensus that is technically up to the implementer, but since the pseudocode in previous sections reference it, it is documented here for clarity.

The Chain Tips Manager is responsible for tracking all live tips of the Filecoin blockchain, and tracking what the current 'best' tipset is.

```go
// Returns the ticket that is at round 'r' in the chain behind 'head'
func TicketFromRound(head Tipset, r Round) {}

// Returns the tipset that contains round r (Note: multiple rounds' worth of tickets may exist within a single block due to losing tickets being added to the eventually successfully generated block)
func TipsetFromRound(head Tipset, r Round) {}

// GetBestTipset returns the best known tipset. If the 'best' tipset hasn't changed, then this
// will return the previous best tipset.
func GetBestTipset()

// Adds the losing ticket to the chaintips manager so that blocks can be mined on top of it
func AddLosingTicket(parent Tipset, t Ticket)
```
