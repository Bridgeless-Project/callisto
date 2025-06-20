## Version v3.3.1

### Changes

#### Gov Module
- ([\#11](https://github.com/Bridgeless-Project/callisto/pull/11)) Added fetching proposal metadata via a link specified in the proposal's metadata field. As `x/gov` v1.MsgSubmitProposal does not support metadata fields directly, they are listed in a metadata file accessible via the provided link.


## Version v3.3.0
- ([\#702](https://github.com/forbole/bdjuno/pull/702)) Add `message_type` module and store msg types inside `message_type` table, add `messages_by_type` function to allow to query messages by their types

### Changes

#### CI
- ([\#508](https://github.com/forbole/bdjuno/pull/508)) Upgrade workflow golangci version to v1.50.1
- ([\#2](https://github.com/Bridgeless-Project/callisto/pull/2)) Added CI/CD on develop branch
- ([\#3](https://github.com/Bridgeless-Project/callisto/pull/3)) Fixed Dockerfile
- ([\#6](https://github.com/Bridgeless-Project/callisto/pull/6)) Fixed Dockerfile.vendor

#### Parse Command
- ([\#492](https://github.com/forbole/bdjuno/pull/492)) Add parse command for periodic tasks: `x/bank` total supply, `x/distribution` community pool, `x/mint` inflation, `pricefeed` token price and price history, `x/staking` staking pool

#### Upgrade Module
- ([\#467](https://github.com/forbole/bdjuno/pull/467)) Store software upgrade plan and refresh data at upgrade height

#### Staking Module
- ([\#443](https://github.com/forbole/bdjuno/pull/443)) Remove tombstone status from staking module(already stored in slashing module)
- ([\#455](https://github.com/forbole/bdjuno/pull/455)) Added `unbonding_tokens` and `staked_not_bonded_tokens` values to staking pool table
- ([\#536](https://github.com/forbole/bdjuno/pull/536) Fix `PoolSnapshot` tokens type from  `sdk.Int` to `sdkmath.Int`

#### Gov Module
- ([\#461](https://github.com/forbole/bdjuno/pull/461)) Parse `x/gov` genesis with `genesisDoc.InitialHeight` instead of the hard-coded height 1
- ([\#465](https://github.com/forbole/bdjuno/pull/465)) Get open proposal ids in deposit or voting period by block time instead of current time
- ([\#489](https://github.com/forbole/bdjuno/pull/489)) Remove block height foreign key from proposal_vote and proposal_deposit tables and add column timestamp
- ([\#499](https://github.com/forbole/bdjuno/pull/499)) Check if proposal has passed voting end time before marking it invalid
- ([\#523](https://github.com/forbole/bdjuno/pull/523)) Update proposal snapshots handling on block
- ([\#681](https://github.com/forbole/bdjuno/pull/681)) Handle proposal status change from deposit to voting
- ([\#2](https://github.com/Bridgeless-Project/callisto/pull/2)) Gov params update. Use the core app instead simd app
- ([\#8](https://github.com/Bridgeless-Project/callisto/pull/8)) Gov msg handler update. Added support for `v1` and `v1beta1` message types

#### Bridge Module
- ([\#4](https://github.com/Bridgeless-Project/callisto/pull/4)) Updated module to handle updated transaction struct and genesis
- ([\#6](https://github.com/Bridgeless-Project/callisto/pull/6)) Updated module to handle transactions_submissions

#### Mint Module
- ([\#7](https://github.com/Bridgeless-Project/callisto/pull/7)) Updated module to update params each block

#### Daily refetch
- ([\#454](https://github.com/forbole/bdjuno/pull/454)) Added `daily refetch` module to refetch missing blocks every day

#### Hasura
- ([\#473](https://github.com/forbole/bdjuno/pull/473)) Improved Hasura permissions
- ([\#491](https://github.com/forbole/bdjuno/pull/491)) Add host address to Hasura actions
- ([\#1](https://github.com/Bridgeless-Project/callisto/pull/1)) Updated the relevant Hasura migrations.


### Database and migrations
- ([\#1](https://github.com/Bridgeless-Project/callisto/pull/1)) Adjusted migration scripts to simplify future database updates
- ([\#4](https://github.com/Bridgeless-Project/callisto/pull/4)) Update `x/bridge` module database schema
- ([\#6](https://github.com/Bridgeless-Project/callisto/pull/6)) Added `transactions_submissions` table to `x/bridge` module database schema
- ([\#8](https://github.com/Bridgeless-Project/callisto/pull/8)) Updated `x/gov` proposals table schema

### Dependencies
- ([\#462](https://github.com/forbole/bdjuno/pull/462)) Updated Juno to `v3.4.0`
- ([\#542](https://github.com/forbole/bdjuno/pull/542)) Updated Juno to `v4.1.0`,  BDJuno to `v4` and Golang version to `1.19`
- ([\#3](https://github.com/Bridgeless-Project/callisto/pull/3)) Bump cosmos-sdk version to `v0.46.30`
- ([\#4](https://github.com/Bridgeless-Project/callisto/pull/4)) Fixed broken `rosetta-sdk` import
- ([\#4](https://github.com/Bridgeless-Project/callisto/pull/4)) Updated bridgeless-core version.
- ([\#5](https://github.com/Bridgeless-Project/callisto/pull/5)) Bump cosmos-sdk version to `v0.46.32`
- ([\#6](https://github.com/Bridgeless-Project/callisto/pull/6)) Bump ibc-go version to `v6.1.8`
- ([\#7](https://github.com/Bridgeless-Project/callisto/pull/7)) Updated cosmos-sdk version to `v0.46.33`
- ([\#7](https://github.com/Bridgeless-Project/callisto/pull/7)) Updated go version to `1.23`
- ([\#7](https://github.com/Bridgeless-Project/callisto/pull/7)) Updated ibc-go version to `v6.1.9`
- ([\#7](https://github.com/Bridgeless-Project/callisto/pull/7)) Updated juno version to `v4.0.2`
- ([\#8](https://github.com/Bridgeless-Project/callisto/pull/8)) Updated bridgeless-core version to `v12.1.19`

### Modules
- ([\#1](https://github.com/Bridgeless-Project/callisto/pull/1)) Implemented `x/bridge`, `x/nft`, `x/accumulator` modules
- ([\#2](https://github.com/Bridgeless-Project/callisto/pull/2)) Implemented `x/multisig` module

### Fixes
- ([\#1](https://github.com/Bridgeless-Project/callisto/pull/1)) Resolved an issue with the deprecated package go.tmz.dev/musttag
- ([\#2](https://github.com/Bridgeless-Project/callisto/pull/2)) Fixed hasura metadata error

## Version v3.2.0
### Changes
#### Mint module
- ([\#432](https://github.com/forbole/bdjuno/pull/432)) Update inflation rate when mint param change proposal is passed

#### Gov module
- ([\#401](https://github.com/forbole/bdjuno/pull/401)) Update the proposal status to the latest in `bdjuno parse gov proposal [id]` command
- ([\#430](https://github.com/forbole/bdjuno/pull/430)) Update the proposals that have invalid status but can still be in voting or deposit periods 

### Dependencies
- ([\#440](https://github.com/forbole/bdjuno/pull/440)) Updated Juno to `v3.3.0`

## Version v3.1.0
### Dependencies
- Updated Juno to `v3.2.0`

### Changes 
#### Hasura
- ([\#395](https://github.com/forbole/bdjuno/pull/395)) Remove time label from Hasura Prometheus monitoring

#### Bank module
- ([\#410](https://github.com/forbole/bdjuno/pull/410)) Change total supply query from only 1 page to all pages

## Version v3.0.1
### Dependencies
- Updated Juno to `v3.1.1`

## Version v3.0.0
### Notes
This version introduces breaking changes to `transaction` and `message` PostgreSQL tables. It implements PostgreSQL table partitioning to fix slow data retrieval from database that stores large amount of transactions and messages. Read more details about [migrating to v3.0.0](https://docs.bigdipper.live/cosmos-based/parser/migrations/v2.0.0)

### New features 
#### CLI
- ([\#356](https://github.com/forbole/bdjuno/pull/356)) Implemented `migrate` command to perform easy migration to higher BDJuno versions
- ([\#356](https://github.com/forbole/bdjuno/pull/356)) Updated `parse-genesis` command to parse genesis file without accessing the node

#### Database
- ([\#356](https://github.com/forbole/bdjuno/pull/356)) Added PostgreSQL table partition to `transaction` and `message` table
- ([\#356](https://github.com/forbole/bdjuno/pull/356)) Created new `messages_by_address` function

### Changes 
#### Hasura
- ([\#377](https://github.com/forbole/bdjuno/pull/377)) Updated Hasura metadata
- ([\#381](https://github.com/forbole/bdjuno/pull/381)) Hasura actions are now a module 

### Dependencies
- ([\#356](https://github.com/forbole/bdjuno/pull/356)) Updated Juno to `v3.0.0`

## Version v2.0.0
### Notes
This version introduces breaking changes to certain address-specific data that is no longer periodically parsed from the node and stored in the database. Instead, the data is now obtained directly from the node when needed using Hasura Actions. Read more details about [migrating to v2.0.0](https://docs.bigdipper.live/cosmos-based/parser/migrations/v2.0.0)

### New features
#### CLI
- ([\#257](https://github.com/forbole/bdjuno/pull/257)) Added `parse-genesis` command to parse the genesis file
- ([\#228](https://github.com/forbole/bdjuno/pull/228)) ([\#248](https://github.com/forbole/bdjuno/pull/248)) Added `fix` command:
  - `auth`: fix vesting accounts details
  - `blocks`: fix missing blocks and transactions from given start height
  - `gov`: fix proposal with given proposal ID  
  - `staking`: fix validators info at the latest height  

#### Hasura Actions
- ([\#329](https://github.com/forbole/bdjuno/pull/329)) Implemented Hasura Actions service to replace periodic queries. If you are using GraphQL queries on your application, you should updated the old queries to use the below new actions instead.  
  Here's a list of data acquired through Hasura Actions:
    - Of a certain address/delegator:
      - Account balance (`action_account_balance`)
      - Delegation rewards (`action_delegation_reward`)
      - Delegator withdraw address (`action_delegator_withdraw_address`)
      - Delegations (`action_delegation`)
      - Total delegations amount (`action_delegation_total`)
      - Unbonding delegations (`action_unbonding_delegation`)
      - Total unbonding delegations amount (`action_unbonding_delegation_total`)
      - Redelegations (`action_redelegation`)
    - Of a certain validator:
      - Commission amount (`action_validator_commission_amount`)
      - Delegations to this validator (`action_validator_delegations`)
      - Redelegations from this validator (`action_validator_redelegations_from`)
      - Unbonding delegations (`action_validator_unbonding_delegations`)
- ([\#352](https://github.com/forbole/bdjuno/pull/352)) Added prometheus monitoring to hasura actions

#### Local node support
- Added the support for `node.type = "local"` for parsing a static local node without the usage gRPC queries: [config reference](https://docs.bigdipper.live/cosmos-based/parser/config/config#node).

#### Modules
- ([\#232](https://github.com/forbole/bdjuno/pull/232)) Updated the `x/auth` module support to handle and store `vesting accounts` and `vesting periods` inside the database. 
- ([\#276](https://github.com/forbole/bdjuno/pull/276)) Added the support for the `x/feegrant` module (v0.44.x)

### Changes 

#### CLI
- ([\#351](https://github.com/forbole/bdjuno/pull/351)) Fixed version display for `bdjuno version` cmd 

#### Database
- ([\#300](https://github.com/forbole/bdjuno/pull/300)) Changed `bonded_tokens` and `not_bonded_tokens` type inside `staking_pool` table  to `TEXT` to avoid value overflow
- ([\#275](https://github.com/forbole/bdjuno/pull/275)) Added `tombstoned` column inside `validator_status` table
- ([\#232](https://github.com/forbole/bdjuno/pull/232)) Added `vesting_account` and `vesting_period` table
- ([\#276](https://github.com/forbole/bdjuno/pull/276)) Added `fee_grant_allowance` table (v0.44.x)

#### Modules
- ([\#353](https://github.com/forbole/bdjuno/pull/353)) Removed the support for the `history` module
