CREATE TABLE IF NOT EXISTS trades
(
    id                       UUID     DEFAULT generateUUIDv4(),
    chain_id                 Int32 NOT NULL,
    pool_type                String,
    initiator_address        String,
    agent_address            String,
    from_token_address       String NOT NULL,
    to_token_address         String NOT NULL,
    from_token_amount        Decimal(64, 20),
    to_token_amount          Decimal(64, 20),
    eth_value                Decimal(64, 20),
    price                    Decimal(64, 20),
    current_base_pool_amount Decimal(64, 20),
    tx_hash                  String NOT NULL,
    block_number             Int64,
    contract_address         String,
    tx_at                    DateTime NOT NULL,
    created_at               DateTime DEFAULT now(),
    updated_at               DateTime DEFAULT now(),
    PRIMARY                  KEY (id)
) ENGINE = MergeTree()
      ORDER BY id;

