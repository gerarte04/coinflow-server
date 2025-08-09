CREATE TABLE coinflow_ch_db.transactions (
    id              UUID,
    user_id         UUID,

    type            String,
    target          String,
    description     String,
    category        String,
    cost            Float64,

    timestamp       DateTime
)
ENGINE = MaterializedPostgreSQL('postgres:5432', 'coinflow_pg_db', 'transactions', 'admin', 'adminpass')
ORDER BY (user_id, timestamp)
SETTINGS allow_experimental_materialized_postgresql_table = 1;

CREATE ROW POLICY cdc_policy ON coinflow_ch_db.transactions FOR SELECT USING _peerdb_is_deleted = 0 TO admin;
