create table if not exists transactions
(
    id               bigserial primary key,
    chain_id         int          not null,
    tx_hash          varchar(255) not null,
    block_number     bigint       not null,
    type             varchar(255) not null,
    contract_address varchar(255),
    status           varchar(255)          default 'FAILED',
    tx_at            timestamp    not null,
    processed_at     timestamp    not null default now(),
    created_at       timestamp    not null default now(),
    updated_at       timestamp    not null default now()
);

