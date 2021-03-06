# baseApp

`baseApp` requires stores to be mounted via capabilities keys - handlers can only access stores they're given the key to. The `baseApp` ensures all stores are properly loaded, cached, and committed. One mounted store is considered the "main" - it holds the latest block header, from which we can find and load the most recent state.

`baseApp` distinguishes between two handler types - the `AnteHandler` and the `MsgHandler`. The former is a global validity check (checking nonces, sigs and sufficient balances to pay fees,
e.g. things that apply to all transaction from all modules), the later is the full state transition function.
During CheckTx the state transition function is only applied to the checkTxState and should return
before any expensive state transitions are run (this is up to each developer). It also needs to return the estimated
gas cost.

During DeliverTx the state transition function is applied to the blockchain state and the transactions
need to be fully executed.

BaseApp is responsible for managing the context passed into handlers - it makes the block header available and provides the right stores for CheckTx and DeliverTx. BaseApp is completely agnostic to serialization formats.